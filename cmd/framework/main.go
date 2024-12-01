package main

import (
	"container/heap"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/fatih/color"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/mattn/go-sqlite3"
	"github.com/patrickmn/go-cache"
)

// Структура шага
type Step struct {
	ID            string
	Text          string
	Buttons       []string
	ExpectedInput string // "text", "number", "file", "image"
	Callback      func(*tgbotapi.BotAPI, *tgbotapi.Message)
	NextSteps     map[string]string // ответ -> следующий шаг
	Command       string            // Название команды, например "/start"
}

var steps = map[string]*Step{}

// Кеш для запросов
var requestCache = cache.New(5*time.Minute, 10*time.Minute)

// Мьютекс для синхронизации доступа к общим ресурсам
var mu sync.Mutex

// Структура для хранения активности пользователя
type UserActivity struct {
	UserID   int64
	Activity int
	Index    int // Индекс в heap
}

// Приоритетная очередь для топ-10 пользователей
type ActivityHeap []*UserActivity

func (h ActivityHeap) Len() int           { return len(h) }
func (h ActivityHeap) Less(i, j int) bool { return h[i].Activity > h[j].Activity }
func (h ActivityHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].Index = i
	h[j].Index = j
}

func (h *ActivityHeap) Push(x interface{}) {
	n := len(*h)
	userActivity := x.(*UserActivity)
	userActivity.Index = n
	*h = append(*h, userActivity)
}

func (h *ActivityHeap) Pop() interface{} {
	old := *h
	n := len(old)
	userActivity := old[n-1]
	userActivity.Index = -1 // для безопасности
	*h = old[0 : n-1]
	return userActivity
}

var activityHeap = &ActivityHeap{}
var userActivities = make(map[int64]*UserActivity)

var db *sql.DB

func main() {
	bot, err := tgbotapi.NewBotAPI("8132285126:AAFsXf01JsowcbJoiAOMARphAIlmgYHPbWo")
	if err != nil {
		log.Panic(err)
	}

	// Настраиваем цветные логи
	infoLog := color.New(color.FgGreen).PrintfFunc()
	errorLog := color.New(color.FgRed).PrintfFunc()
	debugLog := color.New(color.FgCyan).PrintfFunc()

	bot.Debug = true

	infoLog("Авторизация бота %s успешно выполнена\n", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	// Инициализируем базу данных
	initializeDatabase()

	// Инициализируем шаги сценария
	initializeSteps()

	userStates := sync.Map{} // для хранения состояний пользователей в потокобезопасной карте

	heap.Init(activityHeap)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		go handleUpdate(bot, update, &userStates, infoLog, errorLog, debugLog)
	}
}

func initializeDatabase() {
	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(color.RedString("Ошибка подключения к базе данных: %v", err))
	}

	// Создаем таблицу records, если она не существует
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS records (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        subject TEXT NOT NULL,
        text TEXT NOT NULL
    );`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(color.RedString("Ошибка при создании таблицы records: %v", err))
	}
}

func handleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update, userStates *sync.Map,
	infoLog func(format string, a ...interface{}), errorLog func(format string, a ...interface{}),
	debugLog func(format string, a ...interface{})) {

	chatID := update.Message.Chat.ID

	debugLog("Получено сообщение от пользователя %d: %s\n", chatID, update.Message.Text)

	// Обновляем активность пользователя
	updateUserActivity(chatID)

	// Проверяем, если это команда
	if update.Message.IsCommand() {
		command := update.Message.Command()

		// Ищем шаг, соответствующий команде
		stepFound := false
		for _, step := range steps {
			if step.Command == "/"+command {
				userStates.Store(chatID, step.ID)
				sendStep(bot, chatID, step)
				infoLog("Пользователь %d начал команду %s, переход на шаг %s\n", chatID, command, step.ID)
				stepFound = true
				break
			}
		}
		if !stepFound {
			msg := tgbotapi.NewMessage(chatID, "Неизвестная команда.")
			bot.Send(msg)
			errorLog("Пользователь %d ввел неизвестную команду: %s\n", chatID, command)
		}
		return
	}

	// Получаем текущее состояние пользователя
	var currentStepID string
	state, exists := userStates.Load(chatID)
	if !exists {
		// Если пользователь не начал диалог, игнорируем сообщение или можно отправить приветствие
		msg := tgbotapi.NewMessage(chatID, "Пожалуйста, введите команду /start, чтобы начать.")
		bot.Send(msg)
		return
	} else {
		currentStepID = state.(string)
	}
	currentStep := steps[currentStepID]

	// Обрабатываем коллбек текущего шага
	if currentStep.Callback != nil {
		currentStep.Callback(bot, update.Message)
	}

	// Переходим к следующему шагу
	var nextStepID string

	if len(currentStep.Buttons) > 0 {
		// Если ожидаем нажатие кнопки
		response := update.Message.Text
		nextStepID = currentStep.NextSteps[response]
	} else {
		// Если ожидаем ввод
		switch currentStep.ExpectedInput {
		case "number":
			_, err := strconv.Atoi(update.Message.Text)
			if err != nil {
				msg := tgbotapi.NewMessage(chatID, "Пожалуйста, введите число.")
				bot.Send(msg)
				return
			}
			nextStepID = currentStep.NextSteps["valid_number"]
		case "text":
			nextStepID = currentStep.NextSteps["any_text"]
		default:
			nextStepID = currentStep.NextSteps["default"]
		}
	}

	// Сохраняем ввод пользователя для последующего использования
	if currentStepID == "enter_subject" || currentStepID == "enter_text" {
		userInputKey := fmt.Sprintf("user_input_%d_%s", chatID, currentStepID)
		requestCache.Set(userInputKey, update.Message.Text, cache.DefaultExpiration)
	}

	// Обновляем состояние пользователя
	if nextStepID != "" {
		userStates.Store(chatID, nextStepID)
		sendStep(bot, chatID, steps[nextStepID])
		infoLog("Пользователь %d перешел на шаг %s\n", chatID, nextStepID)
	} else {
		msg := tgbotapi.NewMessage(chatID, "Некорректный ввод. Попробуйте снова.")
		bot.Send(msg)
		errorLog("Пользователь %d ввел некорректные данные на шаге %s\n", chatID, currentStepID)
	}
}

func initializeSteps() {
	steps["start"] = &Step{
		ID:      "start",
		Text:    "Добро пожаловать! Выберите действие:",
		Buttons: []string{"Показать запись по ID", "Список записей", "Добавить запись"},
		NextSteps: map[string]string{
			"Показать запись по ID": "enter_id",
			"Список записей":        "show_list",
			"Добавить запись":       "enter_subject",
		},
		Command: "/start", // Ассоциируем шаг с командой /start
	}

	steps["enter_id"] = &Step{
		ID:            "enter_id",
		Text:          "Пожалуйста, введите ID записи:",
		ExpectedInput: "number",
		NextSteps: map[string]string{
			"valid_number": "show_record",
		},
	}

	steps["show_record"] = &Step{
		ID:       "show_record",
		Text:     "",
		Callback: showRecordByID,
		NextSteps: map[string]string{
			"default": "start",
		},
	}

	steps["show_list"] = &Step{
		ID:       "show_list",
		Text:     "",
		Callback: showRecordList,
		NextSteps: map[string]string{
			"default": "start",
		},
	}

	// Новые шаги для добавления записи
	steps["enter_subject"] = &Step{
		ID:            "enter_subject",
		Text:          "Пожалуйста, введите тему записи:",
		ExpectedInput: "text",
		NextSteps: map[string]string{
			"any_text": "enter_text",
		},
	}

	steps["enter_text"] = &Step{
		ID:            "enter_text",
		Text:          "Пожалуйста, введите текст записи:",
		ExpectedInput: "text",
		NextSteps: map[string]string{
			"any_text": "confirm_add_record",
		},
	}

	steps["confirm_add_record"] = &Step{
		ID:       "confirm_add_record",
		Callback: addRecord,
		NextSteps: map[string]string{
			"default": "start",
		},
	}

	// Добавляем команду /add для добавления записи
	steps["add_record"] = &Step{
		ID:      "add_record",
		Command: "/add",
		NextSteps: map[string]string{
			"default": "enter_subject",
		},
	}
}

func sendStep(bot *tgbotapi.BotAPI, chatID int64, step *Step) {
	// Если шаг не содержит текст, не отправляем сообщение
	if step.Text != "" {
		msg := tgbotapi.NewMessage(chatID, step.Text)

		if len(step.Buttons) > 0 {
			var buttons [][]tgbotapi.KeyboardButton
			row := []tgbotapi.KeyboardButton{}
			for i, btnText := range step.Buttons {
				row = append(row, tgbotapi.NewKeyboardButton(btnText))
				// Переходим на новую строку после каждых 2 кнопок (или по вашему усмотрению)
				if (i+1)%2 == 0 {
					buttons = append(buttons, row)
					row = []tgbotapi.KeyboardButton{}
				}
			}
			if len(row) > 0 {
				buttons = append(buttons, row)
			}
			keyboard := tgbotapi.NewReplyKeyboard(buttons...)
			msg.ReplyMarkup = keyboard
		} else {
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		}

		bot.Send(msg)
	}

	// Если у шага есть коллбек и он не был вызван ранее, вызываем его
	if step.Callback != nil {
		// Передаем исходное сообщение пользователя в коллбек
		step.Callback(bot, &tgbotapi.Message{Chat: &tgbotapi.Chat{ID: chatID}})
	}
}

func showRecordByID(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	chatID := message.Chat.ID
	id, _ := strconv.Atoi(message.Text)

	cacheKey := fmt.Sprintf("record_%d", id)
	var subject, text string

	// Получаем индекс активности пользователя для TTL
	ttlMultiplier := getUserActivityIndex(chatID)

	// Проверяем кеш
	if cachedData, found := requestCache.Get(cacheKey); found {
		record := cachedData.(map[string]string)
		subject = record["subject"]
		text = record["text"]
	} else {
		err := db.QueryRow("SELECT subject, text FROM records WHERE id = ?", id).Scan(&subject, &text)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(chatID, "Запись не найдена."))
			return
		}

		// Сохраняем в кеш с учетом TTL и индекса активности пользователя
		ttl := time.Duration(ttlMultiplier) * time.Minute
		record := map[string]string{"subject": subject, "text": text}
		requestCache.Set(cacheKey, record, ttl)
	}

	response := fmt.Sprintf("Запись #%d\nТема: %s\nТекст: %s", id, subject, text)
	bot.Send(tgbotapi.NewMessage(chatID, response))
}

func showRecordList(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	chatID := message.Chat.ID

	cacheKey := "record_list"
	var response string

	// Получаем индекс активности пользователя для TTL
	ttlMultiplier := getUserActivityIndex(chatID)

	// Проверяем кеш
	if cachedData, found := requestCache.Get(cacheKey); found {
		response = cachedData.(string)
	} else {
		rows, err := db.Query("SELECT id, subject, text FROM records")
		if err != nil {
			bot.Send(tgbotapi.NewMessage(chatID, "Ошибка при получении списка записей."))
			return
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var subject string
			var text string
			rows.Scan(&id, &subject, &text)
			response += fmt.Sprintf("Запись #%d\nТема: %s\nТекст: %s", id, subject, text)
		}

		if response == "" {
			response = "Нет доступных записей."
		}

		// Сохраняем в кеш с учетом TTL и индекса активности пользователя
		ttl := time.Duration(ttlMultiplier) * time.Minute
		requestCache.Set(cacheKey, response, ttl)
	}

	bot.Send(tgbotapi.NewMessage(chatID, response))
}

func addRecord(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	chatID := message.Chat.ID

	// Получаем введенные пользователем данные из кеша
	subjectKey := fmt.Sprintf("user_input_%d_enter_subject", chatID)
	textKey := fmt.Sprintf("user_input_%d_enter_text", chatID)

	subjectData, foundSubject := requestCache.Get(subjectKey)
	textData, foundText := requestCache.Get(textKey)

	if !foundSubject || !foundText {
		bot.Send(tgbotapi.NewMessage(chatID, "Ошибка: не удалось получить данные для сохранения записи."))
		return
	}

	subject := subjectData.(string)
	text := textData.(string)

	// Вставляем новую запись в базу данных
	insertQuery := "INSERT INTO records (subject, text) VALUES (?, ?)"
	_, err := db.Exec(insertQuery, subject, text)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(chatID, "Ошибка при добавлении записи в базу данных."))
		log.Println(color.RedString("Ошибка при добавлении записи: %v", err))
		return
	}

	// Очищаем сохраненные данные пользователя
	requestCache.Delete(subjectKey)
	requestCache.Delete(textKey)

	// Отправляем сообщение об успешном добавлении
	bot.Send(tgbotapi.NewMessage(chatID, "Запись успешно добавлена!"))

	// Запускаем таймер на 3 секунды, после которого отправим начальное меню
	time.AfterFunc(3*time.Second, func() {
		sendStep(bot, chatID, steps["start"])
	})
}

// Функция для обновления активности пользователя
func updateUserActivity(userID int64) {
	mu.Lock()
	defer mu.Unlock()

	if ua, exists := userActivities[userID]; exists {
		ua.Activity++
		heap.Fix(activityHeap, ua.Index)
	} else {
		ua := &UserActivity{
			UserID:   userID,
			Activity: 1,
		}
		heap.Push(activityHeap, ua)
		userActivities[userID] = ua
	}

	// Поддерживаем размер heap до 10 элементов
	if activityHeap.Len() > 10 {
		removed := heap.Pop(activityHeap).(*UserActivity)
		delete(userActivities, removed.UserID)
	}
}

// Функция для получения индекса активности пользователя
func getUserActivityIndex(userID int64) int {
	mu.Lock()
	defer mu.Unlock()

	if ua, exists := userActivities[userID]; exists {
		// Индекс активности влияет на TTL (например, от 1 до 5)
		index := ua.Activity / 10 // настраиваем коэффициент по желанию
		if index < 1 {
			index = 1
		} else if index > 5 {
			index = 5
		}
		return index
	}
	return 1
}

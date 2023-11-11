package app

import (
	"fmt"
	"tg_bot/models"
	"time"

	"github.com/davecgh/go-spew/spew"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	open_title       = "Открытые заявки"
	open_callback    = "show_open"
	pending_title    = "Заявки на оценку"
	pending_callback = "show_pending"
	update_title     = "Выберите заявку, которую необходимо обновить"
	update_callback  = "update_request"
)

func addCheckStatusButton(buttons []models.Button, title string, number int, callback string) []models.Button {
	if number > 0 {
		buttons = append(buttons, models.Button{Title: fmt.Sprintf("%s (%d)", title, number), Callback: callback})
	}
	return buttons
}

func start(callback *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	Logger.Debug("start")
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "")
	msg.Text = "\n\nВы можете создать, обновить или проверить статус вашей заявки. Для остановки бота просто напишите /stop, для повторного запуска нажмите /start"

	msg.ReplyMarkup = GetInitialKeyboard()
	bot.Send(msg)
}

func newRequest(callback *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI, userData *models.UserState) {
	Logger.Debug("newRequest")
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Введите тему заявки")
	bot.Send(msg)
	userData.CurrentState = "waiting_for_request_topic"
}

func checkStatus(callback *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI, userData *models.UserState) {
	Logger.Debug("checkStatus")
	spew.Dump(userData)
	state_type_count := getStateTypeCount(userData.CustomerUserLogin)

	buttons := []models.Button{}
	buttons = addCheckStatusButton(buttons, open_title, state_type_count.CountOpen, open_callback)
	buttons = addCheckStatusButton(buttons, pending_title, state_type_count.CountPendAuto, pending_callback)

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "")
	msg.Text = "Нет открытых заявок"
	if state_type_count.CountOpen > 0 || state_type_count.CountPendAuto > 0 {
		msg.Text = "Выберите"
		msg.ReplyMarkup = getInlineKeyboard(buttons)
	}
	bot.Send(msg)
}

func listTickets(callback *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI, userData *models.UserState) {
	Logger.Debug("listTickets")

	tickets := getTickets(userData.UserName)

	var title string
	switch userData.Action {
	case "show_open":
		title = open_title
		tickets = getOpenTickets(tickets)
	case "show_pending":
		title = pending_title
		tickets = getPendingTickets(tickets)
	case "update_request":
		title = update_title
	default:
		title = ""
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, title)
	buttons := convertTicketsToButtons(tickets)
	buttons = append(buttons, models.Button{Title: "Назад", Callback: "start"})
	msg.ReplyMarkup = getInlineKeyboard(buttons)
	bot.Send(msg)
}

func attachFile(callback *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Приложите файл до 20Mb")
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Создать без вложения", "create"),
		),
	)
	msg.ReplyMarkup = inlineKeyboard
	bot.Send(msg)
}

func create(callback *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI, userData *models.UserState) {
	Logger.Debug("create")

	if userData.TicketID != "" {
		_, err := otrsRequest("update", models.OtrsRequest{
			TicketID: &userData.TicketID,
			// Ticket:   nil,
			Article: &models.Article{
				CommunicationChannel: "Internal",
				SenderType:           "customer",
				Charset:              "utf-8",
				MimeType:             "text/plain",
				From:                 userData.CustomerUserLogin,
				Subject:              "Telegram message",
				Body:                 userData.Description,
			},
		})

		if err != nil {
			Logger.Warn(err)
			bot.Send(tgbotapi.NewMessage(callback.Message.Chat.ID, "Ошибка при создании заявки."))
			return
		}
		userData = &models.UserState{}

		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Вашe обращение обновлено")
		buttons := []models.Button{
			{Title: "Назад", Callback: "start"},
		}
		msg.ReplyMarkup = getInlineKeyboard(buttons)
		bot.Send(msg)
	} else {
		new_ticket, err := otrsRequest("create", models.OtrsRequest{
			Ticket: &models.Ticket{
				Title:        &userData.Topic,
				QueueID:      &QueueID,
				TypeID:       &TypeID, // Request for service
				CustomerUser: &userData.CustomerUserLogin,
				StateID:      &StateID,    // new
				PriorityID:   &PriorityID, // normal
				OwnerID:      &OwnerID,    // admin
				LockID:       &LockID,     // unlock
			},
			Article: &models.Article{
				CommunicationChannel: "Internal",
				SenderType:           "customer",
				Charset:              "utf-8",
				MimeType:             "text/plain",
				From:                 userData.CustomerUserLogin,
				Subject:              userData.Topic,
				Body:                 userData.Description,
			},
		})

		if err != nil {
			Logger.Warn(err)
			bot.Send(tgbotapi.NewMessage(callback.Message.Chat.ID, "Ошибка при создании заявки."))
			return
		}
		userData = &models.UserState{}

		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, fmt.Sprintf("Вашe обращение принято. Номер заявки #%s", new_ticket.TicketNumber))
		buttons := []models.Button{
			{Title: "Назад", Callback: "start"},
		}
		msg.ReplyMarkup = getInlineKeyboard(buttons)
		bot.Send(msg)
	}

}

func preview_ticket(callback *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI, userData *models.UserState) {
	Logger.Debug("preview_ticket")
	var ticket models.TgTicket

	db.First(&ticket, "TicketID = ?", callback.Data)

	var formattedTime string
	if ticket.SolutionTimeDestinationDate > 0 {
		formattedTime = time.Unix(int64(ticket.SolutionTimeDestinationDate), 0).Format("2006-01-02 15:04:05")
	}

	preview := fmt.Sprintf("Номер заявки:\t#%s\nTип:\t%s\nИсполнитель:\t%s %s\nСтатус:\t%s\nПредельное время решения:\t%s\n\nОписание:\t%s",
		ticket.TicketNumber,
		ticket.Type,
		ticket.UserFirstname,
		ticket.UserLastname,
		Translate(ticket.State),
		formattedTime,
		*ticket.Body,
	)

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, preview)

	buttons := []models.Button{
		{
			Title:    "Назад",
			Callback: "start",
		},
	}

	if userData.Action == pending_callback {
		vote_buttons := []models.Button{
			{Title: "5", Callback: "vote_5"},
			{Title: "4", Callback: "vote_4"},
			{Title: "3", Callback: "vote_3"},
			{Title: "2", Callback: "vote_2"},
			{Title: "Вернуть на доработку", Callback: "vote_"},
		}
		buttons = append(vote_buttons, buttons...)
	} else if userData.Action == update_callback {
		update_buttons := []models.Button{
			{Title: "Добавить комментарий", Callback: "add_comment"},
		}
		buttons = append(update_buttons, buttons...)
	}

	msg.ReplyMarkup = getInlineKeyboard(buttons)
	bot.Send(msg)
}

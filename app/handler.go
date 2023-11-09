package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"sync"
	"tg_bot/models"

	"github.com/davecgh/go-spew/spew"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var (
	welcome     string = "Добро пожаловать в систему постановки заявок EFSOL"
	ticketRegex        = regexp.MustCompile(`^\d+$`)
	voteRegex          = regexp.MustCompile(`^vote_?(\d)?$`)
	userStates  sync.Map
	QueueID     int = 3
	TypeID      int = 3
	StateID     int = 1
	PriorityID  int = 3
	OwnerID     int = 1
	LockID      int = 1
)

func handleCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI, userData *models.UserState, txn *newrelic.Transaction) {
	Logger.Debug("handleCommand")
	Logger.Debug("update.Message.Chat.ID: ", update.Message.Chat.ID)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	switch update.Message.Command() {
	case "start":
		msg.Text = welcome + "\n\nВы можете создать, обновить или проверить статус вашей заявки." +
			"Для остановки бота просто напишите /stop, для повторного запуска нажмите /start"

		msg.ReplyMarkup = getInitialKeyboard()
		bot.Send(msg)
	case "stop":
		msg.Text = "До встречи."
		bot.Send(msg)
		// delete(userStates, update.Message.From.ID)
	case "dump":
		msg.Text = "userStates"
		bot.Send(msg)
	default:
		msg.Text = "Unknown command."
		bot.Send(msg)
	}
}

func HandleDocument(update tgbotapi.Update, bot *tgbotapi.BotAPI, userData *models.UserState, txn *newrelic.Transaction) {
	Logger.Debug("handleDocument")

	doc := update.Message.Document

	fileID := doc.FileID

	file, err := bot.GetFile(tgbotapi.FileConfig{FileID: fileID})
	if err != nil {
		Logger.Warn(err)
		return
	}

	response, err := http.Get(file.Link(bot.Token))
	if err != nil {
		Logger.Warn(err)
		return
	}
	defer response.Body.Close()

	documentBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		Logger.Warn("Error getting file: ", err)
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Произошла ошибка при получении файла."))
		return
	}

	var text = ""
	var new_ticket models.OtrsResponse
	if Mock_OTRS == "1" {
		new_ticket = models.OtrsResponse{TicketNumber: "123"}
		text = fmt.Sprintf("Вашe обращение принято. Номер заявки #%s", new_ticket.TicketNumber)
	} else {
		if userData.TicketID != "" {
			new_ticket, err = otrsRequest("update", models.OtrsRequest{
				TicketID: &userData.TicketID,
				Article: &models.Article{
					CommunicationChannel: "Internal",
					SenderType:           "customer",
					Charset:              "utf-8",
					MimeType:             "text/plain",
					From:                 userData.CustomerUserLogin,
					Subject:              "Telegram message",
					Body:                 userData.Description,
				},
				Attachment: &models.Attachment{
					Content:     string(documentBytes),
					ContentType: doc.MimeType,
					Filename:    doc.FileName,
				},
			})

			if err != nil {
				Logger.Warn(err)
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка при обновлении заявки."))
				return
			}
			text = fmt.Sprintf("Вашe обращение принято. Номер заявки #%s", new_ticket.TicketNumber)
		} else {
			new_ticket, err = otrsRequest("create", models.OtrsRequest{
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
				Attachment: &models.Attachment{
					Content:     string(documentBytes),
					ContentType: doc.MimeType,
					Filename:    doc.FileName,
				},
			})

			if err != nil {
				Logger.Warn(err)
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка при создании заявки."))
				return
			}
			text = "Вашe обращение обновлено"
		}
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	buttons := []models.Button{
		{Title: "Назад", Callback: "start"},
	}
	msg.ReplyMarkup = getInlineKeyboard(buttons)
	bot.Send(msg)

	// clean user
	userData.Topic = ""
	userData.Description = ""
}

func handleMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI, userData *models.UserState, txn *newrelic.Transaction) {
	Logger.Debug("handleMessage")
	switch userData.CurrentState {
	case "waiting_for_request_topic":
		userData.Topic = update.Message.Text
		userData.CurrentState = "waiting_for_request_description"
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите описание заявки")
		bot.Send(msg)

	case "waiting_for_request_description":
		userData.Description = update.Message.Text
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Прикрепить файл?")
		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Да", "attach_file"),
				tgbotapi.NewInlineKeyboardButtonData("Нет", "create"),
			),
		)
		msg.ReplyMarkup = inlineKeyboard
		bot.Send(msg)
	case "waiting_for_comment":
		userData.Description = update.Message.Text
		spew.Dump(userData)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Прикрепить файл?")
		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Да", "attach_file"),
				tgbotapi.NewInlineKeyboardButtonData("Нет", "create"),
			),
		)
		msg.ReplyMarkup = inlineKeyboard
		bot.Send(msg)
	}
}

func handleCallbackQuery(update tgbotapi.Update, bot *tgbotapi.BotAPI, userData *models.UserState, txn *newrelic.Transaction) {
	Logger.Debug("handleCallbackQuery: " + update.CallbackQuery.Data)
	callback := update.CallbackQuery

	switch callback.Data {
	case "start":
		start(callback, bot)
	case "new_request":
		newRequest(callback, bot, userData)
	case "check_status":
		checkStatus(callback, bot, userData)
	case "show_open":
		userData.Action = "show_open"
		listTickets(callback, bot, userData)
	case "show_pending":
		userData.Action = "show_pending"
		listTickets(callback, bot, userData)
	case "update_request":
		userData.Action = "update_request"
		listTickets(callback, bot, userData)
	case "add_comment":
		userData.Action = "add_comment"
		// listTickets(callback, bot, userData)
		// userData.Topic = update.Message.Text
		userData.CurrentState = "waiting_for_comment"
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Введите комментарий")
		bot.Send(msg)

	case "attach_file":
		attachFile(callback, bot)
	case "create":
		create(callback, bot, userData)
	default:
		if ticketRegex.MatchString(callback.Data) {
			// userData.TicketID, _ = strconv.Atoi(callback.Data)
			userData.TicketID = callback.Data
			preview_ticket(callback, bot, userData)
		} else if voteRegex.MatchString(callback.Data) {
			Logger.Debug("default:is_vote")
			match := voteRegex.FindStringSubmatch(callback.Data)
			spew.Dump(match)
			closedStateID := 2
			if match[1] == "" {
				closedStateID = 4 // open
			}

			updated_ticket, err := otrsVoteRequest("update", models.OtrsVoteRequest{
				TicketID: &userData.TicketID,
				Ticket: &models.VoteTicket{
					StateID: &closedStateID, //closed successful
				},
				DynamicField: models.DynamicField{
					Name:  "TicketVote",
					Value: match[1],
				},
			})

			spew.Dump(updated_ticket)
			spew.Dump(err)

			msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Обращение оценено")
			buttons := []models.Button{
				{Title: "Назад", Callback: "start"},
			}
			msg.ReplyMarkup = getInlineKeyboard(buttons)
			bot.Send(msg)
			userData.Action = ""
		}
	}
}

func HandleUpdate(update tgbotapi.Update, bot *tgbotapi.BotAPI, userData *models.UserState, txn *newrelic.Transaction) {
	segment := txn.StartSegment("HandleUpdate")
	defer segment.End()

	if update.Message != nil {
		if update.Message.IsCommand() {
			handleCommand(update, bot, userData, txn)
		} else if update.Message.Document != nil {
			HandleDocument(update, bot, userData, txn)
		} else {
			handleMessage(update, bot, userData, txn)
		}
	} else if update.CallbackQuery != nil {
		handleCallbackQuery(update, bot, userData, txn)
	}
}

func handleAuthoriseMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI, userData *models.UserState, txn *newrelic.Transaction) {
	segment := txn.StartSegment("HandleUpdate")
	defer segment.End()

	if userData.CurrentState != "" {
		switch userData.CurrentState {
		case "waiting_for_auth_email":
			Email := update.Message.Text
			userData.Email = Email
			Code := generateSixDigitNumber()
			userData.Code = int(Code)

			resp, err := otrsConfirm("confirm_account", models.OtrsConfirmRequest{
				Code:  Code,
				Email: Email,
			})

			var template = "Адрес %s не найден в OTRS!\nПопробуйте другой адрес"
			if err != nil {
				template = "Произошла ошибка при подтверждении %s, попробуйте позднее"
			} else if resp.Sent == 1 {
				template = "Проверочный код отправлен на %s\nДля подтверждения адреса электронной почты введите код из письма"
				userData.CurrentState = "waiting_for_auth_email_code"
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf(template, Email))
			bot.Send(msg)
		case "waiting_for_auth_email_code":
			Code := update.Message.Text
			text := "Код неверный, попробуйте еще раз"
			CodeInt, _ := strconv.Atoi(Code)

			if userData.Code == CodeInt {
				text = "Код верный"
				// resp, err
				otrsConfirm("confirm_account", models.OtrsConfirmRequest{
					Email:         userData.Email,
					TelegramLogin: update.Message.From.UserName,
				})
				// clean user_data
				userData.CurrentState = ""
				userData.Code = 0
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

			if userData.Code == CodeInt {
				msg.ReplyMarkup = getInitialKeyboard()
			}
			bot.Send(msg)

		default:
		}
	} else {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, welcome+"\n\nВведите свой адрес электронный почты, который зарегистрирован в OTRS")
		userData.CurrentState = "waiting_for_auth_email"
		bot.Send(msg)
	}
}

func Handle(wg *sync.WaitGroup, update tgbotapi.Update, bot *tgbotapi.BotAPI, txn *newrelic.Transaction) {
	defer wg.Done()

	segment := txn.StartSegment("Handle")
	defer segment.End()

	userID := getUserID(update)
	userName := getUserName(update)
	userData := getUserData(userID, userName)

	isAuthorized, CustomerUserLogin := IsAuthorized(userName)
	userData.CustomerUserLogin = CustomerUserLogin

	if isAuthorized {
		HandleUpdate(update, bot, &userData, txn)
	} else {
		handleAuthoriseMessage(update, bot, &userData, txn)
	}
	userStates.Store(userID, userData)
}

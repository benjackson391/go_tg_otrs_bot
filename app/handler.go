package app

import (
	"fmt"
	"regexp"
	"strconv"
	"sync"
	"tg_bot/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var (
	welcome     string = "Добро пожаловать в систему постановки заявок EFSOL"
	ticketRegex        = regexp.MustCompile(`^\d+$`)
	voteRegex          = regexp.MustCompile(`^vote_?(\d)?$`)
	userStates  sync.Map
)

func HandleCommand(update tgbotapi.Update, bot models.BotAPI, userData *models.UserState, txn *newrelic.Transaction) {
	Logger.Debug("HandleCommand")
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	switch update.Message.Command() {
	case "start":
		msg.Text = WelcomeMessage + "\n\n" + WelcomeDescription
		msg.ReplyMarkup = GetInitialKeyboard()
		bot.Send(msg)
	case "stop":
		msg.Text = LeaveMessage
		bot.Send(msg)
		CleanUpUserData(userData)
	default:
		msg.Text = UnknownCommand
		bot.Send(msg)
	}
}

func HandleDocument(update tgbotapi.Update, bot models.BotAPI, userData *models.UserState, txn *newrelic.Transaction) {
	Logger.Debug("HandleDocument")
	doc := update.Message.Document

	if doc.FileSize > FileSizeLimit {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, BigFileMessage))
	} else {
		documentBytes, err := GetFileContent(doc, bot)

		if err != nil {
			Logger.Warn("Error getting file: ", err)
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, ErrorMsq1))
		} else {
			_, text := CreateOrUpdateTicket(bot, userData, doc, &documentBytes)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			buttons := []models.Button{
				{Title: "Назад", Callback: "start"},
			}
			msg.ReplyMarkup = getInlineKeyboard(buttons)
			bot.Send(msg)
			CleanUpUserData(userData)
		}
	}

}

func HandleMessage(update tgbotapi.Update, bot models.BotAPI, userData *models.UserState, txn *newrelic.Transaction) {
	Logger.Debug("HandleMessage")
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

func HandleCallbackQuery(update tgbotapi.Update, bot models.BotAPI, userData *models.UserState, txn *newrelic.Transaction) {
	Logger.Debug("HandleCallbackQuery: " + update.CallbackQuery.Data)
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
		// userData.Topic = update.Message.Text
		userData.CurrentState = "waiting_for_comment"
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Введите комментарий")
		bot.Send(msg)

	case "attach_file":
		if userData.Action == "add_comment" || (userData.Topic != "" && userData.Description != "" && userData.Action != "attach_file") {
			userData.Action = "attach_file"
			attachFile(callback, bot)
		}
	case "create":
		create(callback, bot, userData)
	default:
		if ticketRegex.MatchString(callback.Data) {
			userData.TicketID = callback.Data
			preview_ticket(callback, bot, userData)
		} else if voteRegex.MatchString(callback.Data) {
			Logger.Debug("default:is_vote")
			match := voteRegex.FindStringSubmatch(callback.Data)

			userData.Vote = &match[1]

			if userData.TicketID != "" {
				_, text := CreateOrUpdateTicket(bot, userData, nil, nil)
				msg := tgbotapi.NewMessage(callback.Message.Chat.ID, text)
				buttons := []models.Button{
					{Title: "Назад", Callback: "start"},
				}
				msg.ReplyMarkup = getInlineKeyboard(buttons)
				bot.Send(msg)
				CleanUpUserData(userData)
			} else {
				msg := tgbotapi.NewMessage(callback.Message.Chat.ID, NotEnoughData)
				bot.Send(msg)
			}
		}
	}
}

func HandleUpdate(update tgbotapi.Update, bot *tgbotapi.BotAPI, userData *models.UserState, txn *newrelic.Transaction) {
	segment := txn.StartSegment("HandleUpdate")
	defer segment.End()

	if update.Message != nil {
		if update.Message.IsCommand() {
			HandleCommand(update, bot, userData, txn)
		} else if update.Message.Document != nil {
			HandleDocument(update, bot, userData, txn)
		} else {
			HandleMessage(update, bot, userData, txn)
		}
	} else if update.CallbackQuery != nil {
		HandleCallbackQuery(update, bot, userData, txn)
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
				msg.ReplyMarkup = GetInitialKeyboard()
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

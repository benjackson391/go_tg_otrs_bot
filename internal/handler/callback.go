package handler

import (
	"fmt"
	"strings"
	"tg_bot/config"
	"tg_bot/internal/common"
	"tg_bot/internal/database"
	"tg_bot/internal/logger"
	"tg_bot/internal/models"
	"tg_bot/internal/otrs"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	open_title       = "Открытые заявки"
	open_callback    = "show_open"
	pending_title    = "Заявки на оценку"
	pending_callback = "show_pending"
	update_callback  = "update_request"
)

func HandleCallbackQuery(update tgbotapi.Update, bot models.BotAPI, userData *models.UserState, logger *logger.Logger) {
	userData.Trace = append(userData.Trace, "HandleCallbackQuery")
	callback := update.CallbackQuery

	logger.Info(callback.Data)

	switch callback.Data {
	case "start":
		start(callback, bot, userData)
	case "new_request":
		userData.TicketID = ""
		userData.State = 2
		newRequest(callback, bot, userData)
	case "check_status":
		userData.State = 3
		checkStatus(callback, bot, userData)
	case "show_open":
		listTickets(callback, bot, userData)
	case "show_pending":
		listTickets(callback, bot, userData)
	case "update_request":
		userData.State = 4
		listTickets(callback, bot, userData)
	case "add_comment":
		if userData.State == 41 {
			userData.State = 42
		}

		userData.Action = "waiting_for_comment"
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Введите комментарий")
		bot.Send(msg)

	case "attach_file":
		if userData.State == 43 {
			userData.State = 431
		}
		userData.Action = "attach_file"
		attachFile(callback, bot, userData)
	case "create":
		if userData.State == 43 {
			userData.State = 432
		}
		create(callback, bot, userData, logger)
	default:
		if ticketRegex.MatchString(callback.Data) {
			userData.State = 41
			userData.TicketID = callback.Data
			preview_ticket(callback, bot, userData, logger)
		} else if voteRegex.MatchString(callback.Data) {
			logger.Debug("default:is_vote")
			match := voteRegex.FindStringSubmatch(callback.Data)

			userData.Vote = &match[1]

			if userData.TicketID != "" {
				_, text := otrs.CreateOrUpdateTicket(bot, userData, nil, nil, nil, logger)
				msg := tgbotapi.NewMessage(callback.Message.Chat.ID, text)
				buttons := []models.Button{
					{Title: "Назад", Callback: "start"},
				}
				msg.ReplyMarkup = common.GetInlineKeyboard(buttons)
				bot.Send(msg)
				common.CleanUpUserData(userData)
			} else {
				msg := tgbotapi.NewMessage(callback.Message.Chat.ID, config.NotEnoughData)
				bot.Send(msg)
			}
		}
	}
}

func addCheckStatusButton(buttons []models.Button, title string, number int, callback string, userData *models.UserState) []models.Button {
	userData.Trace = append(userData.Trace, "addCheckStatusButton")
	if number > 0 {
		buttons = append(buttons, models.Button{Title: fmt.Sprintf("%s (%d)", title, number), Callback: callback})
	}
	return buttons
}

func start(callback *tgbotapi.CallbackQuery, bot models.BotAPI, userData *models.UserState) {
	userData.Trace = append(userData.Trace, "start")
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, config.WelcomeDescription)
	msg.ReplyMarkup = common.GetInitialKeyboard()
	bot.Send(msg)
}

func newRequest(callback *tgbotapi.CallbackQuery, bot models.BotAPI, userData *models.UserState) {
	userData.Trace = append(userData.Trace, "newRequest")
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Введите тему заявки")
	bot.Send(msg)
	userData.Action = "waiting_for_title"
}

func checkStatus(callback *tgbotapi.CallbackQuery, bot models.BotAPI, userData *models.UserState) {
	userData.Trace = append(userData.Trace, "checkStatus")

	state_type_count := database.GetStateTypeCount(userData.CustomerUserLogin)

	buttons := []models.Button{}
	buttons = addCheckStatusButton(buttons, open_title, state_type_count.CountOpen, open_callback, userData)
	buttons = addCheckStatusButton(buttons, pending_title, state_type_count.CountPendAuto, pending_callback, userData)

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "")
	msg.Text = "Нет открытых заявок"
	if state_type_count.CountOpen > 0 || state_type_count.CountPendAuto > 0 {
		msg.Text = "Выберите"
		msg.ReplyMarkup = common.GetInlineKeyboard(buttons)
	}
	bot.Send(msg)
}

func listTickets(callback *tgbotapi.CallbackQuery, bot models.BotAPI, userData *models.UserState) {
	userData.Trace = append(userData.Trace, "listTickets")

	tickets := database.GetTickets(userData.UserName)

	var title string
	switch callback.Data {
	case "show_open":
		title = open_title
		tickets = common.GetOpenTickets(tickets)
	case "show_pending":
		title = pending_title
		tickets = common.GetPendingTickets(tickets)
	case "update_request":
		title = config.UpdateTitle
	default:
		title = ""
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, title)
	buttons := common.ConvertTicketsToButtons(tickets)
	buttons = append(buttons, models.Button{Title: "Назад", Callback: "start"})
	msg.ReplyMarkup = common.GetInlineKeyboard(buttons)
	bot.Send(msg)
}

func attachFile(callback *tgbotapi.CallbackQuery, bot models.BotAPI, userData *models.UserState) {
	userData.Trace = append(userData.Trace, "attachFile")
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Приложите файл до 20Mb")
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Создать без вложения", "create"),
		),
	)
	msg.ReplyMarkup = inlineKeyboard
	bot.Send(msg)
}

func create(callback *tgbotapi.CallbackQuery, bot models.BotAPI, userData *models.UserState, logger *logger.Logger) {
	userData.Trace = append(userData.Trace, "create")
	if userData.State == 431 {
		userData.State = 45
	}

	if userData.State == 43 || userData.State == 431 || userData.State == 432 || userData.State == 22 {
		if userData.Description == "" {
			return
		}

		_, text := otrs.CreateOrUpdateTicket(bot, userData, nil, nil, nil, logger)
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, text)
		buttons := []models.Button{
			{Title: "Назад", Callback: "start"},
		}
		msg.ReplyMarkup = common.GetInlineKeyboard(buttons)
		bot.Send(msg)

		common.CleanUpUserData(userData)
	}
}

func preview_ticket(callback *tgbotapi.CallbackQuery, bot models.BotAPI, userData *models.UserState, logger *logger.Logger) {
	userData.Trace = append(userData.Trace, "preview_ticket")

	ticket := database.GetTicket(callback.Data)
	articles := database.GetTicketArticles(callback.Data)

	var formattedTime string
	if ticket.SolutionTimeDestinationDate > 0 {
		formattedTime = time.Unix(int64(ticket.SolutionTimeDestinationDate), 0).Format("2006-01-02 15:04:05")
	}

	preview := fmt.Sprintf("Номер заявки:\t#%s\nTип:\t%s\nИсполнитель:\t%s %s\nСтатус:\t%s\nПредельное время решения:\t%s\n\nОписание:\t%s",
		ticket.TicketNumber,
		ticket.Type,
		ticket.UserFirstname,
		ticket.UserLastname,
		common.Translate(ticket.State),
		formattedTime,
		*ticket.Body,
	)

	if len(articles) > 0 {
		preview += "\nСообщения по заявке:"
		for _, article := range articles {
			dateParsed, _ := time.Parse("2006-01-02 15:04:05", article.CreateTime)
			dateFormatted := dateParsed.Format("02.01.2006 15:04")

			articleText := fmt.Sprintf(
				"\n\n📝 Тема: %s\n"+
					"👤 От: %s\n"+
					"📅 Дата: %s\n"+
					"\n"+
					"%s",
				article.Subject,
				article.From,
				dateFormatted,
				article.Body,
			)
			preview += articleText
		}
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, preview)
	// msg.ParseMode = "MarkdownV2"

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
	} else if userData.Action == update_callback || userData.State == 41 {
		update_buttons := []models.Button{
			{Title: "Добавить комментарий", Callback: "add_comment"},
		}
		buttons = append(update_buttons, buttons...)
	}

	msg.ReplyMarkup = common.GetInlineKeyboard(buttons)
	bot.Send(msg)

}

func escapeMarkdown(text string) string {
	var replacer = strings.NewReplacer(
		"_", "\\_",
		"*", "\\*",
		"[", "\\[",
		"]", "\\]",
		"(", "\\(",
		")", "\\)",
		"~", "\\~",
		"`", "\\`",
		">", "\\>",
		"#", "\\#",
		"+", "\\+",
		"-", "\\-",
		"=", "\\=",
		"|", "\\|",
		"{", "\\{",
		"}", "\\}",
		".", "\\.",
		"!", "\\!",
	)
	return replacer.Replace(text)
}

package app

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"tg_bot/models"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var TRANSLATION = map[string]string{
	"new":                 "новая",
	"open":                "открыта",
	"pending reminder":    "ожидает напоминания",
	"pending auto":        "ожидает автозакрытия",
	"pending auto close+": "ожидает автозакрытия(+)",
	"pending auto close-": "ожидает автозакрытия(-)",
}

func getUserName(update tgbotapi.Update) string {
	if update.Message != nil {
		return update.Message.From.UserName
	} else if update.CallbackQuery != nil {
		return update.CallbackQuery.From.UserName
	}
	return ""
}

func getUserID(update tgbotapi.Update) int {
	if update.Message != nil {
		return update.Message.From.ID
	} else if update.CallbackQuery != nil {
		return update.CallbackQuery.From.ID
	}
	return 0
}

func getUserData(userID int, userName string) models.UserState {
	userData, ok := userStates.Load(userID)
	if ok {
		return userData.(models.UserState)
	}
	return models.UserState{UserName: userName}
}

func InitAPI(url string, user string, pass string) {
	URL = url
	USER = user
	PASS = pass
}

func Translate(status string) string {
	translated_status, ok := TRANSLATION[status]
	if !ok {
		return status
	}
	return translated_status
}

func generateSixDigitNumber() int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(900000) + 100000
}

func getInlineKeyboard(buttons []models.Button) tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton

	for _, button := range buttons {
		row := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(button.Title, button.Callback),
		)
		rows = append(rows, row)
	}

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(rows...)
	return inlineKeyboard
}

func convertTicketsToButtons(tickets []models.TgTicket) []models.Button {
	var buttons []models.Button // создаем массив пар строк

	for _, ticket := range tickets {
		Title := fmt.Sprintf("#%s: %s", ticket.TicketNumber, *ticket.Title)
		Callback := fmt.Sprint(ticket.TicketID)
		buttons = append(buttons, models.Button{Title: Title, Callback: Callback})
	}

	return buttons
}

func GetInitialKeyboard() *tgbotapi.InlineKeyboardMarkup {
	buttons := []models.Button{
		{
			Title:    "Новая заявка",
			Callback: "new_request",
		},
		{
			Title:    "Проверить статус",
			Callback: "check_status",
		},
		{
			Title:    "Обновить заявку",
			Callback: "update_request",
		},
	}

	keyboard := getInlineKeyboard(buttons)
	return &keyboard
}

func getOpenTickets(tickets []models.TgTicket) []models.TgTicket {
	var openTickets []models.TgTicket
	for _, ticket := range tickets {
		if ticket.StateType != "pending auto" {
			openTickets = append(openTickets, ticket)
		}
	}
	return openTickets
}

func getPendingTickets(tickets []models.TgTicket) []models.TgTicket {
	var pendingTickets []models.TgTicket
	for _, ticket := range tickets {
		if ticket.StateType == "pending auto" {
			pendingTickets = append(pendingTickets, ticket)
		}
	}
	return pendingTickets
}
func GetFileContent(doc *tgbotapi.Document, bot models.BotAPI) ([]byte, error) {
	// file, err := bot.GetFile(tgbotapi.FileConfig{FileID: doc.FileID})
	if err != nil {
		Logger.Warn(err)
		return nil, err
	}

	fileLink, err := bot.GetFileDirectURL(doc.FileID)
	if err != nil {
		Logger.Warn(err)
		return nil, err
	}

	response, err := http.Get(fileLink)
	if err != nil {
		Logger.Warn(err)
		return nil, err
	}
	defer response.Body.Close()

	return io.ReadAll(response.Body)
}

func CleanUpUserData(userData *models.UserState) {
	*userData = models.UserState{
		UserName:          userData.UserName,
		CustomerUserLogin: userData.CustomerUserLogin,
	}
}

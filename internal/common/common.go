package common

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"sync"
	"tg_bot/internal/logger"
	"tg_bot/internal/models"
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

func GetUserName(update tgbotapi.Update) string {
	if update.Message != nil {
		return update.Message.From.UserName
	} else if update.CallbackQuery != nil {
		return update.CallbackQuery.From.UserName
	}
	return ""
}

func GetUserID(update tgbotapi.Update) int {
	if update.Message != nil {
		return update.Message.From.ID
	} else if update.CallbackQuery != nil {
		return update.CallbackQuery.From.ID
	}
	return 0
}

func GetUserData(userID int, userName string, userStates *sync.Map) models.UserState {
	userData, ok := userStates.Load(userID)
	if ok {
		return userData.(models.UserState)
	}
	return models.UserState{UserName: userName}
}

func Translate(status string) string {
	translated_status, ok := TRANSLATION[status]
	if !ok {
		return status
	}
	return translated_status
}

func GenerateSixDigitNumber() int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(900000) + 100000
}

func GetInlineKeyboard(buttons []models.Button) tgbotapi.InlineKeyboardMarkup {
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

func ConvertTicketsToButtons(tickets []models.TgTicket) []models.Button {
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

	keyboard := GetInlineKeyboard(buttons)
	return &keyboard
}

func GetOpenTickets(tickets []models.TgTicket) []models.TgTicket {
	var openTickets []models.TgTicket
	for _, ticket := range tickets {
		if ticket.StateType != "pending auto" {
			openTickets = append(openTickets, ticket)
		}
	}
	return openTickets
}

func GetPendingTickets(tickets []models.TgTicket) []models.TgTicket {
	var pendingTickets []models.TgTicket
	for _, ticket := range tickets {
		if ticket.StateType == "pending auto" {
			pendingTickets = append(pendingTickets, ticket)
		}
	}
	return pendingTickets
}
func GetFileContent(fileID string, bot models.BotAPI) ([]byte, error) {
	fileLink, err := bot.GetFileDirectURL(fileID)

	if err != nil {
		logger.Warning(err.Error())
		return nil, err
	}

	response, err := http.Get(fileLink)
	if err != nil {
		logger.Warning(err.Error())
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("bad return code: %d", response.StatusCode)
	}

	return io.ReadAll(response.Body)
}

func CleanUpUserData(userData *models.UserState) {
	*userData = models.UserState{
		UserName:          userData.UserName,
		CustomerUserLogin: userData.CustomerUserLogin,
		Action:            userData.Action,
	}
}

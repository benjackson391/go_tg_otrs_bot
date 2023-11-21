package handler

import (
	"fmt"
	"testing"
	"tg_bot/config"
	"tg_bot/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleMessageAttachment(t *testing.T) {
	bot := new(MockBotAPI)

	csenario := []struct {
		UserData     *models.UserState
		FileSize     int
		OtrsPath     string
		OtrsResponse string
		AnswerText   string
		Code         int
	}{
		{
			UserData:     &models.UserState{TicketID: ""},
			FileSize:     config.FileSizeLimit + config.FileSizeLimit*0.5,
			OtrsPath:     "/create",
			OtrsResponse: "",
			Code:         200,
			AnswerText:   "Файл слишком большой. Приложите другой файл",
		},
		{
			UserData:     &models.UserState{TicketID: ""},
			FileSize:     config.FileSizeLimit,
			OtrsPath:     "/create",
			OtrsResponse: `{"ArticleID":"1","TicketID":"1","TicketNumber":"1"}`,
			Code:         200,
			AnswerText:   fmt.Sprintf(config.TicketCreatedTemplate, "1"),
		},
		{
			UserData:     &models.UserState{TicketID: ""},
			FileSize:     config.FileSizeLimit * 0.5,
			OtrsPath:     "/create",
			OtrsResponse: `{"ArticleID":"1","TicketID":"2","TicketNumber":"2"}`,
			Code:         200,
			AnswerText:   fmt.Sprintf(config.TicketCreatedTemplate, "2"),
		},
		{
			UserData:     &models.UserState{TicketID: "1"},
			FileSize:     config.FileSizeLimit * 0.5,
			OtrsPath:     "/update",
			OtrsResponse: `{"ArticleID":"1","TicketID":"2","TicketNumber":"3"}`,
			Code:         404,
			AnswerText:   config.ErrorMsq1,
		},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for test_number, test_case := range csenario {
		httpmock.RegisterResponder("POST", test_case.OtrsPath,
			httpmock.NewStringResponder(test_case.Code, test_case.OtrsResponse))

		httpmock.RegisterResponder("GET", "https://api.telegram.org/file/bot/",
			httpmock.NewStringResponder(test_case.Code, "context"))

		document := tgbotapi.Document{
			FileID:   "file",
			FileName: "example.pdf",
			MimeType: "application/pdf",
			FileSize: test_case.FileSize,
		}

		message := &tgbotapi.Message{
			Document: &document,
			Chat:     &tgbotapi.Chat{ID: 12345},
		}
		update := tgbotapi.Update{
			Message: message,
		}

		bot.On("GetFile", mock.Anything).Return(tgbotapi.File{}, nil)
		bot.On("Send", mock.Anything).Return(tgbotapi.Message{}, nil)
		bot.On("GetFileDirectURL", mock.Anything).Return("", nil)

		HandleMessage(update, bot, test_case.UserData)

		if len(bot.SentMessages) != 1 {
			t.Errorf("[%d] Expected one message to be sent for start command, got: %v", test_number, len(bot.SentMessages))
		} else {
			answer, ok := bot.SentMessages[0].(tgbotapi.MessageConfig)
			if !ok {
				t.Errorf("[%d] Expected message to be of type tgbotapi.MessageConfig", test_number)
			} else {
				assert.Equal(t, test_case.AnswerText, answer.Text, fmt.Sprintf("[%d] Text should be equal", test_number))
			}
		}
		bot.SentMessages = []tgbotapi.Chattable{}
	}
}

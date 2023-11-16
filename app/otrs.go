package app

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"tg_bot/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	URL        string
	USER       string
	PASS       string
	Mock_OTRS  string
	stateId    int
	QueueID    int = 3
	TypeID     int = 3
	PriorityID int = 3
	OwnerID    int = 1
	LockID     int = 1
)

var States = map[string]int{
	"new":               1,
	"open":              4,
	"closed_successful": 2,
}

func CreateOrUpdateTicket(bot models.BotAPI, userData *models.UserState, doc *tgbotapi.Document, documentBytes *[]byte) (models.OtrsResponse, string) {
	var ticket models.OtrsResponse
	var path = "update"
	var otrs_request models.OtrsRequest

	template := TicketUpdatedMessage

	if Mock_OTRS == "1" {
		ticket = models.OtrsResponse{TicketNumber: "123"}
		template = TicketCreatedTemplate
	} else {
		// otrs_request := models.OtrsRequest{
		// 	TicketID: &userData.TicketID,
		// 	Article: &models.Article{
		// 		CommunicationChannel: "Internal",
		// 		SenderType:           "customer",
		// 		Charset:              "utf-8",
		// 		MimeType:             "text/plain",
		// 		From:                 userData.CustomerUserLogin,
		// 		Subject:              "Telegram message",
		// 		Body:                 userData.Description,
		// 	},
		// }

		if userData.TicketID == "" { // create
			path = "create"
			stateId = States["new"]
			otrs_request = models.OtrsRequest{
				Ticket: &models.Ticket{
					Title:        &userData.Topic,
					QueueID:      &QueueID,
					TypeID:       &TypeID, // Request for service
					CustomerUser: &userData.CustomerUserLogin,
					StateID:      &stateId,    // new
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
			}
			template = TicketCreatedTemplate
		} else { //update
			template = TicketUpdatedMessage
			var state int
			otrs_request = models.OtrsRequest{
				TicketID: &userData.TicketID,
			}

			if userData.Description != "" {
				otrs_request.Article = &models.Article{
					CommunicationChannel: "Internal",
					SenderType:           "customer",
					Charset:              "utf-8",
					MimeType:             "text/plain",
					From:                 userData.CustomerUserLogin,
					Subject:              "Telegram message",
					Body:                 userData.Description,
				}
			}

			if userData.Vote != nil {
				if *userData.Vote != "" {
					state = States["closed_successful"]
				} else {
					state = States["open"]
				}
				otrs_request.Ticket = &models.Ticket{
					StateID: &state,
				}

				otrs_request.DynamicField = &models.DynamicField{
					Name:  "TicketVote",
					Value: userData.Vote,
				}
				template = VoteTemplate
			}
		}

		if documentBytes != nil && len(*documentBytes) > 0 {
			otrs_request.Attachment = &models.Attachment{
				Content:     string(*documentBytes),
				ContentType: doc.MimeType,
				Filename:    doc.FileName,
			}
		}

		ticket, err = otrsRequest(path, otrs_request)
		if err != nil {
			Logger.Warn(err)
			return models.OtrsResponse{}, "Ошибка при создании заявки: " + err.Error()
		}
	}

	return ticket, fmt.Sprintf(template, ticket.TicketNumber)
}

func performHttpRequest(path string, encodedJson []byte) ([]byte, error) {
	request, err := http.NewRequest("POST", URL+"/"+path, bytes.NewBuffer(encodedJson))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	if transport, ok := http.DefaultTransport.(*http.Transport); ok {
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

func otrsRequest(path string, jsonData models.OtrsRequest) (models.OtrsResponse, error) {
	var responseJson models.OtrsResponse

	jsonData.UserLogin = USER
	jsonData.Password = PASS

	encodedJson, err := json.Marshal(jsonData)
	if err != nil {
		Logger.Warn("json.Marshal:" + err.Error())
		return responseJson, err
	}

	response, err := performHttpRequest(path, encodedJson)
	if err != nil {
		Logger.Warn("performHttpRequest:" + err.Error())
		Logger.Warn("response: ", response)
		return responseJson, err
	}

	err = json.Unmarshal(response, &responseJson)
	if err != nil {
		Logger.Warn("json.Unmarshal:" + err.Error())
		return responseJson, err
	}

	return responseJson, nil
}

func otrsConfirm(path string, jsonData models.OtrsConfirmRequest) (models.OtrsConfirmResponse, error) {
	var responseJson models.OtrsConfirmResponse

	jsonData.UserLogin = USER
	jsonData.Password = PASS

	encodedJson, err := json.Marshal(jsonData)
	if err != nil {
		Logger.Warn("json.Marshal:" + err.Error())
		return responseJson, err
	}

	response, err := performHttpRequest(path, encodedJson)
	if err != nil {
		Logger.Warn("performHttpRequest:" + err.Error())
		Logger.Warn("response: ", response)
		return responseJson, err
	}

	err = json.Unmarshal(response, &responseJson)
	if err != nil {
		Logger.Warn("json.Unmarshal:" + err.Error())
		return responseJson, err
	}

	return responseJson, nil
}

func SetupOTRSConnector() {
	USER = os.Getenv("OTRS_USER")
	PASS = os.Getenv("OTRS_PASSWORD")
	URL = os.Getenv("OTRS_URL")
	Mock_OTRS = os.Getenv("MOCK_OTRS")
}

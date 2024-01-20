package otrs

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"tg_bot/config"
	"tg_bot/internal/logger"
	"tg_bot/internal/models"
)

var (
	URL         string
	USER        string
	PASS        string
	Mock_OTRS   string
	stateId     int
	QueueIDprod int = 2
	QueueIDdev  int = 3
	TypeID      int = 3
	PriorityID  int = 3
	OwnerID     int = 1
	LockID      int = 1
	err         error
)

var States = map[string]int{
	"new":               1,
	"open":              4,
	"closed_successful": 2,
}

func GetQueueID() *int {
	mode := os.Getenv("MODE")
	if mode == "DEV" {
		return &QueueIDdev
	}
	return &QueueIDprod
}

func CreateOrUpdateTicket(bot models.BotAPI, userData *models.UserState, mimetype *string, filename *string, documentBytes *string, logger *logger.Logger) (models.OtrsResponse, string) {
	userData.Trace = append(userData.Trace, "CreateOrUpdateTicket")

	var ticket models.OtrsResponse
	var path = "update"
	var otrs_request models.OtrsRequest

	template := config.TicketUpdatedMessage

	if Mock_OTRS == "1" {
		ticket = models.OtrsResponse{TicketNumber: "123"}
		template = config.TicketCreatedTemplate
	} else {
		if userData.TicketID == "" { // create
			path = "create"
			stateId = States["new"]
			otrs_request = models.OtrsRequest{
				Ticket: &models.Ticket{
					Title:        &userData.Topic,
					QueueID:      GetQueueID(),
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
			template = config.TicketCreatedTemplate
		} else { //update
			template = config.TicketUpdatedMessage
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
				template = config.VoteTemplate
			}
		}

		if documentBytes != nil && len(*documentBytes) > 0 && mimetype != nil && filename != nil {
			otrs_request.Attachment = &models.Attachment{
				Content:     *documentBytes,
				ContentType: *mimetype,
				Filename:    *filename,
			}
		}

		ticket, err = otrsRequest(path, otrs_request, logger)
		if err != nil {
			logger.Warn(err.Error())
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

func otrsRequest(path string, jsonData models.OtrsRequest, logger *logger.Logger) (models.OtrsResponse, error) {
	var responseJson models.OtrsResponse

	jsonData.UserLogin = USER
	jsonData.Password = PASS

	encodedJson, err := json.Marshal(jsonData)
	if err != nil {
		logger.Warn("json.Marshal:" + err.Error())
		return responseJson, err
	}

	response, err := performHttpRequest(path, encodedJson)
	if err != nil {
		// logger.Warn("performHttpRequest:" + err.Error())
		// logger.Warn("response: ", response)
		return responseJson, err
	}

	err = json.Unmarshal(response, &responseJson)
	if err != nil {
		logger.Warn("json.Unmarshal:" + err.Error())
		return responseJson, err
	}

	return responseJson, nil
}

func OtrsConfirm(path string, jsonData models.OtrsConfirmRequest, logger *logger.Logger) (models.OtrsConfirmResponse, error) {
	var responseJson models.OtrsConfirmResponse

	jsonData.UserLogin = USER
	jsonData.Password = PASS

	encodedJson, err := json.Marshal(jsonData)
	if err != nil {
		logger.Warn("json.Marshal:" + err.Error())
		return responseJson, err
	}

	response, err := performHttpRequest(path, encodedJson)
	if err != nil {
		// logger.Warn("performHttpRequest:" + err.Error())
		// logger.Warn("response: ", response)
		return responseJson, err
	}

	err = json.Unmarshal(response, &responseJson)
	if err != nil {
		logger.Warn("json.Unmarshal:" + err.Error())
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

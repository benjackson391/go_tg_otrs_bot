package app

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"tg_bot/models"
)

var (
	URL       string
	USER      string
	PASS      string
	Mock_OTRS string
)

func performHttpRequest(path string, jsonData interface{}, responseJson interface{}) error {
	jsonDataMap, err := StructToMap(jsonData)
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	jsonDataMap["UserLogin"] = USER
	jsonDataMap["Password"] = PASS

	jsonValue, err := json.Marshal(jsonDataMap)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", URL+"/"+path, bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &responseJson)
	if err != nil {
		return err
	}

	return nil
}

func otrsRequest(path string, jsonData models.OtrsRequest) (models.OtrsResponse, error) {
	var responseJson models.OtrsResponse

	if Mock_OTRS == "1" {
		return models.OtrsResponse{TicketNumber: "123"}, nil
	}

	err := performHttpRequest(path, jsonData, &responseJson)
	return responseJson, err
}

func otrsVoteRequest(path string, jsonData models.OtrsVoteRequest) (models.OtrsResponse, error) {
	var responseJson models.OtrsResponse

	if Mock_OTRS == "1" {
		return models.OtrsResponse{TicketNumber: "123"}, nil
	}

	err := performHttpRequest(path, jsonData, &responseJson)
	return responseJson, err
}

func otrsConfirm(path string, jsonData models.OtrsConfirmRequest) (models.OtrsConfirmResponse, error) {
	var responseJson models.OtrsConfirmResponse

	// if Mock_OTRS == "1" {
	// 	return models.OtrsConfirmResponse{Sent: 1}, nil
	// }

	err := performHttpRequest(path, jsonData, &responseJson)
	return responseJson, err
}

func SetupOTRSConnector() {
	USER = os.Getenv("OTRS_USER")
	PASS = os.Getenv("OTRS_PASSWORD")
	URL = os.Getenv("OTRS_URL")
	Mock_OTRS = os.Getenv("MOCK_OTRS")
}

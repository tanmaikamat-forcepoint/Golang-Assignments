package helper

import (
	"encoding/json"
	"errors"
	"net/http"
)

type RequestParseOutput struct {
	Body interface{}
}

type ResponseBody struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data"`
}

func ParseRequestBody(r *http.Request, body any) (RequestParseOutput, error) {
	tempRequestParseOutput := RequestParseOutput{}
	tempRequestParseOutput.Body = body
	err := json.NewDecoder(r.Body).Decode(tempRequestParseOutput.Body)
	if err != nil {
		return RequestParseOutput{}, errors.New("Please Enter a Valid Json Body")
	}
	return tempRequestParseOutput, nil
}

func PackRequestBody(w http.ResponseWriter, statusCode int, message string, body interface{}) error {

	err := json.NewEncoder(w).Encode(
		ResponseBody{
			StatusCode: statusCode,
			Message:    message,
			Data:       body})
	if err != nil {
		return errors.New("Internal Server Error")
	}

	return nil
}

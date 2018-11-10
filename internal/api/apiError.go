package api

import (
	"encoding/json"
	"log"
)

type apiError struct {
	Description string `json:"description"`
	Message     string `json:"message"`
	ErrorCode   string `json:"error"`
}

func (e *apiError) Error() string {
	return e.ErrorCode
}

// apiErrorHandle creates apiError and returns JSON serialized object
func apiErrorHandle(message, errorCode string, err error) []byte {
	dbErr := apiError{
		Description: err.Error(),
		Message:     message,
		ErrorCode:   errorCode,
	}
	dbErrData, _ := json.Marshal(dbErr) // should check error
	log.Println(dbErr)
	return dbErrData
}
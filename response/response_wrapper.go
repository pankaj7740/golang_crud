package response

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type Success struct {
	Status string `json:"status"`
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

// function to send response
func Response (w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func SuccessResponse (w http.ResponseWriter, statusCode int, status string, message string, data interface{}) {
	response := Success{
        Status:  status,
        Code:    statusCode,
        Message: message,
        Data:    data,
    }
	Response(w, statusCode, response)
}

func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
    Response(w, statusCode, Error{Code: statusCode, Message: message})
}


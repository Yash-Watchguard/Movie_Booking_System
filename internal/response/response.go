package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	Data      any    `json:"data"`
}


func SuccessResponse(W http.ResponseWriter,data any,message string,code int){
	resonse:=Response{
		Status: "Success",
		Message: message,
		Data: data,
	}

	W.Header().Set("content-Type","application/json")
	W.WriteHeader(code)
	json.NewEncoder(W).Encode(resonse)
}

func ErrorResponse(w http.ResponseWriter,message string,code int){
	response:=Response{
		Message: message,
		Status: "Fail",
	}

	w.Header().Set("content-Type","application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

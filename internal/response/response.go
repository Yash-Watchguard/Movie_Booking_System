package response

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	Errorcode int    `json:"error_code"`
	Data      any    `json:"data"`
}
type Successresponse struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	Data      any    `json:"data"`
}

func SuccessResponse(W http.ResponseWriter,data any,message string,code int){
	// encode the data

	resonse:=Successresponse{
		Status: "Success",
		Message: message,
		Data: data,
	}

	W.Header().Set("content-Type","application/json")
	W.WriteHeader(code)
	json.NewEncoder(W).Encode(resonse)
}

func ErrorResponse(w http.ResponseWriter,message string,code int,errorcode int){
	response:=response{
		Message: message,
		Status: "Fail",
		Errorcode: errorcode,
	}

	w.Header().Set("content-Type","application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status        int         `json:"status"`
	Message       string      `json:"message"`
	Data          interface{} `json:"data,omitempty"`
	contentType   string
	responseWrite http.ResponseWriter
}

func CreateDefaultResponse(rw http.ResponseWriter) Response {
	return Response{
		Status:        http.StatusOK,
		responseWrite: rw,
		contentType:   "application/json",
	}
}

func (res *Response) Send() {
	res.responseWrite.Header().Set("Content-Type", res.contentType)
	res.responseWrite.WriteHeader(res.Status)

	output, _ := json.Marshal(&res)
	fmt.Fprintln(res.responseWrite, string(output))
}

func SendData(rw http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(rw)
	response.Data = data
	response.Send()
}

func (res *Response) NotFound() {
	res.Status = http.StatusNotFound
	res.Message = "Not Found"
}

func SendNotFound(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.NotFound()
	response.Send()
}

func (res *Response) UnprocessableEntity() {
	res.Status = http.StatusUnprocessableEntity
	res.Message = "Unprocessable Entity"
}

func SendUnprocessableEntity(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.UnprocessableEntity()
	response.Send()
}

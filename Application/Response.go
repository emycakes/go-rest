package Application

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Writer http.ResponseWriter
}

func (response Response) SendAsJson(statusCode int, content interface{}) {
	response.SetHeader("Content-Type", "application/json; charset=utf-8")

	if content != nil {
		body, _ := json.MarshalIndent(content, "", "  ")
		response.Send(statusCode, string(body))
		return
	}

	response.Send(statusCode, "")
}

func (response Response) Send(statusCode int, content string) {
	response.Writer.WriteHeader(statusCode)

	_, err := fmt.Fprint(response.Writer, content)

	if err !=  nil {
		log.Printf("Fprint: %v\n", err)
	}
}

func (response Response) SetHeader(key string, value string) {
	response.Writer.Header().Set(key, value)
}
package Application

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type Request struct {
	Pointer *http.Request
}

func (request Request) Get(key string) string {
	return mux.Vars(request.Pointer)[key]
}

func (request Request) GetBody() []byte {
	requestBody, _ := ioutil.ReadAll(request.Pointer.Body)
	return requestBody
}
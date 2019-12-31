package main

import (
	"github.com/gpmgo/gopm/modules/log"
	"golang_learn/errorhandling/filelistingserver/handler"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errorWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			log.Error("error", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", errorWrapper(handler.FileReadHandler()))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

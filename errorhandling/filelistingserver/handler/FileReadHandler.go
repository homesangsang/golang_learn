package handler

import (
	"io/ioutil"
	"net/http"
	"os"
)

func FileReadHandler() func(writer http.ResponseWriter, request *http.Request) error {
	return func(writer http.ResponseWriter, request *http.Request) error {
		path := request.URL.Path[len("/list/"):]
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		all, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}

		writer.Write(all)
		return nil
	}
}

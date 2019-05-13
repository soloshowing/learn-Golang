package main

import (
	"errorhandling/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			log.Printf("Error handling request: %s", err.Error())

			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}
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

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8889", nil)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	port := kingpin.Flag("port", "Port Number to listen").Default("9999").Short('p').String()
	kingpin.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	http.Handle("/", r)

	logrus.Info(*port)
	http.ListenAndServe("0.0.0.0:"+*port, r)
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello World!")
	logrus.Info(request.RemoteAddr)
}

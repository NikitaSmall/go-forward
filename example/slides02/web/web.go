package web

import (
	"fmt"
	"net/http" // this package has everything required to run a web server
)

// a handler is any function (or a struct) that has corresponding interface
// (response writer and a request object)
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// register handlers for the endpoints
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// initialize the server of the port
	http.ListenAndServe(":8090", nil)
}

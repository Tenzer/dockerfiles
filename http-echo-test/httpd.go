package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func requestHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Method:", request.Method)
	fmt.Fprintln(response, "Host:", request.Host)
	fmt.Fprintln(response, "URL:", request.URL)
	fmt.Fprintln(response, "Protocol:", request.Proto)
	fmt.Fprintln(response, "Client:", request.RemoteAddr)

	fmt.Fprintln(response, "\nHeaders:\n--------")
	for key, value := range request.Header {
		fmt.Fprintf(response, "%v: %v\n", key, strings.Join(value, ", "))
	}

	log.Printf("%v %v (%v)\n", request.Method, request.URL, request.RemoteAddr)
}

func main() {
	// Default listen port
	listenPort := "8080"

	var err error

	if os.Getenv("PORT") != "" {
		listenPort = os.Getenv("PORT")
		if err != nil {
			log.Fatalln(err)
		}
	}

	if len(os.Args) > 1 {
		listenPort = os.Args[1]
		if err != nil {
			log.Fatalln(err)
		}
	}

	log.Printf("Listening on http://0.0.0.0:%v\n", listenPort)
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":"+listenPort, nil)
}

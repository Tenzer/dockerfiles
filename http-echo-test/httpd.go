package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

func requestHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Method:", request.Method)
	fmt.Fprintln(response, "Host:", request.Host)
	fmt.Fprintln(response, "URL:", request.URL)
	fmt.Fprintln(response, "Protocol:", request.Proto)
	fmt.Fprintln(response, "Client:", request.RemoteAddr)

	fmt.Fprintln(response, "\nHeaders:\n--------")
	var headers []string
	for name := range request.Header {
		headers = append(headers, name)
	}
	sort.Strings(headers)
	for _, name := range headers {
		fmt.Fprintf(response, "%v: %v\n", name, strings.Join(request.Header[name], ", "))
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

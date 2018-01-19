package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

var hostname string

func requestHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Server: %v\n\n", hostname)

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

	_ = request.ParseMultipartForm(100000)
	if len(request.PostForm) > 0 {
		fmt.Fprintln(response, "\nForm data:\n----------")
		for key, value := range request.PostForm {
			fmt.Fprintf(response, "%v: %v\n", key, value)
		}
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Println("Error while reading request body:", err)
	}
	if len(body) > 0 {
		fmt.Fprintln(response, "\nRequest body:\n-------------")
		fmt.Fprintln(response, string(body))
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

	hostname, _ = os.Hostname()

	log.Printf("Listening on http://0.0.0.0:%v\n", listenPort)
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":"+listenPort, nil)
}

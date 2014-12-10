package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"sort"
	"strings"
)

type handlerFunc func(http.ResponseWriter, *http.Request)

func flattenHeaders(originalHeader http.Header) string {
	headers := make([]string, len(originalHeader))

	var sortedHeaders []string
	for header := range originalHeader {
		sortedHeaders = append(sortedHeaders, header)
	}
	sort.Strings(sortedHeaders)

	for _, header := range sortedHeaders {
		for _, value := range originalHeader[header] {
			headers = append(headers, fmt.Sprintf("%s=%q", header, value))
		}
	}
	return strings.Join(headers, ", ")
}

func httpLog(loggedHandler handlerFunc) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %q %q %s\n", r.Method, html.EscapeString(r.URL.Path), r.Proto, flattenHeaders(r.Header))
		loggedHandler(w, r)
	}
}

func main() {
	var port = flag.Int("port", 8080, "port")
	flag.Parse()
	serving := fmt.Sprintf("0.0.0.0:%d", *port)
	log.Println("serving on", serving)

	fs := http.FileServer(http.Dir("bigfiles"))
	http.HandleFunc("/", httpLog(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
	log.Fatal(http.ListenAndServe(serving, nil))
}

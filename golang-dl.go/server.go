// type http.HandlerFunc func(http.ResponseWriter, *http.Request)

func httpLog(loggedHandler http.HandlerFunc) http.HandlerFunc {
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

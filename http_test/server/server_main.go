package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Req: http://localhost:1234/upper?word=abc
// Res: ABC
func upperCaseHandler(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid request")
		return
	}
	word := query.Get("word")
	log.Printf("word: %v", word)
	if len(word) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "missing word")
		return
	}
	w.WriteHeader(http.StatusOK)
	result := strings.ToUpper(word)
	fmt.Fprintf(w, result)
	log.Printf("result: %v", result)

}

func main() {
	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	http.HandleFunc("/upper", upperCaseHandler)
	log.Printf("server started on :1234")
	log.Fatal(http.ListenAndServe(":1234", nil))
}

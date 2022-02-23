package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// middleware/hello_with_time_elapse.go
var logger = log.New(os.Stdout, "", 0)

func helloHandler(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("your friends is tom and alex 1"))
}

func showInfoHandler(wr http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()
	wr.Write([]byte("your friends is tom and alex 2"))
	timeElapsed := time.Since(timeStart)
	logger.Println(timeElapsed)
}

func showEmailHandler(wr http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()
	wr.Write([]byte("your friends is tom and alex 3"))
	timeElapsed := time.Since(timeStart)
	logger.Println(timeElapsed)
}
func showFriendsHandler(wr http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()
	wr.Write([]byte("your friends is tom and alex 4"))
	timeElapsed := time.Since(timeStart)
	logger.Println(timeElapsed)
}

func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		// next handler
		next.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)
		logger.Println(timeElapsed)
	})
}

func main() {
	http.Handle("/", timeMiddleware(http.HandlerFunc(helloHandler)))
	http.HandleFunc("/info/show", showInfoHandler)
	http.HandleFunc("/email/show", showEmailHandler)
	http.HandleFunc("/friends/show", showFriendsHandler)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)

}

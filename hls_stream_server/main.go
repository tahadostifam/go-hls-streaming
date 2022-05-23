package main

import (
	"fmt"
	"log"
	"net/http"
)

const songsDir = "./songs"
const httpPort = 8080

func main() {
	http.Handle("/",
		allowCors(
			http.FileServer(http.Dir(songsDir)),
		),
	)

	fmt.Printf("Server has listening on :%d\n", httpPort)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), nil))
}

func allowCors(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}

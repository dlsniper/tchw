package main

import (
	"net/http"

	"github.com/dlsniper/tchw/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8085", handler)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "White Thighhighs")
}

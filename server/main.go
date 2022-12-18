package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8085", nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "White Thighhighs")
}

package main

import (
	"io"
	"net/http"
)

func writeData(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Writing Response to Server...")
}

func main() {
	//handle routing
	//http://localhost:8080/
	http.HandleFunc("/", writeData)

	//Listen on Port 8080
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"net/http"
)

func main() {
	// Tu http.Handle llama aquí
	http.ListenAndServe("localhost:4000", nil)
}

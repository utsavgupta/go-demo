package main

import (
	"net/http"

	"github.com/utsavgupta/go-demo/calc"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", calc.CalculateHandler)
	http.ListenAndServe(":8080", r)
}

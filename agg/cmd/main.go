package main

import (
	"net/http"

	"github.com/utsavgupta/go-demo/agg"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", agg.NewAggrHandler(agg.NewAsyncCalcRequester(http.DefaultClient)))
	http.ListenAndServe(":8081", r)
}

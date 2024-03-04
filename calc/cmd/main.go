package main

import (
	"net/http"
	"os"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/utsavgupta/go-demo/calc"
)

func main() {

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("Calc Service"),
		newrelic.ConfigLicense(os.Getenv("nr_license")),
	)

	if err != nil {
		panic(err)
	}

	r := http.NewServeMux()
	r.HandleFunc(newrelic.WrapHandleFunc(app, "/calc", calc.CalculateHandler))
	http.ListenAndServe(":8080", r)
}

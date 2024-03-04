package main

import (
	"net/http"
	"os"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/utsavgupta/go-demo/agg"
)

func main() {

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("Agg Service"),
		newrelic.ConfigLicense(os.Getenv("nr_license")),
	)

	if err != nil {
		panic(err)
	}

	client := http.DefaultClient
	client.Transport = newrelic.NewRoundTripper(client.Transport)

	r := http.NewServeMux()
	r.HandleFunc(newrelic.WrapHandleFunc(app, "/agg", agg.NewAggrHandler(agg.NewAsyncCalcRequester(client))))
	http.ListenAndServe(":8081", r)
}

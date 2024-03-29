package agg

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"encoding/json"

	"github.com/utsavgupta/go-demo/calc"
)

type ApiAggRequest struct {
	Operand1 int `json:"a"`
	Operand2 int `json:"b"`
}

type ApiAggResponse struct {
	Results []calc.ApiResponse `json:"results"`
}

type AsyncCalcRequester func(*calc.ApiRequest, chan calc.ApiResponse)

func NewAggrHandler(requester AsyncCalcRequester) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var req ApiAggRequest

		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		c := make(chan calc.ApiResponse)

		/// Call calc service for all 4 operations
		go requester(&calc.ApiRequest{Operand1: req.Operand1, Operand2: req.Operand2, Operation: "+"}, c)
		go requester(&calc.ApiRequest{Operand1: req.Operand1, Operand2: req.Operand2, Operation: "-"}, c)
		go requester(&calc.ApiRequest{Operand1: req.Operand1, Operand2: req.Operand2, Operation: "*"}, c)
		go requester(&calc.ApiRequest{Operand1: req.Operand1, Operand2: req.Operand2, Operation: "/"}, c)

		results := make([]calc.ApiResponse, 0, 4)
		timeout := time.After(50 * time.Millisecond)

		for i := 0; i < 4; i++ {
			select {
			case <-timeout:
				break
			case result := <-c:
				results = append(results, result)
			}
		}

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&ApiAggResponse{results})
	}
}

func NewAsyncCalcRequester(client *http.Client) AsyncCalcRequester {

	return func(request *calc.ApiRequest, c chan calc.ApiResponse) {

		b, _ := json.Marshal(request)
		br := bytes.NewReader(b)

		resp, err := client.Post("http://localhost:8080/", "appplication/json", br)

		if err != nil {
			log.Printf("Could not complete request, error %v", err.Error())
			c <- calc.ApiResponse{Message: "Unable to connect to backend service"}
			return
		}

		var apiResponse calc.ApiResponse
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
			log.Printf("Could decode response, due to error %v", err.Error())
		}
		c <- apiResponse
	}
}

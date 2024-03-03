package calc

import (
	"encoding/json"
	"log"
	"net/http"
)

type ApiRequest struct {
	Operand1  int    `json:"a"`
	Operand2  int    `json:"b"`
	Operation string `json:"op"`
}

type ApiResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Operand1   int    `json:"a,omitempty"`
	Operand2   int    `json:"b,omitempty"`
	Operation  string `json:"op,omitempty"`
	Result     *int   `json:"result,omitempty"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	var req ApiRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var res ApiResponse

	res.Operand1 = req.Operand1
	res.Operand2 = req.Operand2
	res.Operation = req.Operation

	switch req.Operation {
	case "+":
		res.StatusCode = http.StatusOK
		res.Message = "Success"
		val := req.Operand1 + req.Operand2
		res.Result = &val
	case "-":
		res.StatusCode = http.StatusOK
		res.Message = "Success"
		val := req.Operand1 - req.Operand2
		res.Result = &val
	case "*":
		res.StatusCode = http.StatusOK
		res.Message = "Success"
		val := req.Operand1 * req.Operand2
		res.Result = &val
	case "/":
		res.StatusCode = http.StatusOK
		res.Message = "Success"
		val := req.Operand1 / req.Operand2
		res.Result = &val
	default:
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Invalid operation"
		log.Printf("Encountered Invalid Operation: %s", req.Operation)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	json.NewEncoder(w).Encode(res)
}

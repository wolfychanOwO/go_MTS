package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// API
const version = "v1.0.0"

type DecodeRequest struct {
	InputString string `json:"inputString"`
}

type DecodeResponse struct {
	OutputString string `json:"outputString"`
}

// /version
func VersionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, version)
}

// /decode
func DecodeHandler(w http.ResponseWriter, r *http.Request) {
	var request DecodeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(request.InputString)
	if err != nil {
		http.Error(w, "Invalid base64 string", http.StatusBadRequest)
		return
	}

	response := DecodeResponse{OutputString: string(decoded)}
	json.NewEncoder(w).Encode(response)
}

// /hard-op
func HardOpHandler(w http.ResponseWriter, r *http.Request) {
	// Симуляция работы с случайной задержкой от 10 до 20 секунд
	sleepTime := 10 + rand.Intn(10)
	time.Sleep(time.Duration(sleepTime) * time.Second)

	if rand.Float32() < 0.5 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "500 Internal Server Error")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "200 OK")
	}
}

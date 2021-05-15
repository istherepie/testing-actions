package service

import (
	"encoding/json"
	"net/http"
)

func Index(instance string) func(w http.ResponseWriter, r *http.Request) {

	type Response struct {
		Instance string `json:"instance"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		response := &Response{instance}
		body, err := json.Marshal(response)

		if err != nil {
			http.Error(w, "Unable to create response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}

func New(instance string) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index(instance))

	return mux
}

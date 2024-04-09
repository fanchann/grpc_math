package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteToResponseBody[X any](w http.ResponseWriter, statusCode int, res X) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(res); err != nil {
		log.Fatalf(err.Error())
	}
}

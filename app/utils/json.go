package json

import (
	"encoding/json"
	"net/http"
)

func WriteJsonHTTP(w http.ResponseWriter, httpCode int, data interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(&data)
}

package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(endpoint func(w http.ResponseWriter, r *http.Request) (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := endpoint(w, r)
		if err != nil {
			handleError(err, w)
			return
		}

		jsonResult, err := json.Marshal(result)
		if err != nil {
			handleError(err, w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(jsonResult)
	}
}

func handleError(err error, w http.ResponseWriter) {
	log.Println(err)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(400)
	w.Write([]byte(err.Error()))
}

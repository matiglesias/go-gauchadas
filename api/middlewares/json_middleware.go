package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(endpoint func(r *http.Request) (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Accept", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		result, err := endpoint(r)
		if err != nil {
			handleError(err, w)
			return
		}

		jsonResult, err := json.Marshal(result)
		if err != nil {
			handleError(err, w)
			return
		}

		w.WriteHeader(200)
		w.Write(jsonResult)
	}
}

func handleError(err error, w http.ResponseWriter) {
	log.Println(err)
	w.WriteHeader(400)
	w.Write([]byte(err.Error()))
}

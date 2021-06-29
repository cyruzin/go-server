package middleware

import (
	"log"
	"net/http"
	"os"
)

func LogMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.SetOutput(os.Stdout)
		log.Println(r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}

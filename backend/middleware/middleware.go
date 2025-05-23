package middleware

import (
	"log"
	"net/http"
	"time"
)

var headers = map[string]string{
	"Access-Control-Allow-Origin": "*",
	"Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, OPTIONS",
	"Access-Control-Allow-Headers": "Authorization, Content-Type",
}

func CorsMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s, %s, %s", r.Method, r.RequestURI, time.Now())
		for key, value := range headers {
			w.Header().Set(key, value)
		}

		if (r.Method == http.MethodOptions) {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
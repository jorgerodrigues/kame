package routes

import (
	"net/http"
	"strings"
)

var trustedOrigins = []string{
	"http://localhost:3000",
	"http://localhost:3001",
}
var allowedMethods = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
	http.MethodOptions,
}

var allowedHeaders = []string{
	"Accept",
	"Accept-Language",
	"Content-Type",
	"Content-Language",
	"Origin",
}

func (r *Routes) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Vary", "Origin")
		w.Header().Set("Vary", "Access-Control-Request-Method")

		origin := r.Header.Get("Origin")

		if origin != "" {
			for i := range trustedOrigins {
				if origin == trustedOrigins[i] {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
						w.Header().Set("Access-Control-Allow-Methods", strings.Join(allowedMethods, ", "))
						w.Header().Set("Access-Control-Allow-Headers", strings.Join(allowedHeaders, ", "))
						w.WriteHeader(http.StatusOK)
						return
					}
					break
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}

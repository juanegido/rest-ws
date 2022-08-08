package middleware

import (
	"net/http"
	"rest-ws/helpers"
	"rest-ws/server"
)

var (
	NO_AUTH_NEEDED = []string{"signup", "login"}
)

func shouldCheckToken(route string) bool {
	for _, r := range NO_AUTH_NEEDED {
		if r == route {
			return false
		}
	}
	return true
}

func CheckAuthMiddlware(s server.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !shouldCheckToken(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}
			_, err := helpers.GetJWTAuthorizationToken(s, w, r)
			if err != nil {
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

package middleware

import (
	"github.com/golang-jwt/jwt"
	"net/http"
	"rest-ws/models"
	"rest-ws/server"
	"strings"
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
			tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
			_, err := jwt.ParseWithClaims(tokenString, models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(s.Config().JWTSecret), nil
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

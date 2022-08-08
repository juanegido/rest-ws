package helpers

import (
	"github.com/golang-jwt/jwt"
	"net/http"
	"rest-ws/models"
	"rest-ws/server"
	"strings"
)

func GetJWTAuthorizationToken(s server.Server, w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
	token, err := jwt.ParseWithClaims(tokenString, models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Config().JWTSecret), nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return nil, err
	}
	return token, nil
}

package http

import (
	"errors"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"

	log "github.com/sirupsen/logrus"
)

func JWTAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			w.WriteHeader(http.StatusUnauthorized)
			log.Error("an unauthorized request has been made")
			return
		}

		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			log.Error("authorization header could not be parsed")
			return
		}

		if validateToken(authHeaderParts[1]) {
			original(w, r)
		} else {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

	}
}

func validateToken(accessToken string) bool {
	var mySigningKey = []byte("secret-key")

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("could not validate auth token")
		}
		return mySigningKey, nil
	})
	if err != nil {
		return false
	}

	return token.Valid
}

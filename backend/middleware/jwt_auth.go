package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func JwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Missing auth token"))
			return
		}

		tokenParts := strings.Split(tokenHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Invalid auth token"))
			return
		}

		token, err := jwt.Parse(tokenParts[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("jwtSecret")), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Invalid auth token"))
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Invalid auth token"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

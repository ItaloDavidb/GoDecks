package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/italodavidb/goCrud/internal/utils/jwtUtils"
)

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		if tokenString == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &jwtUtils.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de assinatura inválido")
			}
			return jwtUtils.JwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		if claims, ok := token.Claims.(*jwtUtils.UserClaims); ok && token.Valid {
			r.Header.Set("Username", claims.Username)
		}

		next.ServeHTTP(w, r)
	})
}

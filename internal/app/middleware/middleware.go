package middleware

import (
	config "CRUD/internal/app"
	logger "CRUD/internal/app/logs"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

func IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenCookie, err := r.Cookie("token")
		if err != nil {
			logger.Error.Print("Error occurred while reading cookie")
		}
		if tokenCookie != nil {
			claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(tokenCookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte((*config.GetConfig()).JWT.Key), nil
			})

			if token.Valid {
				if claims["Role"] == "admin" {
					next.ServeHTTP(w, r)
					return
				}
				logger.Info.Println("Insufficient access rights")
			} else if errors.Is(err, jwt.ErrTokenMalformed) {
				logger.Info.Println("That's not even a token")
			} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
				logger.Info.Println("Timing is everything")
			} else {
				logger.Info.Println("Couldn't handle this token:", err)
			}
		}
		http.Redirect(w, r, "/login", 301)
	})
}

package service

import (
	config "CRUD/internal/app"
	logger "CRUD/internal/app/logs"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

var (
	signingKey          = []byte((*config.GetConfig()).JWT.Key)
	ttl                 = 2 * time.Minute
	microsecondInSecond = 1000000000
)

type tokenClaims struct {
	jwt.StandardClaims
	Role string
}

func CreateTokenWithRole(userRole string) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(ttl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userRole,
	})
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		logger.Error.Println("Error signing JWT key!")
		panic(err)
	}

	return tokenString
}

func SetTokenInCookies(w http.ResponseWriter, token string) {
	cookie := &http.Cookie{Name: "token", Value: token, MaxAge: int(ttl) / microsecondInSecond, HttpOnly: true}
	http.SetCookie(w, cookie)
}

func UnsetTokenInCookies(w http.ResponseWriter) {
	cookie := &http.Cookie{Name: "token", Value: "", MaxAge: int(ttl) / microsecondInSecond, HttpOnly: true}
	http.SetCookie(w, cookie)
}

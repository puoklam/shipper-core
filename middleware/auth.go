package middleware

import (
	"net/http"
	"os"

	"github.com/go-chi/jwtauth/v5"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte(os.Getenv("HS256_SECRET")), nil)
}

func Authenticator(next http.Handler) http.Handler {
	next = jwtauth.Verifier(tokenAuth)(next)
	next = jwtauth.Authenticator(next)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

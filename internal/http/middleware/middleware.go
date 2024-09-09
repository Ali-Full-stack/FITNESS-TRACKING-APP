package middleware

import (
	"net/http"

	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/internal/auth/token"
)

func ConfirmTokenMiddleware(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userToken :=r.Header.Get("token")
		if userToken == " " {
			http.Error(w, "missing token !", http.StatusUnauthorized)
			return
		}
		if err :=token.VerifyToken(userToken); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
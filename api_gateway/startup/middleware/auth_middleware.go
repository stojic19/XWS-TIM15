package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
)

func AuthMiddleware(w http.ResponseWriter, req *http.Request, next func(http.ResponseWriter, *http.Request)) {
	token := req.Header.Get("token")
	fmt.Println(token)
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("jwtsecret"), nil
	})
	for key, value := range claims {
		if key == "sub" {
			req.Header.Set("sub", value.(string))
			break
		}
	}
	next(w, req)
}

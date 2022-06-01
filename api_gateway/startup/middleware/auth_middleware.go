package middleware

import (
	"fmt"
	"net/http"
)

func AuthMiddleware(w http.ResponseWriter, req *http.Request, next func(http.ResponseWriter, *http.Request)) {
	fmt.Println("AuthMiddleware")
	token := req.Header.Get("token")
	fmt.Println(token)
	if token == "token" {
		w.WriteHeader(401)
		w.Write([]byte("Unauthorized"))
		return
	}
	req.Header.Set("Server", "MiddlewareMux")
	next(w, req)
}

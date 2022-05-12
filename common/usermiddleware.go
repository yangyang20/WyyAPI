package common

import (
	"fmt"
	"net/http"
)

type UserMiddleware struct {
}

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

func (m *UserMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Global_Cookie = r.Cookies()
		fmt.Println("gCookie", Global_Cookie)
		// Passthrough to next handler if need
		next(w, r)
	}
}

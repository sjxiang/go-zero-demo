package middleware

import (
	"fmt"
	"net/http"
)


type SetGlobalMiddleware struct {
}

func NewGlobalMiddleware() *SetGlobalMiddleware {
	return &SetGlobalMiddleware{}
}

func (m *SetGlobalMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("global before")
		
		next(w, r)

		fmt.Println("golbal end")
	}
}
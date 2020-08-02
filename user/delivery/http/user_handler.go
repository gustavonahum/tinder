package http

import (
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
}

func NewUserHandler(w http.ResponseWriter, r *http.Request) {
	handler := &UserHandler{}
}

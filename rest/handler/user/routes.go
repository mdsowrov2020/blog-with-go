package user

import (
	"net/http"

	"blog/rest/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /users", manager.With(http.HandlerFunc(h.Create)))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(h.Login)))
}

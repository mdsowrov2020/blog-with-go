package post

import (
	"net/http"

	"blog/rest/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /posts",
		manager.With(
			http.HandlerFunc(h.Create),
		),
	)

	mux.Handle("GET /posts",
		manager.With(
			http.HandlerFunc(h.List),
		),
	)
	mux.Handle("GET /posts/{id}",
		manager.With(
			http.HandlerFunc(h.Get),
		),
	)

	mux.Handle("DELETE /posts/{id}",
		manager.With(
			http.HandlerFunc(h.Delete),
		),
	)

	mux.Handle("PUT /posts/{id}",
		manager.With(
			http.HandlerFunc(h.Update),
		),
	)
}

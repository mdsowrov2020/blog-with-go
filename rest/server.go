package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"blog/config"
	"blog/rest/handler/post"
	"blog/rest/middleware"
)

type Server struct {
	postHandler *post.Handler
}

func NewServer(postHandler *post.Handler) *Server {
	return &Server{
		postHandler: postHandler,
	}
}

func (server *Server) Start(cnf *config.Config) {
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)
	mux := http.NewServeMux()

	server.postHandler.RegisterRoutes(mux, manager)

	wrapedMux := manager.WrapWithMux(mux)

	port := ":" + strconv.Itoa(cnf.HTTPPort)

	err := http.ListenAndServe(port, wrapedMux)
	if err != nil {
		fmt.Println("Server connection failed")
		return
	}
}

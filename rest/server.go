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
	cnf         *config.Config
	postHandler *post.Handler
}

func NewServer(
	cnf *config.Config,
	postHandler *post.Handler,
) *Server {
	return &Server{
		cnf:         cnf,
		postHandler: postHandler,
	}
}

func (server *Server) Start() {
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)
	mux := http.NewServeMux()

	server.postHandler.RegisterRoutes(mux, manager)

	wrapedMux := manager.WrapWithMux(mux)

	port := ":" + strconv.Itoa(server.cnf.HTTPPort)

	fmt.Println("Server running on port: ", port)
	err := http.ListenAndServe(port, wrapedMux)
	if err != nil {
		fmt.Println("Server connection failed")
		return
	}
}

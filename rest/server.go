package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"blog/config"
	"blog/rest/handler/post"
	"blog/rest/handler/user"
	"blog/rest/middleware"
)

type Server struct {
	cnf         *config.Config
	postHandler *post.Handler
	userHandler *user.Handler
}

func NewServer(
	cnf *config.Config,
	postHandler *post.Handler,
	userHandler *user.Handler,
) *Server {
	return &Server{
		cnf:         cnf,
		postHandler: postHandler,
		userHandler: userHandler,
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
	server.userHandler.RegisterRoutes(mux, manager)

	wrapedMux := manager.WrapWithMux(mux)

	port := ":" + strconv.Itoa(server.cnf.HTTPPort)

	fmt.Println("Server running on port: ", port)
	err := http.ListenAndServe(port, wrapedMux)
	if err != nil {
		fmt.Println("Server connection failed")
		return
	}
}

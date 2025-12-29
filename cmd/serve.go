package cmd

import (
	"blog/config"
	"blog/rest"
	"blog/rest/handler/post"
	"blog/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()

	middlewares := middleware.NewMiddlewares(cnf)

	postHandler := post.NewHandler(middlewares)

	server := rest.NewServer(cnf, postHandler)
	server.Start()
}

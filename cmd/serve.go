package cmd

import (
	"blog/config"
	"blog/repo"
	"blog/rest"
	"blog/rest/handler/post"
	"blog/rest/handler/user"
	"blog/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()

	middlewares := middleware.NewMiddlewares(cnf)

	postRepo := repo.NewPostRepo()
	userRepo := repo.NewUserRepo()

	postHandler := post.NewHandler(postRepo, middlewares)
	userHandler := user.NewHandler(userRepo, cnf)

	server := rest.NewServer(cnf, postHandler, userHandler)
	server.Start()
}

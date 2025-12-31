package cmd

import (
	"fmt"
	"os"

	"blog/config"
	"blog/infra/db"
	"blog/repo"
	"blog/rest"
	"blog/rest/handler/post"
	"blog/rest/handler/user"
	"blog/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	middlewares := middleware.NewMiddlewares(cnf)

	postRepo := repo.NewPostRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	postHandler := post.NewHandler(postRepo, middlewares)
	userHandler := user.NewHandler(userRepo, cnf)

	server := rest.NewServer(cnf, postHandler, userHandler)
	server.Start()
}

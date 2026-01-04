package cmd

import (
	"fmt"
	"os"

	"blog/config"
	"blog/infra/db"
	"blog/post"
	"blog/repo"
	"blog/rest"
	pstHandler "blog/rest/handler/post"
	usrHandler "blog/rest/handler/user"
	"blog/rest/middleware"
	"blog/user"
)

func Serve() {
	cnf := config.GetConfig()

	// fmt.Printf("%+v", cnf.DB)

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println("Database migration failed", err)
		os.Exit(1)
	}

	middlewares := middleware.NewMiddlewares(cnf)

	// Repo
	postRepo := repo.NewPostRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	// Domains
	userService := user.NewService(userRepo)
	postService := post.NewService(postRepo)

	// Handler
	postHandler := pstHandler.NewHandler(postService, middlewares)
	userHandler := usrHandler.NewHandler(userService, cnf)

	server := rest.NewServer(cnf, postHandler, userHandler)
	server.Start()
}

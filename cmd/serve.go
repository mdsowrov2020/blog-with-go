package cmd

import (
	"blog/config"
	"blog/rest"
	"blog/rest/handler/post"
)

func Serve() {
	cnf := config.GetConfig()

	postHandler := post.NewHandler()

	server := rest.NewServer(postHandler)
	server.Start(cnf)
}

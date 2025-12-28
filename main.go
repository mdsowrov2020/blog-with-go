package main

import (
	"fmt"

	"blog/util"
)

// "blog/cmd"

func main() {
	// cmd.Serve()

	jwt, err := util.CreateJWT("my-secret", util.Payload{
		Sub:       1,
		FirstName: "Md Sowrov",
		LastName:  "Khadem",
		IsAdmin:   false,
	})
	if err != nil {
		fmt.Println("JWT is invalid")
		return
	}

	fmt.Println("YOUR JWT: ", jwt)
}

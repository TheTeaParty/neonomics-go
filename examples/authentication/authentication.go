package main

import (
	"context"
	"fmt"
	"github.com/TheTeaParty/neonomics-go"
	"os"
)

func main() {

	c := neonomics.NewSandbox(&neonomics.Config{
		ClientID: os.Getenv("CLIENT_ID"),
		SecretID: os.Getenv("SECRET_ID"),
	})

	rsp, err := c.TokenRequest(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)

}

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

	ctx := context.WithValue(context.Background(), neonomics.CtxKeyDeviceID, "111111")

	rsp, err := c.GetSupportedBanks(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)

}

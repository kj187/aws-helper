package main

import (
	"fmt"
	"os"
	"github.com/kj187/aws-inspector/commands"
)

func main() {
	fmt.Println(`
____ _ _ _ ____    _ _  _ ____ ___  ____ ____ ___ ____ ____ 
|__| | | | [__     | |\ | [__  |__] |___ |     |  |  | |__/ 
|  | |_|_| ___]    | | \| ___] |    |___ |___  |  |__| |  \ 
© Julian Kleinhans - @kj187	
	`)

	/*
	https://docs.aws.amazon.com/de_de/sdk-for-go/v1/developer-guide/configuring-sdk.html
	
	os.Setenv("AWS_PROFILE", test-account)

	if os.Getenv("AWS_ACCESS_KEY_ID") == "" {
		color.Red("No AWS_ACCESS_KEY_ID env var available!")
		return
	}
	if os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
		color.Red("No AWS_SECRET_ACCESS_KEY env var available!")
		return
	}*/

	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
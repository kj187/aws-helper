package main

import (
	"fmt"
	"github.com/kj187/aws-inspector/commands"
)

func main() {
	fmt.Println(`
____ _ _ _ ____    _ _  _ ____ ___  ____ ____ ___ ____ ____ 
|__| | | | [__     | |\ | [__  |__] |___ |     |  |  | |__/ 
|  | |_|_| ___]    | | \| ___] |    |___ |___  |  |__| |  \ 
© Julian Kleinhans - @kj187	
	`)

	commands.Execute()
}
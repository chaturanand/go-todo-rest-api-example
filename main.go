package main

import (
	"github.com/chaturanand/go-todo-rest-api-example/app"
	"github.com/chaturanand/go-todo-rest-api-example/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
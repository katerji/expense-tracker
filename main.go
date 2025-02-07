package main

import (
	"github.com/katerji/expense-tracker/env"
	"github.com/katerji/expense-tracker/webserver"
)

func main() {
	env.InitEnv()
	webserver.InitWebServer()
}

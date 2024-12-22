package main

import (
	"calc_go_anatoliy/pkg/application"
)

func main() {
	app := application.New()
	app.RunServer()
}

package main

import (
	"github.com/ad-07/calc_go_anatoliy/pkg/application"
)

func main() {
	app := application.New()
	// app.Run()
	app.RunServer()
}

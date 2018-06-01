package main

import (
	"github.com/Leomn138/Widget-Factory/blob/master/app"
	"github.com/Leomn138/Widget-Factory/blob/master/config"
)
func main() {

	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(config.Port)
}
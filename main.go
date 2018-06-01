package main

import (
	//"github.com/Leomn138/Widget-Factory/app"
	"widgetFactory/app"
	"widgetFactory/config"
	//"github.com/Leomn138/Widget-Factory/config"
)
func main() {

	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(config.Port)
}
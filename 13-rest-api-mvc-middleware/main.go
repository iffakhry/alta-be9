package main

import (
	_config "be9/restmvc/config"
	_middlewares "be9/restmvc/middlewares"
	_routes "be9/restmvc/routes"
)

func main() {
	_config.InitDB()
	e := _routes.New()
	_middlewares.LogMiddlewares(e)

	e.Logger.Fatal(e.Start(":80"))
}

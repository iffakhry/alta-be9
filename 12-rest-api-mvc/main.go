package main

import (
	_config "be9/restmvc/config"
	_routes "be9/restmvc/routes"
)

func main() {
	_config.InitDB()
	e := _routes.New()
	e.Logger.Fatal(e.Start(":80"))
}

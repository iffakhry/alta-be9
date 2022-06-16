package main

import (
	"be9/restclean/routes"
)

func main() {
	e := routes.New()
	e.Logger.Fatal(e.Start(":80"))
}

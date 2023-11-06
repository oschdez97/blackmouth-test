package main

import (
	"github.com/oschdez97/blackmouth-test/app"
	_ "github.com/oschdez97/blackmouth-test/docs"
)

// @title           Blackmouth Test Service
// @version         1.0
// @description     A game management service API in Go using Gin framework.
// @contact.name   	Oscar Hernandez
// @contact.email  	oschdez97@gmail.com
// @host      		localhost:8080
// @BasePath  		/api/v1
func main() {
	var a app.App
	a.CreateConnection()
	a.Migrate()
	a.CreateRoutes()
	a.Run()
}
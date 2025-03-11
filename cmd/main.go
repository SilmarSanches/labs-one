package main

import (
	"fmt"
	"labs-one/internal/infra/web/webserver"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "labs-one/docs" 
)

// @title Tudo Azul API
// @version 1.0
// @description Tudo Azul Temperaturas
// @BasePath /
func main() {

	getTemp := NewGetTempoHandler()

	httpServer := webserver.NewWebServer(NewConfig())
	httpServer.AddHandler("GET", "/swagger/*", httpSwagger.WrapHandler)
	httpServer.AddHandler("GET", "/get-temp", getTemp.HandleLabsOne)
	fmt.Println("HTTP server running at port 8080")
	httpServer.Start()
}

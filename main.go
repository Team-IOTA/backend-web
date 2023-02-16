package main

import (
	"classifie/pkg/apis"
	"classifie/pkg/connections"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {

}

func init() {
	fmt.Println("Connecting to Database")
	connections.ConnectMongo()
	fmt.Println("Initializing HTTP Server")
	initEchoServer()
}

func initEchoServer() {
	e := echo.New()
	apis.EchoManager(e)
	e.Logger.Fatal(e.Start(":8000"))
}


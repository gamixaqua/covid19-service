package main

import (
	"fmt"

	"github.com/covid19-service/src/main/config"
	"github.com/labstack/echo"
)

func main() {
	//connecting database
	config.SetUpDatabase()
	//connecting Redis
	config.SetUpRedis()

	e := echo.New()
	v1 := e.Group("/api")
	config.InitializeApplicationConfig(v1)
	e.Logger.Fatal(e.Start("localhost:8000"))
	fmt.Println("Server has started at 8000 !")
}

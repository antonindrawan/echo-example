package main

import (
	"net/http"

	_ "echo-example/docs"
	handlers "echo-example/handlers"

	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var index = `
	<!DOCTYPE html>
	<html lang="en">
	<body>
		<h1>Hello World! <a href=swagger/index.html>See the swagger doc</a></h1>
		<p>
			Build with <a href=https://echo.labstack.com>echo</a>
			and <a href=https://github.com/swaggo/echo-swagger>echo-swagger</a>
		</p>
	</body>
`

// @title Echo Example
// @description This is an echo application
// @version 1.0
// @host localhost:8080
// @BasePath /v1
func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `<h1>Hello World! <a href=swagger/index.html>See the swagger doc</a></h1>
			<br/> Build with <a href=https://echo.labstack.com>echo</a>
			and <a href=https://github.com/swaggo/echo-swagger>echo-swagger</a>`)
	})

	v1 := e.Group("/v1")
	{
		v1.GET("/user/:id", handlers.GetUser)
		v1.POST("/user", handlers.CreateUser)
		v1.PUT("/user", handlers.UpdateUser)
		v1.DELETE("/user/:id", handlers.DeleteUser)
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

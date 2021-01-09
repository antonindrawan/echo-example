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

var login = `
	<script src="https://apis.google.com/js/platform.js" async defer>
	</script>
	<script>
		function onSignIn(googleUser) {
			var profile = googleUser.getBasicProfile();
			console.log('ID: ' + profile.getId()); // Do not send to your backend! Use an ID token instead.
			console.log('Name: ' + profile.getName());
			console.log('Image URL: ' + profile.getImageUrl());
			console.log('Email: ' + profile.getEmail()); // This is null if the 'email' scope is not present.
		}
		function signOut() {
			var auth2 = gapi.auth2.getAuthInstance();
			auth2.signOut().then(function () {
				console.log('User signed out.');
			});
		}

	</script>
	<html lang="en">
	<meta name="google-signin-client_id" content="913549998475-dggku238g8m65v0rpumofeidu6n8s5qr.apps.googleusercontent.com">
	<div class="g-signin2" data-onsuccess="onSignIn"></div>
	<a href="#" onclick="signOut();">Sign out</a>
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
		return c.HTML(http.StatusOK, index)
	})

	e.GET("/login", func(c echo.Context) error {
		return c.HTML(http.StatusOK, login)
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

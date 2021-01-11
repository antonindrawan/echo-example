package main

import (
	"net/http"
	"os"

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
			fetch('/v1/auth', {
				method: 'POST',
				headers: {
					'Authorization': 'Bearer ' + googleUser.getAuthResponse().id_token
				}
			}).then(function (response) {
				if (response.ok) {
					return response.text();
				}
				return Promise.reject(response);
			}).then(function (data) {
				console.log(data);
			}).catch(function (error) {
				console.warn('Something went wrong.', error);
			});
		}
		function signOut() {
			var auth2 = gapi.auth2.getAuthInstance();
			auth2.signOut().then(function () {
				console.log('User signed out.');
			});
		}

	</script>
	<html lang="en">
	<meta name="google-signin-client_id" content="` + os.Getenv("APPLICATION_CLIENT_ID") + `">
	<div class="g-signin2" data-onsuccess="onSignIn"></div>
	<a href="#" onclick="signOut();">Sign out</a>
`

// @title Echo Example
// @description This is an echo application
// @version 1.0
// @license.name MIT

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

	v1 := e.Group("/v1")
	{
		v1.GET("/login", func(c echo.Context) error {
			return c.HTML(http.StatusOK, login)
		})
		v1.POST("/auth", handlers.Login)

		user := v1.Group("/user")
		{
			user.GET("/:id", handlers.GetUser)
			user.POST("", handlers.CreateUser)
			user.PUT("", handlers.UpdateUser)
			user.DELETE("/:id", handlers.DeleteUser)
		}
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

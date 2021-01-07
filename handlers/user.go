package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetUser godoc
// @Summary Get a user
// @Description Get a user based on an {id}.
// @Param id path string true "Get User"
// @Success 200 {string} string
// @Failure 400 {string} string "Error"
// @Router /user/{id} [get]
func GetUser(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]string{
		"id": id,
	})
}

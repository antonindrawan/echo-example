package handlers

import (
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"

	"echo-example/models"
)

var (
	currentID int64 = 1
	users           = map[int64]*models.User{}
)

// GetUser godoc
// @Summary Get a user based on an {id}.
// @Description Get a user based on an {id}.
// @Param id path string true "User ID"
// @Success 200 {object} object
// @Failure 404 {object} object "Error message"
// @Router /user/{id} [get]
func GetUser(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if value, ok := users[id]; ok {
		return c.JSON(http.StatusOK, value)
	}

	errorMsg := map[string]string{
		"error_msg": "The ID does not exist",
	}
	return c.JSON(http.StatusNotFound, errorMsg)
}

// CreateUser godoc
// @Summary Create a user
// @Description Create a user.
// @Accept json
// @Produce json
// @Param user body models.User true "A user object"
// @Success 201 {object} object "Created"
// @Failure 400 {object} object "Error message"
// @Router /user [post]
func CreateUser(c echo.Context) error {

	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	u.ID = currentID
	users[currentID] = u
	currentID++
	return c.JSON(http.StatusCreated, u)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user.
// @Accept json
// @Produce json
// @Param user body models.User true "A user object"
// @Success 200 {object} object "Created"
// @Failure 400 {object} object "Error message"
// @Router /user [put]
func UpdateUser(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	id := u.ID

	if _, ok := users[id]; ok {
		users[id] = u
		return c.JSON(http.StatusOK, u)
	}

	errorMsg := map[string]string{
		"error_msg": "The ID does not exist",
	}
	return c.JSON(http.StatusNotFound, errorMsg)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user based on an {id}.
// @Param id path string true "Delete User"
// @Success 204
// @Failure 400 {object} object "Error message"
// @Router /user/{id} [delete]
func DeleteUser(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		errorMsg := map[string]string{
			"error_msg": "The ID is not an int64",
		}
		c.JSON(http.StatusBadRequest, errorMsg)
	}

	delete(users, id)

	return c.NoContent(http.StatusNoContent)

}

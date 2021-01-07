package handlers

import (
	"echo-example/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	context := e.NewContext(req, rec)
	context.SetPath("/v1/user/:id")
	context.SetParamNames("id")
	context.SetParamValues("42")

	if assert.NoError(t, handlers.GetUser(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"id":"42"}`+"\n", rec.Body.String())
	}
}

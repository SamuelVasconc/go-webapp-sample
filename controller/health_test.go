package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SamuelVasconc/go-webapp-sample/test"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetHealthCheck(t *testing.T) {
	router, container := test.Prepare()

	health := NewHealthController(container)
	router.GET(APIHealth, func(c echo.Context) error { return health.GetHealthCheck(c) })

	req := httptest.NewRequest("GET", APIHealth, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `"healthy"`, rec.Body.String())
}

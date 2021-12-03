package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SamuelVasconc/go-webapp-sample/model"
	"github.com/SamuelVasconc/go-webapp-sample/test"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetFormatList(t *testing.T) {
	router, container := test.Prepare()

	format := NewFormatController(container)
	router.GET(APIFormats, func(c echo.Context) error { return format.GetFormatList(c) })

	req := httptest.NewRequest("GET", APIFormats, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	data := [...]*model.Format{
		{ID: 1, Name: "Paper Book"},
		{ID: 2, Name: "e-Book"},
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

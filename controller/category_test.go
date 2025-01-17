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

func TestGetCategoryList(t *testing.T) {
	router, container := test.Prepare()

	category := NewCategoryController(container)
	router.GET(APICategories, func(c echo.Context) error { return category.GetCategoryList(c) })

	req := httptest.NewRequest("GET", APICategories, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	data := [...]*model.Category{
		{ID: 1, Name: "Technical Book"},
		{ID: 2, Name: "Magazine"},
		{ID: 3, Name: "Novel"},
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

package controllers

import (
	"dot-hiring-go/models"
	"dot-hiring-go/tests"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

var bookController = BookController{}

func TestBookGetAll(t *testing.T) {
	models.DB.Create(models.User{
		Name: "John",
	})

	mockResponse := `{"message":"Welcome..."}`
	r := tests.SetUp()
	r.GET("/", bookController.GetAll)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

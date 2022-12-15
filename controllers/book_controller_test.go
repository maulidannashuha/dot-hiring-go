package controllers

import (
	"dot-hiring-go/models"
	"dot-hiring-go/tests"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

var bookController = BookController{}

func TestBookGetAll(t *testing.T) {
	r := tests.SetUp()

	user := models.User{
		Name:       "John",
		Email:      "johndoe@admin.com",
		DefaultField: models.DefaultField{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	models.DB.Create(&user)

	mockResponse := `{"data":{"id":2,"created_at":"time","updated_at":"2022-12-15T04:13:25.191563084+07:00","deleted_at":null,"name":"maulidan nashuha","email":"maulidannashuha@gmail.com","books":[]}`
	r.GET("/", bookController.GetAll)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

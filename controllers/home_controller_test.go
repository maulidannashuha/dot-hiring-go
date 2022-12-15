package controllers

import (
	"dot-hiring-go/tests"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

var homeController = HomeController{}

func TestHomepageHandler(t *testing.T) {
	mockResponse := `{"message":"Welcome..."}`
	r := tests.SetUp()
	r.GET("/", homeController.Welcome)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

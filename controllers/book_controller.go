package controllers

import (
	"dot-hiring-go/models"
	"dot-hiring-go/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookController struct{}

var cache utils.Cache
var CACHE_KEY = "BookController.GetAll"

func (ctrl BookController) GetAll(c *gin.Context) {
	var user *models.User
	var data = cache.Get(CACHE_KEY)

	if data == "" {
		if err := models.DB.Where("id = ?", c.Param("userId")).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}

		data, err := json.Marshal(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		cache.Set(CACHE_KEY, data)
	} else {
		err := json.Unmarshal([]byte(data), &user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func (ctrl BookController) Store(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("userId")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title}
	models.DB.Create(&book)

	cache.Remove(CACHE_KEY)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (ctrl BookController) Update(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateBook := models.Book{Title: input.Title}

	models.DB.Model(&book).Updates(updateBook)

	cache.Remove(CACHE_KEY)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func (ctrl BookController) Delete(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	cache.Remove(CACHE_KEY)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

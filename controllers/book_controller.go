package controllers

import (
	"dot-hiring-go/models"
	"dot-hiring-go/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct{}

var BOOK_CACHE_KEY = "BookController.GetAll"

func (ctrl BookController) GetAll(c *gin.Context) {
	var user *models.User
	var data = cache.Get(BOOK_CACHE_KEY)

	if data == "" {
		if err := models.DB.Preload("Books").Where("id = ?", c.Param("userId")).First(&user).Error; err != nil {
			utils.Logger(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}

		data, err := json.Marshal(user)
		if err != nil {
			utils.Logger(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		cache.Set(BOOK_CACHE_KEY, data)
	} else {
		err := json.Unmarshal([]byte(data), &user)
		if err != nil {
			utils.Logger(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

type CreateBookInput struct {
	Title string `json:"title" binding:"required"`
}

func (ctrl BookController) Store(c *gin.Context) {
	var userId = c.Param("userId")
	var user models.User
	if err := models.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		utils.Logger(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Logger(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	intUserId, err := strconv.Atoi(userId)
	if err != nil {
		utils.Logger(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, UserID: uint(intUserId)}
	models.DB.Create(&book)

	cache.Remove(BOOK_CACHE_KEY)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (ctrl BookController) Update(c *gin.Context) {
	var userId = c.Param("userId")
	var book models.Book
	if err := models.DB.Where("id = ?", userId).First(&book).Error; err != nil {
		utils.Logger(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Logger(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateBook := models.Book{Title: input.Title}

	models.DB.Model(&book).Updates(updateBook)

	cache.Remove(BOOK_CACHE_KEY)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func (ctrl BookController) Delete(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("userId")).First(&book).Error; err != nil {
		utils.Logger(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	cache.Remove(BOOK_CACHE_KEY)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

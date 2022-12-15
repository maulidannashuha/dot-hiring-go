package controllers

import (
	"dot-hiring-go/models"
	"dot-hiring-go/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var USER_CACHE_KEY = "UserController.GetAll"

func (ctrl UserController) GetAll(c *gin.Context) {
	var users []*models.User
	var data = cache.Get(BOOK_CACHE_KEY)

	if data == "" {
		models.DB.Find(&users)

		data, err := json.Marshal(users)
		if err != nil {
			utils.Logger(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		cache.Set(USER_CACHE_KEY, data)
	} else {
		err := json.Unmarshal([]byte(data), &users)
		if err != nil {
			utils.Logger(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

type CreateUserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func (ctrl UserController) Store(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Logger(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: input.Name, Email: input.Email}
	models.DB.Create(&user)

	cache.Remove(USER_CACHE_KEY)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

type UpdateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (ctrl UserController) Update(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("userId")).First(&user).Error; err != nil {
		utils.Logger(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Logger(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateUser := models.User{Name: input.Name, Email: input.Email}

	models.DB.Model(&user).Updates(updateUser)

	cache.Remove(USER_CACHE_KEY)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (ctrl UserController) Delete(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("userId")).First(&user).Error; err != nil {
		utils.Logger(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&user)

	cache.Remove(USER_CACHE_KEY)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

package controllers

import (
	"dot-hiring-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (ctrl UserController) GetAll(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

type CreateUserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func (ctrl UserController) Store(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: input.Name, Email: input.Email}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

type UpdateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (ctrl UserController) Update(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateUser := models.User{Name: input.Name, Email: input.Email}

	models.DB.Model(&user).Updates(updateUser)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (ctrl UserController) Delete(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

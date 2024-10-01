package controllers

import (
	"net/http"
	"trilune/database"
	"trilune/models"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Succesfuly": category})
	}

}

func GetAllCategories(c *gin.Context) {
	var category []models.Category

	database.DB.Find(&category)
	c.JSON(http.StatusOK, category)

}

func GetCategoriesByID(c *gin.Context) {
	var category models.Category

	id := c.Param("id")

	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, category)

}

func UpdateCategories(c *gin.Context) {
	var category models.Category

	id := c.Param("id")

	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	if err := database.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, category)
}

func DeleteCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
	}

	if err := database.DB.Delete(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Expense deleted successfully"})

}

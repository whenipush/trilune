package controllers

import (
	"net/http"
	"trilune/database"
	"trilune/models"

	"github.com/gin-gonic/gin"
)

func CreateIncomes(c *gin.Context) {
	var incomes models.Incomes
	var category models.Category

	if err := c.ShouldBindJSON(&incomes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if incomes.CategoryID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category is required"})
		return
	}

	if err := database.DB.First(&category, incomes.CategoryID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category does not exist"})
		return
	}

	if err := database.DB.Create(&incomes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Succesfuly": incomes})
	}
}

func GetAllIncomes(c *gin.Context) {
	var incomes []models.Incomes

	database.DB.Find(&incomes)
	c.JSON(http.StatusOK, incomes)

}

func GetIncomesByID(c *gin.Context) {
	var incomes models.Incomes

	id := c.Param("id")

	if err := database.DB.First(&incomes, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Incomes not found"})
		return
	}
	c.JSON(http.StatusOK, incomes)

}

func UpdateIncomes(c *gin.Context) {
	var incomes models.Incomes

	id := c.Param("id")

	if err := database.DB.First(&incomes, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
	}

	if err := c.ShouldBindJSON(&incomes); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	if err := database.DB.Save(&incomes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, incomes)
}

func DeleteIncomes(c *gin.Context) {
	var incomes models.Incomes
	id := c.Param("id")

	if err := database.DB.First(&incomes, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
	}

	if err := database.DB.Delete(&incomes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Expense deleted successfully"})

}

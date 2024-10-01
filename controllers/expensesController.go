package controllers

import (
	"net/http"
	"trilune/database"
	"trilune/models"

	"github.com/gin-gonic/gin"
)

func CreateExpenses(c *gin.Context) {
	var expenses models.Expenses
	var category models.Category

	if err := c.ShouldBindJSON(&expenses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if expenses.CategoryID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category is required"})
		return
	}

	if err := database.DB.First(&category, expenses.CategoryID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category does not exist"})
		return
	}

	if err := database.DB.Create(&expenses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Succesfuly": expenses})
	}

}

func GetAllExpenses(c *gin.Context) {
	var expenses []models.Expenses

	database.DB.Find(&expenses)
	c.JSON(http.StatusOK, expenses)

}

func GetExpensesByID(c *gin.Context) {

	id := c.Param("id")

	var expense models.Expenses

	// Ищем расходы по ID в базе данных
	if err := database.DB.First(&expense, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}
	c.JSON(http.StatusOK, expense)

}

func UpdateExpenses(c *gin.Context) {
	var expenses models.Expenses

	id := c.Param("id")

	if err := database.DB.First(&expenses, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
	}

	if err := c.ShouldBindJSON(&expenses); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	if err := database.DB.Save(&expenses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, expenses)
}

func DeleteExpenses(c *gin.Context) {
	var expenses models.Expenses
	id := c.Param("id")

	if err := database.DB.First(&expenses, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
	}

	if err := database.DB.Delete(&expenses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Expense deleted successfully"})

}

// var Category = []models.Category{{Name: "Первая категория"}, {Name: "Вторая категория"}}

// expenses := models.Expenses{Name: "Вторая трата", Money: 100, Categorys: Category}

// database.DB.Create(&expenses)

package main

import (
	"trilune/controllers"
	"trilune/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	r := gin.Default()
	r.GET("/expenses", controllers.GetAllExpenses)
	r.GET("/expenses/:id", controllers.GetExpensesByID)
	r.PUT("/expenses/:id", controllers.UpdateExpenses)
	r.POST("/expenses", controllers.CreateExpenses)
	r.DELETE("/expenses/:id", controllers.DeleteExpenses)

	r.GET("/incomes", controllers.GetAllIncomes)
	r.GET("/incomes/:id", controllers.GetIncomesByID)
	r.PUT("/incomes/:id", controllers.UpdateIncomes)
	r.POST("/incomes", controllers.CreateIncomes)
	r.DELETE("/incomes/:id", controllers.DeleteIncomes)

	r.GET("/category", controllers.GetAllCategories)
	r.GET("/category/:id", controllers.GetCategoriesByID)
	r.PUT("/category/:id", controllers.UpdateCategories)
	r.POST("/category", controllers.CreateCategory)
	r.DELETE("/category/:id", controllers.DeleteCategory)

	r.Run()

}

package main

import (
	"trilune/database"
)

func main() {
	database.Connect()

	// Example insert values
	// var Category = []models.Category{{Name: "Первая категория"}, {Name: "Вторая категория"}}

	// expenses := models.Expenses{Name: "Вторая трата", Money: 100, Categorys: Category}

	// database.DB.Create(&expenses)
}

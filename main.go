package main

import (
	"github.com/gin-gonic/gin"
	"github.com/szlezak/recipe-api/database"
	"github.com/szlezak/recipe-api/handlers"
)

func main() {
	r := gin.Default()

	// Initialize Database
	database.ConnectDatabase()

	// Routes
	r.GET("/recipes", handlers.FindRecipes)
	r.POST("/recipes", handlers.CreateRecipe)
	r.DELETE("/recipes/:id", handlers.DeleteRecipe)
	r.Run()
}
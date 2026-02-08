package main

import (
	"github.com/gin-gonic/gin"
	"github.com/szlezak/recipe-api/database"
	"github.com/szlezak/recipe-api/handlers"
	"github.com/szlezak/recipe-api/middleware"
)

func main() {
	r := gin.Default()

	// Initialize Database
	database.ConnectDatabase()
	r.Use(middleware.MyLogger())
	// Routes
	r.GET("/health", handlers.HealthCheck)
	r.GET("/recipes", handlers.FindRecipes)
	r.GET("/recipes/:id", handlers.FindRecipe)
	r.POST("/recipes", handlers.CreateRecipe)
	r.PUT("/recipes/:id", handlers.UpdateRecipe)
	r.DELETE("/recipes/:id", handlers.DeleteRecipe)
	r.GET("/recipes/search", handlers.SearchRecipes)
	
	r.Run()
}
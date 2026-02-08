package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/szlezak/recipe-api/database"
	"github.com/szlezak/recipe-api/models"
)

func HealthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status":  "available",
        "time":    time.Now().Format(time.RFC3339),
        "version": "1.0.0",
    })
}

func FindRecipes(c *gin.Context) {
	var recipes []models.Recipe

	database.DB.Find(&recipes)
	c.JSON(http.StatusOK, recipes)
}

func FindRecipe(c *gin.Context) {
    recipeId := c.Param("id")
    var recipe models.Recipe
    // Instead of .First(&recipe, recipeId)
    // Use .Where() to be 100% explicit about the column and the value
    if err := database.DB.Where("id = ?", recipeId).First(&recipe).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
        return
    }

    c.JSON(http.StatusOK, recipe)
}

func CreateRecipe(c *gin.Context) {
	var input models.Recipe
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&input)
	c.JSON(http.StatusCreated, input)
}

func UpdateRecipe(c *gin.Context) {
    recipeId := c.Param("id")
    var recipe models.Recipe      // To hold the existing record from DB
    var input models.Recipe       // To hold the new data from JSON

    // 1. Find the existing recipe first
    if err := database.DB.Where("id = ?", recipeId).First(&recipe).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
        return
    }

    // 2. Bind the incoming JSON to our 'input' variable
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 3. Update the existing record with the input data
    database.DB.Model(&recipe).Updates(input)

    c.JSON(http.StatusOK, recipe)
}

func DeleteRecipe(c *gin.Context) {
    recipeId := c.Param("id")
    var recipe models.Recipe

    if err := database.DB.Where("id = ?", recipeId).First(&recipe).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
        return
    }

    // Delete it
    database.DB.Delete(&recipe)

    c.JSON(http.StatusOK, gin.H{"message": "Recipe #" + recipeId + " deleted"})
}

func SearchRecipes(c *gin.Context) {
    // Get the "title" from the URL query string
    title := c.Query("title") 
    var recipes []models.Recipe

    // We use a "LIKE" query with % to find partial matches (e.g., "Pas" finds "Pasta")
    // The ? is a placeholder to prevent SQL Injection (safety first!)

    if result := database.DB.Where("title LIKE ?", "%"+title+"%").Find(&recipes); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        return
    }

    c.JSON(http.StatusOK, recipes)
}
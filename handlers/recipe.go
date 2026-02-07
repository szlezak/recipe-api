package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/szlezak/recipe-api/database"
	"github.com/szlezak/recipe-api/models"
)

func FindRecipes(c *gin.Context) {
	var recipes []models.Recipe

	database.DB.Find(&recipes)
	c.JSON(http.StatusOK, recipes)
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

func DeleteRecipe(c *gin.Context) {
	var recipeId string
	if err := c.ShouldBindUri(&recipeId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	database.DB.Delete(&models.Recipe{}, recipeId)
	c.Status(http.StatusNoContent)
}
package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecipeHandler struct {
	collection *mongo.Collection
	ctx context.Context
}


func newRecipeHandler(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest ,gin.H{"error": err.Error()})
		return
	}
	recipe.Id = primitive.NewObjectID()
	recipe.PublishedAt = time.Now()
	_, err := client.Database("demo").Collection("recipes").InsertOne(ctx, recipe)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not insert a new recipe"})
		return
	}
	c.JSON(http.StatusOK, recipe)
}
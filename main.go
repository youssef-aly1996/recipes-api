package main

import (
	"context"
	"log"
	"os"
	"recipes-api/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ctx context.Context
var err error
var client *mongo.Client
var recipesHandler *handlers.RecipesHandler
func init()  {
	
	// recipes = make([]Recipe, 0)
	// file, _ := ioutil.ReadFile("recipes.json")
	// _ = json.Unmarshal([]byte(file), &recipes)
	// for _, recipe := range recipes {
	// 	listOfRecipes = append(listOfRecipes, recipe)
	// }
	ctx = context.Background()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err := client.Ping(context.TODO(),  readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to Mongodb")
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")
	recipesHandler = handlers.NewRecipesHandler(ctx, collection)

	// insertManyResult, err := collection.InsertMany(ctx, listOfRecipes)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("many recipes inserted", len(insertManyResult.InsertedIDs))
}


func updateRecipeHandler(c *gin.Context)  {
// 	id := c.Param("id")
// 	var recipe Recipe
// 	if err := c.ShouldBindJSON(&recipe); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error()})
// 	return
// 	}
// 	index := -1
// 	for i:=0; i<len(recipes); i++ {
// 		if recipes[i].Id == id {
// 			index = i
// 		}
// 	}
// 	if index == -1 {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "recipe not found"})
// 		return
// 	}
// 	recipes[index] = recipe
// 	c.JSON(http.StatusOK, recipe)
}

func deleteRecipeHandler(c *gin.Context)  {
	// id := c.Param("id")
	// index := -1
	// for i := 0; i < len(recipes); i++ {
	// 	if recipes[i].Id == id {
	// 		index = i
	// 	}
	// }
	// if index == -1 {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "recipe not found"})
	// }
	// recipes = append(recipes[:index], recipes[index+1:]...)
	// c.JSON(http.StatusOK, gin.H{"message": "Recipe has been deleted"})
}

func main() {
	router := gin.Default()
	router.POST("/recipes", recipesHandler.ListRecipesHandler)
	router.GET("/recipes", recipesHandler.NewRecipeHandler)
	router.PUT("/recipes/:id", updateRecipeHandler)
	router.DELETE("/recipes/:id", deleteRecipeHandler)
	router.Run()
}
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ctx context.Context
var err error
var client *mongo.Client
var recipes [] Recipe
var listOfRecipes []interface{}

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
	// collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")
	// insertManyResult, err := collection.InsertMany(ctx, listOfRecipes)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("many recipes inserted", len(insertManyResult.InsertedIDs))
}
type Recipe struct {
	Id primitive.ObjectID `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Tags []string `json:"tags" bson:"tags"`
	Ingredients []string `json:"ingredients" bson:"ingredients"`
	Instructions []string `json:"instructions" bson:"instructions"`
	PublishedAt time.Time `json:"publishedAt" bson:"publishedAt"`
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

func listRecipesHandler(c *gin.Context) {
	cur, err := client.Database("demo").Collection("recipes").Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cur.Close(ctx)
	recipes = make([]Recipe, 0)
	for cur.Next(ctx) {
		var recipe Recipe
		cur.Decode(&recipe)
		recipes = append(recipes, recipe)
	}
	c.JSON(http.StatusOK, recipes)
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
	router.POST("/recipes", newRecipeHandler)
	router.GET("/recipes", listRecipesHandler)
	router.PUT("/recipes/:id", updateRecipeHandler)
	router.DELETE("/recipes/:id", deleteRecipeHandler)
	router.Run()
}
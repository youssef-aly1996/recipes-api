package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

var recipes [] Recipe

func init()  {
	recipes = make([]Recipe, 0)
	file, _ := ioutil.ReadFile("recipes.json")
	_ = json.Unmarshal([]byte(file), &recipes)
	// f, err := os.Open("recipes.json") 
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer f.Close()
	// recipes,err := ioutil.ReadAll(f)
	// if err := json.Unmarshal(recipes, &recipes); err != nil {
	// 	fmt.Println(err)
	// }
}
type Recipe struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Tags []string `json:"tags"`
	Ingredients []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
	PublishedAt time.Time `json:"publishedAt"`
}

func newRecipeHandler(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest ,gin.H{"error": err.Error()})
		return
	}
	recipe.Id = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)
}

func listRecipesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}

func updateRecipeHandler(c *gin.Context)  {
	id := c.Param("id")
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
	return
	}
	index := -1
	for i:=0; i<len(recipes); i++ {
		if recipes[i].Id == id {
			index = i
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "recipe not found"})
		return
	}
	recipes[index] = recipe
	c.JSON(http.StatusOK, recipe)
}

func main() {
	router := gin.Default()
	router.POST("/recipes", newRecipeHandler)
	router.GET("/recipes", listRecipesHandler)
	router.PUT("/recipes/:id", updateRecipeHandler)
	router.Run()
}
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"practice.com/http/pkg/repository/recipes"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func homePage(c *gin.Context) {
	c.String(http.StatusOK, "This is my home page")
}

type RecipesHandler struct {
	store recipeStore
	s3    *recipes.S3Store
	redis *recipes.Redis
}

func NewRecipesHandler(s recipeStore, s3 *recipes.S3Store, redis *recipes.Redis) *RecipesHandler {
	return &RecipesHandler{
		store: s,
		s3:    s3,
		redis: redis,
	}
}

type recipeStore interface {
	Add(name string, recipe recipes.Recipe) error
	Get(name string) (recipes.Recipe, error)
	List() (map[string]recipes.Recipe, error)
	Update(name string, recipe recipes.Recipe) error
	Remove(name string) error
}

// CreateRecipe handles the creation of a new recipe.
//
// It takes a gin.Context object as a parameter and expects the request body to be a JSON
// representation of a recipes.Recipe object. The function converts the request body into
// a recipes.Recipe object and creates a URL-friendly name for the recipe using the slug.Make
// function. The recipe is then added to the store using the h.store.Add method. If there is an
// error during the process, the function returns a JSON response with the error message. If
// everything is successful, the function returns a JSON response with a "status" field set to
// "success".
func (h RecipesHandler) CreateRecipe(c *gin.Context) {
	// Get request body and convert it to recipes.Recipe
	var recipe recipes.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create a url friendly name
	id := slug.Make(recipe.Name)

	// add to the store
	CheckError(h.store.Add(id, recipe))

	// return success payload
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// UploadImage handles the upload of an image file.
//
// It takes a gin.Context object as a parameter and expects the request to contain a file
// with the key "file". The function retrieves the file from the request, saves it to a
// temporary directory, uploads it to an S3 bucket using the h.s3.UploadToS3 method, and
// then removes the temporary file. If there is an error during the process, the function
// returns a JSON response with the error message. If everything is successful, the
// function returns a JSON response with a "status" field set to "success".
func (h RecipesHandler) UploadImage(c *gin.Context) {

	// single file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err))
		return
	}
	log.Println(file.Filename)

	dst := "./tmp/" + file.Filename

	// Upload the file to specific dst.
	err = c.SaveUploadedFile(file, dst)

	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err))
		return
	}

	h.s3.UploadToS3(dst)

	err = os.Remove(dst)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("remove file err: %s", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h RecipesHandler) GetImage(c *gin.Context) {

	var tempURL recipes.Image
	image := c.Param("image")

	cache, err := h.redis.GetImageURL(image)

	if err != nil {
		tempURL = h.s3.GenerateURL(image)
		CheckError(h.redis.AddImageURL(image, tempURL.URL))
		log.Println("cache miss")
	} else {
		tempURL = cache
		log.Println("cache hit")
	}

	fmt.Println(tempURL)
	//send url string
	c.JSON(http.StatusOK, tempURL)

}

func (h RecipesHandler) ListRecipes(c *gin.Context) {
	r, err := h.store.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(200, r)
}

func (h RecipesHandler) GetRecipe(c *gin.Context) {
	id := c.Param("id")

	recipe, err := h.store.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(200, recipe)
}

func (h RecipesHandler) UpdateRecipe(c *gin.Context) {
	// Get request body and convert it to recipes.Recipe
	var recipe recipes.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	err := h.store.Update(id, recipe)
	if err != nil {
		if err == recipes.ErrFoo {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return success payload
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h RecipesHandler) DeleteRecipe(c *gin.Context) {

	// var recipe recipes.Recipe
	// if err := c.ShouldBindJSON(&recipe); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	id := c.Param("id")

	err := h.store.Remove(id)
	if err != nil {
		if err == recipes.ErrFoo {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return success payload
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

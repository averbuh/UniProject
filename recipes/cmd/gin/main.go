package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"practice.com/http/pkg/repository/recipes"
)

func main() {
	// Create Gin router
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))
	// Instantiate recipe Handler and provide a data store

	s3 := recipes.NewS3Store("us-east-1", "test-images-vue")

	var psqlconn string
	host, exist := os.LookupEnv("POSTGRES_HOST")
	if !exist {
		log.Print("POSTGRES_HOST not set")
	}
	port := 5432
	user, exist := os.LookupEnv("POSTGRES_USER")
	if !exist {
		log.Print("POSTGRES_USER not set")
	}
	password, exist := os.LookupEnv("POSTGRES_PASSWORD")
	if !exist {
		log.Print("POSTGRES_PASSWORD not set")
	}
	dbname, exist := os.LookupEnv("POSTGRES_DB")
	if !exist {
		log.Print("POSTGRES_DB not set")
	}
	Addr, exist := os.LookupEnv("REDIS_HOST")
	if !exist {
		log.Print("REDIS_HOST not set")
	}
	DB := 0
	Password, exist := os.LookupEnv("REDIS_PASSWORD")
	if !exist {
		log.Print("REDIS_PASSWORD not set")
	}

	// Create new redis object
	redis, err := recipes.NewRedis(Addr, Password, DB)
	if err != nil {
		log.Print("Failed to connect to redis: ", err)
	} else {
		log.Print("Connected to redis")
	}

	// Create new postgres object
	psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		log.Print("Failed to connect to database: ", err)
	}
	store, err := recipes.NewPostgres(db)

	store.CreateTestTable(db)

	defer store.CloseDB()
	if err != nil {
		log.Print("Failed to connect to database: ", err)
	} else {
		log.Print("Connected to database")
	}
	// Create new handler
	recipesHandler := NewRecipesHandler(store, &s3, redis)

	recipesRoutes := map[string]string{
		"id": "/recipes/:id",
	}

	// Register Routes
	router.GET("/", homePage)
	router.GET("/recipes", recipesHandler.ListRecipes)
	router.POST("/recipes", recipesHandler.CreateRecipe)
	router.POST("/recipes/upload", recipesHandler.UploadImage)
	router.GET("/recipes/:id/:image", recipesHandler.GetImage)
	router.GET(recipesRoutes["id"], recipesHandler.GetRecipe)
	router.PUT(recipesRoutes["id"], recipesHandler.UpdateRecipe)
	router.DELETE(recipesRoutes["id"], recipesHandler.DeleteRecipe)

	// Start the server
	router.Run()
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
	h.store.Add(id, recipe)

	// return success payload
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

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

	var tempUrl recipes.Image
	image := c.Param("image")

	cache, err := h.redis.GetImageURL(image)

	if err != nil {
		tempUrl = h.s3.GenerateUrl(image)
		h.redis.AddImageURL(image, tempUrl.Url)
		log.Println("cache miss")
	} else {
		tempUrl = cache
		log.Println("cache hit")
	}

	fmt.Println(tempUrl)
	//send url string
	c.JSON(http.StatusOK, tempUrl)

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
		if err == recipes.NotFoundErr {
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
		if err == recipes.NotFoundErr {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return success payload
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

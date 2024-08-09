package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"practice.com/http/pkg/repository/recipes"
)

func main() {
	// Create Gin router
	router := gin.Default()
	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"averbuchpro.com"}
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
	CheckError(err)

	CheckError(store.CreateTestTable(db))

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
	CheckError(router.Run())
}

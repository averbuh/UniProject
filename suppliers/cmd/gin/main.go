package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"

	"practice.com/http/pkg/repository/suppliers"
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

	// s3 := suppliers.NewS3Store("us-east-1", "test-images-vue")

	var psqlconn string
	host, exist := os.LookupEnv("POSTGRES_HOST")
	if !exist {
		panic("POSTGRES_HOST not set")
	}
	port := 5432
	user, exist := os.LookupEnv("POSTGRES_USER")
	if !exist {
		panic("POSTGRES_USER not set")
	}
	password, exist := os.LookupEnv("POSTGRES_PASSWORD")
	if !exist {
		panic("POSTGRES_PASSWORD not set")
	}
	dbname, exist := os.LookupEnv("POSTGRES_DB")
	if !exist {
		panic("POSTGRES_DB not set")
	}
	// Addr, exist := os.LookupEnv("REDIS_HOST")
	// if !exist {
	// 	panic("REDIS_HOST not set")
	// }
	// DB := 0
	// Password, exist := os.LookupEnv("REDIS_PASSWORD")
	// if !exist {
	// 	panic("REDIS_PASSWORD not set")
	// }

	psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		log.Print("Failed to connect to database: ", err)
	}
	store, err := suppliers.NewPostgres(db)

	store.CreateTestTable(db)

	defer store.CloseDB()
	if err != nil {
		log.Print("Failed to connect to database: ", err)
	} else {
		log.Print("Connected to database")
	}
	// Instantiate supplier Handler and provide a data store
	suppliersHandler := NewSuppliersHandler(store)
	// defer store.CloseDB()

	suppliersRoutes := map[string]string{
		"id": "/suppliers/:id",
	}

	// Register Routes
	router.GET("/suppliers", suppliersHandler.ListSuppliers)
	router.POST("/suppliers", suppliersHandler.CreateSupplier)
	router.GET(suppliersRoutes["id"], suppliersHandler.GetSupplier)
	router.PUT(suppliersRoutes["id"], suppliersHandler.UpdateSupplier)
	router.DELETE(suppliersRoutes["id"], suppliersHandler.DeleteSupplier)
	//TODO: Get recommented suppliers based on today recipes ingredients
	//TODO: Get favourite suppliers

	// Start the server
	router.Run()
}

func homePage(c *gin.Context) {
	c.String(http.StatusOK, "This is my home page")
}

type SuppliersHandler struct {
	store supplierStore
}

func NewSuppliersHandler(s supplierStore) *SuppliersHandler {
	return &SuppliersHandler{
		store: s,
	}
}

type supplierStore interface {
	Add(name string, supplier suppliers.Supplier) error
	Get(name string) (suppliers.Supplier, error)
	List() (map[string]suppliers.Supplier, error)
	Update(name string, supplier suppliers.Supplier) error
	Remove(name string) error
}

func (h SuppliersHandler) CreateSupplier(c *gin.Context) {
	// Get request body and convert it to suppliers.Supplier
	var supplier suppliers.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create a url friendly name
	id := slug.Make(supplier.Name)

	// add to the store
	h.store.Add(id, supplier)

	// return success payload
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h SuppliersHandler) ListSuppliers(c *gin.Context) {
	r, err := h.store.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(200, r)
}

func (h SuppliersHandler) GetSupplier(c *gin.Context) {
	id := c.Param("id")

	supplier, err := h.store.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(200, supplier)
}

func (h SuppliersHandler) UpdateSupplier(c *gin.Context) {
	// Get request body and convert it to suppliers.Supplier
	var supplier suppliers.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	err := h.store.Update(id, supplier)
	if err != nil {
		if err == errors.New("not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return success payload
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h SuppliersHandler) DeleteSupplier(c *gin.Context) {
	id := c.Param("id")

	err := h.store.Remove(id)
	if err != nil {
		if err == errors.New("not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return success payload
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

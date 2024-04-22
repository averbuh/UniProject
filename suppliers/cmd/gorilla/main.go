package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"practice.com/http/pkg/repository/recipes"
)

func main() {

	store := recipes.NewMemStore()

	router := mux.NewRouter()
	s := router.PathPrefix("/recipes").Subrouter()

	NewRecipesHandler(store, s)

	http.ListenAndServe(":8010", router)
}

func NewRecipesHandler(s recipeStore, router *mux.Router) *RecipesHandler {

	handler := &RecipesHandler{
		store: s,
	}

	router.HandleFunc("/", handler.ListRecipes).Methods("GET")
	router.HandleFunc("/", handler.CreateRecipe).Methods("POST")
	router.HandleFunc("/{id}", handler.GetRecipe).Methods("GET")
	router.HandleFunc("/{id}", handler.UpdateRecipe).Methods("PUT")
	router.HandleFunc("/{id}", handler.DeleteRecipe).Methods("DELETE")

	return handler
}

type recipeStore interface {
	Add(name string, recipe recipes.Recipe) error
	Get(name string) (recipes.Recipe, error)
	List() (map[string]recipes.Recipe, error)
	Update(name string, recipe recipes.Recipe) error
	Remove(name string) error
}

type RecipesHandler struct {
	store recipeStore
}

func (h RecipesHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe recipes.Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)

}
func (h RecipesHandler) ListRecipes(w http.ResponseWriter, r *http.Request)  {}
func (h RecipesHandler) GetRecipe(w http.ResponseWriter, r *http.Request)    {}
func (h RecipesHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {}
func (h RecipesHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}

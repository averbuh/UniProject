package main

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/gosimple/slug"
	"practice.com/http/pkg/repository/recipes"
)

func main() {
	store := recipes.NewMemStore()
	RecipesHandler := NewRecipesHandler(store)

	mux := http.NewServeMux()

	mux.Handle("/", &homeHandler{})
	mux.Handle("/recipes", RecipesHandler)
	mux.Handle("/recipes/", RecipesHandler)

	http.ListenAndServe(":8080", mux)

}

type homeHandler struct{}

var (
	RecipeRe       = regexp.MustCompile(`^/recipes/*$`)
	RecipeReWithID = regexp.MustCompile(`^/recipes/([a-z0-9]+(?:-[a-z0-9]+)+)$`)
)

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home page!"))
}

type recipeStore interface {
	Add(name string, recipe recipes.Recipe) error
	Get(name string) (recipes.Recipe, error)
	Update(name string, recipe recipes.Recipe) error
	List() (map[string]recipes.Recipe, error)
	Remove(name string) error
}

type RecipesHandler struct {
	store recipeStore
}

func NewRecipesHandler(s recipeStore) *RecipesHandler {
	return &RecipesHandler{
		store: s,
	}
}

func (h *RecipesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && RecipeRe.MatchString(r.URL.Path):
		h.CreateRecipe(w, r)

	case r.Method == http.MethodGet && RecipeRe.MatchString(r.URL.Path):
		h.ListRecipes(w, r)

	case r.Method == http.MethodGet && RecipeReWithID.MatchString(r.URL.Path):
		h.GetRecipe(w, r)

	case r.Method == http.MethodPut && RecipeReWithID.MatchString(r.URL.Path):
		h.UpdateRecipe(w, r)

	case r.Method == http.MethodDelete && RecipeReWithID.MatchString(r.URL.Path):
		h.DeleteRecipe(w, r)
	}
}

func (h *RecipesHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe recipes.Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	// Convert the name of the recipe into URL friendly string
	resourceID := slug.Make(recipe.Name)
	// Call the store to add the recipe
	if err := h.store.Add(resourceID, recipe); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	// Set the status code to 200
	w.WriteHeader(http.StatusOK)
}

func (h *RecipesHandler) ListRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := h.store.List()
	jsonBytes, err := json.Marshal(recipes)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *RecipesHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	matches := RecipeReWithID.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		InternalServerErrorHandler(w, r)
		return
	}
	recipe, err := h.store.Get(matches[1])
	if err != nil {
		if err == recipes.NotFoundErr {
			NotFoundHandler(w, r)
			return
		}
	}

	jsonBytes, err := json.Marshal(recipe)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}
func (h *RecipesHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	matches := RecipeReWithID.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		InternalServerErrorHandler(w, r)
		return
	}

	var recipe recipes.Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		InternalServerErrorHandler(w, r)
	}

	if err := h.store.Update(matches[1], recipe); err != nil {
		if err == recipes.NotFoundErr {
			NotFoundHandler(w, r)
			return
		}
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}
func (h *RecipesHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	matches := RecipeReWithID.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		InternalServerErrorHandler(w, r)
		return
	}

	if err := h.store.Remove(matches[1]); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}

package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"practice.com/http/pkg/repository/recipes"
)

// type Recipe struct {
// 	Name        string
// 	IsToday     bool
// 	Ingredients []string
// 	Description string
// 	Image       string
// }

func TestAdd(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	pg, err := recipes.NewPostgres(db)
	if err != nil {
		t.Fatal(err)
	}
	recipe := recipes.Recipe{
		Name:        "Test Recipe",
		IsToday:     true,
		Ingredients: []string{"ingredient1", "ingredient2"},
		Description: "Test Description",
		Image:       "test.jpg",
	}

	mock.ExpectExec("INSERT INTO recipes").
		WithArgs(recipe.Name, recipe.IsToday, pq.Array(recipe.Ingredients), recipe.Description, recipe.Image).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = pg.Add(recipe.Name, recipe)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	pg, err := recipes.NewPostgres(db)
	if err != nil {
		t.Fatal(err)
	}
	recipe := recipes.Recipe{
		Name:        "Test Recipe",
		IsToday:     true,
		Ingredients: []string{"ingredient1", "ingredient2"},
		Description: "Test Description",
		Image:       "test.jpg",
	}

	rows := sqlmock.NewRows([]string{"name", "istoday", "ingredients", "description", "image"}).
		AddRow(recipe.Name, recipe.IsToday, pq.Array(recipe.Ingredients), recipe.Description, recipe.Image)

	mock.ExpectQuery("SELECT name, istoday, ingredients, description, image FROM recipes WHERE name = \\$1").
		WithArgs(recipe.Name).
		WillReturnRows(rows)

	result, err := pg.Get(recipe.Name)
	assert.NoError(t, err)
	assert.Equal(t, recipe, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	pg, err := recipes.NewPostgres(db)
	if err != nil {
		t.Fatal(err)
	}
	recipe := recipes.Recipe{
		Name:        "Test Recipe",
		IsToday:     true,
		Ingredients: []string{"ingredient1", "ingredient2"},
		Description: "Test Description",
		Image:       "test.jpg",
	}

	mock.ExpectExec("UPDATE recipes SET istoday = \\$2, ingredients = \\$3, description = \\$4, image = \\$5 WHERE name = \\$1").
		WithArgs(recipe.Name, recipe.IsToday, pq.Array(recipe.Ingredients), recipe.Description, recipe.Image).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = pg.Update(recipe.Name, recipe)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRemove(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	pg, err := recipes.NewPostgres(db)
	if err != nil {
		t.Fatal(err)
	}
	recipeName := "Test Recipe"

	mock.ExpectExec("DELETE FROM recipes WHERE name = \\$1").
		WithArgs(recipeName).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = pg.Remove(recipeName)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestList(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	pg, err := recipes.NewPostgres(db)
	if err != nil {
		t.Fatal(err)
	}
	recipe := recipes.Recipe{
		Name:        "Test Recipe",
		IsToday:     true,
		Ingredients: []string{"ingredient1", "ingredient2"},
		Description: "Test Description",
		Image:       "test.jpg",
	}

	rows := sqlmock.NewRows([]string{"name", "istoday", "ingredients", "description", "image"}).
		AddRow(recipe.Name, recipe.IsToday, pq.Array(recipe.Ingredients), recipe.Description, recipe.Image)

	mock.ExpectQuery("SELECT name, istoday, ingredients, description, image FROM recipes").
		WillReturnRows(rows)

	result, err := pg.List()
	assert.NoError(t, err)
	assert.Equal(t, map[string]recipes.Recipe{recipe.Name: recipe}, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

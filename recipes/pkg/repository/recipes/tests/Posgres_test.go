package tests

import (
	"errors"
	"testing"

	"practice.com/http/pkg/repository/recipes"

	"github.com/lib/pq"
)

func TestAdd(t *testing.T) {
	// Create a mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()
	// Create a new Postgres instance with the mock database
	p := recipes.Postgres{db: db}
	// Test case 1: Successful insertion
	mock.ExpectExec("INSERT INTO recipes VALUES").WithArgs("test", true, pq.Array([]string{"ingredient1", "ingredient2"}), "description", "image").WillReturnResult(sqlmock.NewResult(1, 1))
	err = p.Add("test", recipes.Recipe{IsToday: true, Ingredients: []string{"ingredient1", "ingredient2"}, Description: "description", Image: "image"})
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	// Test case 2: Insertion failure
	//test
	mock.ExpectExec("INSERT INTO recipes VALUES").WithArgs("test", true, pq.Array([]string{"ingredient1", "ingredient2"}), "description", "image").WillReturnError(errors.New("insertion failed"))
	err = p.Add("test", recipes.Recipe{IsToday: true, Ingredients: []string{"ingredient1", "ingredient2"}, Description: "description", Image: "image"})
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	// Test case 3: Empty name parameter
	err = p.Add("", recipes.Recipe{IsToday: true, Ingredients: []string{"ingredient1", "ingredient2"}, Description: "description", Image: "image"})
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	// Test case 4: Empty recipe parameter
	err = p.Add("test", recipes.Recipe{})
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	// Check if all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

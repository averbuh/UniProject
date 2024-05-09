package suppliers 

import (
	"database/sql"
	"fmt"

	"context"

	"github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) (*Postgres, error) {

	err := db.Ping()

	return &Postgres{
		db: db,
	}, CheckError(err)
}

func CheckError(err error) error {
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}
	return nil
}

func (p Postgres) CloseDB() {
	p.db.Close()
}

func (p Postgres) CreateTestTable(db *sql.DB) error {
	// delete table if it exists
	_, e := db.Exec("CREATE TABLE IF NOT EXISTS suppliers(name varchar not null primary key, istoday boolean, ingredients text[] , description varchar, image varchar)")
	CheckError(e)
	// new suppliers
	ins := "INSERT INTO suppliers (name, istoday, ingredients, description, image) VALUES ($1, $2, $3, $4, $5)"

	ingredients := []string{"eggs", "milk", "flour", "sugar"}

	_, e = db.Exec(ins, "toast", true, pq.Array(ingredients), "some description", "")

	CheckError(e)
	return e
}

// Add adds a new supplier to the Postgres database.
//
// Parameters:
// - name: the name of the supplier (string).
// - supplier: the supplier object containing the supplier details (Supplier).
//
// Returns:
// - error: an error if the insertion fails.
func (p Postgres) Add(name string, supplier Supplier) error {
	_, e := func() (sql.Result, error) {
		var args []any = []any{name, supplier.Description, supplier.IsFavorite, pq.Array(supplier.Products) pq.Array(supplier.Ingredients), supplier.Description, supplier.Image}
		return p.db.ExecContext(context.Background(), "INSERT INTO suppliers VALUES ($1, $2, $3, $4, $5)", args...)
	}()
	CheckError(e)
	return e
}

func (p Postgres) Get(name string) (Supplier, error) {
	fmt.Println("GET called for " + name)
	rows, e := p.db.Query("SELECT name, istoday, ingredients, description, image FROM suppliers WHERE name = $1", name)
	CheckError(e)
	var r Supplier
	for rows.Next() {
		e := rows.Scan(&r.Name, &r.IsToday, pq.Array(&r.Ingredients), &r.Description, &r.Image) // array with str ingredients
		CheckError(e)
	}
	return r, e
}

func (p Postgres) List() (map[string]Supplier, error) {
	fmt.Println("List called")
	sel := "SELECT name, istoday, ingredients, description, image FROM suppliers"
	rows, err := p.db.Query(sel)
	CheckError(err)
	defer rows.Close()
	suppliers := make(map[string]Supplier)

	for rows.Next() {
		var r Supplier
		err := rows.Scan(&r.Name, &r.IsToday, pq.Array(&r.Ingredients), &r.Description, &r.Image) // array with str ingredients
		CheckError(err)
		suppliers[r.Name] = r
	}

	// for rows.Next() {
	// 	var r Supplier
	// 	var json_bytes []byte
	// 	err := rows.Scan(&r.Name, &json_bytes) // array with str ingredients
	// 	CheckError(err)
	// 	err = json.Unmarshal(json_bytes, &r.Ingredients)
	// 	CheckError(err)
	// 	suppliers[r.Name] = r
	// }

	return suppliers, nil
}

func (p Postgres) Update(name string, supplier Supplier) error {
	_, e := func() (sql.Result, error) {
		var args []any = []any{name, supplier.IsToday, pq.Array(supplier.Ingredients), supplier.Description, supplier.Image}
		return p.db.ExecContext(context.Background(), "UPDATE suppliers SET istoday = $2, ingredients = $3, description = $4, image = $5 WHERE name = $1", args...)
	}()
	CheckError(e)
	return e
}

func (p Postgres) Remove(name string) error {
	_, e := p.db.Exec("DELETE FROM suppliers WHERE name = $1", name)
	CheckError(e)
	return e
}

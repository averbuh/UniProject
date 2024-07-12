package suppliers

import (
	"database/sql"
	"fmt"

	"context"
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
	_, e := db.Exec("CREATE TABLE IF NOT EXISTS suppliers(name varchar not null primary key, description varchar, isfavorite boolean, image varchar)")
	CheckError(e)
	// new suppliers
	ins := "INSERT INTO suppliers (name, description, isfavorite, image) VALUES ($1, $2, $3, $4)"

	_, e = db.Exec(ins, "test", "test", true, "https://img.freepik.com/free-photo/colorful-design-with-spiral-design_188544-9588.jpg")

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
	query := "INSERT INTO suppliers (name, description, isfavorite, image) VALUES ($1, $2, $3, $4)"
	args := []any{name, supplier.Description, supplier.IsFavorite, supplier.Image}
	result, err := p.db.ExecContext(context.Background(), query, args...)
	if err != nil {
		return fmt.Errorf("error adding supplier: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		return fmt.Errorf("error adding supplier: %v", err)
	}
	return nil
}

func (p Postgres) Get(name string) (Supplier, error) {
	fmt.Println("GET called for " + name)
	rows, e := p.db.Query("SELECT name, description, isfavorite, image FROM suppliers WHERE name = $1", name)
	CheckError(e)
	var s Supplier
	for rows.Next() {
		e := rows.Scan(&s.Name, &s.Description, &s.IsFavorite, &s.Image) // array with str ingredients
		CheckError(e)
	}
	return s, e
}

func (p Postgres) List() (map[string]Supplier, error) {
	fmt.Println("List called")
	sel := "SELECT name, description, isfavorite, image FROM suppliers"
	rows, err := p.db.Query(sel)
	CheckError(err)
	defer rows.Close()
	suppliers := make(map[string]Supplier)

	for rows.Next() {
		var s Supplier
		err := rows.Scan(&s.Name, &s.Description, &s.IsFavorite, &s.Image) // array with str ingredients
		CheckError(err)
		suppliers[s.Name] = s
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
		var args []any = []any{name, supplier.Description, supplier.IsFavorite, supplier.Image}
		return p.db.ExecContext(context.Background(), "UPDATE suppliers SET description = $2, isfavorite = $3, image = $4 WHERE name = $1", args...)
	}()
	CheckError(e)
	return e
}

func (p Postgres) Remove(name string) error {
	_, e := p.db.Exec("DELETE FROM suppliers WHERE name = $1", name)
	CheckError(e)
	return e
}

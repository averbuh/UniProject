package recipes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var environment = "prod"

const (
	host     = "172.17.0.3"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "postgres"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres() *Postgres {
	var psqlconn string
	if environment == "prod" {
		host := "postgres-postgresql.default.svc.cluster.local"
		port := 5432
		user := os.Getenv("POSTGRES_USER")
		password := os.Getenv("POSTGRES_PASSWORD")
		dbname := os.Getenv("POSTGRES_DB")
		psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	} else {
		psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	}

	fmt.Println(os.Getenv("POSTGRES_DB"))
	fmt.Println(os.Getenv("POSTGRES_PASSWORD"))

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!") // insert
	// hardcoded

	// dynamic
	//	insertDynStmt := `insert into "Students"("Name", "Roll_Number") values($1, $2)`
	//	_, e = db.Exec(insertDynStmt, "Jack", 21)
	//	CheckError(e)
	return &Postgres{
		db: db,
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func (p Postgres) CloseDB() {
	p.db.Close()
}

func (p Postgres) Add(name string, recipe Recipe) error {
	jsonData, e := json.Marshal(recipe.Ingredients)
	CheckError(e)
	_, e = p.db.Exec("INSERT INTO recipes VALUES ($1, $2)", recipe.Name, jsonData)
	CheckError(e)
	return e
}

func (p Postgres) Get(name string) (Recipe, error) {
	fmt.Println("GET called for " + name)
	rows, e := p.db.Query("SELECT name, ingredients FROM recipes WHERE name = $1", name)
	CheckError(e)
	var r Recipe
	for rows.Next() {
		var json_bytes []byte
		e := rows.Scan(&r.Name, &json_bytes) // array with str ingredients
		CheckError(e)
		e = json.Unmarshal(json_bytes, &r.Ingredients)
		CheckError(e)
	}
	return r, e
}

func (p Postgres) List() (map[string]Recipe, error) {
	fmt.Println("List called")
	rows, err := p.db.Query("SELECT name, ingredients FROM recipes")
	CheckError(err)
	defer rows.Close()
	recipes := make(map[string]Recipe)

	for rows.Next() {
		var r Recipe
		var json_bytes []byte
		err := rows.Scan(&r.Name, &json_bytes) // array with str ingredients
		CheckError(err)
		err = json.Unmarshal(json_bytes, &r.Ingredients)
		CheckError(err)
		recipes[r.Name] = r
	}

	return recipes, nil
}

func (p Postgres) Update(name string, recipe Recipe) error {
	return nil
}

func (p Postgres) Remove(name string) error {
	return nil
}

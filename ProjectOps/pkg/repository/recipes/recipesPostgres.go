package recipes

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

const (
	host     = "172.17.0.2"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "postgres"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres() *Postgres {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()
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

func (p Postgres) Add(name string, recipe Recipe) error {
	insertStmt := `insert into recipes (name, ingredients) values ($1, $2)`
	var ing_names []string
	for _, ing := range recipe.Ingredients {
		ing_names = append(ing_names, ing.Name)
	}
	_, e := p.db.Exec(insertStmt, recipe.Name, pq.Array(ing_names))
	CheckError(e)
	return nil
}

func (p Postgres) Get(name string) (Recipe, error) {
	fmt.Println("GET called")
	return Recipe{}, nil
}

func (p Postgres) List() (map[string]Recipe, error) {
	fmt.Println("List called")
	return nil, nil
}

func (p Postgres) Update(name string, recipe Recipe) error {
	return nil
}

func (p Postgres) Remove(name string) error {
	return nil
}

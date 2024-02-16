package main

import (
	"log"

	"github.com/goberman15/sandbox-l4/server"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func initDB(db *sqlx.DB) error {
	schema := `
		CREATE TABLE IF NOT EXISTS item (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL UNIQUE,
			status VARCHAR(24) DEFAULT 'Empty',
			amount INTEGER DEFAULT 0 CHECK (amount >= 0)
		);

		CREATE TABLE IF NOT EXISTS item_detail (
			id SERIAL PRIMARY KEY,
			item_id INTEGER NOT NULL,
			name VARCHAR(255) NOT NULL,
			CONSTRAINT fk_item FOREIGN KEY(item_id) REFERENCES item(id) ON UPDATE CASCADE ON DELETE CASCADE
		);
	`

	_, err := db.Exec(schema)
	return err
}

func main() {
	db, err := sqlx.Connect("pgx", "host=localhost user=postgres password=unbroken dbname=sandbox sslmode=disable")
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}
	defer db.Close()

	err = initDB(db)
	if err != nil {
		log.Fatalf("fail to initialize database: %v", err)
	}
	
	s := server.NewServer(db)
	s.RegisterRouter()

	log.Fatal(s.Start())

}

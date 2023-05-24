package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgressConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgressConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func main() {
	cfg := PostgressConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}
	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	// creating a table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT UNIQUE NOT NULL
	);
		CREATE TABLE IF NOT EXISTS orders (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			amount INT,
			description TEXT
	);
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("tables created")

	//insert data
	// name := "Jennifer White"
	// email := "TEST2@gmail.com"
	// //_, err = db.Exec(`
	// //	INSERT INTO users(name, email)
	// //	VALUES ($1, $2);
	// //`, name, email)
	// // Executing to return data
	// row := db.QueryRow(`
	// 	INSERT INTO users(name, email)
	// 	VALUES	($1, $2) RETURNING id;
	// `, name, email)
	// var id int
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("User created. id", id)

	// query for user with specific id
	id := 1
	row := db.QueryRow(`
		SELECT name, email
		FROM users
		WHERE id = $1;
	`, id)
	var name, email string
	err = row.Scan(&name, &email)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("User information: name=%s, email=%s", name, email)

	// userID := 1
	// for i := 1; i < 5; i++ {
	// 	amount := i * 100
	// 	desc := fmt.Sprintf("Fake order #%d", i)
	// 	_, err := db.Exec(`
	// 		INSERT INTO orders(user_id, amount, description)
	// 		VALUES ($1, $2, $3)
	// 	`, userID, amount, desc)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// fmt.Println("created fake orders")

	type Orders struct {
		ID          int
		UserID      int
		Amount      int
		Description string
	}
	var orders []Orders
	userID := 1
	rows, err := db.Query(`
		SELECT id, amount, description
		FROM orders
		WHERE user_id = $1;
	`, userID)

	defer rows.Close()

	for rows.Next() {
		var order Orders
		order.UserID = userID
		err := rows.Scan(&order.ID, &order.Amount, &order.Description)
		if err != nil {
			panic(err)
		}
		orders = append(orders, order)
	}
	if rows.Err() != nil {
		panic(rows.Err())
	}

	fmt.Println("Orders", orders)
	//check for and error
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type user struct {
	id      int
	name    string
	balance int
}

type product struct {
	id      int
	product string
	price   int
}

type balance struct {
	balance int
}

func (u *user) updateBalance(db *sql.DB, newCash int) {
	_, err := db.Exec("update buyers set balance = $1 where id = $2", newCash, u.id)
	if err != nil {
		log.Fatalf("log in func updateBalance")
	}
}

func (u *user) getBalance(db *sql.DB) (newBalance int) {
	row := db.QueryRow("select balance from buyers where id = $1", u.id)
	b := balance{}
	err := row.Scan(&b.balance)
	if err != nil {
		log.Fatalf("log in func getUserCash")
	}
	return b.balance
}

func getProduct(db *sql.DB, id int) product {
	row := db.QueryRow("select * from  shop where id = $1", id)
	product := product{}
	err := row.Scan(&product.id, &product.product, &product.price)
	if err != nil {
		log.Fatalf("log in func getProduct")
	}
	print(" you bought : ", product.product)
	return product
}

func insertInShoppingList(db *sql.DB, name string) {
	_, err := db.Exec("insert into shopping_list (name) values ($1)", name)
	if err != nil {
		log.Fatalf("log in func insertInShoppingList")
	}
}

func getUser(db *sql.DB) *user {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5) + 1
	row := db.QueryRow("select * from  buyers where id = $1", r)
	user := user{}
	err := row.Scan(&user.id, &user.name, &user.balance)
	if err != nil {
		log.Fatalf("log in func getUser")
	}
	fmt.Println("buyer : ", user.name, "  balance : ", user.balance)
	return &user
}

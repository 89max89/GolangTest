package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	"time"
)

const (
	salary = 30
)

/*
	func buy(db *sql.DB) {
		rand.Seed(time.Now().UnixNano())
		var sliceRandom []int
		var _ []int
		random := rand.Intn(11) + 1
		for i := 0; i < random; i++ {
			newRandom := rand.Intn(10) + 1
			sliceRandom = append(sliceRandom, newRandom)
		}
		sort.Ints(sliceRandom)
		for i := 0; i < len(sliceRandom)-1; i++ {
			if sliceRandom[i] == sliceRandom[i+1] {
				sliceRandom = append(sliceRandom[0:i], sliceRandom[i+1:]...)
				i--
			}
		}
		for i := 0; i < len(sliceRandom); i++ {
			product := getProduct(db, sliceRandom[i])
			deleteFromShop(db, sliceRandom[i])
			inserInShopingList(db, product.product)
		}
	}
*/

func (u *user) makePurchase(db *sql.DB) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(10) + 1
	product := getProduct(db, r)
	insertInShoppingList(db, product.product)
	newCash := u.balance - product.price
	u.updateBalance(db, newCash)
	print("new balance : ", newCash)
}

func (u *user) works(db *sql.DB) {
	for {
		cash := u.getBalance(db)
		if cash <= 40 {
			print(u.name, " works ")
			Cash := cash + salary
			u.updateBalance(db, Cash)
			println(" amount of money after work :", Cash)
		}
		break
	}
}

func main() {
	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("log in func main")
	}
	defer db.Close()
	user := getUser(db)
	user.works(db)
	user.makePurchase(db)
}

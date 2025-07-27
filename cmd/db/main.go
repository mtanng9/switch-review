package main

import "github.com/mtanng9/switch-review/store"

func main() {
	db := store.InitDB()
	store.SeedDB(db)
}

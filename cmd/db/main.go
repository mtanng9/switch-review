package main

import (
	"github.com/mtanng9/switch-review/store"
	"github.com/mtanng9/switch-review/store/seed"
)

func main() {
	db := store.InitDB()
	seed.SeedDB(db)
}

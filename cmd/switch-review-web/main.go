package main

import (
	"github.com/mtanng9/switch-review/server"
	"github.com/mtanng9/switch-review/store"
)

func main() {
	db := store.InitDB()
	server.StartServer(db)
}

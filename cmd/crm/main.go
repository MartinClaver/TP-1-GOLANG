package main

import (
	"awesomeProject/internal/app"
	"awesomeProject/internal/storage"
)

func main() {

	var store storage.Storer = storage.NewMemoryStore()

	app.Run(store)

}

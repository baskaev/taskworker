package main

import (
	"github.com/baskaev/db/datab"
)

func main() {
	// Инициализация базы данных
	if err := datab.InitDB(); err != nil {
		panic(err)
	}
}

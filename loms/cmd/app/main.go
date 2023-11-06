package main

import (
	"log"
	"route256/loms/internal/app"
)

func main() {
	app := new(app.App)
	if err := app.Run(); err != nil {
		log.Printf("[ERROR] service exited with err: %v", err)
	}
}

package main

import (
	"log"

	"github.com/KRR19/number-search/cmd/app"
)

func main() {
	app := app.NewApplication()

	if err := app.ServeHTTP(); err != nil {
		log.Fatalf("failed to serve http: %v", err)
	}
}

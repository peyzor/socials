package main

import (
	"github.com/joho/godotenv"
	"github.com/peyzor/socials/internal/env"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	app := &application{config: cfg}

	log.Printf("server has started at %s", app.config.addr)

	mux := app.mount()
	err = app.run(mux)
	if err != nil {
		log.Fatal(err)
	}
}

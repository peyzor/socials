package main

import (
	"github.com/joho/godotenv"
	"github.com/peyzor/socials/internal/db"
	"github.com/peyzor/socials/internal/env"
	"github.com/peyzor/socials/internal/store"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:               env.GetString("DB_ADDR", "postgres://postgres:@localhost/socials"),
			maxOpenConnections: env.GetInt("DB_MAX_OPEN_CONNECTIONS", 30),
			maxIdleConnections: env.GetInt("DB_MAX_IDLE_CONNECTIONS", 30),
			maxIdleTime:        env.GetString("DB_MAX_IDLE_TIME", "15min"),
		},
	}

	database, err := db.New(cfg.db.addr, cfg.db.maxOpenConnections, cfg.db.maxIdleConnections, cfg.db.maxIdleTime)
	if err != nil {
		log.Fatal(err)
	}

	s := store.NewStorage(database)
	app := &application{
		config: cfg,
		store:  s,
	}

	log.Printf("server has started at %s", app.config.addr)

	mux := app.mount()
	err = app.run(mux)
	if err != nil {
		log.Fatal(err)
	}
}

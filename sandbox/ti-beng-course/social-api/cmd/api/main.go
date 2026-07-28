package main

import (
	"log"

	"github.com/Yuseonr/gin-labs/sandbox/ti-beng-course/social-api/internal/env"
	"github.com/Yuseonr/gin-labs/sandbox/ti-beng-course/social-api/internal/store"
	"github.com/Yuseonr/gin-labs/sandbox/ti-beng-course/social-api/internal/db"
)

func main (){
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			dsn: env.GetString("DB_DSN", "host=localhost port=5432 user=admin password=adminpassword dbname=social sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30), // higher number = higher concurent query (atleast what i understand)
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30), // number of open connection in connection pool
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(cfg.db.dsn, 
		cfg.db.maxOpenConns, 
		cfg.db.maxIdleConns, 
		cfg.db.maxIdleTime)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	log.Println("database conn established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store: store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
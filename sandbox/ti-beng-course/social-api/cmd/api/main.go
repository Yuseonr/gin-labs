package main

import (
	"log"

	"github.com/Yuseonr/gin-labs/sandbox/ti-beng-course/social-api/internal/env"
	"github.com/Yuseonr/gin-labs/sandbox/ti-beng-course/social-api/internal/store"
)

func main (){
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store: store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
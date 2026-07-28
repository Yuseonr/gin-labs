package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Yuseonr/gin-labs/sandbox/ti-beng-course/social-api/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store store.Storage
}

type config struct {
	addr string
	db dbConfig
}

type dbConfig struct {
	dsn string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime	string
}

func (app *application) mount() http.Handler {

	r := chi.NewRouter()

	// digunakan untuk generate request id untuk setiap request yang masuk ke server
	r.Use(middleware.RequestID)

	// ngelog setiap req yang masuk ke server
	// jadi -> log nya akan seperti 
	// 2026/07/26 19:31:48 "GET http://localhost:8080/v1/ HTTP/1.1" from [::1]:62206 - 200 19B in 157.709µs
	// digunakan untuk debugging misal error ketika hit endpoint tertentu, bisa dilihat dari log nya
	r.Use(middleware.Logger)

	// middleware.Recoverer digunakan untuk menangkap panic yang terjadi di handler
	// misal kita bikin handler yang sengaja panic, maka middleware ini akan menangkapnya dan mengembalikan 500 internal server error
	// jadi server tidak akan crash dan tetap berjalan
	r.Use(middleware.Recoverer)


	// timeout value on request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped
	r.Use(middleware.Timeout(time.Second * 60))
	
	// // inline routing
	// // nameless func
	// r.Get("/", func(w http.ResponseWriter, r *http.Request){
	// 	w.Write([]byte("Hello this from / !"))
	// })

	// // pass to handfunc
	// r.Get("/health", app.healthCheckHandler)

	// below using group or route in chi
	// curl http://localhost:8080/v1/
	// curl http://localhost:8080/v1/health
	r.Route("/v1", func(r chi.Router){
		r.Get("/", func(w http.ResponseWriter, r *http.Request){w.Write([]byte("Hello this from / !"))})
		r.Get("/health", app.healthCheckHandler)
		// assigning new handler doesnt need comma because 
		// Go separates statements by line breaks, not commas.
	})

	return  r

}

// Note : we can change entire *chi.Mux to http.Handler with the power of interface
// dimana si chi mux ini sendiri mengimplementasikan http.handler sehingga bisa diganti langsung ke http.handler
// and result -> we dont really need chimux and can provide our own router if want to
func (app *application) run(mux http.Handler) error{

	srv := http.Server{
		Addr: app.config.addr,
		Handler: mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute ,
	}

	log.Printf("server started at : %s", app.config.addr)

	return srv.ListenAndServe()
}
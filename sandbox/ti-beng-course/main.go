package main

import (
	"net/http"
	"log"
)

type api struct {
	addres string
}

// func (sv *api) ServeHTTP (w http.ResponseWriter, r *http.Request){
// 	w.Write([]byte("Hello Wolrld"))
// }

// sekarang Api (prev;server) bertindak bukan sebagai routing 
// atau harus implement ServeHTTP tapi penyedia methode untuk handler saja
func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("nih list of users"))
}

func (a *api) createUserHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("user dibuat ! :>"))
}

func main() {
	api := &api{addres: ":8080"}

	// init servemux
	// mux sudah implement ServeHTTP sehingga dia valid jadi http.handler
	// nah karena itu kita tidak perlu membuat api sebagai http.handler
	// tpi kita bisa borrow function / methode si api buat digunakan ServeMux 
	// untuk handle incoming request
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr: api.addres,
		Handler: mux,
		// dengan menggunakan http server bisa set berbagai config
	}

	// disini mux bertindak sebagai routing yang kita bisa register methode seta resource nya
	// untuk menuju ke handler yang ditentukan 

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUserHandler)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}

// to test
// curl -X POST http://localhost:8080/users	
// curl http://localhost:8080/users

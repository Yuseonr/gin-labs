package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type user struct {
	Name string `json:"name"` // fungsi json tag yaitu : auto marshal and marshaling buat encode decode
	Age int		`json:"age"`
}

type api struct {
	addres string
}

func beforeInsertUser(u user) error {
	// user validation sebelum dimasukan ke slice
	if (u.Name) == "" {
		return errors.New("Nama tidak boleh kosong")
	}
	if (u.Age) < 0 {
		return errors.New("Umur tidak bisa negatif")
	}
	if (u.Age) > 100 {
		return errors.New("Yang benar saja")
	}

	// tidak boleh ada user yang sama
	for _, user := range users{
		if u.Name == user.Name && u.Age == user.Age {
			return  errors.New("User sudah ada !")
		}
	}
	return nil 
}

var users = []user{}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request){
	// menuliskan header ke rsponse writter
	w.Header().Set("Content-Type", "application/json")

	/* 
	content-type : application/json artinya ngasi tau ke client
	"body dari response ini adalah json, jadi tolong di parse"
	tanpa ini, client bisa nggap response nya itu plain text ato malah html dan
	bakal bikin error / bug jika frontend nyoba buat make response.json() parsing
	*/

	// encode users slice to json
	// json.NewEncoder(w) akan otomatis encode ke json dan write ke response writer
	// encode itu apa ? jadi encode itu proses convert data ke format lain, misal dari struct ke json ato dari json ke struct
	// dalam hal ini kita convert dari struct user ke json, jadi nanti response nya bisa di parse di frontend ato postman
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (a *api) createUserHandler(w http.ResponseWriter, r *http.Request){
	// decode request body to user struct
	// payload user ini adalah data yang dikirim dari client, misal dari postman ato frontend
	var payload user

	// decorder nya akan take request body dan decode ke struct user
	// decode disini akan otomatis mapping field dari json ke struct user, misal json name -> struct name
	// makanya perlu payload sebagai struct user, biar bisa di mapping ke struct user
	// jika ada error, misal json nya invalid ato field nya ga sesuai, maka akan return error
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// hasil dari decode payload akan di append ke slice users, jadi nanti bisa di get semua user nya
	u := user{
		Name: payload.Name,
		Age: payload.Age,
	}

	// disini bisa kita masukan validation atau fungsi lain untuk memvalidasi input
	if err := beforeInsertUser(u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users = append(users, u)
	w.WriteHeader(http.StatusCreated)

}

func main() {
	api := &api{addres: ":8080"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr: api.addres,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUserHandler)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
// To Test : 
/*
curl -X POST -v 'http://localhost:8080/users' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{"name" : "nandra", "age" : 21}'
*/

// curl http://localhost:8080/users
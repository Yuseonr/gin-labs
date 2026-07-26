#  Complete Backend Engineering Course in Go
```
source : https://youtu.be/h3fqD6IprIA?si=mM7ftkpqU8OYoLeA
```
---

## The `net/http` Package

> **`net/http`** adalah standard library bawaan Go untuk membangun HTTP client dan HTTP server secara native tanpa framework

---

```Go
package main

import (
	"log"
	"net/http"
)

type server struct {
	addres string
}

// test: curl http://localhost:8080
// server akan membalas "Hello world" ke semua request
func (sv *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func main() {
	s := &server{addres: ":8080"}
	if err := http.ListenAndServe(s.addres, s); err != nil {
		log.Fatal(err)
	}
}
```

---

### Flow Request → Response

```
curl http://localhost:8080/apa-aja
        │
        ▼
  http.ListenAndServe(":8080", server)
        │
        ▼
  server.ServeHTTP(w, r)
        │
        ▼
  w.Write([]byte("Hello world"))
        │
        ▼
  client terima: "Hello world"
```
---
### Claude - Gimana Cara Server Kirim Data ke Client?

**1. HTTP itu pada dasarnya Cuma "Ngirim Surat"**

Client (browser / curl) dan server itu kayak dua orang yang cuma bisa komunikasi lewat surat. 

**2. Client Kirim Surat (HTTP Request)**

```
┌─────────────────────────────────────────┐
│  Dari: browser (192.168.1.5)            │
│  Ke:   server (localhost:8080)          │
│                                          │
│  GET /hello HTTP/1.1                     │
│  Host: localhost:8080                    │
│                                          │
│  (body kosong)                           │
└─────────────────────────────────────────┘
```

Client nulis surat. Isinya: "Aku mau minta halaman `/hello`". Surat ini dikirim lewat **TCP connection** — ibarat tukang pos yang nganterin surat lewat jalur yang udah dibuka dari rumah client ke rumah server.

**3. Server Terima Surat — `http.ListenAndServe`**

```
   surat masuk ──► http.ListenAndServe
                         │
                         │  "Ada surat nih, baca..."
                         │
                         ▼
                  pecah jadi 2 benda
                  ┌────────┐  ┌──────────┐
                  │   w    │  │    r     │
                  │ Writer │  │ Request  │
                  └────────┘  └──────────┘
```

`http.ListenAndServe` ibarat **penjaga pintu** rumah server. Dia:
- Nunggu di depan pintu (port 8080) 24/7
- Begitu ada surat datang, dia buka amplopnya
- Amplopnya dipecah jadi dua:
  - **r (Request)** — isi surat dari client (dari siapa, minta apa, lewat mana)
  - **w (ResponseWriter)** — kertas kosong buat balesan

Terus dia manggil kamu: "Heh, ini ada surat. `r` isi suratnya, `w` kertas kosong buat bales."

**4. Kamu (ServeHTTP) Nulis Balasan**

```go
func (sv *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello world"))
}
```

Kamu sebagai `ServeHTTP` itu ibarat **orang yang ditugasin bales surat**. Kamu dikasih:
- **r**: surat dari client (boleh dibaca, boleh diabaikan)
- **w**: kertas kosong

Lalu kamu nulis di kertas kosong itu:

```
┌─────────────────────────────────────────┐
│  Dari: server (localhost:8080)          │
│  Ke:   client                           │
│                                          │
│  HTTP/1.1 200 OK                         │
│  Content-Length: 11                      │
│                                          │
│  Hello world                             │
└─────────────────────────────────────────┘
```

`w.Write` itu **kamu megang pulpen, nulis langsung di kertas balasan**. Ga pake perantara.

**5. Balik ke Client**

```
   w.Write selesai ──► http.ListenAndServe
                              │
                              │  "Balasan udah siap, kirim!"
                              │
                              ▼
                         Tukang pos (TCP)
                         anter balik surat
                              │
                              ▼
                         Client terima:
                         "Hello world"
```

Begitu kamu selesai nulis di `w`, server langsung nyuruh TCP anter balik surat ke client. Client buka amplop, baca: "Hello world".

**Flow Lengkap**

```
Browser (client)                              Server Go
    │                                              │
    │  "GET /" ──────── TCP Connection ──────────► │
    │                                              │
    │                                    ListenAndServe
    │                                    terima surat
    │                                    pecah jadi r & w
    │                                          │
    │                                    ServeHTTP(w, r)
    │                                          │
    │                                   w.Write("Hello world")
    │                                          │
    │  "Hello world" ◄──── TCP Connection ──────┘
    │                                              │
    ▼                                              ▼
  tampil di browser                        balik nunggu request
```

<!-- claude generated end -->

---
### Native Way to Handle HTTP Request

```Go
package main

import (
	"log"
	"net/http"
	"fmt"
)

type server struct {
	addres string
}

// pada contoh dibawah kita gunakan arsitektur simple
// mapping routing methode then resource
// namun pada best practice nya better resource then methode karena 
// resource itu benda utama. method cuma operasi atas benda itu. jadi natural-nya: tentuin dulu benda mana, baru tentuin operasinya.
func (sv *server) ServeHTTP (w http.ResponseWriter, r *http.Request){
	switch r.Method {
	// bisa "GET" tapi hardcoded string is not good, use http.MethodGet instead
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			// write merupakan native way to send messg ke client
			// dimana emang langsung dikirim byte nya seperti itu dan client menerima apa adanya
			// why not fmt?, bisa pakai fmt tapi tetap fmt yang nerima Response writer untuk nulis reponse nya
			w.Write([]byte("index ni"))

			// set return biar ga bablas, bablas disini misal kita ada case lanjutan setelah keluar dari swtich
			return

		case "/users":
			// contoh pake fmt
			fmt.Fprint(w, "users data ni")
			return

		default:
			// tidak ada Get dengan resource ini
			w.Write([]byte("404 resource not found"))
			return
		}

	default:
		// tidak ada methode ini
		w.Write([]byte("404 methode not found"))
		return
	}
}

func main() {
	s := &server{addres: ":8080"}
	if err := http.ListenAndServe(s.addres, s); err != nil {
		log.Fatal(err)
	}
}

// to test
// curl http://localhost:8080/ 				-> index
// curl http://localhost:8080/users 		-> users
// curl http://localhost:8080/random		-> 404 resource not found
// curl -X POST http://localhost:8080/users -> 404 methode not found
```
---
### One Abstraction using `http.HandlerFunc` and `http.ServeMux`
>ServeMux (Service Multiplexer) adalah HTTP router bawaan Go. Fungsinya: nentuin handler mana yang jalan berdasarkan path dan method dari request yang masuk. Singkatnya: request multiplexer.

```
Manual (satu handler, switch-case):

                     ┌──────────────┐
  semua request ──►  │  ServeHTTP    │
  GET /              │              │
  GET /users         │  switch path │
  POST /users        │  switch meth │
  GET /random        │  semua disini│
                     └──────────────┘
                           │
                    satu function nangani semuanya

ServeMux (router + banyak handler):

                     ┌──────────────┐
  GET / ──────────►  │              │──► indexHandler
  GET /users ──────► │  ServeMux    │──► usersHandler
  POST /users ─────► │  (router)    │──► usersHandler (jalan juga kalo POST)
  GET /random ─────► │              │──► 404 (default, otomatis)
                     └──────────────┘
                           │
                    mux milihin handler
```
---
```Go
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
```

---

## Encoding and Decoding JSON
> fungsinya untuk mengubah data dari struct ke JSON (encoding) dan dari JSON ke struct (decoding)

```Go
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
```

---

## Setting up development
```
Go 1.22
Docker
Postgres on Docker
Swagger for docs
Golang migrate for migrations
```

### Folder Structure
> membagi busniess logic dan routing handler agar lebih clean dan maintainable
```
├── bin
├── cmd
│   ├── api
│   └── migrate
│       └── migrations
├── docs
├── internal
└── scripts
```

1. **Separation of concern** : Tiap level pada program harus di pisahkan dengan barrier yang clear, transport layer, service layer, storage layer
2. **Dependency Inversion Principle (DIP)** : Kamu menginject dependency ke layer mu, bukan langsung memanggilnya agar membuat loose coupling dan memudahkan testing
3. **Adaptability to change** : dengan mengorganisasi kode kita secara modular kita bisa dengan mudah menambah fitur, refactor kode, atau merespon terhadap perubahan rbusiness requirment dengan lebih mudah. Sistem harus dirancang supaya dapat dengan mudah dirubah tanpa mengganggu bagian lain dari kode
4. **Focus on Busineess Value** : focus on adding or delivering value to user as they are the one responsible for paying your bills. 

### Layers
> Pada struktur ini kita akan membagi ke tiga layer yaitu Transport, Service dan Storage, namun kita bisa nge ommit service karena hanya simple nanti
1. Transport layer : bagaimana cara kita deliver message to user, misal sekarang HTTP
2. Service Layer : business logic, misal kita mau bikin user baru, maka service layer yang akan handle logic nya
3. Storage Layer : bagaimana cara kita menyimpan data, misal kita mau simpan user, abstraksi antara service layer dengan misal database

> Injection 
```
Repository -> Service -> Transport
```

![layer](layer.png)

---

## Setting up HTTP Server
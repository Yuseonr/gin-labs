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
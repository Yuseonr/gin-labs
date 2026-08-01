# Complete Backend Engineering Course in Go 2
```
source : udemy
```
---
> Part 1    : [Complete Backend Engineering Course in Go](../youtube/TI-Beng-Course-1.md)

> Repo      : [PROJECT REPOSITORY](https://github.com/Yuseonr/social-go)

## Posts CRUD
### Marshalling
```GO
package main

import (
	"encoding/json"
	"net/http"
)

// jadi fungsi dibawah digunakan untuk menulis response json ke client,
// agar tidak perlu nulis ulang-ulang
func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}


// fungsi ini akan membaca request body, lalu mendecode json nya ke struct yang kita tentukan
// gimana decode nya ini bekerja dan mengapa tidak return struct tapi error ?
// jadi fungsi ini menerima parameter data any, yang artinya bisa struct apapun
// lalu fungsi ini akan mendecode json dari request body ke struct yang kita tentukan
// misal kita punya struct Post, maka kita bisa panggil fungsi ini dengan data &post
// jadi fungsi ini akan mengisi field-field struct post dengan data dari json request body
// jika terjadi error saat decode, maka fungsi ini akan mengembalikan error
func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	// max req size = 1mb
	maxBytes := 1_048_578
	r.Body = http.MaxBytesReader(w,r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

// when building api response should be consistant in a way that consumer expect everything to be the same format or structure
// not limited to only success response but also the error one
// basically cuma model error nya harus konsisten
func writeJSONError (w http.ResponseWriter, status int, message string) error {
	type envelope struct {
		Error string `json:"error"`
	}
	return writeJSON(w, status, &envelope{Error: message})
}
```
so the health handler becomes :
```GO
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	// mengapa map[string]string?
	// jadi map string string itu map key nya string, value nya string
	// nah ini nge mimic data json key value pair yang key dan value nya string
	// sehingga ketika di encode ke json, akan menjadi seperti ini
	// {
	// 	"status": "ok",
	// 	"env": "developement"
	//	"version": "0.0.1"
	// }
	data := map[string]string{
		"status": "ok",
		"env":    "development",
		"version": "0.0.1",
	}

	if err := writeJSON(w, http.StatusOK, data); err != nil {
		// error
		// log.Print(err.Error())
		writeJSONError(w, http.StatusInternalServerError, err.Error())
	}
}
```


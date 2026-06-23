# Gin Gonic Web Framework - 1 [Bahasa Indonesia]
```
source : https://youtu.be/Mx5fdAlZSwI?si=nAtBBG9d3Eowpipu
```
---

## PENDAHULUAN
#### *inisialisasi GO MOD*
```
go mod init <github.com/Yuseonr/gin-labs> atau bisa <gin-labs>
```
#### *instalasi gin-gonic*
```
go get -u github.com/gin-gonic-gin
```

## Membuat simple endpoint
```Go
package main

import "github.com/gin-gonic/gin"

func main (){
	app := gin.Default()
	route := app
	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message" : "Lets Go OK!"})

        // Dibawah sini biasanya diberi return
        // return

        // return biasa digunakan untuk membuat if statement dapat memberhentikan kode,
        // agar misal suatu kondisi tertentu tercapai, kirim X dan jangan lanjutkan dibawahnya.

	})

    app.Run(":8080")
}
```

```
go run main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.main.func1 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://github.com/gin-gonic/gin/blob/master/docs/doc.md#dont-trust-all-proxies for details.
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2026/06/23 - 08:01:55 | 200 | 171.291µs |             ::1 | GET      "/"
[GIN] 2026/06/23 - 08:01:56 | 404 |    791ns |             ::1 | GET      "/favicon.ico"
```
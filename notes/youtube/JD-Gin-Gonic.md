# Tutorial 2023! 
>Penamaan file dan folder bukan patokan

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

# Gin Gonic Web Framework - 2 [Bahasa Indonesia]
```
source : https://youtu.be/0rC9zfbRt78?si=VUNikrRy4_FGTVB9
```
---

## Bootstrap App
*Melakukan Abstraksi dan enkapsulasi serta pemisahan tanggung jawab, supaya app dapat lebih termaintain*
```
.
├── bootstrap
│   └── index.go
├── config
│   └── app_config
│       └── index.app_config.go
├── controller
│   ├── book_controller
│   │   └── index.book_controller.go
│   └── user_controller
│       └── index.user_controller.go
├── main.go
└── routes
    └── index.route.go
```
```
main        -> manggil bootstrap.
bootstrap   -> appnya yang dari init gin, manggil initroute, dan run di port yang diambil dari config.
app_config  -> nyimpan config app agar di satu tempat, kaya port
routes      -> inisialisasi semua ruting yang dibutuhkan, kaya GET /user atau GET/book terus dikasi function controller
controller  -> menyediakan function sesuai data nya, misal user_controller dan punya func GetAllUser() handle memberikan user
```

```Go
// sandbox/jd-gin-gonic/main.go
package main

import b "github.com/Yuseonr/gin-labs/sandbox/jd-gin-gonic/bootstrap"

func main (){
	b.BootstrapApp()
}
```


```Go
// sandbox/jd-gin-gonic/bootstrap/index.go
package bootstrap

import (
	"github.com/Yuseonr/gin-labs/sandbox/jd-gin-gonic/config/app_config"
	"github.com/Yuseonr/gin-labs/sandbox/jd-gin-gonic/routes"
	"github.com/gin-gonic/gin"
)

func BootstrapApp (){
	app := gin.Default()
	routes.InitRoute(app)
	app.Run(app_config.Port)
}
```

```Go
// sandbox/jd-gin-gonic/controller/user_controller/index.user_controller.go

package user_controller

import "github.com/gin-gonic/gin"

func GetALLUser(ctx *gin.Context) {

	isValidate := true

	if isValidate {
		ctx.JSON(200, gin.H{
			"message": "Lets Go OK!",
			"user": "{user1, user2, user3}",
		})

		return

	} else {
		ctx.JSON(400, gin.H{
			"message": "Not validated!",
		})
	}
}

```

```Go
// sandbox/jd-gin-gonic/routes/index.route.go

package routes

import (
	"github.com/Yuseonr/gin-labs/sandbox/jd-gin-gonic/controller/book_controller"
	"github.com/Yuseonr/gin-labs/sandbox/jd-gin-gonic/controller/user_controller"
	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app
	route.GET("/user",user_controller.GetALLUser)
	route.GET("/book", book_controller.GetAllBook)
}	

```

## POSTMAN
> Postman dapat digunakan untuk analisa endpoint

![postman](postman.png)

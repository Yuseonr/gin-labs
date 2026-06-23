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
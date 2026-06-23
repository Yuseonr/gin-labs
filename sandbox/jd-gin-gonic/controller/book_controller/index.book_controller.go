package book_controller

import "github.com/gin-gonic/gin"

func GetAllBook(ctx *gin.Context){

	isValidate := true

	if isValidate {
		ctx.JSON(200, gin.H{
			"message": "Lets Go OK!",
			"book": "{book1, book2, book3}",
		})

		return

	} else {
		ctx.JSON(400, gin.H{
			"message": "Not validated!",
		})
	}
}
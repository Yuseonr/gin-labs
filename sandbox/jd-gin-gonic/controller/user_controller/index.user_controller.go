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

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

package bootstrap

import (
	"log"

	"github.com/Yuseonr/gin-labs/sandbox/jd-gin-gonic/config/app_config"
	"github.com/Yuseonr/gin-labs/sandbox/jd-gin-gonic/database"
	"github.com/Yuseonr/gin-labs/sandbox/jd-gin-gonic/routes"
	"github.com/gin-gonic/gin"
)

func BootstrapApp (){
	err := database.ConnectDatabase(); if err != nil {
		log.Fatal(err)
	}
	app := gin.Default()
	routes.InitRoute(app)
	app.Run(app_config.Port)
}
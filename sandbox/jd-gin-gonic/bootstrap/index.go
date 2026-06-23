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
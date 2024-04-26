package routes

import (
	"blog_server/global"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRoutes() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	routerGroup := router.Group("api")
	routerGroupApp := RouterGroup{routerGroup}

	// 系统配置API
	routerGroupApp.SettingsRoutes()

	return router
}
/*
* @Time    : 2020-11-17 11:35
* @Author  : CoderCharm
* @File    : router.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    : 初始化路由
**/
package initialize

import (
	_ "gin_study/docs"
	"gin_study/global"
	"gin_study/middleware"
	"gin_study/router"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func Routers() *gin.Engine {
	var Router = gin.Default()

	// 跨域
	Router.Use(middleware.CORSMiddleware())
	if global.GIN_CONFIG.System.Env == "debug" {
		Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// 方便统一添加路由组前缀 多服务器上线使用
	V1ApiGroup := Router.Group("/api/v1")

	router.InitUserRouter(V1ApiGroup)    // 注册用户路由
	router.InitArticleRouter(V1ApiGroup) // 注册文章路由

	return Router
}

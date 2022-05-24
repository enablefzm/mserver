package ginserver

import (
	"log"

	"github.com/enablefzm/mserver/ginserver/middlewares"
	"github.com/enablefzm/mserver/routers"

	"github.com/gin-gonic/gin"
)

func Run() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// 初始化Gin的中间件
	authMiddleware, err := middlewares.InitMiddlewares(r)
	if err != nil {
		log.Println("加载授权中间件出现错误:", err.Error())
		return
	}
	// 设定登入路由
	r.POST("/auth/login", authMiddleware.LoginHandler)
	// 指定路由组
	authApi := r.Group("/api")
	// authApi路由组使用授权认证的中间件
	authApi.Use(authMiddleware.MiddlewareFunc())
	// 使用Casbin权限认证中间件
	// authApi.Use(middlewares.CasbinMiddleWare)
	// 设定刷Token路由
	authApi.GET("/refresh_token", authMiddleware.RefreshHandler)
	// 加载业务路由
	routers.Init(authApi)
	// 运行GinServer并侦听指定的端口
	r.Run(":8080")
}

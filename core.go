package mserver

import (
	"log"

	"github.com/enablefzm/mserver/casbin"
	"github.com/enablefzm/mserver/dbs"
	"github.com/enablefzm/mserver/ginserver/middlewares"
	"github.com/enablefzm/mserver/routers"
	_ "github.com/enablefzm/mserver/routers/users"

	"github.com/gin-gonic/gin"
)

// Web服务器引擎
var ObServer *gin.Engine

// 标准需要验证用户认证路由组
var AuthApiGroup *gin.RouterGroup

func init() {
	ObServer = gin.Default()
	// 初始化Gin的用户认证中间件
	authMiddleware, err := middlewares.InitMiddlewares(ObServer)
	if err != nil {
		log.Println("加载授权中间件出错:", err.Error())
		panic("加载授权中间件出错!")
	}
	// 设定登入路由处理Handler
	ObServer.POST("/auth/login", authMiddleware.LoginHandler)
	// 指定路由组
	AuthApiGroup = ObServer.Group("/api")
	// 使用用户认证中间件
	AuthApiGroup.Use(authMiddleware.MiddlewareFunc())
	// 刷新当前的Token
	AuthApiGroup.GET("/refresh_token", authMiddleware.RefreshHandler)

	// 初始化路由
	routers.Init(AuthApiGroup)
}

// 运行服务端
//	isUseCasbin 是否需要启用权限认证
func Run(cfg *ServerCfg, dbCfg *dbs.Cfg) {
	// 加载数据库
	log.Print("连接 ", dbCfg.IpAddress, ":", dbCfg.Port, " (", dbCfg.DbName, ")数据库...")
	err := dbs.LinkDb(dbCfg)
	if err != nil {
		log.Println("出错:", err.Error())
		panic("连接数据库服务器出错:")
	} else {
		log.Println("连接成功!")
	}
	if cfg.IsUseCasbin {
		casbin.LinkDbAndLoadRabc(dbCfg)
		// 初始化Casbin权限认证机制
		AuthApiGroup.Use(middlewares.CasbinMiddleWare)
	}

	log.Println("运行MServer Web服务端...")
	ObServer.Run(":" + cfg.Port)
}

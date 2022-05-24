package middlewares

import (
	"github.com/enablefzm/mserver/auth"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 初始化Gin的中间件
func InitMiddlewares(r *gin.Engine) (*jwt.GinJWTMiddleware, error) {
	// 使用解决跨域的中间件
	r.Use(NewCorsMiddleware())
	// 构造Auth JWT中间件
	authMiddleware, err := auth.NewAuthMiddleWare()
	if err != nil {
		// 初始JWT 认证服务出错
		return authMiddleware, err
	}
	// 初始化Auth JWT认证服务
	err = authMiddleware.MiddlewareInit()
	return authMiddleware, err
}

package customers

import (
	"github.com/enablefzm/mserver/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	routers.Include(func(r *gin.RouterGroup) {
		// 注册客户信息
		r.GET("/customers/customer/info", pRouterCustomer.InfoHandler)
	})
}

var pRouterCustomer = &RouterCustomer{}

type RouterCustomer struct{}

func (p *RouterCustomer) InfoHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"custmoerId":   10001,
		"custmoerName": "吉米游戏",
	})
}

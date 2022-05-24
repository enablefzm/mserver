package routers

import (
	"net/http"

	"github.com/enablefzm/mserver/dtos/commdtos"

	"github.com/gin-gonic/gin"
)

type Option func(engine *gin.RouterGroup)

var options = make([]Option, 0, 50)

// 加载需要注册到路由的钩子函数
func Includes(opts ...Option) {
	options = append(options, opts...)
}

func Include(opt Option) {
	options = append(options, opt)
}

// 将路由注册到服务器的路由表里
func Init(r *gin.RouterGroup) {
	for _, opt := range options {
		opt(r)
	}
}

// 将http request的数据绑定指定的数据上
func ShouldBindJSON(c *gin.Context, obj interface{}) error {
	err := c.ShouldBindJSON(obj)
	if err != nil {
		c.JSON(http.StatusBadRequest, commdtos.NewCommError(err.Error()))
	}
	return err
}

func NewCommInputGet(c *gin.Context) (commdtos.CommInputGet, error) {
	var commget commdtos.CommInputGet
	err := ShouldBindJSON(c, &commget)
	return commget, err
}

func CheckErrorResult(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusOK, commdtos.NewCommError(err.Error()))
	} else {
		c.JSON(http.StatusOK, commdtos.NewCommResult("操作成功"))
	}
}

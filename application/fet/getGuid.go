package fet

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/enablefzm/mserver/dtos/commdtos"

	"github.com/gin-gonic/gin"
)

// 获取指定的API带的Guid参数
func GetGuid(c *gin.Context) (string, error) {
	guid := strings.Trim(c.Param("guid"), " ")
	if len(guid) < 10 {
		c.JSON(http.StatusBadRequest, commdtos.NewCommError("请求数据错误Guid错误"))
		return guid, fmt.Errorf("请求数据Guid错误")
	}
	return guid, nil
}

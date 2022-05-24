package middlewares

import (
	"fmt"
	"net/http"

	"github.com/enablefzm/mserver/auth"
	"github.com/enablefzm/mserver/casbin"
	"github.com/enablefzm/mserver/dtos/commdtos"

	"github.com/gin-gonic/gin"
)

// 权限认证中间件
func CasbinMiddleWare(c *gin.Context) {
	ob, ok := c.Keys["User"]
	if !ok {
		c.JSON(http.StatusMethodNotAllowed, commdtos.NewCommError("当前用户权限不足，用户不存在"))
		c.Abort()
		return
	}
	obUser, ok := ob.(auth.User)
	if !ok {
		c.JSON(http.StatusMethodNotAllowed, commdtos.NewCommError("用户无法被转换有效用户"))
		c.Abort()
		return
	}
	// fmt.Println("权限认证中间件:", obUser.Guid, obUser.Uid, obUser.Name)
	// 如果用户是admin则不受权限限制的影响
	if obUser.Uid != "admin" {
		p := c.Request.URL.Path
		m := c.Request.Method
		// 获取用户角色
		arr, err := casbin.Enforcer.GetRolesForUser(obUser.Guid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, commdtos.NewCommError(err.Error()))
			c.Abort()
			return
		}
		blnCheck := false
		for _, role := range arr {
			res, err := casbin.Enforcer.EnforceSafe(role, p, m)
			if err == nil && res == true {
				blnCheck = true
				fmt.Println("角色验证通过!")
				break
			}
		}
		if !blnCheck {
			c.JSON(http.StatusUnauthorized, commdtos.NewCommError("当前用户权限不足"))
			c.Abort()
			return
		}
		// 判断用户权限
		//res, err := casbin.Enforcer.EnforceSafe(obUser.Guid, p, m)
		//if err != nil {
		//	fmt.Println("当前用户没有权限 Err:", err.Error())
		//	c.JSON(http.StatusUnauthorized, commdtos.NewCommError("当前用户权限不足"))
		//	c.Abort()
		//	return
		//}
		//if !res {
		//	fmt.Println("用户权限验证失败:", res)
		//	c.JSON(http.StatusUnauthorized, commdtos.NewCommError("当前用户权限不足"))
		//	c.Abort()
		//	return
		//}
	}
	c.Next()
}

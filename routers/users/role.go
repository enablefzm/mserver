package users

import (
	"fmt"
	"net/http"

	"github.com/enablefzm/mserver/application/userapps"
	"github.com/enablefzm/mserver/dtos/commdtos"
	"github.com/enablefzm/mserver/dtos/userdtos"
	"github.com/enablefzm/mserver/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	routers.Include(func(r *gin.RouterGroup) {
		r.PUT("/users/role/list", pRouterRole.ListHandler)
		r.POST("/users/role/create", pRouterRole.CreateHandler)
		r.GET("/users/role/listpermissioninfos", pRouterRole.ListPermissionInfosHandler)
	})
}

var pRouterRole = &RouterRole{}

type RouterRole struct {
}

func (r *RouterRole) ListHandler(c *gin.Context) {
	var commGet commdtos.CommInputGet
	if err := routers.ShouldBindJSON(c, &commGet); err != nil {
		return
	}
	result, _ := userapps.AppRole.List(commGet)
	c.JSON(http.StatusOK, result)
}

func (r *RouterRole) CreateHandler(c *gin.Context) {
	var createRoleDto userdtos.CreateUpdateRoleDto
	if err := routers.ShouldBindJSON(c, &createRoleDto); err != nil {
		return
	}
	if err := userapps.AppRole.Create(createRoleDto); err != nil {
		c.JSON(http.StatusOK, commdtos.NewCommError(err.Error()))
	} else {
		c.JSON(http.StatusOK, commdtos.NewCommResult("成功添加角色"))
	}
}

func (r *RouterRole) ListPermissionInfosHandler(c *gin.Context) {
	fmt.Println(userapps.AppRole.ListPermissionInfos())
	c.JSON(http.StatusOK, userapps.AppRole.ListPermissionInfos())
}

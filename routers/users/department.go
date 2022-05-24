package users

import (
	"net/http"

	"github.com/enablefzm/mserver/application/userapps"
	"github.com/enablefzm/mserver/dtos/commdtos"
	"github.com/enablefzm/mserver/dtos/userdtos"
	"github.com/enablefzm/mserver/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	routers.Include(func(r *gin.RouterGroup) {
		// 获取部门列表信息
		r.PUT("/users/department/list", pRouterDepartment.ListInfosHandler)
		// 添加部门
		r.POST("/users/department/create", pRouterDepartment.Create)
	})
}

var pRouterDepartment = &RouterDepartment{}

type RouterDepartment struct {
}

// 列出部门信息
func (r *RouterDepartment) ListInfosHandler(c *gin.Context) {
	var commGet commdtos.CommInputGet
	if err := routers.ShouldBindJSON(c, &commGet); err != nil {
		return
	}
	result, err := userapps.AppDepartment.ListInfo(commGet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, commdtos.NewCommError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result)
}

// 创建一个新部门
func (r *RouterDepartment) Create(c *gin.Context) {
	var createDto userdtos.CreateUpdateDepartmentDto
	if err := routers.ShouldBindJSON(c, &createDto); err != nil {
		return
	}
	// 执行添加用户对象
	err := userapps.AppDepartment.Create(createDto)
	if err != nil {
		c.JSON(http.StatusOK, commdtos.NewCommError(err.Error()))
	} else {
		c.JSON(http.StatusOK, commdtos.NewCommResult("添加部门成功"))
	}
}

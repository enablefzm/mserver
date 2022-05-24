package users

import (
	"net/http"

	"github.com/enablefzm/mserver/application/fet"
	"github.com/enablefzm/mserver/application/userapps"
	"github.com/enablefzm/mserver/dtos/commdtos"
	"github.com/enablefzm/mserver/dtos/userdtos"
	"github.com/enablefzm/mserver/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	routers.Include(func(r *gin.RouterGroup) {
		r.PUT("/users/company/list", pRouterCompany.ListHandler)
		r.POST("/users/company/create", pRouterCompany.CreateHandler)
		r.PUT("/users/company/update/:guid", pRouterCompany.UpdateHandler)
		r.DELETE("/users/company/delete/:guid", pRouterCompany.DeleteHandler)
	})
}

var pRouterCompany = &RouterCompany{}

type RouterCompany struct {
}

func (r *RouterCompany) ListHandler(c *gin.Context) {
	var commget commdtos.CommInputGet
	if err := routers.ShouldBindJSON(c, &commget); err != nil {
		return
	}
	result, err := userapps.AppCompany.List(commget)
	if err != nil {
		c.JSON(http.StatusOK, commdtos.NewCommError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result)
}

func (r *RouterCompany) CreateHandler(c *gin.Context) {
	var createUpdateDtos userdtos.CreateUpdateCompanyDto
	if err := routers.ShouldBindJSON(c, &createUpdateDtos); err != nil {
		return
	}
	err := userapps.AppCompany.Create(createUpdateDtos)
	if err != nil {
		c.JSON(http.StatusOK, commdtos.NewCommError(err.Error()))
	} else {
		c.JSON(http.StatusOK, commdtos.NewCommResult("添加公司成功"))
	}
}

func (r *RouterCompany) DeleteHandler(c *gin.Context) {
	guid, err := fet.GetGuid(c)
	if err != nil {
		return
	}
	err = userapps.AppCompany.Delete(guid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, commdtos.NewCommError(err.Error()))
	} else {
		c.JSON(http.StatusOK, "删除成功")
	}
}

func (r *RouterCompany) UpdateHandler(c *gin.Context) {
	guid, err := fet.GetGuid(c)
	if err != nil {
		return
	}
	var updateDtos userdtos.CreateUpdateCompanyDto
	if err = routers.ShouldBindJSON(c, &updateDtos); err != nil {
		return
	}
	err = userapps.AppCompany.Update(guid, updateDtos)
	if err != nil {
		c.JSON(http.StatusOK, commdtos.NewCommError(err.Error()))
	} else {
		c.JSON(http.StatusOK, commdtos.NewCommResult("修改公司信息成功"))
	}
}

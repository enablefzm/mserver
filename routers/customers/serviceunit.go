package customers

import (
	"net/http"

	"github.com/enablefzm/mserver/application/customerapps"
	"github.com/enablefzm/mserver/application/fet"
	"github.com/enablefzm/mserver/dtos/commdtos"
	customerdtos "github.com/enablefzm/mserver/dtos/customerDtos"
	"github.com/enablefzm/mserver/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	routers.Include(func(r *gin.RouterGroup) {
		r.PUT("/customer/serviceunit/list", pRouterServiceUnit.ListHandler)
		r.POST("/customer/serviceunit/create", pRouterServiceUnit.CreateHandler)
		r.PUT("/customer/serviceunit/update/:guid", pRouterServiceUnit.UpdateHandler)
	})

}

var pRouterServiceUnit = &RouterServiceUnit{}

type RouterServiceUnit struct{}

func (r *RouterServiceUnit) ListHandler(c *gin.Context) {
	commget, err := routers.NewCommInputGet(c)
	if err != nil {
		return
	}
	result, err := customerapps.AppServiceUnit.List(commget)
	if err != nil {
		c.JSON(http.StatusInternalServerError, commdtos.NewCommError(err.Error()))
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func (r *RouterServiceUnit) CreateHandler(c *gin.Context) {
	var createUpdateDto customerdtos.CreateUpdateCustomerDto
	if err := routers.ShouldBindJSON(c, &createUpdateDto); err != nil {
		return
	}
	err := customerapps.AppServiceUnit.Create(createUpdateDto)
	routers.CheckErrorResult(c, err)
}

func (r *RouterServiceUnit) UpdateHandler(c *gin.Context) {
	guid, err := fet.GetGuid(c)
	if err != nil {
		return
	}
	var updateDto customerdtos.CreateUpdateCustomerDto
	if err = routers.ShouldBindJSON(c, &updateDto); err != nil {
		return
	}
	err = customerapps.AppServiceUnit.Update(guid, updateDto)
	routers.CheckErrorResult(c, err)
}

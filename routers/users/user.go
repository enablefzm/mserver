package users

import (
	"net/http"

	"github.com/enablefzm/mserver/application/fet"
	"github.com/enablefzm/mserver/application/userapps"
	"github.com/enablefzm/mserver/auth"
	"github.com/enablefzm/mserver/dtos/commdtos"
	"github.com/enablefzm/mserver/dtos/userdtos"
	"github.com/enablefzm/mserver/routers"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func init() {
	routers.Include(func(r *gin.RouterGroup) {
		// 注册获取用户信息
		r.GET("/users/user/info", pRouterUser.InfoHandler)
		r.PUT("/users/user/list", pRouterUser.ListInfosHandler)
		r.POST("/users/user/create", pRouterUser.CreateHandler)
		r.PUT("/users/user/update/:guid", pRouterUser.EditHandler)
	})
}

var pRouterUser = &RouterUser{}

type RouterUser struct {
}

// 查看用户信息
func (r *RouterUser) InfoHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	// user, _ := c.Get(auth.IdentityKey)
	c.JSON(200, gin.H{
		"guid":     claims[auth.IdentityKey],
		"userName": claims["name"],
		"text":     "Hello World.",
	})
}

// 列表查看用户信息
func (r *RouterUser) ListInfosHandler(c *gin.Context) {
	var commGet commdtos.CommInputGet
	if err := routers.ShouldBindJSON(c, &commGet); err != nil {
		return
	}
	// 查询数据
	result, err := userapps.AppUser.ListInfo(commGet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, commdtos.NewCommError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result)
}

// 创建新用户
func (r *RouterUser) CreateHandler(c *gin.Context) {
	var createUserDto userdtos.CreateUpdateUserDto
	if err := routers.ShouldBindJSON(c, &createUserDto); err != nil {
		return
	}
	// 执行添加对象
	err := userapps.AppUser.Create(createUserDto)
	if err != nil {
		c.JSON(http.StatusOK, commdtos.NewCommError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, commdtos.NewCommResult("成功添加用户"))
}

// 更新数据
func (r *RouterUser) EditHandler(c *gin.Context) {
	guid, err := fet.GetGuid(c)
	if err != nil {
		return
	}
	var updateUserDto userdtos.CreateUpdateUserDto
	if err = routers.ShouldBindJSON(c, &updateUserDto); err != nil {
		return
	}
	err = userapps.AppUser.Edit(guid, updateUserDto)
	if err != nil {
		c.JSON(http.StatusOK, commdtos.NewCommError(err.Error()))
	} else {
		c.JSON(http.StatusOK, commdtos.NewCommResult("修改用户信息成功"))
	}
}

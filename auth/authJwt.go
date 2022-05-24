package auth

import (
	"fmt"
	"time"

	"github.com/enablefzm/mserver/application/userapps"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var IdentityKey = "guid"

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type User struct {
	Guid string
	Uid  string
	Name string
}

func NewAuthMiddleWare() (*jwt.GinJWTMiddleware, error) {
	authMiddleWare, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "Hitek_Zone",
		Key:         []byte("FZM_hitek_20220405"),
		Timeout:     time.Hour * 48,
		MaxRefresh:  time.Hour * 48,
		IdentityKey: IdentityKey,

		// 打包Auth PayloadFunc 信息函数
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			// fmt.Println("Auth PayloadFunc ... 打包Payload信息的函数")
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					IdentityKey: v.Guid,
					"uid":       v.Uid,
					"name":      v.Name,
				}
			}
			return jwt.MapClaims{}
		},

		// 检验Auth IdentityHandler 里的Token和Payload信息
		IdentityHandler: func(c *gin.Context) interface{} {
			// fmt.Println("Auth IdentityHandler ExtractClaims ... 验证通过，分拆Token里的Payload信息")
			claims := jwt.ExtractClaims(c)
			pUser := &User{
				Guid: claims[IdentityKey].(string),
				Uid:  claims["uid"].(string),
				Name: claims["name"].(string),
			}
			c.Keys["User"] = *pUser
			return pUser
		},

		// 检验Auth Authenticator用户名和密码逻辑
		Authenticator: func(c *gin.Context) (interface{}, error) {
			// fmt.Println("Auth Authenticator ...执行验证用户和密码的逻辑")
			var loginVales Login
			if err := c.ShouldBind(&loginVales); err != nil {
				fmt.Println("ShouldBind:", err.Error())
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVales.Username
			password := loginVales.Password
			// 执行逻辑判断
			obUser, err := userapps.AppUser.Login(userID, password)
			if err != nil {
				return nil, err
			}
			// 检验成功
			return &User{
				Guid: obUser.Guid,
				Uid:  obUser.UID,
				Name: obUser.Name,
			}, nil
		},

		// 检验失败返回的消息格式
		Unauthorized: func(c *gin.Context, code int, message string) {
			// fmt.Println("Auth Unauthorized ... 验证失败")
			c.JSON(code, gin.H{
				"code":              code,
				"error_description": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	return authMiddleWare, err
}

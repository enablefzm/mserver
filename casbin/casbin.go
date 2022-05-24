package casbin

import (
	"fmt"

	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
	"github.com/enablefzm/mserver/dbs"
	"github.com/jinzhu/gorm"
)

var O *gorm.DB
var PO *gormadapter.Adapter
var Enforcer *casbin.Enforcer

func LinkDbAndLoadRabc(c *dbs.Cfg) {
	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		c.User,
		c.Password,
		c.IpAddress,
		c.Port,
		c.DbName)
	var err error
	O, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Casbin link Db error:", err.Error())
		panic(err)
	}
	// 将数据库连接对象给代理
	PO = gormadapter.NewAdapterByDB(O)
	// 初始化权限认证对象
	Enforcer = casbin.NewEnforcer("./configs/rabc_model.conf", PO)
	// 开启权限认证日志
	Enforcer.EnableLog(true)
	// 加载数据库中的策略
	err = Enforcer.LoadPolicy()
	if err != nil {
		// 加载策略失败
		fmt.Println("加载数据库策略失败:", err.Error())
		panic(err)
	}
}

package mserver

import "github.com/enablefzm/gotools/vaini"

func NewServerCfg() *ServerCfg {
	cfg := &ServerCfg{
		IsUseCasbin: true,
		Port:        "8080",
	}
	C := vaini.NewConfig("./cfg.ini")
	if mp, ok := C.GetNode("webserver"); ok {
		for k, v := range mp {
			switch k {
			case "port":
				cfg.Port = v
			}
		}
	}
	return cfg
}

type ServerCfg struct {
	IsUseCasbin bool   // 是否启用权限
	Port        string // 监听端口
}

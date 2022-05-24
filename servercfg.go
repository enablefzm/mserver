package mserver

func NewServerCfg() *ServerCfg {
	return &ServerCfg{
		IsUseCasbin: true,
		Port:        "8080",
	}
}

type ServerCfg struct {
	IsUseCasbin bool   // 是否启用权限
	Port        string // 监听端口
}

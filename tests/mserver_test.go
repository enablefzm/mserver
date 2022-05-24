package tests

import (
	"testing"

	"github.com/enablefzm/mserver"
	"github.com/enablefzm/mserver/dbs"
)

func TestMServer(t *testing.T) {

	serverCfg := &mserver.ServerCfg{
		Port: "3316",
	}
	dbCfg := &dbs.Cfg{
		User:      "root",
		Password:  "fzmzsp520",
		IpAddress: "localhost",
		Port:      "3316",
		DbName:    "uk_invoice",
	}
	t.Log(serverCfg.Port, dbCfg.DbName)
}

package tests

import (
	"testing"

	"github.com/enablefzm/mserver/casbin"
)

func TestPolicy(t *testing.T) {
	// ok := casbin.Enforcer.AddPolicy("系统管理员", "/api/users/user/list", "*")
	ok := casbin.Enforcer.AddPermissionForUser("系统管理员", "/api/users/user/info", "*")
	// ok := casbin.Enforcer.AddRoleForUser("bb1ff81b-631c-47d7-9290-2f308a9c9727", "系统管理员")
	t.Log("result:", ok)
}

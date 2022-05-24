package tests

import (
	"testing"

	"github.com/enablefzm/mserver/application/userapps"
)

func TestUserInfo(t *testing.T) {
	obUserInfoDto, _ := userapps.AppUser.GetOnUid("dd753b31-231d-4a23-8fde-49fa8d9a2729")
	t.Log(obUserInfoDto)
}

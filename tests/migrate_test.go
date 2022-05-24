package tests

import (
	"testing"

	"github.com/enablefzm/mserver/application"
)

func TestMigrate(t *testing.T) {
	err := application.Migration()
	t.Log(err)
}

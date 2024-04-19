package test

import (
	"github.com/joho/godotenv"
	"levi-api/ops/user"
	"testing"
)

func TestOps(t *testing.T) {
	_ = godotenv.Load()
	var a user.AccountOps
	if err := a.AddAccount("firstofzen", "test", "test"); err != nil {
		t.Fatal(err)
	} else {
		t.Log("added")
	}
}

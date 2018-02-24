package supersaas

import (
	"testing"

	"../supersaas"
)

func TestUsersList(t *testing.T) {
	client := supersaas.NewClient("accnt", "pwd")
	client.DryRun = true
	_, err := client.Users.List(true, 10, 5)
	if err != nil {
		t.Error(err)
	}
}

func TestUsersGet(t *testing.T) {
	client := supersaas.NewClient("accnt", "pwd")
	client.DryRun = true
	_, err := client.Users.Get(12345)
	if err != nil {
		t.Error(err)
	}
}

func TestUsersGetFk(t *testing.T) {
	client := supersaas.NewClient("accnt", "pwd")
	client.DryRun = true
	_, err := client.Users.Get("12345fk")
	if err != nil {
		t.Error(err)
	}
}

func TestUsersCreate(t *testing.T) {
	client := supersaas.NewClient("accnt", "pwd")
	client.DryRun = true
	a := map[string]interface{}{}
	_, err := client.Users.Create(a, "", false)
	if err != nil {
		t.Error(err)
	}
}

func TestUsersCreateFk(t *testing.T) {
	client := supersaas.NewClient("accnt", "pwd")
	client.DryRun = true
	a := map[string]interface{}{}
	_, err := client.Users.Create(a, "12345fk", false)
	if err != nil {
		t.Error(err)
	}
}

func TestUsersUpdate(t *testing.T) {
	client := supersaas.NewClient("accnt", "pwd")
	client.DryRun = true
	a := map[string]interface{}{}
	_, err := client.Users.Update(12345, a, false)
	if err != nil {
		t.Error(err)
	}
}

func TestUsersDelete(t *testing.T) {
	client := supersaas.NewClient("accnt", "pwd")
	client.DryRun = true
	err := client.Users.Delete(12345)
	if err != nil {
		t.Error(err)
	}
}

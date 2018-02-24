package supersaas

import (
	"testing"

	"../supersaas"
)

func TestApi(t *testing.T) {
	client := supersaas.NewClient("accnt", "pwd")
	if client.Appointments == nil {
		t.Errorf("Appointments is nil")
	}
	if client.Forms == nil {
		t.Errorf("Forms is nil")
	}
	if client.Schedules == nil {
		t.Errorf("Schedules is nil")
	}
	if client.Users == nil {
		t.Errorf("Users is nil")
	}
}

func TestInstanceConfiguration(t *testing.T) {
	supersaas.GetInstance().Configure("accnt", "pwd", "", true, true)
	client := supersaas.GetInstance()
	if client.AccountName != "accnt" {
		t.Error("For", client, "expected", "accnt", "got", client.AccountName)
	}
	if client.Password != "pwd" {
		t.Error("For", client, "expected", "pwd", "got", client.Password)
	}
	if client.Host != "http://localhost:3000" {
		t.Error("For", client, "expected", "http://localhost:3000", "got", client.Host)
	}
	if client.DryRun != true {
		t.Error("For", client, "expected", true, "got", client.DryRun)
	}
	if client.Verbose != true {
		t.Error("For", client, "expected", true, "got", client.Verbose)
	}
}

func TestRequestMethods(t *testing.T) {
	client := supersaas.GetInstance()
	client.DryRun = true
	q := map[string]interface{}{}
	p := map[string]interface{}{}
	out := supersaas.Model{}

	err := client.Get("/test", q, out)
	if err != nil {
		t.Error(err)
	}
	err = client.Post("/test", q, p, out)
	if err != nil {
		t.Error(err)
	}
	err = client.Put("/test", q, p, out)
	if err != nil {
		t.Error(err)
	}
	err = client.Delete("/test", q, p)
	if err != nil {
		t.Error(err)
	}
}

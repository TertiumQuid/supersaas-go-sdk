package supersaas

import (
	"testing"
	"time"

	"../supersaas"
)

func TestFormsList(t *testing.T) {
	client := supersaas.NewClient("accnt", "pwd")
	client.DryRun = true
	_, err := client.Forms.List(12345, time.Now())
	if err != nil {
		t.Error(err)
	}
}

func TestFormsGet(t *testing.T) {
	client := supersaas.NewClient("accnt", "pwd")
	client.DryRun = true
	_, err := client.Forms.Get(12345)
	if err != nil {
		t.Error(err)
	}
}

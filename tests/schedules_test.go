package supersaas

import (
	"testing"

	"../supersaas"
)

func TestSchedules(t *testing.T) {
	client := supersaas.NewClient("accnt", "pwd")
	client.DryRun = true
	_, err := client.Schedules.List()
	if err != nil {
		t.Error(err)
	}
}

func TestResources(t *testing.T) {
	client := supersaas.NewClient("accnt", "pwd")
	client.DryRun = true
	_, err := client.Schedules.Resources(12345)
	if err != nil {
		t.Error(err)
	}
}

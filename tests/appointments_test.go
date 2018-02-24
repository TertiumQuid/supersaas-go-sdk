package supersaas

import (
	"testing"
	"time"

	"../supersaas"
)

func TestAppointmentsGet(t *testing.T) {
	client := supersaas.NewClient("accnt", "pwd")
	client.DryRun = true
	_, err := client.Appointments.Get(12345, 67890)
	if err != nil {
		t.Error(err)
	}
}

func TestAppointmentsList(t *testing.T) {
	client := supersaas.NewClient("accnt", "pwd")
	client.DryRun = true
	_, err := client.Appointments.List(12345, true, time.Now(), 10)
	if err != nil {
		t.Error(err)
	}
}

func TestAppointmentsCreate(t *testing.T) {
	client := supersaas.NewClient("accnt", "pwd")
	client.DryRun = true
	a := map[string]interface{}{}
	_, err := client.Appointments.Create(12345, 0, a, true, false)
	if err != nil {
		t.Error(err)
	}
}

func TestAppointmentsUpdate(t *testing.T) {
	client := supersaas.NewClient("accnt", "pwd")
	client.DryRun = true
	a := map[string]interface{}{}
	_, err := client.Appointments.Update(12345, 67890, a, true, true)
	if err != nil {
		t.Error(err)
	}
}

func TestAppointmentsDelete(t *testing.T) {
	client := supersaas.NewClient("accnt", "pwd")
	client.DryRun = true
	err := client.Appointments.Delete(12345)
	if err != nil {
		t.Error(err)
	}
}

func TestAgenda(t *testing.T) {
	client := supersaas.NewClient("accnt", "pwd")
	client.DryRun = true
	_, err := client.Appointments.Agenda(12345, 67890, time.Now())
	if err != nil {
		t.Error(err)
	}
}

func TestAgendaSlots(t *testing.T) {
	client := supersaas.NewClient("accnt", "pwd")
	client.DryRun = true
	_, err := client.Appointments.AgendaSlots(12345, 67890, time.Now())
	if err != nil {
		t.Error(err)
	}
}

func TestAvailable(t *testing.T) {
	client := supersaas.NewClient("accnt", "pwd")
	client.DryRun = true
	_, err := client.Appointments.Available(12345, time.Now(), 60, "", true, 10)
	if err != nil {
		t.Error(err)
	}
}

func TestChanges(t *testing.T) {
	client := supersaas.NewClient("accnt", "pwd")
	client.DryRun = true
	_, err := client.Appointments.Changes(12345, time.Now())
	if err != nil {
		t.Error(err)
	}
}

func TestChangesSlots(t *testing.T) {
	client := supersaas.NewClient("accnt", "pwd")
	client.DryRun = true
	_, err := client.Appointments.ChangesSlots(12345, time.Now())
	if err != nil {
		t.Error(err)
	}
}

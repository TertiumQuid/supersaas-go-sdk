package supersaas

import (
	"fmt"
	"time"
)

// Appointments ...
type Appointments struct {
	API
}

// Slot model ...
type Slot struct {
	Count       int       `json:"count"`
	Description string    `json:"description"`
	Finish      time.Time `json:"finish"`
	ID          int64     `json:"id"`
	Name        int64     `json:"name"`
	Start       time.Time `json:"start"`
	Title       time.Time `json:"title"`

	Bookings []Appointment `json:"bookings"`
	Errors   []interface{} `json:"errors"`
}

// Appointment model ...
type Appointment struct {
	Address      string    `json:"address"`
	Country      string    `json:"country"`
	CreatedBy    string    `json:"created_by"`
	CreatedOn    time.Time `json:"created_on"`
	Deleted      bool      `json:"deleted"`
	Description  string    `json:"description"`
	Email        string    `json:"email"`
	Field1       string    `json:"field_1"`
	Field2       string    `json:"field_2"`
	Field1R      string    `json:"field_1_r"`
	Field2R      string    `json:"field_2_r"`
	Finish       time.Time `json:"finish"`
	FormID       int64     `json:"form_id"`
	FullName     string    `json:"full_name"`
	ID           int64     `json:"id"`
	Mobile       string    `json:"mobile"`
	Name         string    `json:"name"`
	Phone        string    `json:"phone"`
	Price        int       `json:"price"`
	ResName      string    `json:"res_name"`
	ResourceID   int64     `json:"resource_id"`
	ScheduleID   int64     `json:"schedule_id"`
	ScheduleName string    `json:"schedule_name"`
	ServiceID    int64     `json:"service_id"`
	ServiceName  string    `json:"service_name"`
	SlotID       int64     `json:"slot_id"`
	Start        time.Time `json:"start"`
	Status       string    `json:"status"`
	SuperField   string    `json:"super_field"`
	UpdatedBy    string    `json:"updated_by"`
	UpdatedOn    time.Time `json:"updated_on"`
	UserID       int64     `json:"user_id"`
	Waitlisted   bool      `json:"waitlisted"`

	Form   *Form         `json:"form"`
	Slot   *Slot         `json:"slot"`
	Errors []interface{} `json:"errors"`
}

// Agenda returns a list of all appointments for a schedule and user
func (a Appointments) Agenda(scheduleID int64, userID int64, fromTime time.Time) ([]Appointment, error) {
	var appointments []Appointment
	path := "/agenda/" + fmt.Sprint(scheduleID)
	q := map[string]interface{}{"user_id": fmt.Sprint(userID)}
	if !fromTime.IsZero() {
		q["from"] = a.FormatTime(fromTime)
	}
	err := a.client.Get(path, q, &appointments)
	return appointments, err
}

// AgendaSlots returns a list of all slots for a schedule and user
func (a Appointments) AgendaSlots(scheduleID int64, userID int64, fromTime time.Time) ([]Slot, error) {
	var slots []Slot
	path := "/agenda/" + fmt.Sprint(scheduleID)
	q := map[string]interface{}{"user_id": fmt.Sprint(userID), "slot": "true"}
	if !fromTime.IsZero() {
		q["from"] = a.FormatTime(fromTime)
	}
	err := a.client.Get(path, q, &slots)
	return slots, err
}

// Available returns a list of appointments for a schedule
func (a Appointments) Available(scheduleID int64, fromTime time.Time, lengthMinutes int, resource string, full bool, limit int) ([]Appointment, error) {
	var appointments []Appointment
	path := "/available/" + fmt.Sprint(scheduleID)
	q := map[string]interface{}{"from": a.FormatTime(fromTime)}
	if lengthMinutes > 0 {
		q["length"] = fmt.Sprint(lengthMinutes)
	}
	if resource != "" {
		q["resource"] = resource
	}
	if full {
		q["full"] = "true"
	}
	if limit > 0 {
		q["maxresults"] = fmt.Sprint(limit)
	}
	err := a.client.Get(path, q, &appointments)
	return appointments, err
}

// List returns a list of all appointments for a schedule
func (a Appointments) List(scheduleID int64, form bool, startTime time.Time, limit int) ([]Appointment, error) {
	var appointments []Appointment
	path := "/bookings" + fmt.Sprint(scheduleID)
	q := map[string]interface{}{"schedule_id": fmt.Sprint(scheduleID)}
	if !startTime.IsZero() {
		q["start"] = a.FormatTime(startTime)
	}
	if form {
		q["form"] = "true"
	}
	if limit > 0 {
		q["limit"] = fmt.Sprint(limit)
	}
	err := a.client.Get(path, q, &appointments)
	return appointments, err
}

// Get returns a schedule appointment by id
func (a Appointments) Get(scheduleID int64, appointmentID int64) (Appointment, error) {
	appointment := Appointment{}
	path := "/bookings/" + fmt.Sprint(appointmentID)
	q := map[string]interface{}{"schedule_id": fmt.Sprint(scheduleID)}
	err := a.client.Get(path, q, &appointment)
	return appointment, err
}

// Create makes a new appointment record
func (a Appointments) Create(scheduleID int64, userID int64, attributes map[string]interface{}, form bool, webhook bool) (Appointment, error) {
	appointment := Appointment{}
	path := "/bookings"
	b := map[string]interface{}{
		"start":       attributes["start"],
		"finish":      attributes["finish"],
		"name":        attributes["name"],
		"email":       attributes["email"],
		"full_name":   attributes["full_name"],
		"address":     attributes["address"],
		"mobile":      attributes["mobile"],
		"phone":       attributes["phone"],
		"country":     attributes["country"],
		"field_1":     attributes["field_1"],
		"field_2":     attributes["field_2"],
		"field_1_r":   attributes["field_1_r"],
		"field_2_r":   attributes["field_2_r"],
		"super_field": attributes["super_field"],
		"resource_id": attributes["resource_id"],
		"slot_id":     attributes["slot_id"]}
	p := map[string]interface{}{
		"schedule_id": fmt.Sprint(scheduleID),
		"user_id":     fmt.Sprint(userID),
		"booking":     b}
	if webhook {
		p["webhook"] = "true"
	}
	if form {
		p["form"] = "true"
	}
	err := a.client.Post(path, p, nil, &appointment)
	return appointment, err
}

// Update modifies an appointment record
func (a Appointments) Update(scheduleID int64, appointmentID int64, attributes map[string]interface{}, form bool, webhook bool) (Appointment, error) {
	return Appointment{}, nil
}

// Delete removes an appointment record
func (a Appointments) Delete(appointmentID int64) error {
	path := "/bookings/" + fmt.Sprint(appointmentID)
	return a.client.Delete(path, nil, nil)
}

// Changes returns a list of changed appointments for a schedule
func (a Appointments) Changes(scheduleID int64, fromTime time.Time) ([]Appointment, error) {
	var appointments []Appointment
	path := "/changes/" + fmt.Sprint(scheduleID)
	q := map[string]interface{}{"from": a.FormatTime(fromTime)}
	err := a.client.Get(path, q, &appointments)
	return appointments, err
}

// ChangesSlots returns a list of changed slots for a schedule
func (a Appointments) ChangesSlots(scheduleID int64, fromTime time.Time) ([]Slot, error) {
	var slots []Slot
	path := "/changes/" + fmt.Sprint(scheduleID)
	q := map[string]interface{}{"from": a.FormatTime(fromTime), "slot": "true"}
	err := a.client.Get(path, q, &slots)
	return slots, err
}

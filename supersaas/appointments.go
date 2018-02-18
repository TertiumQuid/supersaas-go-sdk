package supersaas

import "time"

// Appointments ...
type Appointments API

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

	Form *Form `json:"form"`
	Slot *Slot `json:"slot"`
}

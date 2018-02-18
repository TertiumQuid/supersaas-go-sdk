package supersaas

import "time"

// Users ...
type Users API

// User model ...
type User struct {
	Address    string    `json:"address"`
	Country    string    `json:"country"`
	CreatedOn  time.Time `json:"created_on"`
	Credit     int       `json:"credit"`
	Email      string    `json:"email"`
	Field1     string    `json:"field_1"`
	Field2     string    `json:"field_2"`
	Fk         string    `json:"fk"`
	FullName   string    `json:"full_name"`
	ID         int64     `json:"id"`
	Mobile     string    `json:"mobile"`
	Name       string    `json:"name"`
	Phone      string    `json:"phone"`
	Role       int       `json:"role"`
	SuperField string    `json:"super_field"`

	Form *Form `json:"form"`
}

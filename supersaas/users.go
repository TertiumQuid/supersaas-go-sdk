package supersaas

import (
	"fmt"
	"time"
)

// Users ...
type Users struct {
	API
}

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

	Form   *Form         `json:"form"`
	Errors []interface{} `json:"errors"`
}

// List returns a list of all users
func (u Users) List(form bool, limit int, offset int) ([]User, error) {
	var users []User
	path := "/users"
	q := map[string]interface{}{}
	if limit > 0 {
		q["limit"] = fmt.Sprint(limit)
	}
	if offset > 0 {
		q["offset"] = fmt.Sprint(offset)
	}
	err := u.client.Get(path, q, &users)
	return users, err
}

// Get returns a user by id
func (u Users) Get(userID int) (User, error) {
	user := User{}
	path := "/users/" + fmt.Sprint(userID)
	err := u.client.Get(path, nil, &user)
	return user, err
}

// Create registers a new a user record
func (u Users) Create(attributes map[string]interface{}, userID string, webhook bool) (User, error) {
	user := User{}
	path := "/users/"
	if userID != "" {
		path += fmt.Sprint(userID)
	}
	q := map[string]interface{}{}
	if webhook {
		q["webhook"] = "true"
	}
	p := map[string]interface{}{"user": attributes}
	err := u.client.Post(path, q, p, &user)
	return user, err
}

// Update modifies a user record
func (u Users) Update(userID int64, attributes map[string]interface{}, webhook bool) (User, error) {
	user := User{}
	path := "/users/" + fmt.Sprint(userID)
	q := map[string]interface{}{}
	if webhook {
		q["webhook"] = "true"
	}
	p := map[string]interface{}{"user": attributes}
	err := u.client.Put(path, q, p, &user)
	return user, err
}

// Delete removes a user record
func (u Users) Delete(userID string) error {
	path := "/users/" + fmt.Sprint(userID)
	return u.client.Delete(path, nil, nil)
}

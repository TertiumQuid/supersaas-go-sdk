package supersaas

import (
	"fmt"
	"time"
)

// Forms ...
type Forms struct {
	API
}

// Form model ...
type Form struct {
	Model
	Content              []interface{} `json:"content"`
	CreatedOn            string        `json:"created_on"`
	Deleted              bool          `json:"deleted"`
	ID                   int64         `json:"id"`
	ReservationProcessID int64         `json:"reservation_process_id"`
	SuperFormID          int64         `json:"super_form_id"`
	Uniq                 string        `json:"uniq"`
	UpdatedName          string        `json:"updated_name"`
	UpdatedOn            string        `json:"updated_on"`
	UserID               int64         `json:"user_id"`
}

// List returns forms for a given form template
func (f Forms) List(formID int, fromTime time.Time) ([]Form, error) {
	var forms []Form
	path := "/forms"
	q := map[string]interface{}{"form_id": fmt.Sprint(formID)}
	if !fromTime.IsZero() {
		q["from"] = f.FormatTime(fromTime)
	}
	err := f.Client.Get(path, q, &forms)
	return forms, err
}

// Get returns a form by id
func (f Forms) Get(formID int) (Form, error) {
	form := Form{}
	path := "/forms/" + fmt.Sprint(formID)
	q := map[string]interface{}{"id": fmt.Sprint(formID)}
	err := f.Client.Get(path, q, &form)
	return form, err
}

package supersaas

// Forms ...
type Forms API

// Form model ...
type Form struct {
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

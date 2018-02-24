package supersaas

import (
	"fmt"
	"time"
)

// API ...
type API struct {
	Client *Client
}

// Model ...
type Model struct {
	Errors []interface{} `json:"errors"`
}

// FormatTime returns YYYY-MM-DD HH:MM:SS datetime string
func (a API) FormatTime(t time.Time) string {
	return fmt.Sprint("%d-%02d-%02d %02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}

package supersaas

import (
	"sync"
)

const (
	defaultHost = "http://localhost:3000"
)

// Configuration ...
type Configuration struct {
	AccountName string
	Password    string
	Host        string
	DryRun      bool
	Verbose     bool
}

// Client ...
type Client struct {
	Appointments *Appointments
	Forms        *Forms
	Schedules    *Schedules
	Users        *Users
}

var instance *Client
var once sync.Once

// GetInstance returns thread-safe singleton client instance ...
func GetInstance() *Client {
	once.Do(func() {
		instance = &Client{}
	})
	return instance
}

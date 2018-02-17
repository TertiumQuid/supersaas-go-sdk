package supersaas

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
	Users        *Users
}

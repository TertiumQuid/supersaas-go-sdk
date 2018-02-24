package supersaas

import "fmt"

// Schedules ...
type Schedules struct {
	API
}

// Schedule model ...
type Schedule struct {
	ID   int64 `json:"id"`
	Name int64 `json:"name"`
}

// Resource model ...
type Resource struct {
	ID   int64 `json:"id"`
	Name int64 `json:"name"`
}

// List returns account schedules
func (s Schedules) List() ([]Schedule, error) {
	var schedules []Schedule
	path := "/schedules"
	err := s.Client.Get(path, nil, &schedules)
	return schedules, err
}

// Resources returns the services/resources for a given schedule
func (s Schedules) Resources(scheduleID int) ([]Resource, error) {
	var resources []Resource
	q := map[string]interface{}{"schedule_id": fmt.Sprint(scheduleID)}
	path := "/resources"
	err := s.Client.Get(path, q, &resources)
	return resources, err
}

package models

// Schedule represents the main schedule structure
type Schedule []Course

// Course represents a single course in the schedule
type Course struct {
	Name     string    `json:"nombre"`
	Schedule []Session `json:"horarios"`
}

// Session represents a single class session
type Session struct {
	Room     string `json:"aula"`
	Day      string `json:"dia"`
	TimeSlot string `json:"horario"`
}

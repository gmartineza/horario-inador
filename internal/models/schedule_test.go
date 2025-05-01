package models

import (
	"encoding/json"
	"os"
	"testing"
)

func TestScheduleParsing(t *testing.T) {
	// Test JSON that mirrors our actual data structure
	testJSON := `[
		{
			"nombre": "Desarrollo Multimedia",
			"horarios": [
				{
					"aula": "SALA DE INFORMATICA B",
					"dia": "Jueves",
					"horario": "18:00 a 19:30"
				},
				{
					"aula": "SALA DE INFORMATICA B",
					"dia": "Lunes",
					"horario": "18:00 a 21:40"
				}
			]
		}
	]`

	var schedule Schedule
	err := json.Unmarshal([]byte(testJSON), &schedule)
	if err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	// Verify the structure
	if len(schedule) != 1 {
		t.Errorf("Expected 1 course, got %d", len(schedule))
	}

	course := schedule[0]
	if course.Name != "Desarrollo Multimedia" {
		t.Errorf("Expected course name 'Desarrollo Multimedia', got '%s'", course.Name)
	}

	if len(course.Schedule) != 2 {
		t.Errorf("Expected 2 sessions, got %d", len(course.Schedule))
	}

	// Test first session
	session := course.Schedule[0]
	expectedSession := Session{
		Room:     "SALA DE INFORMATICA B",
		Day:      "Jueves",
		TimeSlot: "18:00 a 19:30",
	}

	if session != expectedSession {
		t.Errorf("Session mismatch.\nExpected: %+v\nGot: %+v", expectedSession, session)
	}
}

func TestRealDataParsing(t *testing.T) {
	// Read the actual input.json file
	data, err := os.ReadFile("../../input.json")
	if err != nil {
		t.Fatalf("Failed to read input.json: %v", err)
	}

	var schedule Schedule
	err = json.Unmarshal(data, &schedule)
	if err != nil {
		t.Fatalf("Failed to parse real JSON data: %v", err)
	}

	// Verify we have all courses (5 based on the input.json you shared)
	expectedCourses := 5
	if len(schedule) != expectedCourses {
		t.Errorf("Expected %d courses, got %d", expectedCourses, len(schedule))
	}

	// Test specific known values from your data
	expectedCourseNames := []string{
		"Desarrollo Multimedia",
		"Ingeniería de Software III",
		"Modelo y Simulación",
		"Tecnología, Ciencia y Responsabilidad Social",
		"Teoría de la Computación",
	}

	for i, expectedName := range expectedCourseNames {
		if i >= len(schedule) {
			t.Errorf("Missing course at index %d, expected '%s'", i, expectedName)
			continue
		}
		if schedule[i].Name != expectedName {
			t.Errorf("Course at index %d: expected '%s', got '%s'", i, expectedName, schedule[i].Name)
		}
	}

	// Test that each course has the correct number of sessions
	expectedSessions := map[string]int{
		"Desarrollo Multimedia":                        2,
		"Ingeniería de Software III":                   2,
		"Modelo y Simulación":                          2,
		"Tecnología, Ciencia y Responsabilidad Social": 2,
		"Teoría de la Computación":                     2,
	}

	for _, course := range schedule {
		expected := expectedSessions[course.Name]
		if len(course.Schedule) != expected {
			t.Errorf("Course '%s': expected %d sessions, got %d",
				course.Name, expected, len(course.Schedule))
		}
	}
}

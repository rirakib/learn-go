package main

import (
	"errors"
	"testing"
)

func newTestService() StudentService {
	return NewStudentService(NewInMemoryStudentRepository())
}

func TestStudentCRUD(t *testing.T) {
	service := newTestService()

	created, err := service.AddStudent("Rahim", "rahim@example.com", 20)
	if err != nil {
		t.Fatalf("AddStudent returned error: %v", err)
	}
	if created.ID != 1 {
		t.Fatalf("expected student ID 1, got %d", created.ID)
	}

	found, err := service.GetStudent(created.ID)
	if err != nil {
		t.Fatalf("GetStudent returned error: %v", err)
	}
	if found.Name != "Rahim" {
		t.Fatalf("expected name Rahim, got %s", found.Name)
	}

	updated, err := service.UpdateStudent(created.ID, "Rahim Uddin", "rahim.uddin@example.com", 21)
	if err != nil {
		t.Fatalf("UpdateStudent returned error: %v", err)
	}
	if updated.Name != "Rahim Uddin" || updated.Age != 21 {
		t.Fatalf("student was not updated correctly: %+v", updated)
	}

	students := service.ListStudents()
	if len(students) != 1 {
		t.Fatalf("expected 1 student, got %d", len(students))
	}

	err = service.DeleteStudent(created.ID)
	if err != nil {
		t.Fatalf("DeleteStudent returned error: %v", err)
	}

	_, err = service.GetStudent(created.ID)
	if !errors.Is(err, ErrStudentNotFound) {
		t.Fatalf("expected ErrStudentNotFound after delete, got %v", err)
	}
}

func TestStudentValidation(t *testing.T) {
	service := newTestService()

	_, err := service.AddStudent("", "student@example.com", 20)
	if !errors.Is(err, ErrInvalidStudent) {
		t.Fatalf("expected ErrInvalidStudent for empty name, got %v", err)
	}

	_, err = service.AddStudent("Student", "", 20)
	if !errors.Is(err, ErrInvalidStudent) {
		t.Fatalf("expected ErrInvalidStudent for empty email, got %v", err)
	}

	_, err = service.AddStudent("Student", "student@example.com", 0)
	if !errors.Is(err, ErrInvalidStudent) {
		t.Fatalf("expected ErrInvalidStudent for invalid age, got %v", err)
	}
}

func TestUpdateAndDeleteMissingStudent(t *testing.T) {
	service := newTestService()

	_, err := service.UpdateStudent(99, "Missing", "missing@example.com", 18)
	if !errors.Is(err, ErrStudentNotFound) {
		t.Fatalf("expected ErrStudentNotFound for update, got %v", err)
	}

	err = service.DeleteStudent(99)
	if !errors.Is(err, ErrStudentNotFound) {
		t.Fatalf("expected ErrStudentNotFound for delete, got %v", err)
	}
}

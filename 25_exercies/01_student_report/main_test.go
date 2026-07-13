package main

import "testing"

func TestAddStudentAndFindByID(t *testing.T) {
	students = nil

	addStudent(NewStudent(StudentInput{ID: "1", Name: "Alice", Department: "CSE", Age: 20, Marks: 85}))

	got := findStudentByID("1")
	if got == nil {
		t.Fatal("expected student to be found")
	}
	if got.Name != "Alice" {
		t.Fatalf("expected Alice, got %s", got.Name)
	}
}

func TestUpdateStudent(t *testing.T) {
	students = nil
	addStudent(NewStudent(StudentInput{ID: "2", Name: "Bob", Department: "EEE", Age: 21, Marks: 75}))

	updated := updateStudent("2", StudentInput{ID: "2", Name: "Bob Updated", Department: "CSE", Age: 22, Marks: 90})
	if !updated {
		t.Fatal("expected update to succeed")
	}

	got := findStudentByID("2")
	if got == nil {
		t.Fatal("expected student to exist after update")
	}
	if got.Name != "Bob Updated" || got.Department != "CSE" {
		t.Fatalf("unexpected updated student values: %+v", got)
	}
}

func TestDeleteStudent(t *testing.T) {
	students = nil
	addStudent(NewStudent(StudentInput{ID: "3", Name: "Carol", Department: "BBA", Age: 23, Marks: 80}))

	deleted := deleteStudent("3")
	if !deleted {
		t.Fatal("expected delete to succeed")
	}

	if got := findStudentByID("3"); got != nil {
		t.Fatalf("expected student to be removed, got %+v", got)
	}
}

func TestFilterStudents(t *testing.T) {
	students = nil
	addStudent(NewStudent(StudentInput{ID: "4", Name: "Rahim", Department: "CSE", Age: 24, Marks: 78}))
	addStudent(NewStudent(StudentInput{ID: "5", Name: "Rafsan", Department: "EEE", Age: 25, Marks: 82}))
	addStudent(NewStudent(StudentInput{ID: "6", Name: "Nadia", Department: "CSE", Age: 22, Marks: 88}))

	byName := filterStudentsByName("rah")
	if len(byName) != 2 {
		t.Fatalf("expected 2 matches by name, got %d", len(byName))
	}

	byDepartment := filterStudentsByDepartment("CSE")
	if len(byDepartment) != 2 {
		t.Fatalf("expected 2 matches by department, got %d", len(byDepartment))
	}
}

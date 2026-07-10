package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

var (
	ErrStudentNotFound = errors.New("student not found")
	ErrInvalidStudent  = errors.New("student name, email, and valid age are required")
)

// Entity / Model
type Student struct {
	ID    int
	Name  string
	Email string
	Age   int
}

// Repository interface describes data access behavior.
type StudentRepository interface {
	Create(student Student) Student
	FindByID(id int) (Student, error)
	FindAll() []Student
	Update(id int, student Student) (Student, error)
	Delete(id int) error
}

// InMemoryStudentRepository stores students in a map.
type InMemoryStudentRepository struct {
	students map[int]Student
	nextID   int
}

func NewInMemoryStudentRepository() *InMemoryStudentRepository {
	return &InMemoryStudentRepository{
		students: make(map[int]Student),
		nextID:   1,
	}
}

func (r *InMemoryStudentRepository) Create(student Student) Student {
	student.ID = r.nextID
	r.students[student.ID] = student
	r.nextID++

	return student
}

func (r *InMemoryStudentRepository) FindByID(id int) (Student, error) {
	student, ok := r.students[id]
	if !ok {
		return Student{}, ErrStudentNotFound
	}

	return student, nil
}

func (r *InMemoryStudentRepository) FindAll() []Student {
	ids := make([]int, 0, len(r.students))
	for id := range r.students {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	students := make([]Student, 0, len(r.students))
	for _, id := range ids {
		students = append(students, r.students[id])
	}

	return students
}

func (r *InMemoryStudentRepository) Update(id int, student Student) (Student, error) {
	if _, ok := r.students[id]; !ok {
		return Student{}, ErrStudentNotFound
	}

	student.ID = id
	r.students[id] = student

	return student, nil
}

func (r *InMemoryStudentRepository) Delete(id int) error {
	if _, ok := r.students[id]; !ok {
		return ErrStudentNotFound
	}

	delete(r.students, id)
	return nil
}

// Service interface describes student business actions.
type StudentService interface {
	AddStudent(name string, email string, age int) (Student, error)
	GetStudent(id int) (Student, error)
	ListStudents() []Student
	UpdateStudent(id int, name string, email string, age int) (Student, error)
	DeleteStudent(id int) error
}

type studentService struct {
	repository StudentRepository
}

func NewStudentService(repository StudentRepository) StudentService {
	return &studentService{repository: repository}
}

func (s *studentService) AddStudent(name string, email string, age int) (Student, error) {
	student, err := newStudent(name, email, age)
	if err != nil {
		return Student{}, err
	}

	return s.repository.Create(student), nil
}

func (s *studentService) GetStudent(id int) (Student, error) {
	return s.repository.FindByID(id)
}

func (s *studentService) ListStudents() []Student {
	return s.repository.FindAll()
}

func (s *studentService) UpdateStudent(id int, name string, email string, age int) (Student, error) {
	student, err := newStudent(name, email, age)
	if err != nil {
		return Student{}, err
	}

	return s.repository.Update(id, student)
}

func (s *studentService) DeleteStudent(id int) error {
	return s.repository.Delete(id)
}

func newStudent(name string, email string, age int) (Student, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if name == "" || email == "" || age <= 0 {
		return Student{}, ErrInvalidStudent
	}

	return Student{
		Name:  name,
		Email: email,
		Age:   age,
	}, nil
}

func printStudents(title string, students []Student) {
	fmt.Println(title)
	for _, student := range students {
		fmt.Printf("ID: %d | Name: %s | Email: %s | Age: %d\n", student.ID, student.Name, student.Email, student.Age)
	}
	fmt.Println()
}

func printStudent(title string, student Student) {
	fmt.Println(title)
	fmt.Printf("ID: %d | Name: %s | Email: %s | Age: %d\n\n", student.ID, student.Name, student.Email, student.Age)
}

func mustAddStudent(service StudentService, name string, email string, age int) Student {
	student, err := service.AddStudent(name, email, age)
	if err != nil {
		panic(err)
	}

	return student
}

func main() {
	repository := NewInMemoryStudentRepository()
	service := NewStudentService(repository)

	// Create
	student1 := mustAddStudent(service, "Rahim", "rahim@example.com", 20)
	student2 := mustAddStudent(service, "Karim", "karim@example.com", 22)
	mustAddStudent(service, "Ayesha", "ayesha@example.com", 21)

	printStudents("After Create:", service.ListStudents())

	// Read one
	foundStudent, err := service.GetStudent(student1.ID)
	if err != nil {
		fmt.Println(err)
	} else {
		printStudent("Read One:", foundStudent)
	}

	// Update
	updatedStudent, err := service.UpdateStudent(student2.ID, "Karim Ahmed", "karim.ahmed@example.com", 23)
	if err != nil {
		fmt.Println(err)
	} else {
		printStudent("After Update:", updatedStudent)
	}

	// Delete
	err = service.DeleteStudent(student1.ID)
	if err != nil {
		fmt.Println(err)
	}

	printStudents("After Delete:", service.ListStudents())
}

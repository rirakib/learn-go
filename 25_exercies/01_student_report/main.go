package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	ID         string
	Name       string
	Department string
	Age        int
	Marks      float32
}

type StudentInput struct {
	ID         string
	Name       string
	Department string
	Age        int
	Marks      float32
}

var (
	students    []*Student
	inputReader = bufio.NewReader(os.Stdin)
)

func NewStudent(input StudentInput) *Student {
	return &Student{
		ID:         input.ID,
		Name:       input.Name,
		Department: input.Department,
		Age:        input.Age,
		Marks:      input.Marks,
	}
}

func addStudent(s *Student) {
	students = append(students, s)
}

func findStudentByID(id string) *Student {
	for _, student := range students {
		if student.ID == id {
			return student
		}
	}
	return nil
}

func updateStudent(id string, input StudentInput) bool {
	for _, student := range students {
		if student.ID == id {
			student.ID = input.ID
			student.Name = input.Name
			student.Department = input.Department
			student.Age = input.Age
			student.Marks = input.Marks
			return true
		}
	}
	return false
}

func deleteStudent(id string) bool {
	for i, student := range students {
		if student.ID == id {
			students = append(students[:i], students[i+1:]...)
			return true
		}
	}
	return false
}

func filterStudentsByName(keyword string) []*Student {
	keyword = strings.ToLower(strings.TrimSpace(keyword))
	if keyword == "" {
		return students
	}

	result := make([]*Student, 0)
	for _, student := range students {
		if strings.Contains(strings.ToLower(student.Name), keyword) {
			result = append(result, student)
		}
	}
	return result
}

func filterStudentsByDepartment(department string) []*Student {
	department = strings.ToLower(strings.TrimSpace(department))
	if department == "" {
		return students
	}

	result := make([]*Student, 0)
	for _, student := range students {
		if strings.Contains(strings.ToLower(student.Department), department) {
			result = append(result, student)
		}
	}
	return result
}

func printStudents(list []*Student) {
	fmt.Println("=================================================================")
	fmt.Printf("%-5s | %-20s | %-15s | %-5s | %-10s\n", "ID", "Name", "Department", "Age", "Marks")
	fmt.Println("=================================================================")

	if len(list) == 0 {
		fmt.Println("No students found.")
		return
	}

	for _, student := range list {
		fmt.Printf("%-5s | %-20s | %-15s | %-5d | %-10.1f\n", student.ID, student.Name, student.Department, student.Age, student.Marks)
	}
}

func readLine(prompt string) string {
	fmt.Print(prompt)
	line, _ := inputReader.ReadString('\n')
	return strings.TrimSpace(line)
}

func readInt(prompt string) int {
	value := readLine(prompt)
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return parsed
}

func readFloat(prompt string) float32 {
	value := readLine(prompt)
	parsed, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0
	}
	return float32(parsed)
}

func readStudentInput() StudentInput {
	return StudentInput{
		ID:         readLine("Enter student ID: "),
		Name:       readLine("Enter student name: "),
		Department: readLine("Enter department: "),
		Age:        readInt("Enter age: "),
		Marks:      readFloat("Enter marks: "),
	}
}

func readStudentInputForUpdate(existing *Student) StudentInput {
	return StudentInput{
		ID:         readLine(fmt.Sprintf("Enter new ID [%s]: ", existing.ID)),
		Name:       readLine(fmt.Sprintf("Enter new name [%s]: ", existing.Name)),
		Department: readLine(fmt.Sprintf("Enter new department [%s]: ", existing.Department)),
		Age:        readInt(fmt.Sprintf("Enter new age [%d]: ", existing.Age)),
		Marks:      readFloat(fmt.Sprintf("Enter new marks [%.1f]: ", existing.Marks)),
	}
}

func seedStudents() {
	addStudent(NewStudent(StudentInput{ID: "101", Name: "Rakibul Islam", Department: "CSE", Age: 26, Marks: 85.4}))
	addStudent(NewStudent(StudentInput{ID: "102", Name: "Sajib Islam", Department: "EEE", Age: 24, Marks: 81.6}))
	addStudent(NewStudent(StudentInput{ID: "103", Name: "Nusrat Jahan", Department: "BBA", Age: 23, Marks: 89.8}))
	addStudent(NewStudent(StudentInput{ID: "104", Name: "Fahim Ahmed", Department: "CSE", Age: 22, Marks: 91.2}))
	addStudent(NewStudent(StudentInput{ID: "105", Name: "Mim Akter", Department: "English", Age: 24, Marks: 78.5}))
	addStudent(NewStudent(StudentInput{ID: "106", Name: "Tanvir Hasan", Department: "Civil", Age: 25, Marks: 83.1}))
	addStudent(NewStudent(StudentInput{ID: "107", Name: "Farzana Rahman", Department: "Architecture", Age: 23, Marks: 88.4}))
	addStudent(NewStudent(StudentInput{ID: "108", Name: "Mahadi Hasan", Department: "CSE", Age: 21, Marks: 95.7}))
	addStudent(NewStudent(StudentInput{ID: "109", Name: "Jannatul Ferdous", Department: "Law", Age: 24, Marks: 84.3}))
	addStudent(NewStudent(StudentInput{ID: "110", Name: "Ashraful Islam", Department: "EEE", Age: 22, Marks: 79.9}))
	addStudent(NewStudent(StudentInput{ID: "111", Name: "Niaz Morshed", Department: "CSE", Age: 23, Marks: 87.2}))
	addStudent(NewStudent(StudentInput{ID: "112", Name: "Tasmia Rahman", Department: "BBA", Age: 22, Marks: 90.5}))
	addStudent(NewStudent(StudentInput{ID: "113", Name: "Sanzida Akter", Department: "English", Age: 24, Marks: 82.0}))
	addStudent(NewStudent(StudentInput{ID: "114", Name: "Rashedul Islam", Department: "EEE", Age: 25, Marks: 76.4}))
	addStudent(NewStudent(StudentInput{ID: "115", Name: "Anika Tabassum", Department: "Civil", Age: 23, Marks: 85.9}))
}

func main() {
	seedStudents()

	for {
		fmt.Println("\n===== Student Report Menu =====")
		fmt.Println("1. View all students")
		fmt.Println("2. Add student")
		fmt.Println("3. Edit student")
		fmt.Println("4. Delete student")
		fmt.Println("5. Filter by name")
		fmt.Println("6. Filter by department")
		fmt.Println("7. Exit")

		choice := readLine("Choose an option: ")
		switch choice {
		case "1":
			printStudents(students)
		case "2":
			addStudent(NewStudent(readStudentInput()))
			fmt.Println("Student added successfully.")
		case "3":
			id := readLine("Enter student ID to edit: ")
			student := findStudentByID(id)
			if student == nil {
				fmt.Println("Student not found.")
			} else {
				updatedInput := readStudentInputForUpdate(student)
				if updatedInput.ID == "" {
					updatedInput.ID = student.ID
				}
				if updatedInput.Name == "" {
					updatedInput.Name = student.Name
				}
				if updatedInput.Department == "" {
					updatedInput.Department = student.Department
				}
				if updatedInput.Age == 0 {
					updatedInput.Age = student.Age
				}
				if updatedInput.Marks == 0 {
					updatedInput.Marks = student.Marks
				}
				if updateStudent(id, updatedInput) {
					fmt.Println("Student updated successfully.")
				} else {
					fmt.Println("Unable to update student.")
				}
			}
		case "4":
			id := readLine("Enter student ID to delete: ")
			if deleteStudent(id) {
				fmt.Println("Student deleted successfully.")
			} else {
				fmt.Println("Student not found.")
			}
		case "5":
			keyword := readLine("Enter name to filter: ")
			printStudents(filterStudentsByName(keyword))
		case "6":
			department := readLine("Enter department to filter: ")
			printStudents(filterStudentsByDepartment(department))
		case "7", "exit", "q":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

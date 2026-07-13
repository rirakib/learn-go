package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Employee struct {
	ID         string
	Name       string
	Department string
	Salary     float64
	Age        int
	Status     bool
}

type EmployeeInput struct {
	Name       string
	Department string
	Salary     float64
	Age        int
	Status     bool
}

var employeeSet = make(map[string]Employee)

func init() {
	rand.Seed(time.Now().UnixNano())
}




func addEmployee(data EmployeeInput) (string, error) {

	var id string

	for {
		id = fmt.Sprintf("%d", rand.Intn(999999)+1000)
		if _, exists := employeeSet[id]; !exists {
			break
		}
	}

	employee := Employee{
		ID:         id,
		Name:       data.Name,
		Department: data.Department,
		Salary:     data.Salary,
		Age:        data.Age,
		Status:     true, // Default status is true
	}

	employeeSet[id] = employee
	return id, nil
}

func showEmployee(id string) {
	employee, exists := employeeSet[id]
	if !exists || employee.ID == "" {
		fmt.Println("Employee not found")
		return
	}

	statusStr := "Active"
	if !employee.Status {
		statusStr = "Inactive"
	}

	fmt.Println("=================================================================================================")
	fmt.Printf("%-10s | %-25s | %-20s | %-8s | %-12s | %-10s\n", "ID", "Name", "Department", "Age", "Salary", "Status")
	fmt.Println("=================================================================================================")
	fmt.Printf("%-10s | %-25s | %-20s | %-8d | %-12.2f | %-10s\n", employee.ID, employee.Name, employee.Department, employee.Age, employee.Salary, statusStr)
	fmt.Println("=================================================================================================")
}

func displayEmployeeList(list []Employee) {
	if len(list) == 0 {
		fmt.Println("No employee found.")
		return
	}

	fmt.Println("=================================================================================================")
	fmt.Printf("%-10s | %-25s | %-20s | %-8s | %-12s | %-10s\n", "ID", "Name", "Department", "Age", "Salary", "Status")
	fmt.Println("=================================================================================================")
	for _, value := range list {
		statusStr := "Active"
		if !value.Status {
			statusStr = "Inactive"
		}
		fmt.Printf("%-10s | %-25s | %-20s | %-8d | %-12.2f | %-10s\n", value.ID, value.Name, value.Department, value.Age, value.Salary, statusStr)
	}
	fmt.Println("=================================================================================================")
}

func employeeList() {
	var list []Employee
	for _, value := range employeeSet {
		list = append(list, value)
	}
	displayEmployeeList(list)
}

func filterEmployees(attribute string, query string) []Employee {
	var filtered []Employee
	query = strings.ToLower(strings.TrimSpace(query))
	attribute = strings.ToLower(strings.TrimSpace(attribute))

	for _, employee := range employeeSet {
		match := false
		switch attribute {
		case "id":
			match = strings.Contains(strings.ToLower(employee.ID), query)
		case "name":
			match = strings.Contains(strings.ToLower(employee.Name), query)
		case "department":
			match = strings.Contains(strings.ToLower(employee.Department), query)
		}

		if match {
			filtered = append(filtered, employee)
		}
	}
	return filtered
}

func updateEmployee(id string, data EmployeeInput) error {
	employee, exists := employeeSet[id]
	if !exists {
		return fmt.Errorf("employee with ID %s not found", id)
	}

	if data.Name != "" {
		employee.Name = data.Name
	}
	if data.Department != "" {
		employee.Department = data.Department
	}
	if data.Salary > 0 {
		employee.Salary = data.Salary
	}
	if data.Age > 0 {
		employee.Age = data.Age
	}
	employee.Status = data.Status

	employeeSet[id] = employee
	return nil
}

func toggleEmployeeStatus(id string) error {
	employee, exists := employeeSet[id]
	if !exists {
		return fmt.Errorf("employee with ID %s not found", id)
	}
	employee.Status = !employee.Status
	employeeSet[id] = employee
	return nil
}

func seedEmployees() {
	employees := []EmployeeInput{
		{Name: "Rakibul Islam", Department: "Backend", Salary: 60000.00, Age: 26},
		{Name: "Sajib Hasan", Department: "Frontend", Salary: 35000.00, Age: 24},
		{Name: "Nusrat Imroz Orni", Department: "Frontend", Salary: 35000.00, Age: 23},
		{Name: "Fahim Ahmed", Department: "QA", Salary: 40000.00, Age: 25},
		{Name: "Mim Akter", Department: "HR", Salary: 45000.00, Age: 27},
	}

	for _, empInput := range employees {
		addEmployee(empInput)
	}
}

func showStatistics() {
	total := len(employeeSet)
	if total == 0 {
		fmt.Println("No employee data available for statistics.")
		return
	}

	activeCount := 0
	inactiveCount := 0
	var totalSalary float64
	var totalAge int
	var highestSalary float64
	var lowestSalary float64

	first := true
	for _, emp := range employeeSet {
		if emp.Status {
			activeCount++
		} else {
			inactiveCount++
		}

		totalSalary += emp.Salary
		totalAge += emp.Age

		if first {
			highestSalary = emp.Salary
			lowestSalary = emp.Salary
			first = false
		} else {
			if emp.Salary > highestSalary {
				highestSalary = emp.Salary
			}
			if emp.Salary < lowestSalary {
				lowestSalary = emp.Salary
			}
		}
	}

	avgSalary := totalSalary / float64(total)
	avgAge := float64(totalAge) / float64(total)

	fmt.Println("\n=======================================")
	fmt.Println("         EMPLOYEE STATISTICS           ")
	fmt.Println("=======================================")
	fmt.Printf("Total Employees:        %d\n", total)
	fmt.Printf("Active Employees:       %d\n", activeCount)
	fmt.Printf("Inactive Employees:     %d\n", inactiveCount)
	fmt.Printf("Average Salary:         $%.2f\n", avgSalary)
	fmt.Printf("Highest Salary:         $%.2f\n", highestSalary)
	fmt.Printf("Lowest Salary:          $%.2f\n", lowestSalary)
	fmt.Printf("Average Age:            %.1f years\n", avgAge)
	fmt.Println("=======================================")
}

func main() {
	seedEmployees()

	fmt.Println("All Seeded Employees:")
	employeeList()

	// Add an employee for update/toggle demonstration
	testId, _ := addEmployee(EmployeeInput{
		Name:       "Test Employee",
		Department: "Design",
		Salary:     30000.00,
		Age:        22,
	})

	fmt.Println("\nBefore Update:")
	showEmployee(testId)

	err := updateEmployee(testId, EmployeeInput{
		Name:       "Test Employee Updated",
		Department: "UI/UX Design",
		Salary:     42000.00,
		Age:        23,
		Status:     true,
	})
	if err != nil {
		fmt.Println("Error updating employee:", err)
	}

	fmt.Println("\nAfter Update:")
	showEmployee(testId)

	fmt.Println("\nToggling Status:")
	toggleEmployeeStatus(testId)
	showEmployee(testId)

	fmt.Println("\n--- Testing Filter (Department: Frontend) ---")
	frontendStaff := filterEmployees("department", "frontend")
	displayEmployeeList(frontendStaff)

	fmt.Println("\n--- Statistics Report ---")
	showStatistics()
}

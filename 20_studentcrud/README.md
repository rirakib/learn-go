# Student CRUD in Go

This example shows a complete **Create, Read, Update, Delete** flow for students using interfaces and a simple layered design.

---

## Design Pattern Used

This example uses the **Repository Pattern** with a small **Service Layer**.

| Layer | Responsibility |
|---|---|
| `Student` | The entity/model that represents student data. |
| `StudentRepository` | Interface for storage operations. |
| `InMemoryStudentRepository` | Map-based repository implementation. |
| `StudentService` | Interface for business actions. |
| `studentService` | Service implementation that validates data and calls the repository. |
| `main` | Demonstrates the CRUD workflow. |

---

## CRUD Operations

| Operation | Method |
|---|---|
| Create | `AddStudent` |
| Read One | `GetStudent` |
| Read All | `ListStudents` |
| Update | `UpdateStudent` |
| Delete | `DeleteStudent` |

---

## Why Interfaces?

Interfaces make the code flexible. Today the app uses `InMemoryStudentRepository`, but later you can create another implementation such as `MySQLStudentRepository` or `PostgresStudentRepository` without changing the service logic.

```go
type StudentRepository interface {
    Create(student Student) Student
    FindByID(id int) (Student, error)
    FindAll() []Student
    Update(id int, student Student) (Student, error)
    Delete(id int) error
}
```

---

## How to Run

```bash
go run main.go
```

## How to Test

```bash
go test
```

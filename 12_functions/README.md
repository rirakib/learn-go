# Functions as First-Class Citizens in Go

In Go, functions are **first-class citizens**. This means functions can be assigned to variables, passed as arguments to other functions, returned from other functions, and stored in data structures (like maps or slices).

This directory demonstrates how to use first-class functions to implement clean, reusable, and extensible code adhering to the **Open-Closed Principle (OCP)** and **DRY (Don't Repeat Yourself)** principles.

---

## 1. Defining Function Types

In Go, you can define a custom type representing a function signature. This makes the code self-documenting and easier to read:

```go
type MathOp func(int, int) float64
```
Any function that takes two `int` parameters and returns a `float64` automatically implements the `MathOp` signature.

---

## 2. Implementing Design Principles

### A. Open-Closed Principle (OCP) via Strategy Pattern
The **Open-Closed Principle** states that software entities should be open for extension but closed for modification. 

In [main.go](file:///c:/Users/rirra/Downloads/learn-go/12_functions/main.go), the `Calculate` function uses the Strategy Pattern by accepting a `MathOp` function as a parameter:
```go
func Calculate(a, b int, op MathOp) float64 {
	return op(a, b)
}
```
* **Closed for Modification**: We never have to change the `Calculate` function.
* **Open for Extension**: We can pass any new mathematical operation (e.g., Modulo or Power) into `Calculate` as long as it matches the `MathOp` signature.

For example, passing an anonymous function:
```go
resultMod := Calculate(10, 3, func(a, b int) float64 {
	return float64(a % b)
})
```

---

### B. DRY & Extensible Lookup via Registry Pattern
Historically, lookup/dispatcher functions are written using large `switch-case` blocks:
```go
// VIOLATION of OCP (must modify this switch whenever we add operations)
func process(opName string, a, b int) float64 {
	switch opName {
	case "add": return Add(a, b)
	// ...
	}
}
```
To keep the code **DRY** and **OCP-compliant**, we replace the `switch` statement with a dynamic map-based registry:
```go
var registry = map[string]MathOp{
	"add":      Add,
	"subtract": Subtract,
	"multiply": Multiply,
	"divide":   Divide,
}

func Process(opName string, a, b int) (float64, error) {
	op, exists := registry[opName]
	if !exists {
		return 0, fmt.Errorf("unsupported operation: %s", opName)
	}
	return op(a, b), nil
}
```
We can now dynamically register new operations at runtime by writing to the map without modifying `Process`:
```go
registry["modulo"] = func(a, b int) float64 { return float64(a % b) }
```

---

## How to Run the Code in this Directory

To execute the example program demonstrating functions:

```bash
go run main.go
```

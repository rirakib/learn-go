# Constants and Functions in Go

This directory covers declaring constants using `const` and defining basic functions with parameters and return types in Go.

## Code Explanation

```go
package main 

import "fmt"

const (
	host = "localhost"
	port = 8080
	db_name = "my_database"
)


func sum(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Host:", host)
	fmt.Println("Port:", port)
	fmt.Println("Database Name:", db_name)


	fmt.Print(sum(10,20))
}
```

### Breakdown of the Code:

1. **`const (...)`**
   - Declares character, string, boolean, or numeric values that cannot be mutated (changed) after compile time.
   - Using parentheses allows grouping multiple constant declarations together.
   - Here, we declare `host` as a string, `port` as an untyped integer constant, and `db_name` as a string.

2. **`func sum(a int, b int) int { ... }`**
   - Defines a custom function named `sum`.
   - It takes two input parameters: `a` (type `int`) and `b` (type `int`).
   - The final `int` before the opening brace `{` specifies the **return type** of the function.
   - It returns the sum of `a` and `b` via the `return` keyword.

3. **`main` Function**:
   - Prints the values of the constant variables.
   - Calls the `sum` function passing `10` and `20` as arguments, and prints the returned result (`30`).

---

## Key Learning Takeaways

* **Constants**:
  - Declared with the `const` keyword.
  - Can only be string, character, boolean, or numeric values.
  - Their values cannot be modified during program execution.
* **Function Definition**:
  - Go functions are declared with the `func` keyword.
  - Types are specified after the parameter names.
  - The return type is defined after the parameter list.

---

## How to Run

To run this example, navigate to this directory in your terminal and execute:

```bash
go run main.go
```

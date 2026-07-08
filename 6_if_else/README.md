# If/Else Conditionals in Go

This directory demonstrates standard conditional branching using `if`, `else if`, and `else` statements, alongside logical operators.

## Code Explanation

```go
package main 

import "fmt"

func main() {

	var age int = 20;

	if age >= 18 {
		fmt.Println("You are an adult.")
	}else if age >= 13 {
		fmt.Println("You are a teenager.")
	}else {
		fmt.Println("You are a child.")
	}

	var username string = "admin"
	var password string = "password123"

	if username == "admin" && password == "password123" {
		fmt.Println("Login successful!")
	}else {
		fmt.Println("Invalid username or password.")
	}

}
```

### Breakdown of the Code:

1. **Age Verification Block**:
   ```go
   if age >= 18 {
       fmt.Println("You are an adult.")
   } else if age >= 13 {
       fmt.Println("You are a teenager.")
   } else {
       fmt.Println("You are a child.")
   }
   ```
   - Standard conditional chain. Go does not require parentheses around the conditions (e.g., `age >= 18`), but braces `{}` are mandatory even for single-line blocks.
   - The code checks if `age` is greater than or equal to 18, then 13, and falls back to else.

2. **Login Verification Block**:
   ```go
   if username == "admin" && password == "password123" {
       fmt.Println("Login successful!")
   } else {
       fmt.Println("Invalid username or password.")
   }
   ```
   - Demonstrates compound conditions using the logical AND (`&&`) operator.
   - The block within `if` only executes if both `username == "admin"` and `password == "password123"` evaluate to `true`.

---

## Key Learning Takeaways

* **Braces Requirement**: Curly braces `{}` are always required for `if`/`else` structures in Go, even if there is only a single statement inside.
* **No Parentheses**: You do not need to wrap your conditions in parentheses: `if age >= 18` (unlike Java, C++, or JavaScript).
* **Logical Operators**:
  - `&&`: Logical AND
  - `||`: Logical OR
  - `!`: Logical NOT

---

## How to Run

To run this example, navigate to this directory in your terminal and execute:

```bash
go run main.go
```

# Simple Values in Go

This directory covers basic data types and simple values in Go. It demonstrates printing integers, floats, strings, booleans, and basic operations.

## Code Explanation

```go
package main

import "fmt"

func main() {
	fmt.Println(1+3)

	fmt.Println("this is string value")

	fmt.Println(true)

	fmt.Println(1.5)

	fmt.Println(7/3.12)
}
```

### Breakdown of the Code:

1. **`fmt.Println(1+3)`**
   - Performs basic addition of two integers (`1` and `3`) and prints the result: `4`.

2. **`fmt.Println("this is string value")`**
   - Prints a basic string literal. In Go, string literals are enclosed in double quotes `"`.

3. **`fmt.Println(true)`**
   - Prints a boolean value (`true`). Booleans represent logical values and can be either `true` or `false`.

4. **`fmt.Println(1.5)`**
   - Prints a floating-point number (decimal value).

5. **`fmt.Println(7/3.12)`**
   - Performs division between an integer (`7`) and a floating-point number (`3.12`). Since one of the operands is a float, the compiler performs floating-point division and prints the result (approx. `2.243589`).

---

## Key Learning Takeaways

* **Data Types**: Go has standard built-in types such as:
  - **Integers**: Whole numbers (e.g., `1`, `3`).
  - **Floats**: Decimal numbers (e.g., `1.5`, `3.12`).
  - **Strings**: Text values (e.g., `"this is string value"`).
  - **Booleans**: True/false values (e.g., `true`).
* **Implicit Conversion**: When operating on mixed types like integers and floats (e.g., `7 / 3.12`), Go converts/promotes types to perform the operation. However, Go is strictly typed, and operations on completely incompatible types will result in a compile error.

---

## How to Run

To run this example, navigate to this directory in your terminal and execute:

```bash
go run main.go
```

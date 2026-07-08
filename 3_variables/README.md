# Variables in Go

This directory demonstrates different ways to declare, initialize, and use variables in Go.

## Code Explanation

```go
package main 

import "fmt"

func main() {
	var name string = "John Doe"
	var age int = 30
	var price float32 = 19.99
	isActive := true
	var isMember = true

	fmt.Println(name,age,price,isActive,isMember)
}
```

### Breakdown of the Code:

1. **`var name string = "John Doe"`**
   - Declares a variable named `name` of type `string` and initializes it to `"John Doe"`.
   - The standard declaration syntax is: `var <variable_name> <type> = <value>`.

2. **`var age int = 30`**
   - Declares an integer variable named `age` and initializes it to `30`.

3. **`var price float32 = 19.99`**
   - Declares a floating-point variable named `price` of type `float32` and initializes it to `19.99`.

4. **`isActive := true`**
   - This is the **short variable declaration** syntax.
   - The type is automatically inferred by the Go compiler (in this case, `bool`).
   - Note: The `:=` syntax is only available *inside* functions.

5. **`var isMember = true`**
   - Declares a variable `isMember` without specifying the type explicitly, but using the `var` keyword.
   - Go infers the type based on the initial value (`bool`).

6. **`fmt.Println(name,age,price,isActive,isMember)`**
   - Prints all the variable values to the console, separated by spaces.

---

## Key Learning Takeaways

* **Explicit Type Declaration**: You can explicitly define a variable's type: `var name string = "John Doe"`.
* **Type Inference**: Go can automatically infer the type of a variable if you omit it: `var isMember = true`.
* **Short Variable Declaration (`:=`)**: Inside functions, you can declare and initialize variables quickly using `:=`. This is the most common way to declare variables in Go when they are local to a function.
* **Static Typing**: Once a variable is declared with a specific type (explicitly or inferred), its type cannot be changed.

---

## How to Run

To run this example, navigate to this directory in your terminal and execute:

```bash
go run main.go
```

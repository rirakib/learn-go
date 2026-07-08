# Hello World in Go

This directory introduces the very basics of a Go program. It contains the simplest Go program that prints a message to the console.

## Code Explanation

```go
package main 

import "fmt"

func main(){
	fmt.Println("Hello, Rakib!")
}
```

### Breakdown of the Code:

1. **`package main`**
   - Every Go file must start with a `package` declaration.
   - The name `main` is a special package name in Go. It tells the Go compiler that this program should compile as an executable command rather than a shared library.

2. **`import "fmt"`**
   - The `import` keyword is used to include external or standard library packages.
   - Here, we import the `"fmt"` (format) package from the standard library, which contains functions for formatting text, including printing to the console.

3. **`func main() { ... }`**
   - The `main` function is the entry point of every executable Go program.
   - When you run the program, execution starts here.
   - It takes no arguments and returns nothing.

4. **`fmt.Println("Hello, Rakib!")`**
   - `Println` is a function from the imported `fmt` package.
   - It prints the string argument to the standard output (terminal), followed by a newline.

---

## Key Learning Takeaways

* **Packages**: Code in Go is organized into packages. Executable programs must use the `main` package.
* **Entry Point**: The `main` function in the `main` package is required to execute the program.
* **Formatting**: Standard input/output tasks are handled by the built-in `fmt` package.

---

## How to Run

To run this example, navigate to this directory in your terminal and execute:

```bash
go run main.go
```

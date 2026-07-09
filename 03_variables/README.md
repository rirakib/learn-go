# Variables in Go

Variables are containers for storing data values during program execution. Go is statically typed, meaning every variable must have a defined type, but Go also provides powerful type inference to keep your code clean and concise.

---

## Variable Declaration Styles

Go provides multiple ways to declare and initialize variables depending on the context and scoping requirements.

### 1. Explicit Type Declaration
This is the standard declaration syntax. It is useful when you want to be completely explicit about a variable's type.
```go
var name string = "John Doe"
var age int = 30
```

### 2. Type Inference with `var`
If you provide an initial value but omit the type, Go will automatically infer the type based on the value on the right-hand side.
```go
var isMember = true // Inferred as bool
var price = 19.99   // Inferred as float64
```

### 3. Short Variable Declaration (`:=`)
This is the most common way to declare variables inside functions. It acts as a shortcut for declaring and initializing a variable.
```go
isActive := true
count := 10
```
> [!WARNING]
> The short declaration operator (`:=`) can **only** be used inside functions. Variables declared at the package level (outside any function) must use the `var` keyword.

### 4. Zero-Value Declaration
If you declare a variable without an initial value, Go automatically initializes it with its default "zero value". This guarantees memory safety.
```go
var score int   // Defaults to 0
var title string // Defaults to ""
```

---

## Advanced Declaration Techniques

### Multiple Variable Declarations
You can declare multiple variables on a single line if they share the same type or initializations:
```go
var a, b, c int = 1, 2, 3
x, y, z := "one", 2, true
```

### Grouped Declaration Block
For grouping related package-level variables or configuration settings:
```go
var (
    firstName string = "Alex"
    lastName  string = "Smith"
    userAge   int    = 25
)
```

---

## Variable Scoping and Naming Conventions

### Scoping Rules
* **Package Level**: Variables declared outside of any function are visible across the entire package. They cannot be declared using `:=`.
* **Local Level**: Variables declared inside a function (block scope) are only visible within that function or block.

### Naming Conventions
* **camelCase**: Standard variable names should use camelCase (e.g., `userAge`, `totalCount`).
* **Capitalization for Exporting (Visibility)**:
  * If a package-level variable starts with an **uppercase letter** (e.g., `ExportedVar`), it is public and can be accessed by other packages.
  * If it starts with a **lowercase letter** (e.g., `privateVar`), it is private and accessible only within the containing package.

---

## How to Run the Code in this Directory

To execute the example program demonstrating variables, run:

```bash
go run main.go
```

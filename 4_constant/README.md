# Constants and Functions in Go

This directory covers **Constants**—immutable values determined at compile time—and introduces **Functions** in Go.

---

## 1. Constants (`const`)

Constants are fixed values that the compiler resolves during compilation. Once declared, their values cannot be changed or reassigned at runtime.

### Declaration and Rules
Constants are declared using the `const` keyword:
```go
const Pi = 3.14159
```

* **Compile-Time Only**: Constants must be evaluable at compile time. This means you cannot initialize a constant with a value that is only known at runtime (such as a function call, database query, or system input).
  ```go
  // Invalid! math.Sqrt(4) is calculated at runtime.
  const root = math.Sqrt(4) 
  ```
* **Supported Types**: Constants can only be character, string, boolean, or numeric (integers, floats, complex) values.

### Grouped Constants
Similar to variables, you can group multiple constants together inside parentheses:
```go
const (
    host    = "localhost"
    port    = 8080
    dbName  = "my_database"
)
```

### Typed vs. Untyped Constants
Go constants are unique because they can be **untyped**:
* **Untyped Constants**: An untyped constant (like `const x = 5`) does not have a fixed size or strict type until it is used in a context that requires one. This provides massive flexibility and avoids tedious conversions.
* **Typed Constants**: A typed constant has a strict type (e.g., `const x int32 = 5`) and can only be used with variables of that exact type.

### The `iota` Constant Generator
Go provides a special pre-declared identifier called `iota` to simplify the definition of incrementing numbers, commonly used to create enums:
```go
const (
    Sunday = iota // 0
    Monday        // 1
    Tuesday       // 2
)
```
Each time `iota` is used in a new constant block, it resets to `0` and increments by `1` for each subsequent constant line.

---

## 2. Basic Functions

Functions are the core building blocks of logic in Go.

### Syntax
```go
func functionName(parameter1 type, parameter2 type) returnType {
    // block of code
    return value
}
```

### Key Rules:
1. **Parameter Types**: You must explicitly declare the type of each parameter. If consecutive parameters share the same type, you can omit the type for all but the last:
   ```go
   func sum(a, b int) int { // both a and b are integers
       return a + b
   }
   ```
2. **Return Type**: The return type is defined after the parameter list. If the function doesn't return a value, the return type is omitted entirely.
3. **Multiple Return Values**: Go natively supports returning multiple values from a single function (useful for returning errors along with results):
   ```go
   func divide(a, b float64) (float64, error) {
       if b == 0 {
           return 0, errors.New("cannot divide by zero")
       }
       return a / b, nil
   }
   ```

---

## How to Run the Code in this Directory

To execute the example program demonstrating constants and functions, run:

```bash
go run main.go
```

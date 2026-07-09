# If/Else Conditionals in Go

Conditionals are used to perform different actions based on different conditions. Go's conditional statements follow a strict, clean syntax that prioritizes readability and safety.

---

## 1. Basic Conditional Syntax

Go supports the standard `if`, `else if`, and `else` control flow structure.

```go
if age >= 18 {
    fmt.Println("You are an adult.")
} else if age >= 13 {
    fmt.Println("You are a teenager.")
} else {
    fmt.Println("You are a child.")
}
```

### Critical Syntax Rules:
1. **Mandatory Braces**: Curly braces `{}` are **always** required, even if there is only a single statement inside the block.
2. **No Parentheses**: You do not need to wrap your conditions in parentheses: `if age >= 18` (instead of `if (age >= 18)`).
3. **Else Placement**: The `else` or `else if` keyword must start on the same line as the closing brace `}` of the preceding block. Writing `else` on a new line will cause a compilation error.
   ```go
   // Correct
   if condition {
   } else {
   }

   // Incorrect - Will cause compiler error!
   if condition {
   }
   else {
   }
   ```

---

## 2. Statement Initialization (Inline Variable Declaration)

Go supports a unique, highly useful feature: declaring a temporary variable directly inside the `if` statement.

```go
if limit := getLimit(); val > limit {
    fmt.Println("Over limit")
}
```
* **Syntax**: `if init; condition { ... }`
* **Scoping**: The variable `limit` is only visible inside the `if` statement blocks (including any `else if` and `else` blocks attached to it). Once the conditional block ends, the variable goes out of scope and is cleaned up. This helps keep local namespace variables clean.

---

## 3. Logical Operators

You can build complex logic checks using logical operators. Go uses **short-circuit evaluation**, meaning if the result is already determined by the first operand, the second operand is not evaluated.

| Operator | Meaning | Example |
| :--- | :--- | :--- |
| **`&&`** | Logical AND (both conditions must be true) | `if username == "admin" && password == "123"` |
| **`||`** | Logical OR (at least one condition must be true) | `if isWeekend || isHoliday` |
| **`!`** | Logical NOT (inverts the boolean check) | `if !isLoggedIn` |

---

## How to Run the Code in this Directory

To execute the example program demonstrating conditionals, run:

```bash
go run main.go
```

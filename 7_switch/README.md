# Switch Statements in Go

A `switch` statement is a conditional statement that matches a value against multiple possible cases and executes the corresponding block. In Go, `switch` is highly flexible, powerful, and has safer defaults than most C-family languages.

---

## 1. Core Switch Characteristics

### Standard Switch
A basic switch matches a single expression against concrete case values:
```go
switch day {
case "Monday":
    fmt.Println("Start of the work week!")
case "Friday":
    fmt.Println("Weekend is near!")
default:
    fmt.Println("A normal day.")
}
```

* **Implicit Break**: Go automatically breaks out of the switch block after executing a matching case. You do not need to end your case blocks with the `break` keyword.
* **Default Case**: The optional `default` block executes if none of the other cases match. It can be placed anywhere in the switch block, but is traditionally put at the end.

### Matching Multiple Values (Comma-Separated Cases)
You can match multiple conditions within a single case statement using a comma-separated list. This acts as a logical **OR**:
```go
switch day {
case "Saturday", "Sunday":
    fmt.Println("It's the weekend!")
default:
    fmt.Println("It's a working day.")
}
```

---

## 2. Advanced Switch Features

### The `fallthrough` Keyword
If you want execution to continue into the next case statement instead of breaking automatically, use the `fallthrough` keyword:
```go
switch status {
case "critical":
    sendAlert()
    fallthrough // execution flows into "warning" case
case "warning":
    logWarning()
}
```
> [!WARNING]
> `fallthrough` will execute the next case block **without** checking its condition expression. Use it sparingly.

### Expressionless Switch (If-Else Alternative)
A switch without an expression is equivalent to `switch true`. This allows you to evaluate complex boolean conditions in each case, making it a highly readable replacement for nested `if`/`else if` chains.
```go
switch {
case marks >= 80:
    fmt.Println("Grade: A")
case marks >= 60:
    fmt.Println("Grade: B")
default:
    fmt.Println("Grade: F")
}
```

### Type Switches
Go allows you to switch on the underlying concrete **type** of an empty interface (`interface{}` or `any`) variable:
```go
func inspectType(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("Integer:", v)
    case string:
        fmt.Println("String:", v)
    default:
        fmt.Println("Unknown type")
    }
}
```

---

## How to Run the Code in this Directory

To execute the example program demonstrating switch structures, run:

```bash
go run main.go
```

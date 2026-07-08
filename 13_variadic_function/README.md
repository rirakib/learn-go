# Variadic Functions in Go

A **Variadic Function** is a function that can accept a variable number of arguments. In Go, the built-in `fmt.Println` and `fmt.Printf` functions are common examples of variadic functions.

---

## 1. Syntax

To declare a variadic function, prefix the type of the final parameter with three dots `...`:

```go
func sum(numbers ...int) int {
    // numbers is treated as a slice of type []int
}
```

> [!IMPORTANT]
> The variadic parameter must be the **last** parameter in the function signature. You cannot have parameters after a variadic parameter.
> For example, `func calculate(name string, numbers ...int)` is valid, but `func calculate(numbers ...int, name string)` is invalid.

---

## 2. Under the Hood

Within the function, the variadic parameter behaves exactly like a **slice** of the specified type.
For example, in `sum(numbers ...int)`:
- The type of `numbers` inside the function is `[]int`.
- You can query its length using `len(numbers)`.
- You can iterate over it using a standard `for range` loop.

---

## 3. Passing Arguments to a Variadic Function

Go supports two ways to supply arguments to a variadic parameter:

### A. Passing Individual Arguments
You can pass zero or more arguments of the specified type as separate values:
```go
total := sum(1, 2, 3, 4, 5)
```

### B. Unpacking a Slice (Spreading)
If you already have a slice of the matching type, you can pass it directly to the variadic function by appending `...` to the slice variable. This is known as **unpacking** or **spreading** the slice:
```go
marks := []int{90, 85, 92}
total := sum(marks...) // Unpacks the slice into 90, 85, 92
```

---

## 4. Key Use Cases

1. **Logging and Formatting**: Functions like `fmt.Println(a ...any)` where you don't know beforehand how many items will be printed.
2. **Aggregations**: Mathematical operations (like sum, minimum, maximum, average, or concatenation) acting on an arbitrary set of inputs.
3. **Pluggable Options**: Passing configuration options or plugins to a constructor dynamically.

---

## How to Run the Code in this Directory

To execute the example program demonstrating variadic functions:

```bash
go run main.go
```

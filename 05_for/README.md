# Loops in Go (The For Loop)

In Go, there is only **one** looping construct: the **`for` loop**. There are no `while` or `do-while` loops. However, Go's single `for` loop is highly versatile and can cover all looping patterns.

---

## 1. Loop Syntax Variations

### The Classic Three-Component Loop
This is the standard C-style loop. It contains three parts separated by semicolons:
1. **Init statement**: executed before the first iteration (usually declaring a counter variable `i := 0`).
2. **Condition expression**: evaluated before every iteration. If `true`, the loop body runs.
3. **Post statement**: executed at the end of each iteration (usually incrementing the counter `i++`).

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

### The Condition-Only Loop (Equivalent to `while`)
You can omit the init and post statements. This acts exactly like a traditional `while` loop in other languages.
```go
i := 1
for i <= 10 {
    fmt.Println(i)
    i++
}
```

### The Infinite Loop
Omitting all components creates a loop that executes forever until a `break` or `return` statement is reached.
```go
for {
    fmt.Println("looping forever...")
    break // stops the loop
}
```

### The Range-Based Loop (Go 1.22+)
Starting in Go 1.22, you can range directly over an integer. The loop runs from `0` up to `n - 1`.
```go
for i := range 5 {
    fmt.Println(i) // prints 0, 1, 2, 3, 4
}
```

---

## 2. Loop Control Flow

Go provides two keywords to control the flow of execution inside loops:

| Keyword | Action | Example |
| :--- | :--- | :--- |
| **`break`** | Exits the loop immediately, skipping all remaining iterations. | ```go for i := 0; i < 10; i++ { if i == 5 { break } } ``` *(Stops looping entirely when i reaches 5)* |
| **`continue`** | Skips the rest of the current iteration and jumps straight to the next evaluation/post statement. | ```go for i := 0; i < 10; i++ { if i == 5 { continue } } ``` *(Skips printing 5, continues with 6)* |

---

## 3. Ranging Over Data Structures

The `range` keyword is also used to loop over elements in slices, arrays, maps, strings, and channels.

* **Slices/Arrays**: Returns the index and element value.
  ```go
  for index, val := range mySlice { ... }
  ```
* **Maps**: Returns the key and value.
  ```go
  for key, val := range myMap { ... }
  ```
* **Strings**: Returns the byte index and the Unicode character (`rune`).
* **Ignoring variables**: You can use the blank identifier `_` to discard the index or value if you don't need it.
  ```go
  for _, val := range mySlice {
      // Index is ignored, only using val
  }
  ```

---

## How to Run the Code in this Directory

To execute the example program demonstrating loop structures, run:

```bash
go run main.go
```

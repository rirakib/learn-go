# Range in Go

The `range` keyword in Go is used in `for` loops to iterate over elements in various data structures: arrays, slices, maps, strings, and channels. It provides a clean, idiomatic way to loop through collections.

---

## 1. Syntax & Iteration Behavior

Depending on the data structure, `range` yields different values:

| Data Type | First Value | Second Value | Description |
| :--- | :--- | :--- | :--- |
| **Arrays & Slices** | Index (`int`) | Element Value | Iterates index by index from `0` to `len-1`. |
| **Maps** | Key | Value | Iterates through all key-value pairs (unordered). |
| **Strings** | Byte Index (`int`) | Rune (`rune`) | Iterates Unicode code points (UTF-8 decoded). |
| **Channels** | Element Value | *None* | Receives values from the channel until closed. |

---

## 2. Using `range` with Different Types

### A. Arrays and Slices
```go
nums := []int{10, 20, 30}
for index, val := range nums {
    fmt.Printf("Index: %d, Value: %d\n", index, val)
}
```

### B. Maps
```go
data := map[string]string{
    "name":  "John Doe",
    "email": "john@gmail.com",
}
for key, val := range data {
    fmt.Printf("%s: %s\n", key, val)
}
```

### C. Strings
When ranging over a string, Go iterates over Unicode code points (runes) rather than raw bytes.
```go
for index, char := range "Go!" {
    fmt.Printf("Index: %d, Char: %c (Unicode: %U)\n", index, char, char)
}
```

---

## 3. Ignoring Loop Variables

If you don't need both the key/index and the value, you can omit or ignore them:

### Ignore Value
Omit the second variable completely:
```go
for index := range mySlice {
    fmt.Println("Index:", index)
}
```

### Ignore Index/Key
Use the blank identifier `_` to discard the index/key:
```go
for _, val := range mySlice {
    fmt.Println("Value:", val)
}
```

### Loop Count Only (Go 1.22+)
If you only want to loop a certain number of times or drain a channel without assigning variables:
```go
for range mySlice {
    // Loops len(mySlice) times without allocating loop variables
}
```

---

## 4. Under the Hood (Actual Work & Mechanics)

Understanding how `range` operates behind the scenes is crucial for writing efficient and bug-free Go code.

### A. Copy-by-Value Behavior
> [!IMPORTANT]
> The `range` loop always **copies** the value of each element into the loop variable. It does not return a direct reference or pointer to the element in the original collection.

Modifying the loop variable will **not** affect the original slice or map:
```go
numbers := []int{1, 2, 3}

for _, val := range numbers {
    val = val * 10 // Only modifies the local copy 'val'
}
// numbers is still [1, 2, 3]
```

**How to fix it:** If you want to modify elements in place, access them using the index:
```go
for i := range numbers {
    numbers[i] = numbers[i] * 10
}
// numbers is now [10, 20, 30]
```

---

### B. UTF-8 Decoding on Strings
A Go string is a read-only slice of bytes. However, `for range` does not simply loop through each byte. 
Instead, it decodes one UTF-8-encoded Unicode character (rune) at a time.
* The loop variable is of type `rune` (alias for `int32`).
* The index represents the **byte offset** where the current rune begins.
* If a character is multi-byte (like emojis or special symbols), the index will skip bytes.

For example, ranging over `"Go💙"`:
* `'G'` takes 1 byte (index 0)
* `'o'` takes 1 byte (index 1)
* `'💙'` takes 4 bytes (index 2)
* The loop ends, but the length of `"Go💙"` in bytes is `6`. Ranging gives indices `0`, `1`, and `2`.

---

### C. Randomized Map Iteration
When you iterate over a map, Go does not guarantee any specific order. In fact, Go **intentionally randomizes** the start index of map iteration on every execution.
* **Why?** To prevent developers from relying on key ordering, as the underlying hash table structure does not preserve insertion order and can shift elements during resizing.

---

### D. Loop Variable Lifetime (Go 1.22 Update)
Before Go 1.22, loop variables in a `for range` loop were allocated **once** and reused across all iterations. This caused issues when using goroutines or closures inside loops:
```go
// Pre-Go 1.22 Gotcha
for _, val := range []int{1, 2, 3} {
    go func() {
        fmt.Println(val) // Might print 3, 3, 3 because all goroutines share the same 'val' variable
    }()
}
```
**From Go 1.22 onwards**, Go creates a new instance of the loop variable for each iteration. The above code now correctly prints `1`, `2`, `3` (in arbitrary order due to goroutine scheduling) without needing manual reassignment.

---

## How to Run the Code in this Directory

To execute the example program demonstrating the `range` keyword, run:

```bash
go run main.go
```

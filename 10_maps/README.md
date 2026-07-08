# Maps in Go

A **Map** is Go's built-in associative data type (often called a hash table, dictionary, or lookup table). It stores unordered collections of key-value pairs, providing fast $O(1)$ lookups, insertions, and deletions.

---

## 1. Creating and Initializing Maps

### Using the `make()` Function
The standard way to initialize an empty, ready-to-use map is via the `make` function:
```go
// syntax: make(map[KeyType]ValueType)
ageMap := make(map[string]int)
```
> [!CAUTION]
> Declaring a map without initializing it (e.g., `var myMap map[string]int`) creates a `nil` map. Writing to a `nil` map will cause a **runtime panic**. Always initialize maps using `make()` or a map literal.

### Using Map Literals
You can declare and populate a map in a single step:
```go
marks := map[string]float64{
    "Math":    90.5,
    "Science": 85.0,
    "English": 92.0,
}
```

---

## 2. Basic Operations

| Operation | Syntax | Description |
| :--- | :--- | :--- |
| **Insert/Update** | `myMap[key] = value` | Adds a new key-value pair or updates the value if the key already exists. |
| **Delete** | `delete(myMap, key)` | Removes the key and its value from the map. If the key doesn't exist, this is a safe no-op. |
| **Clear** | `clear(myMap)` | Removes all entries from the map, leaving it empty but initialized. |
| **Length** | `len(myMap)` | Returns the number of active key-value pairs in the map. |

---

## 3. Important Map Concepts

### A. Missing Keys & The Zero-Value Behavior
If you request a key that does not exist in a map, Go does **not** throw an error. Instead, it returns the **zero value** of the map's value type.
```go
score := marks["History"] // "History" is missing, returns 0.0 (float64 zero value)
```

### B. The Comma-Ok Idiom
Because requesting a missing key returns a zero value, you need a way to check if a key actually exists or is just missing. Use the **comma-ok** syntax:
```go
value, ok := marks["Math"]

if ok {
    fmt.Println("Math mark exists:", value)
} else {
    fmt.Println("Math mark is not present.")
}
```
* `value`: The value associated with the key (or the zero value if missing).
* `ok`: A boolean that is `true` if the key exists in the map, and `false` if it is absent.

### C. Map Key Restrictions
Not all types can be used as keys in a Go map.
* **Allowed Keys**: Any type that is comparable using the comparison operator `==`. This includes integers, floating-point numbers, booleans, strings, interfaces, pointers, channels, and structs (if all struct fields are comparable).
* **Disallowed Keys**: Slices, maps, and functions cannot be map keys because they are not comparable.

### D. Reference Types & Unordered Iteration
* **Reference Type**: Maps are reference types. Passing a map to a function or assigning it to another variable copies the reference. Modifying the map inside a function **will** affect the original map.
* **Unordered**: Iterating over a map using `for range` returns key-value pairs in a randomized/unordered sequence. Go intentionally randomizes map iteration to prevent developers from relying on a specific order.

### E. Map Comparison
To check if two maps contain identical key-value pairs, use the `maps` package:
```go
import "maps"

areEqual := maps.Equal(map1, map2) // returns true/false
```

---

## How to Run the Code in this Directory

To execute the example program demonstrating maps, run:

```bash
go run main.go
```

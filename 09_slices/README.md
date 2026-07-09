# Slices in Go

A **Slice** is a dynamically sized, flexible view into the elements of an array. Slices are far more common and powerful than arrays in Go programming.

---

## 1. Anatomy of a Slice

Under the hood, a slice does not store any data itself. Instead, it is a descriptor for a contiguous segment of an underlying array. A slice consists of three components:

1. **Pointer**: A reference to the first element of the underlying array that the slice accesses.
2. **Length (`len`)**: The number of elements currently present in the slice.
3. **Capacity (`cap`)**: The total number of elements in the underlying array, counting from the first element of the slice.

---

## 2. Creating Slices

Go provides three main ways to create slices:

### A. Slice Literals
Declared like an array literal, but without specifying a fixed size:
```go
s := []int{1, 2, 3} // Length: 3, Capacity: 3
```

### B. Using the `make()` Function
The built-in `make` function allocates an underlying array and returns a slice referencing it.
```go
s := make([]int, 5)       // len=5, cap=5 (all initialized to zero)
s2 := make([]int, 3, 10)  // len=3, cap=10
```

### C. Slicing an Array or Existing Slice
Created by specifying a half-open range `[low : high]` (includes element at `low`, excludes element at `high`):
```go
arr := [5]int{10, 20, 30, 40, 50}
s := arr[1:4] // s will contain [20, 30, 40]
```

---

## 3. Modifying Slices & Capacity Growth

### Dynamic Resizing with `append()`
You can add elements to a slice using the built-in `append` function:
```go
var numbers []int
numbers = append(numbers, 1) // automatically grows
```
When a slice exceeds its current capacity during an `append`, Go:
1. Allocates a new underlying array with larger capacity (typically doubling the size).
2. Copies the existing elements into the new array.
3. Appends the new value.
4. Returns the new slice referencing this new array.

---

## 4. Crucial Slice Gotchas

### A. The Copy Gotcha
The built-in `copy(dst, src)` function copies elements from a source slice to a destination slice. However:
> [!IMPORTANT]
> The number of elements copied is the **minimum** of the two slices' lengths.
* If the destination slice has a length of `0`, **no elements will be copied**.
```go
src := []int{1, 2, 3}
dst := make([]int, 0, 5) // len=0, cap=5

copy(dst, src) // Copies 0 elements! dst remains []
```
* **Solution**: Ensure the destination slice has a length matching the source:
```go
dst := make([]int, len(src))
copy(dst, src) // Copies all elements successfully
```

### B. Shared Underlying Array
Modifying elements of a slice modifies the underlying array, which affects any other slices sharing that same array.
```go
arr := [3]int{1, 2, 3}
s1 := arr[:]
s2 := arr[:]

s1[0] = 99 // Both s2[0] and arr[0] will now be 99!
```

---

## 5. Slice Operations

### Comparing Slices
Unlike arrays, slices cannot be compared directly using `==` or `!=` (you can only compare a slice to `nil`). To compare two slices, use the standard library `slices` package:
```go
import "slices"

isEqual := slices.Equal(slice1, slice2) // returns true/false
```

---

## How to Run the Code in this Directory

To execute the example program demonstrating slices, run:

```bash
go run main.go
```

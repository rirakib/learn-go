# Arrays in Go

An **Array** is a numbered sequence of elements of a single type with a **fixed size**. Once defined, an array's length cannot be resized.

---

## 1. Array Declarations and Syntax

### Standard Declaration (Zero-Valued)
Declares an array of size `5` initialized with zero values:
```go
var numbers [5]int // Outputs: [0 0 0 0 0]
```

### Initializing with Values (Array Literals)
```go
var fruits = [4]string{"Apple", "Banana", "Mango", "Orange"}
```

### Inferred Size (`...`)
If you let the compiler count the number of elements automatically, use the ellipsis `...` operator:
```go
numbers := [...]int{10, 20, 30} // size is automatically set to 3
```

### Specifying Values at Particular Indices
You can initialize specific elements while leaving others as zero values:
```go
numbers := [5]int{1: 99, 4: 100} // Outputs: [0 99 0 0 100]
```

---

## 2. Key Characteristics of Go Arrays

### Fixed-Size Constraints
An array's size is a fundamental part of its data type. This means:
* `[5]int` and `[10]int` are entirely different types.
* You cannot assign a `[5]int` to a variable of type `[10]int`.
* You cannot resize an array (for dynamically sized arrays, use **Slices**).

### Array Copy Behavior (Value Types)
> [!IMPORTANT]
> Unlike arrays in languages like Java or JavaScript, which are reference types, Go arrays are **value types**.
* If you assign one array to another, Go **copies all elements** instead of creating a pointer.
* If you pass an array to a function, the function receives a copy of the array. Modifications inside the function will **not** affect the original array.
```go
arr1 := [3]int{1, 2, 3}
arr2 := arr1 // Creates a completely new copy
arr2[0] = 99 // arr1 remains [1, 2, 3]
```

---

## 3. Advanced Array Structures

### Multi-Dimensional Arrays
You can nest arrays to create matrices:
```go
var matrix [2][3]int = [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}
```

### Mixed Type Arrays (Interface Arrays)
By utilizing the empty interface (`interface{}` or `any`), you can create a fixed-size array that stores different types:
```go
var mixedArray [3]interface{} = [3]interface{}{1, "Hello", 3.14}
```

---

## How to Run the Code in this Directory

To execute the example program demonstrating arrays, run:

```bash
go run main.go
```

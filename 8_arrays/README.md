# Arrays in Go

This directory covers **Arrays** in Go. In Go, an array is a numbered sequence of elements of a single type with a **fixed size**. Once declared, their size cannot be changed.

## Code Explanation

```go
package main 

import "fmt"

func main() {

	var numbers[5]int = [5]int{}

	numbers[3] = 99

	fmt.Println(numbers)


	var fruites[4]string = [4]string{"Apple", "Banana", "Mango", "Orange"}

	fmt.Println(fruites)


	var marks[5]float64 = [5]float64{75.5, 80.0, 90.0, 85.0, 95.0}

	fmt.Println(marks)

	var matrix[2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println(matrix)

	var mixedArray[3]interface{} = [3]interface{}{1, "Hello", 3.14}

	fmt.Println(mixedArray)

}
```

### Breakdown of the Code:

1. **`var numbers[5]int = [5]int{}`**
   - Declares an integer array of size `5`.
   - By default, arrays in Go are initialized with the **zero-value** of their elements (for `int`, this is `0`).
   - `numbers[3] = 99` sets the element at index `3` (the 4th element) to `99`.
   - Printing this array outputs: `[0 0 0 99 0]`.

2. **`var fruites[4]string = [4]string{"Apple", "Banana", "Mango", "Orange"}`**
   - Declares and initializes a string array of size `4` with initial values.

3. **`var marks[5]float64 = [5]float64{75.5, 80.0, 90.0, 85.0, 95.0}`**
   - Declares and initializes a float64 array of size `5`.

4. **`var matrix[2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}`**
   - Declares a **multi-dimensional** (2D) array.
   - It represents a matrix of 2 rows and 3 columns.

5. **`var mixedArray[3]interface{} = [3]interface{}{1, "Hello", 3.14}`**
   - Go is statically typed, meaning standard arrays cannot hold different types of elements.
   - However, by using the empty interface `interface{}` (or `any` in modern Go), you can create an array that holds values of **any type** (e.g., `int`, `string`, `float64`).

---

## Key Learning Takeaways

* **Fixed Size**: The size of an array is part of its type. `[5]int` and `[10]int` are distinct types. You cannot resize arrays.
* **Zero Values**: Uninitialized array elements are automatically set to their type's zero-value (`0` for numeric, `""` for string, `false` for bool, `nil` for interfaces/pointers).
* **Index Access**: Access elements using brackets `array[index]`. Indexes in Go are 0-indexed (starting at `0` to `len(array) - 1`).
* **Empty Interface (`interface{}`)**: Using an array of `interface{}` allows storing mixed types, but requires type assertions or type switches to retrieve the original types safely in more complex programs.

---

## How to Run

To run this example, navigate to this directory in your terminal and execute:

```bash
go run main.go
```

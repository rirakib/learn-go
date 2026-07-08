# Slices in Go

This directory covers **Slices** in Go. Slices are a key data structure in Go, providing a dynamic, flexible, and powerful wrapper around arrays. Unlike arrays, slices can grow and shrink dynamically.

## Code Explanation

```go
package main

import (
	"fmt"
	"slices"
)

func main() {

	var num = make([]int, 0, 5)
	var num2 = make([]int, 0, 5)


	for i := 0; i < 20; i++ {
		num = append(num, i)
	}

	fmt.Println(num)
	fmt.Println(len(num))
	fmt.Println(cap(num))

	fmt.Println(num[5:10])
	fmt.Println(num[:10])

	copy(num2, num)

	fmt.Println(num2)

	var test1 = []int{1, 2, 3, 4, 5}
	var test2 = []int{1, 2, 3, 4, 5}

	fmt.Println(slices.Equal(test1,test2))
}
```

### Breakdown of the Code:

1. **`var num = make([]int, 0, 5)`**
   - Declares a slice `num` of type `[]int`.
   - Uses the built-in `make` function: `make([]type, length, capacity)`.
   - Here, the length is set to `0`, and the initial capacity is set to `5`.

2. **Appending Values**:
   ```go
   for i := 0; i < 20; i++ {
       num = append(num, i)
   }
   ```
   - Uses the built-in `append` function to add numbers `0` through `19` to the slice.
   - Even though the initial capacity was `5`, `append` automatically grows the slice's underlying array when it runs out of space.

3. **Length and Capacity**:
   - `len(num)` returns the number of elements in the slice (which is `20`).
   - `cap(num)` returns the capacity of the slice (which represents the size of the underlying array). Because of the automatic growth factor, capacity will have doubled multiple times and will be `40` (or similar depending on Go's internal slice allocation logic).

4. **Slicing Expressions**:
   - `num[5:10]` creates a sub-slice containing elements from index `5` up to (but excluding) index `10`: `[5 6 7 8 9]`.
   - `num[:10]` creates a sub-slice from the beginning of the slice up to index `10`: `[0 1 2 3 4 5 6 7 8 9]`.

5. **The Copy Gotcha**:
   ```go
   copy(num2, num)
   fmt.Println(num2)
   ```
   - > [!IMPORTANT]
     > **Go Gotcha**: The built-in `copy(dst, src)` copies elements from the source slice to the destination slice. However, it only copies up to the size of the destination's **length**, not its capacity.
   - Since `num2` was initialized with a length of `0` (`make([]int, 0, 5)`), `copy` copies `0` elements!
   - Printing `num2` outputs `[]` (an empty slice). To copy successfully, `num2` must be created with a matching length, e.g., `num2 := make([]int, len(num))`.

6. **Slice Comparison**:
   - `slices.Equal(test1, test2)` uses the modern `slices` standard library package to compare two slices element-by-element and returns a boolean (`true` if they are equal).

---

## Key Learning Takeaways

* **Slices vs Arrays**: Slices do not have a fixed size specified in their type declaration (e.g., `[]int` vs `[5]int`).
* **Length vs Capacity**:
  - **Length**: The number of elements currently in the slice.
  - **Capacity**: The maximum number of elements the slice can hold before the underlying array must be reallocated/resized.
* **`copy()` Behavior**: Always ensure your destination slice has a length greater than `0` (typically matching the source's length) before using `copy()`.

---

## How to Run

To run this example, navigate to this directory in your terminal and execute:

```bash
go run main.go
```

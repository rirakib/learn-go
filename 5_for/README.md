# Loops in Go (For Loop)

This directory covers loops in Go. Go has only one looping construct: the `for` loop. The example demonstrates different ways of writing loops and controlling their execution using `break` and `continue`.

## Code Explanation

```go
package main 

import "fmt"

func main() {

	// i := 1

	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// for i = range 5 {
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 10; i++ {
		
	// 	if i == 5 {
	// 		break
	// 	}
		
	// 	fmt.Println(i)
	// }


	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
```

### Breakdown of the Code:

1. **Commented Loop 1: While-like Loop**
   ```go
   // i := 1
   // for i <= 10 {
   //     fmt.Println(i)
   //     i = i + 1
   // }
   ```
   - Go does not have a `while` loop. Instead, you can write a `for` loop with only a condition to achieve the same result.

2. **Commented Loop 2: Range Loop (Go 1.22+)**
   ```go
   // for i = range 5 {
   //     fmt.Println(i)
   // }
   ```
   - Ranges over an integer. This executes the loop 5 times, with `i` taking values from `0` to `4`.

3. **Commented Loop 3: Classic Loop with `break`**
   ```go
   // for i := 0; i < 10; i++ {
   //     if i == 5 {
   //         break
   //     }
   //     fmt.Println(i)
   // }
   ```
   - Demonstrates the standard C-style loop: initialization (`i := 0`), condition (`i < 10`), and post statement (`i++`).
   - Uses `break` to exit the loop early when `i` reaches `5`.

4. **Active Loop: Classic Loop with `continue`**
   ```go
   for i := 0; i < 10; i++ {
       if i == 5 {
           continue
       }
       fmt.Println(i)
   }
   ```
   - Loops from `0` to `9`.
   - When `i` equals `5`, the `continue` statement is triggered. This skips the rest of the current iteration (so `5` is not printed) and jumps straight to the next iteration (`i++`).

---

## Key Learning Takeaways

* **Single Loop Construct**: Go uses only `for` for all looping structures.
* **Flow Control**:
  - `break`: Exits the loop immediately.
  - `continue`: Skips the rest of the current iteration and goes to the next loop iteration.

---

## How to Run

To run this example, navigate to this directory in your terminal and execute:

```bash
go run main.go
```

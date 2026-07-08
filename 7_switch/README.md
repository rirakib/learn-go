# Switch Statements in Go

This directory covers `switch` statements in Go. Switch statements are a cleaner way to express conditional logic with multiple branches. Go's switch statements are highly expressive and support evaluating conditions in cases, comma-separated multiple matches, and default fallbacks.

## Code Explanation

```go
package main 

import "fmt"
import "time"


func workingDay(day string) string {

	fmt.Println(day)

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("It's a weekend.")
	default:
		fmt.Println("It's a working day.")
		return "working day"
	}
	return "weekend"
}

func main() {

	var marks float64 = 75

	switch {
	case marks >= 80 :
		fmt.Println("Grade: A+")
	case marks >= 60 , marks < 80 :
		fmt.Println("Grade: A")

	case marks >= 50 , marks < 60 :
		fmt.Println("Grade: B")

	case marks >= 40 , marks < 50 :
		fmt.Println("Grade: D")	
	default:
		fmt.Println("Grade: F")
	}


	var day string = time.Now().Weekday().String() 
	workingDay(day)

}
```

### Breakdown of the Code:

1. **`workingDay` Function**:
   - Accepts a `day` string parameter.
   - **`switch day`**: Evaluates the value of `day`.
   - **`case "Saturday", "Sunday":`**: Demonstrates matching multiple values in a single case statement by separating them with commas. If `day` is either `"Saturday"` or `"Sunday"`, the code block runs.
   - **`default:`**: If none of the cases match, the default block runs.
   - **No Explicit `break`**: In Go, switch cases break automatically. You do not need to add a `break` at the end of each case block.

2. **`main` Function (Expressionless Switch)**:
   - **`switch` with no expression**: A switch without an expression is equivalent to `switch true`. It acts as a clean alternative to a long `if`/`else if` chain.
   - Each `case` is evaluated as a boolean condition. The first case that evaluates to `true` is executed.
   - **`case marks >= 60 , marks < 80 :`**: Comma separated conditions in an expressionless switch act as logical **OR**. In this case, if `marks >= 60` OR `marks < 80` is true, the case executes. Since `75` is greater than `60` (making the first expression `true`), this case runs and outputs `"Grade: A"`.
   - **`time.Now().Weekday().String()`**: Gets the current day name as a string (e.g., `"Wednesday"`) and passes it to `workingDay` to check whether it's a weekday or weekend.

---

## Key Learning Takeaways

* **Implicit Break**: Go automatically breaks out of a switch block after executing a matching case. If you want to fall through to the next case, you must explicitly use the `fallthrough` keyword.
* **Multiple Values**: Multiple values can be matched in a single case by separating them with commas (e.g., `case "Saturday", "Sunday":`).
* **Expressionless Switch**: Omitting the variable to switch on (`switch { ... }`) allows you to write conditional expressions inside each case, offering a highly readable alternative to nested `if`/`else` structures.

---

## How to Run

To run this example, navigate to this directory in your terminal and execute:

```bash
go run main.go
```

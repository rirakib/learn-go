# Maps in Go

This directory covers **Maps** in Go. A map is Go's built-in associative data type (sometimes called a hash table, dictionary, or lookup table) that associates keys with values.

## Code Explanation

```go
package main

import (
	"fmt"
	"maps"
)

func main() {

	data := make(map[string]string)

	data["name"] = "John Doe"
	data["email"] = "john@gmail.com"
	data["phone"] = "123-456-7890"

	delete(data, "phone")
	fmt.Println("Name:", data["name"], "Email:", data["email"], "Phone:", data["phone"], "Address ", data["address"])

	clear(data)
	fmt.Println(data)

	mark := map[string]float64{
		"Math":    90.5,
		"Science": 85.0,
		"English": 92.0,
	}

	fmt.Println("Math mark:", mark["Math"])

	value, ok := mark["Math"]

	fmt.Println(value)

	if ok {
		fmt.Println("Subject mark is present.")
	} else {
		fmt.Println("Subject mark is not present.")
	}


	test1 := map[string]int{
		"one":   1,"two":   2,
		"three": 3,"four":  4,
		"five":  5,"six":   6,
	}

	test2 := map[string]int{
		"one":   1,"two":   2,
		"three": 3,"four":  4,
		"five":  5,"six":   6,
	}

	fmt.Println(maps.Equal(test1,test2))

}
```

### Breakdown of the Code:

1. **`data := make(map[string]string)`**
   - Creates an empty map using the built-in `make` function.
   - The map structure is defined as `map[keyType]valueType`. In this case, both keys and values are strings.

2. **Basic Operations**:
   - **Adding/Updating**: `data["name"] = "John Doe"` assigns a value to a key.
   - **Deleting**: `delete(data, "phone")` removes the key `"phone"` and its associated value from the map.
   - **Accessing non-existing key**: `data["address"]` returns the value's zero-value (an empty string `""` for string values) instead of throwing an error.

3. **`clear(data)`**:
   - Removes all key-value pairs from the map `data`, leaving it empty but initialized.

4. **Map Literals**:
   ```go
   mark := map[string]float64{
       "Math":    90.5,
       "Science": 85.0,
       "English": 92.0,
   }
   ```
   - Declares and initializes a map inline using key-value pairs.

5. **The Comma-Ok Idiom**:
   ```go
   value, ok := mark["Math"]
   ```
   - Used to distinguish between "a key that is not present in the map" and "a key that is present but has a zero-value".
   - `value` receives the value at that key (if present) or the zero-value (if not present).
   - `ok` is a boolean that is `true` if the key exists in the map, and `false` otherwise.

6. **Map Comparison**:
   - `maps.Equal(test1, test2)` uses the `maps` standard library package to verify if two maps contain identical key-value pairs.

---

## Key Learning Takeaways

* **Key-Value Pairs**: Uniquely maps keys to values.
* **Comma-Ok Idiom**: Always use the `value, ok := myMap[key]` pattern when you need to safely check if a key actually exists in a map.
* **Zero Values**: Accessing a missing key returns the zero-value of the map's value type, not an error.
* **Unordered**: Maps are inherently unordered collections. Iterating over them does not guarantee a specific order of keys.

---

## How to Run

To run this example, navigate to this directory in your terminal and execute:

```bash
go run main.go
```

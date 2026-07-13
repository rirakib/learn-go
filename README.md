# Learn Go

Welcome to the Go learning repository! This project contains step-by-step examples designed to help you learn Go programming from scratch.

Each directory focuses on a specific Go language concept, containing a working `main.go` file and a detailed explanatory `README.md` file.

## Table of Contents

Below is the list of topics covered in this repository. Click on any topic to view its detailed explanation:

| # | Topic | Folder | Description |
|---|-------|--------|-------------|
| 1 | [Hello World](./01_hello_world/README.md) | [01_hello_world](./01_hello_world) | The entry point of Go programs, basic syntax, and printing. |
| 2 | [Simple Values](./02_simple_values/README.md) | [02_simple_values](./02_simple_values) | Standard types like integers, strings, booleans, and floats. |
| 3 | [Variables](./03_variables/README.md) | [03_variables](./03_variables) | Standard variable declaration with `var` and short declaration (`:=`). |
| 4 | [Constants](./04_constant/README.md) | [04_constant](./04_constant) | Defining immutable constants and basic operations. |
| 5 | [Loops](./05_for/README.md) | [05_for](./05_for) | Go's `for` loop structures, including loops, while-style loops, and controls. |
| 6 | [Conditionals](./06_if_else/README.md) | [06_if_else](./06_if_else) | Branching logical decisions using `if`, `else if`, and `else`. |
| 7 | [Switch Statements](./07_switch/README.md) | [07_switch](./07_switch) | Multi-way conditionals, expressionless switches, and case matching. |
| 8 | [Arrays](./08_arrays/README.md) | [08_arrays](./08_arrays) | Working with fixed-size sequences and basic indexing. |
| 9 | [Slices](./09_slices/README.md) | [09_slices](./09_slices) | Dynamic wrappers around arrays, growth mechanics, copying, and slicing. |
| 10 | [Maps](./10_maps/README.md) | [10_maps](./10_maps) | Standard key-value associative arrays, deletion, and existence checks. |
| 11 | [Range](./11_range/README.md) | [11_range](./11_range) | Iterating over slices, maps, strings, and ranges. |
| 12 | [Functions](./12_functions/README.md) | [12_functions](./12_functions) | Defining functions, parameters, return types, and multiple return values. |
| 13 | [Variadic Functions](./13_variadic_function/README.md) | [13_variadic_function](./13_variadic_function) | Declaring and calling functions with a variable number of parameters, and slice unpacking. |
| 14 | [Closures](./14_closures/README.md) | [14_closures](./14_closures) | Anonymous functions that can capture and reference variables from outer scopes. |
| 15 | [Pointers](./15_pointers/README.md) | [15_pointers](./15_pointers) | Understanding memory addresses, pointer dereferencing, and passing by reference. |
| 16 | [Structs](./16_structs/README.md) | [16_structs](./16_structs) | Creating custom grouped types, fields, methods, and constructor patterns. |
| 17 | [Todo With Structs](./17_todo_with_structs/README.md) | [17_todo_with_structs](./17_todo_with_structs) | A simple console-based to-do list application using struct types. |
| 18 | [Embedding](./18_embedding/README.md) | [18_embedding](./18_embedding) | Composition in Go by embedding structures within other structures. |
| 19 | [Interfaces](./19_interfaces/README.md) | [19_interfaces](./19_interfaces) | Contract-based programming, polymorphism, and implicit interface implementation. |
| 20 | [Student CRUD](./20_studentcrud/README.md) | [20_studentcrud](./20_studentcrud) | An in-memory CRUD application to manage student data using structs and slices. |
| 21 | [Enums](./21_enums/README.md) | [21_enums](./21_enums) | Implementing enum patterns in Go using custom types and iota constants. |
| 22 | [Generics](./22_generics/README.md) | [22_generics](./22_generics) | Writing type-independent functions and structures using type parameters. |
| 23 | [Goroutines](./23_goroutines/README.md) | [23_goroutines](./23_goroutines) | Introducing Go's lightweight concurrency execution model. |
| 24 | [WaitGroup](./24_waitgroup/README.md) | [24_waitgroup](./24_waitgroup) | Using `sync.WaitGroup` to coordinate and synchronize multiple running goroutines. |
| 25 | [Exercises](./25_exercies/README.md) | [25_exercies](./25_exercies) | Practical projects including Student Report and Employee Management Systems. |

---

## How to Run Examples

To run the code in any directory:

1. Open your terminal.
2. Navigate to the desired folder, for example:
   ```bash
   cd 01_hello_world
   ```
3. Run the Go file:
   ```bash
   go run main.go
   ```

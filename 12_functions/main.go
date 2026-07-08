package main

import (
	"fmt"
)


type MathOp func(int, int) float64


func Add(a, b int) float64      { return float64(a + b) }
func Subtract(a, b int) float64 { return float64(a - b) }
func Multiply(a, b int) float64 { return float64(a * b) }

func Divide(a, b int) float64 {
	if b == 0 {
		return 0
	}
	return float64(a) / float64(b)
}


func Calculate(a, b int, op MathOp) float64 {
	return op(a, b)
}


var registry = map[string]MathOp{
	"add":      Add,
	"subtract": Subtract,
	"multiply": Multiply,
	"divide":   Divide,
}

func Process(opName string, a, b int) (float64, error) {
	op, exists := registry[opName]
	if !exists {
		return 0, fmt.Errorf("unsupported operation: %s", opName)
	}
	return op(a, b), nil
}

func main() {
	fmt.Println("--- Strategy Pattern ---")
	resultAdd := Calculate(10, 5, Add)
	fmt.Printf("Add (10, 5) = %.1f\n", resultAdd)

	resultMod := Calculate(10, 3, func(a, b int) float64 {
		return float64(a % b)
	})
	fmt.Printf("Modulo (10, 3) = %.1f\n", resultMod)

	fmt.Println("\n--- Registry Pattern ---")
	if resultMult, err := Process("multiply", 10, 5); err == nil {
		fmt.Printf("Multiply (10, 5) = %.1f\n", resultMult)
	}

	registry["modulo"] = func(a, b int) float64 { return float64(a % b) }

	if resultModReg, err := Process("modulo", 10, 3); err == nil {
		fmt.Printf("Registered Modulo (10, 3) = %.1f\n", resultModReg)
	}
}
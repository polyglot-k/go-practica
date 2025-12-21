package main

import (
	"fmt"
	"log"
	"os"

	calc "github.com/polyglot-k/go-practica/calculator/internal"
)

func main() {
	var (
		a, b float64
		op   string
	)

	fmt.Print("operand1: ")
	if _, err := fmt.Scan(&a); err != nil {
		log.Fatalf("error: invalid operand1: %v", err)
	}
	
	fmt.Print("operand2: ")
	if _, err := fmt.Scan(&b); err != nil {
		log.Fatalf("error: invalid operand12: %v", err)
	}

	fmt.Print("operator (+ - * /): ")
	if _, err := fmt.Scan(&op); err != nil {
		log.Fatalf("error: invalid operator: %v", err)
	}

	operator, err := calc.ParseOperator(op)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	result, err := calc.Calculate(calc.Operand(a), calc.Operand(b), operator)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("result:", result)
}
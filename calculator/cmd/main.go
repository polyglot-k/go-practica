package main

import (
	"fmt"
	"log"

	calc "github.com/polyglot-k/go-practica/calculator/internal"
)

func main() {
	var (
		a, b float64
		op   string
	)

	fmt.Print("operand1: ")
	fmt.Scan(&a)

	fmt.Print("operand2: ")
	fmt.Scan(&b)

	fmt.Print("operator (+ - * /): ")
	fmt.Scan(&op)

	operator, err := calc.ParseOperator(op)
	if err != nil {
		log.Fatal(err)
	}

	result, err := calc.Calculate(calc.Operand(a), calc.Operand(b), operator)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result:", result)
}
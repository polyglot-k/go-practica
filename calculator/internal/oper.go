package internal

import "fmt"

type Operand float64

type Operator int

const (
	Add Operator = iota
	Sub
	Mul
	Div
)

func ParseOperator(op string) (Operator, error) {
	switch op {
	case "+":
		return Add, nil
	case "-":
		return Sub, nil
	case "*":
		return Mul, nil
	case "/":
		return Div, nil
	default:
		return 0, fmt.Errorf("invalid operator: %s", op)
	}
}

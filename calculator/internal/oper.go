package internal

import "fmt"

type Operand float64

type Operator int

const (
	ADD Operator = iota
	SUB
	MUL
	DIV
)

// ParseOperator는 문자열로 된 연산자를 Operator 타입으로 변환합니다.
func ParseOperator(op string) (Operator, error) {
	switch op {
	case "+":
		return ADD, nil
	case "-":
		return SUB, nil
	case "*":
		return MUL, nil
	case "/":
		return DIV, nil
	default:
		return 0, fmt.Errorf("invalid operator: %s", op)
	}
}

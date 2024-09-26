package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Operator int

const (
	ADD Operator = iota
	SUB
)

type Expression interface {
	interpret() int
}

type NumberExpression struct {
	value int
}

func (n *NumberExpression) interpret() int {
	return n.value
}

type AdditionExpression struct {
	left  Expression
	right Expression
}

func (a *AdditionExpression) interpret() int {
	return a.left.interpret() + a.right.interpret()
}

type SubtractionExpression struct {
	left  Expression
	right Expression
}

func (s *SubtractionExpression) interpret() int {
	return s.left.interpret() - s.right.interpret()
}

func Parser(input string) Expression {
	tokens := strings.Split(input, " ")

	stack := make([]Expression, 0)
	opStack := make([]Operator, 0)
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" {
			opStack = append(opStack, ADD)
		} else if tokens[i] == "-" {
			opStack = append(opStack, SUB)
		} else {
			value, err := strconv.Atoi(tokens[i])
			if err != nil {
				panic(err)
			}
			stack = append(stack, &NumberExpression{value: value})

			for len(opStack) > 0 && precedence(opStack[len(opStack)-1]) >= precedence(ADD) {
				op := opStack[len(opStack)-1]
				opStack = opStack[:len(opStack)-1]
				right := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				left := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				var expr Expression
				if op == ADD {
					expr = &AdditionExpression{left: left, right: right}
				} else {
					expr = &SubtractionExpression{left: left, right: right}
				}
				stack = append(stack, expr)
			}
		}
	}

	for len(opStack) > 0 {
		op := opStack[len(opStack)-1]
		opStack = opStack[:len(opStack)-1]
		right := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		left := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		var expr Expression
		if op == ADD {
			expr = &AdditionExpression{left: left, right: right}
		} else {
			expr = &SubtractionExpression{left: left, right: right}
		}
		stack = append(stack, expr)
	}

	return stack[0]
}

func precedence(op Operator) int {
	switch op {
	case ADD:
		return 1
	case SUB:
		return 1
	default:
		return 0
	}
}

func main() {
	expression := Parser("7 + 3 - 4 - 1")
	result := expression.interpret()
	fmt.Println("Result:", result)
}

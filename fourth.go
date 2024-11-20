package main

import (
	"awesomeProject1/calculator"
	"awesomeProject1/stack"
	"fmt"
)

func fourth() {

	stk1 := stack.Stk{}

	stk1.Push(22)
	stk1.Push(21)
	stk1.Push(28)

	fmt.Println(stk1.Pop())
	fmt.Println(stk1.Pop())
	fmt.Println(stk1.Pop())
	fmt.Println(stk1.Pop())

	exp := calculator.Calculate{}
	var op1 float64
	var op2 float64

	exp.SetMemory(12, 4)

	op1, op2 = exp.GetMemory()
	fmt.Println("the result by setting operand  is operand 1 = ", op1, " and operand 2 is = ", op2)

	exp.AddtoMemory(12, 54)
	op1, op2 = exp.GetMemory()
	fmt.Println("the result by setting operand  is operand 1 = ", op1, " and operand 2 is = ", op2)

	exp.SubtracttoMemory(12, 54)
	op1, op2 = exp.GetMemory()
	fmt.Println("the result by setting operand  is operand 1 = ", op1, " and operand 2 is = ", op2)

	op1, op2 = exp.GetMemory()
	fmt.Println(op1, "+", op2, "=", exp.Add())
	fmt.Println(op1, "-", op2, "=", exp.Subtract())
	fmt.Println(op1, "*", op2, "=", exp.Product())

	val, er := exp.Division()
	if er == nil {

		fmt.Println(op1, "/", op2, "=", val)
	} else {
		fmt.Println("divsion by zero !!")
	}

	exp.ResetMemmory()

}

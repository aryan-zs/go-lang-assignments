package calculator

import "fmt"

// this structure type is  contains the two variable  i.e the two operands operand1 and operand2
type Calculate struct {
	operand1 float64
	operand2 float64
}

// this function is reseting the two operands  to their default value that is zero for integer and float
func (c *Calculate) ResetMemmory() {

	c.operand1 = 0.0
	c.operand2 = 0.0

}

// this function is  setting the new value to the two operand  as per the parameter passed
func (c *Calculate) SetMemory(a, b float64) {

	c.operand1 = a
	c.operand2 = b

}

// this function is  getting the value of the  two operand  that are present in the stucture
func (c *Calculate) GetMemory() (float64, float64) {

	return c.operand1, c.operand2

}

// this is adding the passed value to the current value  fo the operands
func (c *Calculate) AddtoMemory(a, b float64) {

	c.operand1 += a
	c.operand2 += b

}

// this is the substraction function that is subtracting the passed parameter  from two operands
func (c *Calculate) SubtracttoMemory(a, b float64) {

	c.operand1 -= a
	c.operand2 -= b

}

// does addtion
func (c Calculate) Add() float64 {

	return c.operand1 + c.operand2

}

// does substraction
func (c Calculate) Subtract() float64 {

	return c.operand1 - c.operand2

}

// does muliplication
func (c Calculate) Product() float64 {

	return c.operand1 * c.operand2

}

// does Division
func (c Calculate) Division() (float64, error) {
	if c.operand2 == 0 {
		return -1, fmt.Errorf("the dividend is zero")
	}
	return c.operand1 / c.operand2, nil

}

package calculator

import "fmt"

type Calculate struct {
	operand1 float64
	operand2 float64
}

func (c *Calculate) ResetMemmory() {

	c.operand1 = 0.0
	c.operand2 = 0.0

}

func (c *Calculate) SetMemory(a, b float64) {

	c.operand1 = a
	c.operand2 = b

}

func (c *Calculate) GetMemory() (float64, float64) {

	return c.operand1, c.operand2

}

func (c *Calculate) AddtoMemory(a, b float64) {

	c.operand1 += a
	c.operand2 += b

}

func (c *Calculate) SubtracttoMemory(a, b float64) {

	c.operand1 -= a
	c.operand2 -= b

}

func (c Calculate) Add() float64 {

	return c.operand1 + c.operand2

}
func (c Calculate) Subtract() float64 {

	return c.operand1 - c.operand2

}
func (c Calculate) Product() float64 {

	return c.operand1 * c.operand2

}
func (c Calculate) Division() (float64, error) {
	if c.operand2 == 0 {
		return -1, fmt.Errorf("the dividend is zero")
	}
	return c.operand1 / c.operand2, nil

}

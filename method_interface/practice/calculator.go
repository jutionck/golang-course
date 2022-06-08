package main

import "fmt"

type CalcOps interface {
	DoCalc() float64
}

type Subtraction struct {
	num1 float64
	num2 float64
}

func (a Subtraction) DoCalc() float64 {
	return a.num1 - a.num2
}

func NewSubtraction(num1, num2 float64) CalcOps {
	return Subtraction{
		num1,
		num2,
	}
}

type Addition struct {
	num1 float64
	num2 float64
}

func (a Addition) DoCalc() float64 {
	return a.num1 + a.num2
}

func NewAddition(num1, num2 float64) CalcOps {
	return Addition{
		num1,
		num2,
	}
}

type Calculator struct {
	ops CalcOps
}

func (c Calculator) result() {
	fmt.Println(c.ops.DoCalc())
}

func NewCalculator(ops CalcOps) Calculator {
	return Calculator{ops}
}

func main() {
	c := NewCalculator(NewAddition(1, 2))
	c.result()
	c = NewCalculator(NewSubtraction(1, 2))
	c.result()
}

package main

import (
	"fmt"
	"math"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
*/

func main() {
	a := math.Pow(2, 21)
	b := math.Pow(3, 22)
	bigA := big.NewInt(int64(a))
	bigB := big.NewInt(int64(b))

	fmt.Println(multiply(bigA, bigB))
	fmt.Println(div(bigA, bigB))
	fmt.Println(sum(bigA, bigB))
	fmt.Println(sub(bigA, bigB))

}
func sum(a, b *big.Int) *big.Int {
	res := &big.Int{}
	res.Add(a, b)
	return res
}
func sub(a, b *big.Int) *big.Int {
	res := &big.Int{}
	res.Sub(a, b)
	return res
}
func multiply(a, b *big.Int) *big.Int {
	res := &big.Int{}
	res.Mul(a, b)
	return res
}
func div(a, b *big.Int) *big.Int {
	res := &big.Int{}
	res.Div(b, a)
	return res
}

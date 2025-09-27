package main

import (
	"fmt"
	"math/big"
)

func main() {

	a := 2000000
	b := 3000000

	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("a + b = %d\n", a + b)
	fmt.Printf("a - b = %d\n", a - b)
	fmt.Printf("a * b = %d\n", a * b)
	fmt.Printf("a / b = %d\n", a / b)

	// Числа больше чем int64
	
	A, ok := new(big.Int).SetString("123456789123456789123456789", 10)
	if !ok {
    panic("ошибка при конвертации строки в big.Int")
	}

	B, ok := new(big.Int).SetString("123456789123456789123456789", 10)
	if !ok {
    panic("ошибка при конвертации строки в big.Int")
	}
	

	sum := new(big.Int).Add(A, B)
	diff := new(big.Int).Sub(A, B)
	mul := new(big.Int).Mul(A, B)
	div := new(big.Int).Div(A, B)

	fmt.Println("a =", A)
	fmt.Println("b =", B)
	fmt.Println("a + b =", sum)
	fmt.Println("a - b =", diff)
	fmt.Println("a * b =", mul)
	fmt.Println("a / b =", div)	
}

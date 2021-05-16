package main

import (
	"errors"
	"fmt"
)

func main() {
	println("hello,world!")
	fmt.Println(`new print!`)
	testprint()

	a, _ := testdivide(1, 2)
	fmt.Println("a is", a)

	b, _ := testdivide(2, 1)
	fmt.Println("b is", b)

	c, err := testdivide(1, 0)
	if err != nil {
		fmt.Println("err is", err, c)
	}
}

func testprint() {
	println("hello,world!")
	fmt.Println(`new print!`)
}

func testdivide(a int, b int) (int, error) {
	if b == 0 {
		return 9, errors.New("divided by 0")
	}

	return a / b, nil
}

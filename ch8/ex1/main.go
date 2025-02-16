package main

import "fmt"

type IntegerOrFloat interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func DoubleNumber[T IntegerOrFloat](m T) T {
	return 2 * m
}

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type MyInt int

func (m MyInt) String() string {
	return fmt.Sprintf("MyInt: %d", m)
}

type MyFloat float64

func (m MyFloat) String() string {
	return fmt.Sprintf("MyFloat: %.2f", m)
}

func main() {
	fmt.Println(DoubleNumber(MyInt(1)))
	fmt.Println(DoubleNumber(MyFloat(1.0)))
}

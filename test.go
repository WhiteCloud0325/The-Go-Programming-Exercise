package main

import "reflect"

func main() {
	x := 1
	rx := reflect.ValueOf(&x).Elem()
	rx.SetInt(2)               // OK, x = 2
	rx.Set(reflect.ValueOf(3)) // OK, x = 3
	//rx.SetString("hello")            // panic: string is not assignable to int
	//rx.Set(reflect.ValueOf("hello")) // panic: string is not assignable to int

	var y interface{}
	ry := reflect.ValueOf(&y).Elem()
	//ry.SetInt(2)                     // panic: SetInt called on interface Value
	ry.Set(reflect.ValueOf(3))       // OK, y = int(3)
	ry.SetString("hello")            // panic: SetString called on interface Value
	ry.Set(reflect.ValueOf("hello")) // OK, y = "hello"
}

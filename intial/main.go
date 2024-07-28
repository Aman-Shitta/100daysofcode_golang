package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Hello World !")
	// tour based learning
	// dividing each concept in functions
	values()
	variables()
	constants()
}

func function_call() {
	pc, _, _, _ := runtime.Caller(1)
	fmt.Printf("\n[+] Called || %s || at || %v || [+]\n", runtime.FuncForPC(pc).Name(), time.Now())
}

// values

func values() {
	function_call()
	// Go has various value types including
	// strings, integers, floats, booleans, etc.
	// Here are a few basic examples.

	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

// variables

func variables() {
	function_call()

	// In Go, variables are explicitly declared
	// and used by the compiler
	// e.g. check type-correctness of function calls.

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Printf("b, c : %v, %v\n", b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "Apple"
	fmt.Println(f)

}

// constants

const s string = "constant"

func constants() {
	function_call()

	// Go supports constants of
	// character, string, boolean, and numeric values.

	fmt.Println(s)

	const n = 500000000
	fmt.Println(n)

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(n))

	fmt.Println((math.Sin(n)))
}

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
	for_loops()
	if_else()
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

// for

func for_loops() {
	function_call()

	// for is Go’s only looping construct.
	// Here are some basic types of for loops.

	// The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after for loop.
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// Another way of accomplishing the basic “do this N times”
	// iteration is range over an integer.

	for i := range 3 {
		fmt.Println(i)
	}

	// for without a condition will loop repeatedly
	// until you break out of the loop or return from the enclosing function.

	for {
		fmt.Println("loop")
		break
	}

	// You can also continue to the next iteration of the loop.

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

// If/Else

func if_else() {
	function_call()

	// Branching with if and else in Go is straight-forward.

	// a basic example.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// an if statement without an else.

	if 8%2 == 0 {
		fmt.Println("8 is divisible by 0")
	}

	// Logical operators like && and || are often useful in conditions.

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("Either 8 or 7 is even.")
	}

	// A statement can precede conditionals;
	// any variables declared in this statement are available
	// in the current and all subsequent branches.

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

package main

import (
	"fmt"
	"math"
	"runtime"
	"slices"
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
	switch_statement()
	arrays()
	slicesAdvArr()
}

func function_call() {
	// GPT generated *
	// To check Which function is running for certain outputs/
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

// Switch
func switch_statement() {
	function_call()

	// Switch statements express conditionals across many branches.

	n := 2

	switch n {
	case 1:
		fmt.Println(n, "is called One")
	case 2:
		fmt.Println(n, "is called Two")
	case 3:
		fmt.Println(n, "is called Three")
	}

	// Using commas to separate multiple expressions in the same case statement.
	// We use the optional default case in this example as well.

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's a Weekend !")
	default:
		fmt.Println("It's a Weekday")
	}

	// switch without an expression is an alternate way to express if/else logic.
	// Here we also show how the case expressions can be non-constants.
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("It's before Noon")
	default:
		fmt.Println("Its after Noon")
	}

	// A type switch compares types instead of values.
	// You can use this to discover the type of an interface value.
	// In this example, the variable t will have the type corresponding to its clause.

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a boolean")
		case int:
			fmt.Println("I'm an Integer")
		default:
			fmt.Printf("Don't know Type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

// Arrays
func arrays() {
	function_call()

	// In Go, an array is a numbered sequence of elements of a specific length.
	// In typical Go code, slices are much more common;
	// arrays are useful in some special scenarios.

	// Here we create an array a that will hold exactly 5 ints.
	// The type of elements and length are both part of the array’s type.
	// By default an array is zero-valued, which for ints means 0s.
	var a [5]int
	fmt.Println("emp :", a)

	// We can set a value at an index using the
	// array[index] = value syntax,
	// and get a value with array[index].

	a[4] = 100
	fmt.Println("set :", a)
	fmt.Println("get :", a[4])

	// The builtin len returns the length of an array.
	fmt.Println("len : ", len(a))

	// Use this syntax to declare and initialize an array in one line.
	var b = [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl :", b)

	// You can also have the compiler count the number of elements for you with ...
	c := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println("dcl :", c)

	// If you specify the index with :, the elements in between will be zeroed.
	d := [...]int{10, 3: 5, 20}
	fmt.Println("idx :", d)

	// Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
	var twoD [2][4]int
	for i := 0; i < 2; i++ {
		for j := 3; j > 0; j-- {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("Two dimensional :", twoD)

	// You can create and initialize multi-dimensional arrays at once too.
	twoD = [2][4]int{
		{1, 2, 3, 4},
		{2, 4, 6, 8},
	}
	fmt.Println("Two dimensional dcl :", twoD)
}

// Slices
func slicesAdvArr() {
	function_call()

	// Slices are an important data type in Go,
	// giving a more powerful interface to sequences than arrays.

	// Unlike arrays, slices are typed only by the elements they contain
	// (not the number of elements). An uninitialized slice equals to nil and has length 0.
	var s []string
	fmt.Println("uninit :", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("empt :", s, "len :", len(s), "capacity :", cap(s))

	s[0] = "1"
	s[1] = "2"
	s[2] = "3"

	fmt.Println("set :", s)
	fmt.Println("get :", s[2])

	fmt.Println("len :", s)

	// In addition to these basic operations,
	// slices support several more that make them richer than arrays.
	// One is the builtin append, which returns a slice containing one or more new values.
	// Note that we need to accept a return value from append as we may get a new slice value.
	s = append(s, "a")
	s = append(s, "b")
	s = append(s, "c")
	s = append(s, "d")

	fmt.Println("apd :", s)
	fmt.Println("len apd :", s)
	fmt.Println("cap apd :", cap(s))

	// Slices can also be copy’d.
	// Here we create an empty slice c of the same length as s and copy into c from s.

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println("copied :", c)

	// Slices support a “slice” operator with the syntax slice[low:high]. For example, this gets a slice of the elements s[2], s[3], and s[4].

	// This slices up to (but excluding) s[5].
	l := s[2:5]
	fmt.Println("slice :", l)
	fmt.Println("Slice capacity : ", cap(l))

	// And this slices up from (and including) s[2].
	l = s[2:]
	fmt.Println("slice upto n:", l)
	fmt.Println("Slice capacity upto n : ", cap(l))

	// We can declare and initialize a variable for slice in a single line as well.

	t := []string{"g", "h", "i"}
	fmt.Println("dcl one liner :", t)

	// The slices package contains a number of useful utility functions for slices.
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d :", twoD)

	// Growing slices (the copy and append functions)¶
	var s1 []byte
	s1 = make([]byte, 5, 5)
	t3 := make([]byte, len(s1), cap(s1)+1) // + 1 if s1 length is 0

	for i := range len(s1) {
		t3[i] = s1[i]
	}
	fmt.Println("copied :: ", t3)

	// copy function
	t4 := make([]byte, len(s1), cap(s1)+1)
	copy(t4, s1)
	fmt.Println("copied with internal copy fucntion :: ", t4)

	// A common operation is to append data to the end of a slice.

	appendByte := func(slice []byte, data ...byte) []byte {
		m := len(slice)
		n := m + len(data)

		if n >= cap(slice) { // if necessary, reallocate
			// allocate double what's needed, for future growth.
			newSlice := make([]byte, (n+1)*2)
			copy(newSlice, slice)
			slice = newSlice
		}

		slice = slice[0:n]
		copy(slice[m:n], data)
		return slice

	}
	s1 = appendByte(s1, 23, 22, 43, 11)
	fmt.Println(s1)
	s1 = appendByte(s1, 21)
	fmt.Println(s1)
}

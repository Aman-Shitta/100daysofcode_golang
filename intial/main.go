package main

import (
	"fmt"
	"maps"
	"math"
	"runtime"
	"slices"
	"time"
	"unicode/utf8"
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
	mapss()
	range_fun()
	function()
	multipleReturnValue()
	vardiacFunc()
	closures()
	recursion()
	pointers()
	stringsAndRunes()
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

// maps
func mapss() {
	function_call()

	// Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2

	fmt.Println("map : ", m)

	// Get a value for a key with name[key].
	v1 := m["k1"]
	fmt.Println("v1 : ", v1)

	// If the key doesn’t exist, the zero value of the value type is returned.
	v3 := m["k3"]
	fmt.Println("v3 : ", v3)

	// The builtin len returns the number of key/value pairs when called on a map.
	fmt.Println("len m :", len(m))

	// The builtin delete removes key/value pairs from a map.
	delete(m, "k2")
	fmt.Println("del k2 : ", m)

	// To remove all key/value pairs from a map, use the clear builtin.
	clear(m)
	fmt.Println("clear : m : ", m)

	// The optional second return value when getting a value from a map indicates if the key was present in the map.
	// This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
	// Here we didn’t need the value itself, so we ignored it with the blank identifier _.

	v3, ok := m["k3"]

	if ok {
		fmt.Println("v3 :", v3)
	} else {
		fmt.Println("Key not Present")
	}

	// You can also declare and initialize a new map in the same line with this syntax.
	m1 := map[string]int{
		"k1": 1,
		"k2": 2,
	}

	fmt.Println("single liner : ", m1)

	n2 := map[string]int{"foo": 1, "bar": 2}

	if maps.Equal(m1, n2) {
		fmt.Println("m1 == n2")
	} else {
		fmt.Println("m1 != n2")
	}
}

// Range

func range_fun() {
	function_call()

	// range iterates over elements in a variety of data structures. Let’s see how to use range with some of the data structures we’ve already learned

	nums := []int{
		1, 2, 3, 4, 5,
	}

	// Here we use range to sum the numbers in a slice. Arrays work like this too.
	sum := 0
	for _, i := range nums {
		sum += i
	}
	fmt.Println("sum : ", sum)

	// range on arrays and slices provides both the index and value for each entry.
	// Above we didn’t need the index, so we ignored it with the blank identifier _. Sometimes we actually want the indexes though.

	for idx, num := range nums {
		if num == 3 {
			fmt.Println(num, "is at index : ", idx)
		}
	}

	// range on map iterates over key/value pairs.

	m := map[string]string{
		"a": "A",
		"b": "B",
		"c": "C",
	}

	for key, value := range m {
		fmt.Println(key, " --> ", value)
	}

	// range can also iterate over just the keys of a map.

	for k := range m {
		fmt.Println("KEY :: ", k, "has Value :: ", m[k])
	}

	// range on strings iterates over Unicode code points. The first value is the starting byte index of the rune and the second the rune itself. See Strings and Runes for more details.

	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}

// functions

func function() {
	function_call()
	// Functions are central in Go. We’ll learn about functions with a few different examples.

	plus := func(a int, b int) int {

		return a + b
	}

	fmt.Println("funcrtion call: ", plus(1, 2))

	plusPlus := func(a, b, c int) int {
		return a + b + c
	}

	fmt.Println("Single value type declaration :: ", plusPlus(1, 2, 3))
}

// Multiple Return Values

func multipleReturnValue() {
	function_call()

	// Go has built-in support for multiple return values.
	// This feature is used often in idiomatic Go,
	// for example to return both result and error values from a function.

	vals := func() (int, int) {
		return 1, 2
	}

	a, b := vals()

	fmt.Println("Multiple return values :: ", a, b)
}

// variadic-functions
func vardiacFunc() {
	function_call()

	// Variadic functions can be called with any number of trailing arguments.
	// For example, fmt.Println is a common variadic function.

	sum := func(vals ...int) int {
		sums := 0
		for idx, i := range vals {
			fmt.Println("idx : ", idx)
			sums += i
		}

		return sums
	}

	total := sum(1, 2)
	fmt.Println("total ::", total)

	total = sum(1, 2, 3, 4)
	fmt.Println("total ::", total)
}

// closures

// This function intSeq returns another function, which we define anonymously in the body of intSeq. The returned function closes over the variable i to form a closure.

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func closures() {
	function_call()

	// Go supports anonymous functions, which can form closures.
	// Anonymous functions are useful when you want to define a function inline without having to name it.

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInts := intSeq()
	fmt.Println(nextInts())
}

// Recursion
func fact(n uint64) uint64 {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * fact(n-1)
	}

}
func recursion() {
	function_call()

	fmt.Println("Factorial of 7 :: ", fact((7)))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		} else {
			return fib(n-1) + fib(n-2)
		}
	}
	fmt.Println("fib of 10 :: ", fib(10))
}

// Pointers

// Go supports pointers,
// allowing you to pass references to values and records within your program.

func pointers() {
	zeroval := func(ival int) {
		ival = 0
	}

	zeroptr := func(iptr *int) {
		*iptr = 0
	}

	i := 1

	fmt.Println("initial :", i)

	zeroval(i)
	fmt.Println("zeroval :", i)

	zeroptr(&i)
	fmt.Println("zeroptr :", i)

	fmt.Println("pointer:", &i)
}

// Strings And Runes

func stringsAndRunes() {
	function_call()

	// A Go string is a read-only slice of bytes.
	// The language and the standard library treat strings specially - as containers of text encoded in UTF-8.
	// In other languages, strings are made of “characters”.
	// In Go, the concept of a character is called a rune -
	// A rune is an integer that represents a Unicode code point. This Go blog post(https://go.dev/blog/strings) is a good introduction to the topic.
	const s = "สวัสดี"
	fmt.Println("Len:", len(s))

	// Indexing into a string produces the raw byte values at each index.
	// This loop generates the hex values of all the bytes that constitute the code points in s.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// To count how many runes are in a string, we can use the utf8 package.
	// Note that the run-time of RuneCountInString depends on the size of the string, because it has to decode each UTF-8 rune sequentially.
	// Some Thai characters are represented by UTF-8 code points that can span multiple bytes, so the result of this count may be surprising.
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// A range loop handles strings specially and decodes each rune along with its offset in the string.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// We can achieve the same iteration by using the utf8.DecodeRuneInString function explicitly.

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}

}

// Values enclosed in single quotes are rune literals. We can compare a rune value to a rune literal directly.
func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

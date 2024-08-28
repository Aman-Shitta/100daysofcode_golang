1. Basic Syntax:

    How do you declare and initialize a variable in Go?
        var abc int // initialize  var abc with default zero value (0 -> int)
        // shorthand declaratiion
        abc := 1
    What is the difference between := and var for variable declaration?
        := is a shorthand and require a value at roght signed
        var can declare an zero value varibale, can have value or not assigned
2. Functions:

    How do you define a function in Go that takes two integer arguments and returns their sum?
        func sum(a int, b int) int {
            return a + b
        } 

    What are named return values, and how do you use them?
        func sum(int a, int b) result int {
            result := a + b
        }

        we just define the variable with return type declaration of function 
        and assign a value to function, go runtime will retunr the value from definition
3. Control Structures:

    How do you write an if-else statement in Go?
    if <some-condition-tru> {
        do something
    } else if <another-condition-true> {

    } else {
        so something else
    }
    What are the key differences between switch and select?
    swicth is to run cases based on some value or cinditons
        abc := 2
        switch {
            case abc % 2 == 0:
                fmt.Println("Even")
            case abc%2 != 0:
                fmt.Println("Odd")
            default :
                fmt.Println{"zero"}
        }
    select are specifically designed recieveing data from channel simultaneoulsy
        c1 := make(chan string, 1) // buffered channel

        go func() {
            time.Sleep(2 * time.Second)
            c1 <- "hello"
        }()

        select {
        case res := <-c1:
            println(res)
        case <-time.After(1 * time.Second)
            println("print timeout")
        }


4. Arrays and Slices:

    How do you create an array of integers with 5 elements?
    var abc []int{1,2,3,4,5,}
    or
    abc := []int{1,2,3,4,5}
    What is the difference between an array and a slice in Go, and how do you append elements to a slice?
    array needs size declaration eg var abc [5]int // holds ony 5 elements
    slice is advanced array need no sixe declaratiion eg var abc[]{1}

5. Pointers:

    How do you declare a pointer to an integer in Go?
    var *pointer int

    How do you dereference a pointer to access or modify the value it points to?
    var abc = &pointer


6. Concurrency:

    What is a goroutine, and how do you create one?
    using go keyword before func

    go sum(int a, int b) {
        fmt.Println(a + b)
    }

    sum()

    This will send the sum into go routine and program will exit wihtout printing data
    
    How do you synchronize goroutines using channels?

    I learnt this but I might have forgotten

7. Error Handling:

    How does Go handle errors, and what is the common idiom for handling errors in Go functions?
    Not yet done

8. Structs and Methods:

    How do you define a struct in Go?
    type Human struct {
        legs int
        hands int
        gender string
    }
    How do you define a method on a struct?

    func (h Human) getName() string {
        return h.name
    }

    (h Human) is reciever type

9. Interfaces:

    What is an interface in Go, and how is it used?
    type Abc interface{
        getName()
    }
    func (a Abc) getName() string {
        return a.name
    }

    This will work because Human has function getName makeing feasible with Abc Interface

    (h Human) is reciever type
    Can you explain how the empty interface interface{} works?

10. Packages and Imports:
    not done yet
    How do you create a new package in Go, and how do you import it into your main program?
    What is the significance of the main package and the main() function in a Go program?
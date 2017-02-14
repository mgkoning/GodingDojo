#### Package
`package <name>` (no block)

#### Import
`import "fmt"`  
`import "strconv"`  
 
 *or*  

     import (
        "fmt"
        "strconv"
     )

#### Functions
`func(a int, b int) int {`  
`func(a, b int) (int, int) {`  
`func(a int, b int) (c, d int) {`

#### Variables
`var local string`  
`var local = "123"`  
`local := "123"`  

#### Loops
`for i := 0; i < 10; i++ {`  
`for ; sum < 1000; {`  
`for sum < 1000 {`  
`for {`  

#### Conditionals
`if x < 0 {`  
`if v := math.Pow(x, n); v < lim {`  

    switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

#### Defer
    func main() {
	    defer fmt.Println("world")

	    fmt.Println("hello")
    }

#### Pointers
    var p *int
    // The & operator generates a pointer to its operand.
    i := 42
    p = &i

#### Structs
    type Vertex struct {
        X int
        Y int
    }

    v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex

#### Arrays
`var a [10]int`

#### Slices
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]

    // For the array
    var a [10]int
    // these slice expressions are equivalent:
    a[0:10]
    a[:10]
    a[0:]
    a[:]

`a := make([]int, 5)    // len(a)=5`  
`b := make([]int, 0, 5) // len(b)=0, cap(b)=5`  

    var s []int
	// append works on nil slices.
	s = append(s, 0)
   	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)

#### Ranges
Given: `var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}`  
`for i, v := range pow {`  
`for i := range pow {`  
`for _, value := range pow {  `

#### Maps
    var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

    m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

#### Methods
`func (v Vertex) Abs() float64 {`  

#### Interfaces
    type Abser interface {
        Abs() float64
    }

#### Goroutines
	go say("world")
	say("hello")

#### Channels
    ch := make(chan int)
    ch <- v    // Send v to channel ch.
    v := <-ch  // Receive from ch, and
               // assign value to v.
    // with buffer:
    ch := make(chan int, 100)
    // closing:
    close(ch)
    // iterating:
    for i := range c {
    }

#### Select
    for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
        default:
        // receiving from c would block
        }
	}
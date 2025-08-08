package main

import (
	"fmt"
	"math"
	"strings"
)

// struct
type Vertex struct {
	X int
	Y int
}

type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func createTicTacToeBoard() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendToSlice() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

// function values
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// function closures
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// pointers
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	// struct
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// pointers to structs
	v2 := Vertex{1, 2}
	p2 := &v2
	p2.X = 1e9
	fmt.Println(v2)

	// struct literals
	var (
		structLiteralV1 = Vertex{1, 2}  // has type Vertex
		structLiteralV2 = Vertex{X: 1}  // Y:0 is implicit
		structLiteralV3 = Vertex{}      // X:0 and Y:0
		structLiteralP  = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(structLiteralV1, structLiteralP, structLiteralV2, structLiteralV3)

	// arrays
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// slices: An array has a fixed size. A slice, on the other hand, is a dynamically-sized,
	// flexible view into the elements of an array.
	var s []int = primes[1:4]
	fmt.Println(s)

	// Slices are like references to arrays: A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	nameA := names[0:2]
	nameB := names[1:3]
	fmt.Println(nameA, nameB)

	nameB[0] = "XXX"
	fmt.Println(nameA, nameB)
	fmt.Println(names)

	// slice defaults: When slicing, you may omit the high or low bounds to use their defaults instead.
	// The default is zero for the low bound and the length of the slice for the high bound.
	sd := []int{2, 3, 5, 7, 11, 13}

	sd = sd[1:4]
	fmt.Println(sd)

	sd = sd[:2]
	fmt.Println(sd)

	sd = sd[1:]
	fmt.Println(sd)

	// Slice length and capacity
	newS := []int{2, 3, 5, 7, 11, 13}
	printSlice(newS)

	// Slice the slice to give it zero length.
	newS = newS[:0]
	printSlice(newS)

	// Extend its length.
	newS = newS[:4]
	printSlice(newS)

	// Drop its first two values.
	newS = newS[2:]
	printSlice(newS)


	// Creating a slice with make
	a2 := make([]int, 5)
	printSlice2("a", a2)

	b2 := make([]int, 0, 5)
	printSlice2("b", b2)

	c2 := b2[:2]
	printSlice2("c", c2)

	d2 := c2[2:5]
	printSlice2("d", d2)

	// slices of slices
	createTicTacToeBoard()

	// Appending to a slice
	appendToSlice();

	// range
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// range continued
	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}

	// maps
	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// map literals: Map literals are like struct literals, but the keys are required.
	var m2 = map[string]Vertex2{
		"Bell Labs": Vertex2{
			40.68433, -74.39967,
		},
		"Google": Vertex2{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m2)

	// map literals continued: If the top-level type is just a type name, you can omit it from the elements of the literal.
	var m3 = map[string]Vertex2{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m3)

	// mutating maps
	m4 := make(map[string]int)

	m4["Answer"] = 42
	fmt.Println("The value:", m4["Answer"])

	m4["Answer"] = 48
	fmt.Println("The value:", m4["Answer"])

	delete(m4, "Answer")
	fmt.Println("The value:", m4["Answer"])

	v4, ok4 := m4["Answer"]
	fmt.Println("The value:", v4, "Present?", ok4)	

	// function values: Functions are values too. They can be passed around just like other values.
	// Function values may be used as function arguments and return values.
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// function closures
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
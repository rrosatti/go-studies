package main

import (
	"fmt"
	"image"
	"io"
	"math"
	"strings"
	"time"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

type Abser interface {
	Abs() float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// pointer receivers
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// the empty interface
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// type assertions
func typeAssertions() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)	
}

// type switches
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// errors
type MyError struct {
	When time.Time
	What string
}
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	// methods: Go does not have classes. However, you can define methods on types.
	// A method is a function with a special receiver argument.
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	// methods continued
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// pointer receivers
	v2 := Vertex{3, 4}
	v2.Scale(10)
	fmt.Println(v2.Abs())

	// interfaces: An interface type is defined as a set of method signatures.
	var a Abser
	f3 := MyFloat(-math.Sqrt2)
	v3 := Vertex{3, 4}

	a = f3  // a MyFloat implements Abser
	a = &v3 // a *Vertex implements Abser

	// In the following line, v3 is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v3

	fmt.Println(a.Abs())

	// the empty interface
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	// type assertions
	// typeAssertions()

	// type switches: A type switch is a construct that permits several type assertions in series.
	do(21)
	do("hello")
	do(true)
	
	// errors
	if err := run(); err != nil {
		fmt.Println(err)
	}

	// readers
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	// images
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
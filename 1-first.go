package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// naked return: A return statement without arguments returns the named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var i, j int = 1, 2

const Pi = 3.14

// variadic function
func variadicFunctionSum(nums ...int)  {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}


func main() {
	// 1. importing and using libs
	
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)

	// 2. creating and using functions

	fmt.Println(add(10, 11))

	a, b := swap("hello", "world")
	fmt.Println(a, b);

	fmt.Println(split(17))

	// 3. vars
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	var i int
	var f float64
	var bo bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, bo, s)

	// 4. type conversions
	var x, y int = 3, 4
	var fo float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(fo)
	fmt.Println(x, y, z)

	// 5. constants
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	// variadic function
	variadicFunctionSum(1, 2)
	variadicFunctionSum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	variadicFunctionSum(nums...)
}
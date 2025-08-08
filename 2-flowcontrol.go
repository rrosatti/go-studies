package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// if
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// if with a short statement
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// If and else
func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	sum := 0
	for i := 0; i < 10; i ++ {
		sum += i
	}
	fmt.Println((sum))

	// for continued
	sum2 := 1
	for ; sum2 < 1000; {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// For is Go's "while"
	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}
	fmt.Println(sum3)

	// if
	fmt.Println(sqrt(2), sqrt(-4))
	
	// if with a short statement
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// if and else
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)

	// switch
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// switch with no condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// defer: defers the execution of a function until the surrounding function returns
	defer fmt.Println("world")
	fmt.Println("hello")

	// stacking defers: Deferred function calls are pushed onto a stack. 
	// When a function returns, its deferred calls are executed in last-in-first-out order.
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}
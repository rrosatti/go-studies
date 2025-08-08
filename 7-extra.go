package main

import (
	"cmp"
	"errors"
	"fmt"
	"slices"
	"sync"
	"sync/atomic"
	"time"
)

// enums
type ServerState int
// The possible values for ServerState are defined as constants. 
// The special keyword iota generates successive constant values automatically; in this case 0, 1, 2 and so on.
const (
    StateIdle ServerState = iota
    StateConnected
    StateError
    StateRetrying
)

// By implementing the fmt.Stringer interface, values of ServerState can be printed out or converted to strings.
var stateName = map[ServerState]string{
    StateIdle:      "idle",
    StateConnected: "connected",
    StateError:     "error",
    StateRetrying:  "retrying",
}

// transition emulates a state transition for a server; it takes the existing state and returns a new state.
func transition(s ServerState) ServerState {
    switch s {
    case StateIdle:
        return StateConnected
    case StateConnected, StateRetrying:
        return StateIdle
    case StateError:
        return StateError
    default:
        panic(fmt.Errorf("unknown state: %s", s))
    }
}
	
func (ss ServerState) String() string {
    return stateName[ss]
}

//// generics

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
    for i := range s {
        if v == s[i] {
            return i
        }
    }
    return -1
}
type List[T any] struct {
    head, tail *element[T]
}

type element[T any] struct {
    next *element[T]
    val  T
}

func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

func (lst *List[T]) AllElements() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

func tryGenerics() {
	var s = []string{"foo", "bar", "zoo"}
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	_ = SlicesIndex[[]string, string](s, "zoo")
    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
    fmt.Println("list:", lst.AllElements())
}

////// errors
type argError struct {
    arg     int
    message string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
    if arg == 42 {
		// Return our custom error.
        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}

func tryCustomError() {
	_, err := f(42)
    var ae *argError
    if errors.As(err, &ae) {
        fmt.Println(ae.arg)
        fmt.Println(ae.message)
    } else {
        fmt.Println("err doesn't match argError")
    }
}

/////// channels
func tryChannels() {
	// Create a new channel with make(chan val-type).
	// Channels are typed by the values they convey.
	messages := make(chan string)

	// Send a value into a channel using the channel <- syntax. 
	// Here we send "ping" to the messages channel we made above, from a new goroutine.
	go func() { messages <- "ping" }()

	// The <-channel syntax receives a value from the channel. 
	// Here we’ll receive the "ping" message we sent above and print it out.
	msg := <-messages
	fmt.Println(msg)
}

////// timeouts
func tryTimeouts() {
	// For our example, suppose we’re executing an external call that returns its result on a channel c1 after 2s. 
	// Note that the channel is buffered, so the send in the goroutine is nonblocking. 
	// This is a common pattern to prevent goroutine leaks in case the channel is never read.
	c1 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "result 1"
    }()

	// Here’s the select implementing a timeout. res := <-c1 awaits the result and <-time.After awaits a value to be sent after 
	// the timeout of 1s. Since select proceeds with the first receive that’s ready, we’ll take the timeout case if the 
	// operation takes more than the allowed 1s.
	select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second):
        fmt.Println("timeout 1")
    }

	// If we allow a longer timeout of 3s, then the receive from c2 will succeed and we’ll print the result.
	c2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(3 * time.Second):
        fmt.Println("timeout 2")
    }
}

///// timers
func tryTimers() {
	// Timers represent a single event in the future. You tell the timer how long you want to wait, and it provides
	// a channel that will be notified at that time. This timer will wait 2 seconds.
	timer1 := time.NewTimer(2 * time.Second)

	// The <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer fired.
	<-timer1.C
    fmt.Println("Timer 1 fired")

	// If you just wanted to wait, you could have used time.Sleep. 
	// One reason a timer may be useful is that you can cancel the timer before it fires. Here’s an example of that.
	timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 fired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
	
	// Give the timer2 enough time to fire, if it ever was going to, to show it is in fact stopped.
	time.Sleep(2 * time.Second)
}

////// tickers: for when you want to do something repeatedly at regular intervals
func tryTickers() {
	// Tickers use a similar mechanism to timers: a channel that is sent values. 
	// Here we’ll use the select builtin on the channel to await the values as they arrive every 500ms.
	ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)

	go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()

	// Tickers can be stopped like timers. 
	// Once a ticker is stopped it won’t receive any more values on its channel. We’ll stop ours after 1600ms.
	time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")
}

//////// WaitGroups: To wait for multiple goroutines to finish, we can use a wait group.
func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)
	// Sleep to simulate an expensive task.
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}
func tryWaitGroups() {
	// This WaitGroup is used to wait for all the goroutines launched here to finish. 
	// Note: if a WaitGroup is explicitly passed into functions, it should be done by pointer.
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
        wg.Add(1)

		// Wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done. 
		// This way the worker itself does not have to be aware of the concurrency primitives involved in its execution.
        go func() {
            defer wg.Done()
            worker(i)
        }()
    }

	// Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.
	wg.Wait()
}

////// atomic counters: The primary mechanism for managing state in Go is communication over channels. 
// We saw this for example with worker pools. There are a few other options for managing state though. 
// Here we’ll look at using the sync/atomic package for atomic counters accessed by multiple goroutines.
func tryAtomicCounters() {
	// We’ll use an atomic integer type to represent our (always-positive) counter.
	var ops atomic.Uint64

	// A WaitGroup will help us wait for all goroutines to finish their work.
	var wg sync.WaitGroup

	// We’ll start 50 goroutines that each increment the counter exactly 1000 times.
	for range 50 {
        wg.Add(1)
        go func() {
            for range 1000 {
				// To atomically increment the counter we use Add.
                ops.Add(1)
            }
            wg.Done()
        }()
    }
	// Wait until all the goroutines are done.
	wg.Wait()

	// Here no goroutines are writing to ‘ops’, but using Load it’s safe to atomically read a 
	// value even while other goroutines are (atomically) updating it.
	fmt.Println("ops:", ops.Load())
}

///// Sorting
func trySorting() {
	strs := []string{"c", "a", "b"}
    slices.Sort(strs)
    fmt.Println("Strings:", strs)

    ints := []int{7, 2, 4}
    slices.Sort(ints)
    fmt.Println("Ints:   ", ints)

    s := slices.IsSorted(ints)
    fmt.Println("Sorted: ", s)
}

//// Sorting by Functions
func trySortingByFunctions() {
	fruits := []string{"peach", "banana", "kiwi"}

	// We implement a comparison function for string lengths. cmp.Compare is helpful for this.
	lenCmp := func(a, b string) int {
        return cmp.Compare(len(a), len(b))
    }

	// Now we can call slices.SortFunc with this custom comparison function to sort fruits by name length.
	slices.SortFunc(fruits, lenCmp)
    fmt.Println(fruits)

	// We can use the same technique to sort a slice of values that aren’t built-in types.
	type Person struct {
        name string
        age  int
    }

	people := []Person{
        Person{name: "Jax", age: 37},
        Person{name: "TJ", age: 25},
        Person{name: "Alex", age: 72},
    }

	// Sort people by age using slices.SortFunc
	slices.SortFunc(people,
        func(a, b Person) int {
            return cmp.Compare(a.age, b.age)
        })
	fmt.Println(people)
}

func main()  {
	// enums
	ns := transition(StateIdle)
    fmt.Println(ns)
    ns2 := transition(ns)
    fmt.Println(ns2)

	// generics
	tryGenerics()

	// custom error
	tryCustomError()

	// channels
	tryChannels()

	// timeouts
	tryTimeouts()

	// timers
	tryTimers()

	// tickers
	tryTickers()

	// wait groups
	tryWaitGroups()

	// atomic counters
	tryAtomicCounters()

	// sorting
	trySorting()

	// sorting by functions
	trySortingByFunctions()
}

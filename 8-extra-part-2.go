package main

import (
	"bufio"
	"bytes"
	b64 "encoding/base64"
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	s "strings"
	"syscall"
	"time"
)

///// recover:  Go makes it possible to recover from a panic, by using the recover built-in function.
// A recover can stop a panic from aborting the program and let it continue with execution instead.
func mayPanic() {
    panic("a problem")
}
func tryRecover() {
	// recover must be called within a deferred function. 
	// When the enclosing function panics, the defer will activate and a recover call within it will catch the panic.
	defer func() {
        if r := recover(); r != nil {
			// The return value of recover is the error raised in the call to panic.
            fmt.Println("Recovered. Error:\n", r)
        }
    }()

	mayPanic()

	// This code will not run, because mayPanic panics. 
	// The execution of main stops at the point of the panic and resumes in the deferred closure.
	fmt.Println("After mayPanic()")
}

///// string functions
func tryStringFunctions() {
	var p = fmt.Println

	p("Contains:  ", s.Contains("test", "es"))
    p("Count:     ", s.Count("test", "t"))
    p("HasPrefix: ", s.HasPrefix("test", "te"))
    p("HasSuffix: ", s.HasSuffix("test", "st"))
    p("Index:     ", s.Index("test", "e"))
    p("Join:      ", s.Join([]string{"a", "b"}, "-"))
    p("Repeat:    ", s.Repeat("a", 5))
    p("Replace:   ", s.Replace("foo", "o", "0", -1))
    p("Replace:   ", s.Replace("foo", "o", "0", 1))
    p("Split:     ", s.Split("a-b-c-d-e", "-"))
    p("ToLower:   ", s.ToLower("TEST"))
    p("ToUpper:   ", s.ToUpper("test"))
}

///// regular expressions
func tryRegularExpressions() {
	// This tests whether a pattern matches a string.
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)

	// Above we used a string pattern directly, but for other regexp tasks you’ll need to Compile an optimized Regexp struct.
    r, _ := regexp.Compile("p([a-z]+)ch")

	// Many methods are available on these structs. Here’s a match test like we saw earlier.
    fmt.Println(r.MatchString("peach"))

	// This finds the match for the regexp.
    fmt.Println(r.FindString("peach punch"))

	// This also finds the first match but returns the start and end indexes for the match instead of the matching text.
    fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// The Submatch variants include information about both the whole-pattern matches and the submatches within those matches.
	// For example this will return information for both p([a-z]+)ch and ([a-z]+).
    fmt.Println(r.FindStringSubmatch("peach punch"))

	// Similarly this will return information about the indexes of matches and submatches.
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// The All variants of these functions apply to all matches in the input, not just the first. 
	// For example to find all matches for a regexp.
    fmt.Println(r.FindAllString("peach punch pinch", -1))

	// These All variants are available for the other functions we saw above as well.
    fmt.Println("all:", r.FindAllStringSubmatchIndex(
        "peach punch pinch", -1))

	// Providing a non-negative integer as the second argument to these functions will limit the number of matches.
    fmt.Println(r.FindAllString("peach punch pinch", 2))

	// Our examples above had string arguments and used names like MatchString. 
	// We can also provide []byte arguments and drop String from the function name.
    fmt.Println(r.Match([]byte("peach")))

	// When creating global variables with regular expressions you can use the MustCompile variation of Compile.
	// MustCompile panics instead of returning an error, which makes it safer to use for global variables.
    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println("regexp:", r)

	// The regexp package can also be used to replace subsets of strings with other values.
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// The Func variant allows you to transform matched text with a given function.
    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}

////// time: Go offers extensive support for times and durations
func tryTime() {
    p := fmt.Println

    now := time.Now()
    p(now)

    then := time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())

    p(then.Weekday())

    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

    diff := now.Sub(then)
    p(diff)

    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())

    p(then.Add(diff))
    p(then.Add(-diff))
}

////// Time Formatting / Parsing
func tryTimeFormattingAndParsing() {
    p := fmt.Println

    t := time.Now()
    p(t.Format(time.RFC3339))

    t1, e := time.Parse(
        time.RFC3339,
        "2012-11-01T22:08:41+00:00")
    p(t1)

    p(t.Format("3:04PM"))
    p(t.Format("Mon Jan _2 15:04:05 2006"))
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))
    form := "3 04 PM"
    t2, e := time.Parse(form, "8 41 PM")
    p(t2)

    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())

    ansic := "Mon Jan _2 15:04:05 2006"
    _, e = time.Parse(ansic, "8:41PM")
    p(e)
}

//////// Number Parsing
func tryNumberParsing() {
    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)

    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i)

    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)

    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)

    k, _ := strconv.Atoi("135")
    fmt.Println(k)

    _, e := strconv.Atoi("wat")
    fmt.Println(e)
}

// URL Parsing
func tryURLParsing() {
    s := "postgres://user:pass@host.com:5432/path?k=v#f"

    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

    fmt.Println(u.Scheme)

    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)

    fmt.Println(u.Host)
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)

    fmt.Println(u.Path)
    fmt.Println(u.Fragment)

    fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])
}

////// Base64 Encoding
func tryBase64Encoding() {
    data := "abc123!?$*&()'-=@~"

    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
}

////// Command-Line Arguments
func tryCommandLineArgs() {
    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]

    arg := os.Args[3]

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}
//////// Command-Line Flags
func tryCommandLineFlags() {
	wordPtr := flag.String("word", "foo", "a string")

    numbPtr := flag.Int("numb", 42, "an int")
    forkPtr := flag.Bool("fork", false, "a bool")

    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    flag.Parse()

    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *forkPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}

/////// Environment Variables
func tryEnvVars() {
    os.Setenv("FOO", "1")
    fmt.Println("FOO:", os.Getenv("FOO"))
    fmt.Println("BAR:", os.Getenv("BAR"))

    fmt.Println()
    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        fmt.Println(pair[0])
    }
}

////// HTTP Client
func tryHTTPClient() {
	    resp, err := http.Get("https://gobyexample.com")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("Response status:", resp.Status)

    scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan() && i < 5; i++ {
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }
}

//// HTTP Server
func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}
func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}
func tryHttpServer() {
	http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)
	
	http.ListenAndServe(":8090", nil)
}

////// Exec'ing Processes
func tryExecingProcesses() {
    binary, lookErr := exec.LookPath("ls")
    if lookErr != nil {
        panic(lookErr)
    }

    args := []string{"ls", "-a", "-l", "-h"}

    env := os.Environ()

    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}

func main() {
	// recover
	tryRecover()

	// string functions
	tryStringFunctions()

	// regular expressions
	tryRegularExpressions()

	// time
	tryTime()

	// time formatting / parsing
	tryTimeFormattingAndParsing()

	// number parsing
	tryNumberParsing()

	// URL parsing
	tryURLParsing()

	// Base64 encoding
	tryBase64Encoding()

	// Command-Line args
	tryCommandLineArgs()

	// Command-Line flags
	tryCommandLineFlags()

	// Environment Variables
	tryEnvVars()

	// HTTP client
	tryHTTPClient()

	// Http server
	// tryHttpServer()

	// Exec'ing processes
	tryExecingProcesses()
}

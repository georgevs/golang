# Go by example

[Go by example](https://gobyexample.com)

## Arguments (os, flags)
```
os.Args[i] -> [prog a1 a2 a3...]
val := flag.String/Int/Bool(key, default, help) -> *String
flag.Parse()
flag.Args() -> (tail []T)

./prog -str=abc -int=5 -bool a b c    // tail = [a b c]
```

## Variadic functions
```
func sum(xs ...int) {
  for i,x := range xs { }
}

sum(1)
sum(1, 2, 3, 4)

xs := []int{ 1, 2, 3, 4 }
sum(xs...)
```

## Strings (unicode/utf8)
```
len(s) -> length in bytes, NOT in code points (aka runes)
utf8.RuneCountInString(s) -> length in code points
for i:=0, i<len(s); i++ { b := s[i] }  // iterate bytes 
for i,ch := range s { ch == s[i] } // iterate code points
ch,w := utf8.DecodeRuneInString(s[i:])
```

## Structs, methods, interfaces
```
dog := struct { name string; isGood bool }  // anonymous struct

type person struct { name string; age int }  // type definition
func newPerson() *person { return &person{name:"name", age:42} }  // ctor
func (p *person) foo() { ... } // method

type geometry interface {  // interface
  area() float64
  perimeter() float64
}

type rect struct { w, h float64 }

func (r *rect) area() float64 { return r.w * r.h }   // rect implements geometry
```

## Composifion and embedding
```
type base struct { num int }
type container struct {
  base
  str string
}
co := container{ base{1}, "str" }
co.base.num
```

## Generics
```
type List[T any] struct { head, tail * element[T] }
type element[T any] struct { next element[T]; val T }
func lst := List[int]{} lst.GetAll()
```

## Errors (errors)
```
func foo() (int,error) { return -1, error.New("...reason...") }
type myError struct { what, when string }
func (e *myError) Error() string { return fmt.Sprintf(...) }
if val, ok := e.(*myError); ok { ...e is myError ... }
```

## Waitgroup (sync)
```
var wg sync.WaitGroup
wg.Add(k)
for i:=0; i<k; i++ { got func() { ...; wg.Done() } }
wg.Wait()
```

## Rate limiter (time)
```
limiter := time.Tick(200 * time.Millisecond)
for {
  <- limiter   // 200ms wait
}
```

## Atomic counters (sync/atomic)
```
var counter uint64
atomic.AddUint64(&counter, 1)
```

## Mutex (sync)
```
type container struct { mu sync.Mutex; data int }
func (c *Container) inc() {
  c.mu.Lock(); defer c.mu.Unlock();
  c.data++
}
```

## Stateful goroutines
```
server := func() {
  var state int
  for { switch req := <- ch { ...mutate state } }
}
```

## Sorting (sort)
```
sort.Strings(values)
sort.IntsAreSorted(values)

sortable:
  Len() int
  Swap(i,j int)
  Less(i,j int) bool
```







## Environment (os)
```
os.Setenv(key, val)
os.Getenv(key) -> val
os.Environ() -> []String   "key=val"
```

## Http client (net/http, buffio)
```
resp, err := http.Get(url)
defer res.Body.Close()
res.Status
scanner := buffio.NewScanner(resp.Body)
scanner.Scan()
scanner.Text() 
scanner.Err()
```

## Http server (net/http)
```
func hello(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "hello\n")
}
http.HandlerFunc("/hello", hello)
http.ListenAndServe(":8080", nil)

curl http://localhost:8080/hello
```

## Http request context (net/http)
```
func hello(w http.ResponseWriter, req *http.Request) {
  ctx := req.Context()
  switch {
    case <- ctx.Done(): ...leave; we're done...
  }
}
```

## Spawn process (os/exec)
Read all output
```
proc, err := exec.Command("date")
bytes, err := proc.Output()
fmt.Println(string(bytes))
switch e := err.(type) {
  case *exec.Error: ...
  case *exec.ExitError: ...
}
```
Read through pipes
```
proc, err := exec.Command("grep", "hello")
w, _ := proc.StdinPipe()   // send input to the process
r, _ := proc.StdoutPipe()  // read output from the process

proc.Start()

w.Write([]byte{"...text..."})
w.Close()  // done with input

bytes, _ := io.ReadAll(r)
proc.Wait()   // sync completion
```

## Run and forget (syscall)
```
err := syscall.Exec("/path/to/prog", args, os.Environ())
```

## Signals (os/signal)
```
sigs := make(chan os.Signal, 1)
signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
```

## Exit (os)
```
os.Exit(3)
```

# The Go Programming Language

## Ch1: Tutorial

- Package: Go code is organized into packages, which are similar to libraries or modules
in other languages. A package consits of one or more .go source files in a
single directory that define what the package does. Each source file begins with
a package declaration,that states which package the file belongs to, followed by
a list of other packages that it imports, and then the declarations of the
program that are stored in that file.

- Package main is special. It defines a standalone executable program, not a
library. Within package main the function main is also special - it's where
execution of the program begins.

- A function declaration consists of the keyword func, the name of the function,
a parameter list, a result list, and the body of the function - the statements
that define what id does - enclosed in braces.

- Variable can be declared in any of the following forms:
```
s := "" // can only be used within a function
var s string // default initialization to the zero value for strings
var s = ""
var s string = ""
```
In practice, you should generally use one of the first two forms, with explicit
initialization to say that the initial value is important and implicit
initialization to say that the initial value doesn't matter.

- A goroutine is a conurrent function execution. A channel is a communication
mechanism that allows one goroutine to pass values of a specified type to
another goroutine. The function main runs in a goroutine and the "go" statement
creates additional goroutines.

- Go allows a simple statement such as a local variable declaration to preceded
the if condition, which is particularly useful for error handling as in this
example:
```
err := r.ParseForm()
if err != nil {
		log.Print(err)
}
```
could be written as:
```
if err := r.ParseForm(); err != nil {
		log.Print(err)
}
```
Combining the statements is shorter and reduces the scope of the variable err,
which is a good practice.

## ch2: Program Structure

- In Go, as in any other programming language, one builds large programs from a
small set of basic constructs. Variables store values. Simple expressions are
combined into larger ones with operations like addition and subtraction. Basic
types are collected into aggregates like arrays and structs. Expressions are
used in statements whose execution order is determined by control-flow
statements like if and for. Statements are grouped into functions for isolation
and reuse. Functions are gathered into source files and packages.

### Names
- Go has 25 keywords like "if" and "switch" that may be used only where the
syntax permits; they can't be used as names.
```
break       default         func      interface     select
case        defer           go        map           struct
chan        else            goto      package       switch
const       fallthrough     if        range         type
continue    for             import    return        var
```

- In addition, there are about three dozen predeclared names like int and true
for built-in contants, types, and functions:
```
Constants:    true    false   iota        nil

Types:        int       int8      int16         int32       int64
              uint      uint8     uint16        uint32      uint64 uintptr
              float32   float64   complex128    complex64
              bool      byte      rune          string       error

Functions:    make      len   cap   new   append    copy    close   delete
              complex   real  imag
              panic     recover
```
These names are not reserved, so you may use them in declarations. At some
places redeclaring one of them makes sense, but beware of the potential for
confusion.

- There is no limit on name length. Generally, the larger the scope of a name,
the longer and more meaningful it should be.

- Go programmers use "camel case" when formatting names by combining words;
that is, interior capital letters are preferred over interior underscores. The
letters of acronyms and initialisms like ASCII and HTML are always rendered in
the same case, so a function might be called html-Escape, HTMLEscape, or
escapeHTML, but not escapeHtml.

### Declarations
- There are four major kinds of declarations: var, const, type, and func.

- If an entity is declared within a function, it is local to that function. If
declared outside of a function, however, it is visible in all files of the
package to which it belongs. The case of the first letter of a name determines
its visibility across package boundaries. If the name begins with an upper-case
letter, it is exported, which means that it is visible and accessible outside of
its own package and may be referred to by other parts of the program, as with
Printf in the fmt package. Package names themselves are always in lower case.

### Variables
- A var declaration creates a variable of a particular type, attaches a name to
it, and sets its initial value. Each declaraion has the general form
```
var name type = expression
```
Either the "type" or the "=" expression part may be omitted, but not both. If
the type is omitted, it is determined by the initializer expression. If the
expression is omitted, the initial value is the zero value for the type, which
is 0 for numbers, false for booleans, "" for strings, and nil for interfaces
and reference types (slice, pointer, map, channel, function). The zero value of
an aggregate type like an array or a struct has the zero value of all of its
elements or fields.

The zero-value mechanism ensures that a variable always holds a well-defined
value of its type; in Go there is no such thing as an uninitialized variable.
This simplifies code and often ensures sensible behavior of boundary conditions
without extra work.

It is possible to declare and optionally initialize a set of variables in a
single declarations, with a matching list of expressions.
```
var i, j, k int
var b, f, s = true, 2.3, "four"
```

Initializers may be literal values or arbitrary expressions.

A set of variables can also be initialized by calling a function that returns
multiple values:
```
var f, err = os.Open(name)
```

#### Variable Declaration
- Within a function, an alternate form called a short variable declaration may
be used to declare and initialize local variables. It takes the form
```
name := expression
```
and the type of name is determinted by the type of expression.

- A "var" declaration tends to be reserved for local variables that need an
explicit type that differs from that of the initializer expression, or for when
the variable will be assigned a value later and its initial value is
unimportant.
```
i := 100 // an int
var boiling float64 = 100 // a float64
```
```
var names []string
var err error
var p Point
```

- A short variable declaration must declare at least one new variable, however,
so this code will not compile.
```
f, err := os.Open(infile)
f, err := os.Create(outfile) // compile error: no new variables
```

### Pointers
- A pointer value is the address of a variable. A pointer is thus the location
at which a value is stored.

- If a variable is declared "var x int", the expression &x ("address of x")
yields a pointer to an integer variable, that is, a value of type *int, which is
pronounced "pointer to int". If this value is called p, we say "p points to x",
or equivalentaly "p contains the address of x". The variable to which p points
is written *p. The expression *p yields the value of that variable, an int.
```
x := 1
p := &x         // p, of type *int, points to x
fmt.Println(*p) // "1"
*p = 2          // equivalent to x = 2
fmt.Println(x)  // "2"
```

- The zero value for a pointer of any type is nil.

- Pointers are comparable; two pointers are equal if and only if they point to
the same variable or both are nil.
```
var x, y int
fmt.Println(&x == &x, &x == &y, &x == nil)  // "true false false"
```

- It is perfectly safe for a function to return the address of a local variable.
```
var p = f()
func f() *int {
  v := 1
  return &v
}
```
Each call of f returns a distinct value:
```
fmt.Println(f() == f()) // "false"
```

### New function
- New is only a syntactic convenience, not a fundamental notion. The two newInt
functions below have identical behaviors:
```
func newInt() *int {
  return new(int)
}
```
```
func newInt() *int {
  var dummy int
  return &dummy
}
```

### Lifetime of a variable
- The variable lives on until it becomes unreachable.

- A compiler may choose to allocate local variables on the heap or on the stack
but, perhaps surprisingly, this choice is not determined by whether var or new
was used to declare the variable.
```
var global *int
func f() {
  var x int
  x = 1
  global = &x
}
```
```
func g() {
  y := new(int)
  *y = 1
}
```
Here, x must be heap-allocated because it is still reachable from the vraiable
global after f has returned, despite being declared as a local variable; we say
`x escapes from f`. Conversely, when g returns, the variable *y becomes
unreachable and can be recycled. Since *y does not escape from g, it's safe for
the compiler to allocate *y on the stack, even though it was allocated with new.
In any case, the notion of escaping is not something that you need to worry
about in order to write correct code, though it's good to keep in mind during
performance optimization, since each variable that escapes requires an extra
memory allocation.

- To write efficient programs you still need to be aware of the lifetime of
variables.

### Type declaration
- A type declaration defines a new named type that has the same underlying type
as an existing type. The named type provides a way to separate different and
perhaps incompatible uses of the underlying type so that they can't be mixed
unintentionally.
```
type name underlying-type
```
For example:
```
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
  AbsoluteZeroC Celsius = -273.15
  FreezingC     Celsius = 0
  BoilingC      Celsisu = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32)}

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9)}
```
This package defines two types, Celsius and Fahrenheit, for the two units of
temperature. Even though they have the same underlying type, float64, they are
not the same type, so they cannot be compared or combined in arithmetic
expressions. Distinguishing the types makes it possible to avoid errors like
inadvertently combining temperatures in the two different scales.

- A conversion from one type to another is allowed if both have the same
underlying type, or if both are the unnamed pointer types that point to
variables of the same underlying type; these conversions change the type but not
the representation of the value. If x is assignable to T, a conversion is
permitted but is usually redundant.

- Conversions are also allowed between numeric types, and between string and
some slice types. These conversions may change the representation of the value.
For example converting a floating-point number to an integer discards any
fractional part, and converting to string to a []byte slice allocates a copy of
the string data. In any case, a conversion never fails at run time.

- Named types also make it possible to define new behaviors for values of the
type. These behaviors are expressed as a set of functions associated with the
type, called the type's methods. The declaration below, in which the Celsius
parameter c appears before the function name, associates with the Celsius type
a method named String that returns c's numeric value.
```
func (c Celsius) String() string { return fmt.Sprintf("%gC", c) }
```
```
c := FToC(212.0)
fmt.Println(c.String())
```

### Imports
- Within a Go program, every package is identified by a unique string called its
import path. These are the strings that appear in an import declaration like
"gopl.io/ch2/tempconv". The language specification doesn't define where these
strings come from or what they mean; it's up to the tools to interpret them.
When using the go tool, an import path denotes a directory containing one or
more Go source files that together make up the package.

- It is an error to import a package and then not refer to it. This check helps
eliminate dependencies that become unnecessary as the code evolves, although it
can be a nuisance during debugging.

### Scope
- In the following program
```
if f, err := os.Open(fname); err != nil { // compile error: unused: f
  return err
}
f.Stat() // compile error: undefined f
f.Close() // compile error: undefined f
```
the scope of f is just the if statement, so f is not accessible to the
statements that follow, resulting in compile errors. Depending on the compiler,
you may get an additional error reporting that the local variable f was never
used.

Thus it is often necessary to declare f before the conditions so that it is
accessible after:
```
f, err := os.Open(fname)
if err != nil {
  return err
}
f.Stat()
f.Close()
```
You may be tempted to avoid declaring f and err in the outer block by moving the
calls to Stat and Close inside an else block:
```
if f, err := os.Open(fname); err != nil {
  return err
} else {
  // f and err are visible here too
  f.Stat()
  f.Close()
}
```
but normal practice in Go is to deal with the error in the if block and then
return, so that the successful exection path is not indented.

- Consider the prorgram below, which starts by obtaining its current working
directory and saving it in a package-level variable. This could be done by
calling os.Getwd in function main, but it might be better to separate this
concern from the primary logic, especially if failing to get the directory is a
fatal error. The function log.Fatal prints a message and calls os.Exit(1).
```
var cwd string

func init() {
  cwd, err := os.Getwd() // compile error: unused: cwd
  if err != nil {
    log.Fatal("os.Getwd failed: %v", err)
  }
}
```
Since neither cwd nor err is already declared in the init functions's block, the
:= statement declares both of them as local variables. The inner declaration of
cwd makes the outer one inaccessible, so the statement does not update the
package-level cwd variable as intented.

- Current Go compilers detect that the local cwd variable is never used and
report this as an error, but they are not strictly required to perform this
check. Furthermore, a minor change, such as the addition of a logging statement
that refers to the local cwd would defeat the check.
```
var cwd string

func init() {
  cwd, err := os.Getwd()
  if err != nil {
    log.Fatal("os.Getwd failed: %v", err)
  }
  log.Printf("Working directory = %s", cwd)
}
```
The global cwd variable remains uninitialized, and the apparently normal log
output obfuscates the bug.

There are a number of ways to deal with this potential problem. The most direct
is to avoid := by declaring err in a separate var declaration:
```
var cwd string

func init() {
  var err error
  cwd, err = os.Getwd()
  if err != nil {
    log.Fatal("os.Getwd failed: %v", err)
  }
}
```

## Ch3: Basic Data Types
- The type rune is a synonym for int32 and conventionally indicates that a value
is a Unicode code point. Similarly, the type byte is a synonym for uint8, and
empasizes that the value is a raw data rather than a small numeric quantity.

- Although Go provides unsigned numbers and arithmetic, we tend to use the
signed int form even for quantities that can't be negative, such as the length
of an array, though uint might seem a more obvious choice. Indeed, the built-in
len function returns a signed int, as in this loop which announces prize medals
in reverse order:
```
medals := []string{"gold", "silver", "bronze"}
for i := len(medals) - 1; i >= 0; i-- {
  fmt.Println(medals[i])
}
```
If len returned an unsigned number, then i too would be a uint, and the
condition i >= 0 would always be true by definition. After the third iteration,
in which i == 0, the i-- statement would cause i to become not -1, nit the
maximum uint value (for example, 2**64 - 1), and the evaluation of the medals[i]
would fail at run time, or panic, by attempting to access an element outside the
bounds of the slice.

For this reason, unsigned numbers tend to be used only when their bitwise
operators or peculiar artithmetic operators are required, as when implementing
bit sets, parsing binary file formats, or for hashing and cryptography. They are
typically not used for merely non-negative quantities.

- The natural data type to hold a single rune is int32.

- When a UTF-8 decoder consumes an unexpected input byte, it generates a special
Unicode replacement character, '\uFFFD', which is usually printed as a white
question mark inside a black hexagonal or diamond-like shape �. When a program
encounters this rune value, it's often a sign that some upstream part of the
system that generated the string data has been careless in its treatment of text
encodings.

## ch4: Composite types
### Arrays
- An array is a fixed-length sequence of zero or more elements of a particular
type. Because of their fixed length, arrays are rarely used directly in Go.
```
var a [3]int
fmt.Println(a[0])
fmt.Println(a[len(a)-1])
```

- The size of an array is part of its type, so [3]int and [4]int are different
types.
```
var q [3]int = [3]int{1, 2, 3}
q := [3]int{1, 2, 3}
```

- We can pass a pointer to an array so that any modifications the function makes
to array elements will be visible to the caller. This function zeros the
contents of a [32]byte array.
```
func zero (ptr *[32]byte) {
  for i := range ptr {
    ptr[i] = 0
  }
}
```
Using a pointer to an array is efficient and allows the called function to
mutate the caller's variable, but arrays are still inherently inflexible because
of their fixed size. The zero function will not accept a pointer to a [16]byte
variable, for example, nor is there any way to add or remove any elements. For
these reasons, other than special cases, arrays are seldom used as a function
parameters or results; instead, we use slices.

### Slices
- Slices represent variable-length sequence whose elements all have the same
type. A slice type is written []T, where the elements have type T; it looks like
an array type without a size.

- A slice is a ligthweight data structure that gives access to a subsequence
(or perhaps all) of the elements of an array, which is knowns as the slice's
underlying array.

- Initialization of the slice differs from that for the array.
```
a := [...]int{0, 1, 2, 3, 4, 5}
```
```
s := []int{0, 1, 2, 3, 4, 5}
```
A slice literal looks like an array literal, a sequence of values separated by
commas and surrounded by braces, but the size is not given. This implicitly
creates an array variable of the right size and yields a slice that points to
it. As with array literals, slice literals may specify the values in order, or
give their indices explicitly, or use a mix of the two styles.

- A slice has three components: a pointer, a length, and a capacity. The pointer
points to the first element of the array that is reachable through the slice,
which is not necessarily the array's first element. The length is the number of
slice elements; it can't exceed the capacity, which is usually the number of
elements between the start of the slice and the end of the underlying array.

- Unlike arrays, slices are not comparable. The only legal slice comparison is
against nul, as in
```
if summer == nil {/* ... */}
```

- The zero value of a slice type is nil. A nil slice has no underlying array.
The nil slice has length and capacity zero.
```
var s []int     // len(s) == 0, s == nil
s = nil         // len(s) == 0, s == nil
s = []int(nil)  // len(s) == 0, s == nil
s = []int{}     // len(s) == 0, s != nil
```

- The built-in function make creates a slice of a specified element type,
length, and capacity. The capacity argument may be omitted, in which case the
capacity equals the length.
```
make([]T, len)
make([]T, len, cap)
```
Under the hood, make creates an unnamed array variable and returns a slice of
it; the array is accessible only through the returned slice.

### Maps
- A map is a reference to a hash table, and a map type is written map[K]V, where
K and V are the type of its keys and values. All of the keys in a given map are
of the same type, and all of the values are of the same type, but the keys need
not be of the same type as the values. The key type K must be comparable using
==, so that the map can test whether a given key is equal to one already within
it.

- The zero value for a map type is nil, that is, a reference to no hash table at
all.
```
var ages map[string]int
fmt.Println(ages == nil) // "true"
fmt.Println(len(ages) == 0) // "true"
```
Storing to a nil map causes a panic:
```
ages["carol"] = 21 // panic: assignment to entry in nil map
```
You must allocate the map before you can store into it.

- While accessing a map, if the key is present in the map, you get the
corresponding value; if not, you get the zero value for the element type.

- To check if a key is present in a map or not
```
age, ok := ages["bob"]
if !ok { /* "bob" is not a key in this map; age == 0. */}
```

- As with slices, maps cannot be compared to each other; the only legal
comparison is with nil. To test whether two maps contain the same keys and the
same assoicated values, we must wirte a loop:
```
func equal(x, y map[strint]int) bool {
  if len(x) != len(y) {
    return false
  }
  for k, xv := range x {
    if yv, ok := y[k]; !ok || yv != xv {
      return false
    }
  }
  return true
}
```

- Go does not provide a set type, but since the keys of a map are distinct, a
map can serve this purpose.

- Sometimes we need a map or set whose keys are slices, but because a map's keys
must be comparable, this cannot be expressed directly. However, it can be done
in two steps. First we define a helper function k that maps each key to a
string, with the property that k(x) == k(y) if and only if we consider x and y
equivalent. Then we create a map whose keys are strings, applying helper
function to each key before we access the map.

The example below uses a map to record the number of times Add has been called
with a given list of strings. It uses fmt.Sprintf to convert a slice of strings
into a single string that is a suitable map key.
```
var m = make(map[string]int)

func k(list []string) string {
  return fmt.Sprintf("%q", list)
}

func Add(list []string) {
  m[k(list)]++
}

func Count(list []string) int {
  return m[k(list)]
}
```

### Structs
- These two statements declare a struct type called Employee and a variable
called dilbert that is an instance of an Empoyee:
```
type employee struct {
  id int
  name string
  salary int
}
var dilbert employee
```
Given an employee's unique id, the function employeeById returns a pointer to an
employee struct.
```
func employeeById(id int) *employee { /* ... */ }
fmt.Println(employeeById(dilbert.id).salary)

id := dilbert.id
employeeById(id).salary = 0
```
The last statement updates the employee struct that is pointed to by the result
of the call to employeeByID. If the result type of employeeById were changed to
employee instead of *employee, the assignment statement would not compile since
its left-hand side would not identify a variable, as in:
```
func employeeById(id int) employee { /* ... */}
employeeById(id).salary = 0
```
```
$ go build -o ./bin/ struct-pointer.go
# command-line-arguments
./struct-pointer.go:19:26: cannot assign to employeeById(id).salary
```

- There are two forms of struct literal.
```
p := Point{1, 2}
p := Point{X: 1, Y: 2}
```
More often, the second form is used, in which a struct value is initialized by
listing some or all of the field names and their corresponding values.

- If all the fields of a struct are comparable, the struct itself is comparable,
so two expressions of that type may be compared using == or !=. The == operation
compares the corresponding fields of the two structs in order.

- Comparable struct types, like other comparable struct types, may be used as
the key type of a map.
```
type address struct {
  hostname string
  port     int
}
hits := make(map[address]int)
hits[address{"golang.org", 443}]++
```

- Struct can have a field with a type but no name; such fields are called
anonymous fields. The type of the field must be a named type or a pointer to a
named type. For example:
```
type Point struct {
  X, Y int
}
type Circle struct {
  Point
  Radius int
}
type Wheel struct {
  Circle
  Spokes int
}
```
```
var w Wheel
w.X = 8         // equivalent to w.Circle.Point.X = 8
w.Y = 8         // equivalent to w.Circle.Point.Y = 8
w.Radius = 5    // equivalent to w.Circle.Radius = 5
w.Spokes = 20
```

- Unfortunately, there's no corresponding shorthand for the struct literal
syntax, so neither of these will compile
```
w = Wheel{8, 8, 5, 20}
w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20}
```
The struct literal must follow the shape of the type declaration, so we must use
one of the two forms below
```
w = Wheel{Circle{Point{8, 8}, 5}, 20}
w = Wheel{
  Circle: Circle{
    Point: Point{X:8, Y: 8},
    Radius: 5,
  },
  Spokes: 20
}
```

## ch5: Functions
### Functions
- A function declaration has a name, a list of parameters, an optional list of
results, and a body.
```
func name(parameter-list) (result-list) {
  body
}
```

- The type of a function is sometimes called its signature. Two functions have
the same type or signature if they have the same sequence of parameter types and
the same sequence of result types. The names of parameters and results don't
affect the type, nor does whether or not they are declared using the factored
form.
```
func add(x int, y int) int { return x +y }
func sub(x, y int) (z int) { z = x - y; return }
func first(x int, _ int) (z int) { return x }
func zero(int, int) int { return 0 }

fmt.Printf("%T\n", add) // "func(int, int) int"
fmt.Printf("%T\n", sub) // "func(int, int) int"
fmt.Printf("%T\n", first) // "func(int, int) int"
fmt.Printf("%T\n", zero) // "func(int, int) int"
```

- Named results are treated as variables defined at the top of the function.

- Every function call must provide an argument for each parameter, in the order
in which the parameters were declared. Go has no concept of default values, nor
any way to specify arguments by name, so the names of parameters and results
don't matter to the caller except as documentation.

- Parameters are local variables within the body of the function, with their
initial values set to the arguments supplied by the caller. Function parameters
and named results are variables in the same lexical block as the function's
outermost local variables.

- Arguments are passed by value, so the function receives a copy of each
argument; modifications to the copy do not affect the caller. However, if the
argument contains some kind of reference, like a pointer, slice, map, function,
or channel, then the caller may be affected by any modifications the function
makes to variable indirectly referred to by the argument, for example:
```
func modifySlice(a []string, b string) {
	a[0] = b
}

func main() {
	a := []string{"a", "b"}
	fmt.Println(a) // [a, b]
	modifySlice(a, "b")
	fmt.Println(a) // [b, b]
}
```

### Errors
- Errors are an important part of a package's API or an application's user
interface, and failure are just one of several expected behaviors. This is the
approach Go takes to error handling.

- Go demands that more attention be paid to error-handling logic.

#### Error-handling strategies
There are five strategies for handling errors:

- First, and most common, is to propagate the error, so that a failure in a
subroutine becomes a failure of the calling routine.
```
doc, err := html.Parse(resp.Body)
resp.Body.Close()
if err != nil {
  return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
}
```
The fmt.Errorf function formats an error message using ftm.Sprintf and returns a
new error value.

- Because error messages are frequently chained together, message strings should
not be capitalized and newlines should be avoided. The resulting errors may be
long, but they will be self-contained when found by tools like grep.

- Second strategy is to retry the failed operation, possibly with a delay
between tries, and perhaps with a limit on the number of attempts or the time
spent trying before giving up entirely.

- Third, if progress is impossible, the caller can print the error and stop the
program gracefully, but this course of action should generally be reserved for
the main package of a program. Library functions should usually propagate errors
to the caller, unless the error is a sign of an internal inconsistency - that is
a bug.
```
if err := WaitForServer(url); err != nil {
  fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
}
```
A more convenient way to achieve the same effect is to call log.Fatalf. As with
all the log functions, by default it prefixes the time and data to the error
message.
```
if err := WaitForServer(url); err != nil {
  log.Fatalf("Site is down: %v\n", err)
}
```

- Fourth, in some cases, it's sufficient just to log the error and then
continue, perhaps with reduced functionality.
```
if err := Ping(); err != nil {
  log.Printf("ping failed: %v; networking disabled", err)
}
```

- And fifth and finally, in rare cases we can safely ignore an error entirely.
```
os.RemoveAll(dir) // ignore errors; $TMPDIR is cleaned periodically
```
In such cases the program logic would be the same had we forgotten to deal with
it. Get into the habit of considering errors after every function call, and when
you deliberately ignore one, document your intention clearly.

- After checking an error, failure is usually dealt with before success. If
failure causes the function to return, the logic for success is not indented
within an else block but follows at the outer level. Functions tend to exhibit a
common structure, with a series of initial checks to reject errors, followed by
the substance of the function at the end, minimally indented.

### Function Values
- Functions are first-class values in Go: like other values, functions values
have types, and they may be assigned to variables or passed to or returned from
functions. A function value may be called like any other function.
```
func sqaure(n int) int { return n * n }
func negative(n int) int { return -n }
func product(m, n int) int { return m * n }

f := square
fmt.Println(f(3)) // "9"

f = negative
fmt.Println(f(3)) // "-3"
fmt.Printf("%T\n", f) // "func(int) int"

f = product // compile error: can't assign func(int, int) int to func(int) int
```

- The zero value of a function type is nil. Calling a nil function value causes
a panic.
```
var f func(int) int
f(3) // panic: call of nil function
```

- Function values may be compared with nil:
```
var f func(int) int
if f != nil {
  f(3)
}
```
but they are not comparable, so they may not be compared against each other or
used as keys in a map.

- Function values let us parameterize our functions over not just data, but
behavior too. Meaning we can pass function as an argument to another functions.
```
func add1(r rune) rune { return r + 1 }
fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
```

### Anonymous Functions
- Named functions can be declared only at the package level, but we can use a
function literal to denote a function value within any expression. A function
literal is written like a function declaration, but without a name following the
func keyword. It is an expression, and its value is called an anonymous
function.

- Function literals let us define a function at its point of use. For example:
```
strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")
```
Functions defined in this way have access to the entire lexical environment, so
the inner function can refer to variables from the enclosing function, as this
example shows:
```
func squares() func() int {
  var x int
  return func() int {
    x++
    return x * x
  }
}
func main() {
  f := squares()
  fmt.Println(f()) // "1"
  fmt.Println(f()) // "4"
  fmt.Println(f()) // "9"
}
```
The function squares returns another function, of type func() int. A call to
squares creates a local variable x and returns an anonymous function that, each
time it is called, increments x and returns its square.

The squares example demonstrates that function values are not just code but can
have state. The anonymous inner function can access and update the local
variables of the enclosing function squares. These hidden variable references
are why we classify functions as reference types and why functions are not
comparable. Function values like these are implemented using a technique called
closures.

Here we can see that that the lifetime of variable is not determined by its
scope: the variable x exists after squares has returned within main, even though
x is hidden inside f.

### Variadic Functions
- A variadic function is one that can be called with varying numbers of
arguments.

- To declare a variadic function, the type of the final parameter is preceded by
an ellipsis, "...", which indicates that the function may be called with any
number of arguments of this type.
```
func sum(vals ...int) int {
  total := 0
  for _, val := range vals {
    total += val
  }
  return total
}
```
Within the body of the function, the type of vals is an []int slice.
```
fmt.Println(sum())    // "0"
fmt.Println(sum(3))   // "3"
fmt.Println(sum(1, 2, 3, 4))    // "10"
```
Implicitly, the caller allocates an array, copies the arguments into it, and
passes a slice of the entire array to the function.

- When the arguments are already in a slice, we can place an ellipsis after the
final argument as follows:
```
values := []int{1, 2, 3, 4}
fmt.Println(sum(values...))    // "10"
```

- Although the ...int parameter behaves like a slice within the function body,
the type of a variadic function is distinct from the type of a function with an
ordinary slice parameter.
```
func f(...int) {}
func (g[]int) {}

fmt.Printf("%T\n", f)   // "func(...int)"
fmt.Printf("%T\n", g)   // "func([]int)"
```

### Deferred function calls
- A defer statement is an ordinary function or method call prefixed by the
keyword defer. The function and argument expressions are evaluated when the
statement is executed, but the actual call is deferred until the function that
contains the defer statement has finished, whether normally, by executing a
return statement or falling off the end, or abnormally, by panicking. Any number
of calls may be deferred; they are executed in the reverse of the order in which
they were deferred.

- Each time a "defer" statement executes, the function value and parameters to
the call are evaluated as usual and saved anew but the actual function is not
invoked.
```
defer f(g())()
```
In this case the deferred function is the return value of f, so to evaluate the
deferred function value, f must be called, and to do that g must be called
prior. The returned value of f will be deferred(not called).

- A defer statement could be used to avoid code duplication.

- A defer statement is often used with paired operations like open and close,
connect and disconnect, or lock and unlock to ensure that resources are released
in all cases, no matter how complex the control flow. The right place for a
defer statement that releases a resource is immediately after the resource has
been successfully acquire.

- The defer statement can also be used to pair "on entry" and "on exit" actions
when debugging a complex function. The bigSlowOperation function below calls
trace immediately, which does the "on entry" action then returns a function
value that, when called, does tge corresponding "on exit" action. By deferring a
call to the returned function in this way, we can instrument the entry point and
all exit points of a function in a single statement and even pass values, like
the start time, between the two actions. But don't forget the final parentheses
in the defer statement, or the "on entry" will happen on exit and the on-exit
action won't happen at all!

- Because an anonymous function can access its enclosing function's variables,
including named results, a deferred anonymous function can observe the
function's results.

## ch6: Methods
- An object is simply a value or variable that has methods, and a method is a
function associated with a particular type. An object-oriented program is one
that uses methods to express the properties and operations of each data
structure so that clients need not access the object's representation directly.

- We will show how to define and use methods effectively. We will also cover two
key principles of object-oriented programming, encapsulation and composition.

### Composition
```
type Point  struct{ X, Y float64 }

type ColoredPoint struct {
  Point
  Color color.RGBA
}
```
We could have defined ColoredPoint as a struct of three fields, but instead we
embedded a Point to provide the X and Y fields.

### Encapsulation
- A variable or method of an object is said to be encapsulated if it is
inaccessible to clients of the object. Encapsulation, sometimes called
information hiding, is a key aspect of object-oriented programming.

- Capitalized identifiers are exported from the package in which they are
defined, and uncapitalized names are not. The same mechanism that limits access
to members of a package also limits access to the fields of a struct or the
methods of a type. As a consequence, to encapsulate an object, we must make it a
struct.

- The unit of encapsulation is the package, not the type as in many other
languages. The fields of a struct type are visible to all code within the same
package. Whether the code appears in a function, or a method makes no
difference.

- Encapsulation provides three benefits. First, because clients cannot directly
modify the object's variables, one need inspect fewer statements to understand
the possible values of those variables.

Second, hiding implementation details prevents clients from depending on things
that might change, which gives the designer greater freedom to evolve the
implementation without breaking API compatibility.

```
type Buffer struct {
  buf     []byte
  initial [64]byte
  /* ... */
}

// Grow exapnds the buffer's capacity, if necessary,
// to guarantee space for another n bytes. [...]
func (b *Buffer) Grow(n int) {
  if b.buf == nil {
    b.buf = b.initial[:0] // use preallocated space initially
  }
  if len(b.buf) + n > cap(b.buf) {
    buf := make([]byte, b.Len(), 2*cap(b.buf) + n)
    copy(buf, b.buf)
    b.buf = buf
  }
}
```

The third benefit of encapsulation, is that it prevents clients from setting an
object's variables arbitrarily. Because the object's variables can be set only
by functions in the same package, the author of that package can ensure that all
those functions maintain the object's internal invariants.

## ch7: Interfaces
- Concrete types: A concrete type specifies the exact representation of its
values and exposes the intrinsic operations of that representation, such as
arithmetic for numbers, or indexing, append, and range for slices. A concrete
type may also provide additional behaviors through its methods. When you have a
value of a concrete type, you know exactly what it is and what you can do with
it.

- Interface type: There is another kind of type in Go called an interface type.
An interface is an abstract type. It doesn't expose the representation or
internal structure of its values, or the set of basic operations they support;
it reveals only some of their methods. When you have a value of an interface
type, you know nothing about what it is; you know only what it can do, or more
precisely, what behaviors are provided by its methods.

- An interface type specifies a set of methods that a concrete type must posses
to be considered an instance of that interface.

```
package io

type Writer interface {
  Write(p []byte) (n int, err error)
}

var w  io.Writer
w = os.Stdout         // OK
w = new(bytes.Buffer) // OK
w = time.Second       // compile error: time.Duration lacks Write method
```

- Only the methods revealed by the interface type may be called, even if the
concrete type has others:
```
os.Stdout.Write([]byte("hello"))    // OK
os.Stdout.Close()                   // OK

var w io.Writer
w = os.Stdout
w.Write([]byte("hello"))  // OK
w.Close()                 // compile error: io.Writer lacks Close method
```

- An empty interface type `interface{}`, which has no methods at all, places no
demands on the types that satisfy it, we can assign any value to the empty
interface.
```
var any interface{}
any = true
any = 12.34
any = "hello"
any = map[]string{"one": 1}
any = new(bytes.Buffer)
```

- Each grouping of concrete types based on their shared behaviors can be
expressed as an interface type. Unlike class-based languages, in which the set
of interfaces satisfy by a class is explicit, in Go we can define new
abstractions or groupings of interest when we need them, without modifying the
declaration if the concrete type. This is particularly useful when the concrete
type comes from a package written by a different author. Of course, there do
need to be underlying commonalities in the concrete types.

- Interface values may be compared using == and !=. Two interface values are
equal if both are nil, or if their dynamic types are identical and their
dynamic values are equal according to the usual behavior of == for that type.
Only compare interface values if you are certain that they contain dynamic
values of comparable types.

### Type Assertions
- A type assertion is an operation applied to an interface value. Syntactically,
it looks like x.(T), where x is an expression of an inerface type and T is a
type, called the "asserted" type. A type assertion checks that the dynamic type
of its operand matches the asserted type.

- There are two possibilities. First, if the asserted type T is a concrete type,
then the type assertion checks whether x's dynamic type is identical to T. If
this check succeeds, the result of the type assertion is x's dynamic value,
whose type is of course T. In other words, a type assertion to a concrete type
extracts the concrete value from its operand. If the check fails, then the
operation panics. For example:
```
var w io.Writer
w = os.Stdout
f := w.(*os.File)       // success: f == os.Stdout
c := w.(*bytes.Buffer)  // panic: interface holds *os.File, not *bytes.Buffer
```

- Second, if instead the asserted type T is an interface type, then the type
assertion checks whether x's dynamic type satisfies T. If this check succeeds,
the dynamic value is not extracted; the result is still an interface value with
the same type and value components, but the result has the interface type T. In
other words, a type assertion to an interface type changes the type of the
expression, making a different (and usually larger) set of methods accessible,
but it preserves the dynamic type and value components inside the interface
value.

## ch8: Goroutines and Channels
### Goroutines
- In Go, each concurrently executing activity is called a goroutine.

- When a program starts, its only goroutine is the one that calls the main
function, so we call it the main goroutine.

### Channels
- If goroutines are the activities of a concurrent Go program, channels are the
connections between them. A channel is a communication mechanism that lets one
goroutine send values to another goroutine. Each channel is a conduit for values
of a particular type, called the channel's element type. The type of a channel
whose elements have type int is written `chan int`.

- To create a channel, we use built-in make function:
```
ch := make(chan int) // ch has type 'chan int'
```

- A channel is a reference to the data structure created by make. When we copy a
channel or pass one as an argument to a function, we are copying a reference, so
caller and calle refer to the same data structure. As with other reference
types, the zero value of a channel is nil.

- Two channel of the same type may be compared using ==. The comparison is true
if both are references to the same channel data structure. A channel may also be
compared to nil.

- A channel has two principal operations, send and receive, collectively known
as communications. A send statement transmits a value from one goroutine,
through the channel, to another goroutine executing a corresponding receive
expression. Both operations are written using the <- operator. In a send
statement, the <- separates the channel and value operands. In a receive
expression, <- preceds the channel operand. A receive expression whose result is
not used is a valid statement.
```
ch <- x     // a send statement
x = <- ch   // a receive expression in an assignment statement
<-ch        // a receive statement; result is discarded
```

- Channel supports a third operation, close, which sets a flag indicating that
no more values will ever be sent on this channel; subsequent attempts to send
will panic. Receive operations on a closed channel yield the values that have
been sent until no more values are left; any receive operations thereafter
complete immendiately and yield the zero value of the channel's element type.
```
close(ch)
```

### Unbuffered Channels
- If the capacity is zero or absent, the channel is unbuffered and communication
succeeds only when both a sender and a receiver are ready.

- A send operation on an unbuffered channel blocks the sending goroutine until
another goroutine executes a corresponding receive on the same channel, at
which point the value is transmitted and both goroutines may continue.
Conversely, if the receive operation was attempted first, the receieving
goroutine is blocked until another goroutine performs a send on the same
channel.

- Communiation over an unbuffered channel causes the sending and receiving
goroutines to synchronize. Because of this, unbuffered channels are sometimes
called synchronous channels. When a value is sent on an unbuffered channel, the
receipt of the value happens before the reawakening of the sending goroutine.

- In discussions of concurrency, when we say x happens before y, we don't mean
merely that x occurs earlier in time than y; we mean that it is guaranteed to do
so and that all its prior effects, such as updates to variables, are complete
and that you may rely on them.

- When x neither happens before y nor after y, we say that x is concurrent with
y. This doesn't mean that x and y are necessarily simultaneous, merely that we
cannot assume anything about their ordering.

-
```
func main() {
  conn, err := net.Dial("tcp", "localhost:8000")
  if err != nil {
    log.Fatal(err)
  }
}
done := make(chan struct{})
go func() {
  io.Copy(os.Stdout, conn)
  log.Println("done")
  done <- struct{}{}
}()
mustCopy(conn, os.Stdin)
conn.Close()
<-done
```
When the user closes the standard input stream, mustCopy returns and the main
goroutine calls conn.Close(), closing both halves of the network connection.
Closing the write half of the connection causes the server to see an end-of-file
condition. Closing the read half causes the background goroutine's call io.Copy
to return a "read from closed connection" error.

- Messages sent over channels have two important aspects. Each message has a
value, but sometimes the fact of communication and the moment at which it occurs
are just as important. We call messages events when we wish to stress this
aspect. When the event carries no additional information, that it, its sole
purpose is synchronization, we'll emphasize this by using a channel whose
element type is struct{}, though it's common to use a channel of bool or int for
the same purpose since `done <- 1` is shorter than `done <- struct{}{}`.

### Pipelines
- Channels can be used to connect goroutines together so that the output of one
is the input to another. This is called a pipeline.

- After a channel has been closed, any further send operations on it will panic.
After the closed channel has been drained, that is, after the last sent element
has been received, all subsequent receive operations will proceed without
blocking but will yield a zero value.

- There is no way to test directly whether a channel has been closed, but there
is a variant of the receive operation that produces two results: the received
channel element, plus a boolean value, conventionally called `ok`, which is
true for a successful receive and false for a receive on a closed and drained
channel.
```
fo func() {
  for {
    x, ok := <- naturals
    if !ok {
      break // channel was closed and drained
    }
    squares <- x * x
  }
  close(squares)
}
```
The language lets us use a range loop to iterate over channels too. This is a
more convenient syntax for receiving all the values sent on a channel and
terminating the loop after the last one.

- We needn't close every channel when we've finished with it. It's only
necessary to close a channel when it is important to tell the receiving
goroutines that all data have been sent. A channel that the garbage collector
determines to be unreachable will have its resources reclaimed whether or not it
is closed.

- Attempting to close an already-closed channel causes a panic, as does closing
a nil channel.

### Unidirectional Channel Types
- When a channel is supplied as a function parameter, it is nearly always with
the intent that it be used exclusively for sending or exclusively for receiving.
```
func counter(out chan int)
func suqarer(out, in chan int)
func printer(in chan int)
```
The names `in` and `out` convey the intention that `in`is used to receive the
input and `out` is used to send the output. But still, nothing prevents squarer
from sending to `in` or receiving from out.

To document this intent and prevent misuse, the Go type system provides
`unidirectional` channel types that expose only one or the other of the send
and receive operations. The type `chan<-int`, a `send-only` channel of int,
allows sends but not receives. Conversely, the type `<-chan int`, a
`receive-only` copy channel of int, allows receives but not sends. (The position
of the <- arrow relative to the chan keyword is a mnemonic.) Violations of this
discipline are detected at compile time.

Since the close operation asserts that no more sends will occur on a channel,
only the sending goroutine is in a position to call it, and for this reason it
is a compile-time error to attempt to close a receive-only channel.

- While calling the function with bidirectional channel type as an argument, the
conversion to unidirectional channel type happens implicitly. Conversions from
bidirectional to unidirectional channel types are permitted in any assignment.
There is no going back, however: once you have a value of a unidirectional type
such as `chan<-int`, there is no way to obtain from it a value of type
`chan int` that refers to the same channel data structure.

### Buffered Channels
- A buffered channel has a queue of elements. The queue's maximum size is
determined when it is created, by the capacity argument to make. The statement
below creates a buffered channel capable of holding three string values.
```
ch = make(chan string, 3)
```

- A send operation on a buffered channel inserts an element at the back of the
queue, and a receive operation removes an element from the front. If the channel
is full, the send operation blocks its goroutine until space is made available
by another goroutine's receive. Conversely, if the channel is empty, a receive
operation blocks until a value is sent by another goroutine.

- The example below shows an application of a buffered channel. It makes
parallel requests to three mirrors, that is, equivalent but geographically
distributed servers. It sends their responses over a buffered channel, then
receives and returns only the first response, which is the quickest one to
arrive.
```
func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { respones <- request("asia.gopl.io") }()
	go func() { respones <- request("europe.gopl.io") }()
	go func() { respones <- request("americas.gopl.io") }()
	return <-reponses // return the quickest reponse
}
```
Had we used an unbuffered channel, the two slower goroutines would have gotten
stuck trying to send their responses on a channel from which no goroutine will
ever receive. This situation, called a `goroutine leak`, would be a bug. Unlike
garbage variables, leaked goroutines are not automatically collected, so it is
important to make sure that goroutines terminate themselves when no longer
needed.

- The choice between unbuffered and buffered channels, and the choice of a
buffered channel's capacity, may both affect the correctness of a program.
Unbuffered channels give stronger synchronization guarantees because every send
operation is synchronized with its corresponding receive; with buffered
channels, these operations are decoupled. Also, when we know an upper bound on
the number of values, that will be sent on a channel, it's not unusual to create
a buffered channel of that size and perform all the sends before the first value
is received. Failure to allocate sufficient buffer capacity would cause the
program to deadlock.

### Multiplexing with Select
```
select {
  case <- ch1:
    // ...
  case x := <-ch2:
    // ...
  case ch3 <- y:
    // ...
  default:
    // ...
}
```
- Like a swtich statement, it has a number of cases and an optional default.
Each case specifies a communication (a send or receive operation on some
channel) and an associated block of statements. A receive expression may appear
on its own, as in the first case, or within a short variable declaration, as in
the second case; the second form lets you refer to the received value.

- A select waits until a communication for some case is ready to proceed. It
then performs that communication and executes the case's associated statements;
the other communications do not happen. A `select` with no cases, `select{}`,
waits forever.

- The time.After function immediately returns a channel, and starts a new
goroutine that sends a single value on that channel after the specified time.
The select statement below waits until the first of two events arrives, either
an abort event or the event indicating that 10 seconds have elapsed. If 10
seconds go by with no abort, the launch proceeds.
```
func main() {
  // ...create abort channel...

  fmt.Println("Commencing countdown. Press return to abort.")
  select {
  case <-time.After(10 * time.Second):
    // Do nothing.
  case <-abort:
    fmt.Println("Launch aborted!")
    return
  }
  launch()
}
```

- If multiple cases are ready, select picks one at random, which ensures that
every channel has an equal chance of being selected. Consider the following
example:
```
ch := make(chan int, 1)
for i:=0; i < 10; i++ {
  select {
    case x := <-ch:
      fmt.Println(x)
    case ch <- i:
  }
}
```
in this increasing the buffer size makes its output nondeterministic, because
when the buffer is neither full nor empty, the select statement figuratively
tosses a coin.

- Sometimes we want to try to send or receive on a channel but avoid blocking if
the channel is not ready - a non-blocking communication. A select statement can
do that too. A select may have a default, which specifies what to do when none
of the other communication can proceed immediately.
```
select {
  case <- abort:
    fmt.Printf("Launch aborted!\n")
    return
  default:
    // do nothing
}
```

- The zero value for a channel is nil. Perhaps surprisingly, nil channels are
sometimes useful. Because send and receive operations on a nil channel block
forever, a case in a select statement whose channel is nil is never selected.

## ch9: Concurrency with shared variables
- Consider a function that works correctly in a sequential program. That
function is concurrency-safe if it continues to work correctly even when called
concurrently, that is, from two or more goroutines with no additional
synchronization.

- A race condition is a situation in which the program does not give the correct
result for some interleavings of the operations of multiple goroutines.

- Data race: A data race occurs whenever two goroutines access the same variable
concurrently and at least one the accesses is a write. It follows from this
definition that there are three ways to avoid a data race.

- The first way is not to write the variable. Data structures that are never
modified or are immutable are inherently concurrency-safe and need no
synchronization.

- The second way to avoid a data race is to avoid accessing the variable from
multiple goroutines. Since other goroutines cannot access the variable directly,
they must use a channel to send the confining goroutine a request to query or
update the variable. This is what meant by mantra "Do not communicate by
sharing memory; instead, share memory by communicating".

- The third way to avoid a data race is to allow many goroutines to access the
variable, but only one at a time. This approach is known as mutual exclusion.

## Coding style
- Normal practice in Go is to deal with the error in the if block and then
return, so that the successful execution path is not indented.

## References
- Exercise solutions:
https://github.com/vinceyuan/gopl-solutions
https://github.com/torbiak/gopl

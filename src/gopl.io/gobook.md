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

- Every function call must provide an argument for each parameter, in the order
in which the parameters were declared. Go has no concept of default values, nor
any way to specify arguments by name.

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

### Anonymous Functions
- Named functions can be declared only at the package level, but we can use a
function literal to denote a function value within any expression.
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
time it is called, increments x and returs its square.

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

## Coding style
- Normal practice in Go is to deal with the error in the if block and then
return, so that the successful execution path is not indented.

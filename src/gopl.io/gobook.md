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
as an existing type.

- The named type provides a way to separate different and perhaps incompatible
uses of the underlying type so that they can't be mixed unintentionally.
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
variables of the same underlying type.

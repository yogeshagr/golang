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
another goroutine. The function main runs in a goroutine and the go statement
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

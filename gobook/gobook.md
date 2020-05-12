# The Go Programming Language

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

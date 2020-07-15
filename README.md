# Golang

## Golang best practices
- Avoid nesting by handling errors first. Less nesting means less cognitive load
on the reader.

- Avoid repitition when possible. Deploy one-off utility types for simpler code.

- Use switch to handle special cases.

- Writing everything or nothing.

- Shorter is better or at least longer is not always better.

- Package with multiple files.
  + Avoid very long files.
  + Separate code and tests.
  + Separated package documentation.

- Ask for what you need, don't ask extra.

- Testing using an interface instead of a concrete type makes testing easier.

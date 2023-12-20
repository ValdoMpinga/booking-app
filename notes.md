Absolutely, I'll enhance your notes in Markdown format. Here's the improved version:

# Go (Golang) Crash Course

Go, also known as Golang, was created by Google in 2007. It is open source and designed for efficiency, especially in concurrent programming.

## Introduction

- Go was developed to run on multiple cores and has built-in support for concurrency.
- Concurrency in Go is both cheap and easy.

## Main Use Case

The primary use case for Go is performance-oriented applications.

## Getting Started

To initialize a Go module, use the following command:

```bash
go mod init my-app
```

This command describes the module, including the name/module path and the Go version used in the program.

- All code must belong to a package, and the main package is indicated by `package main` at the beginning of the code.
- Unused variables in Go will result in an error.

## Slices in Go

A slice in Go is an abstraction of an array with a dynamic size. To declare and append to a slice:

```go
var bookings []string // Declaration
bookings = append(bookings, firstName+" "+lastName) // Appending
```

## Loops

In Go, there is only a `for` loop, but it can be used for various loop constructs.

### For-Each Loop

```go
for _, booking := range bookings {
    names := strings.Fields(booking) // Splits the string with white spaces as a separator
    firstNames = append(firstNames, names[0])
}
```

In Go, the underscore `_` is used to identify unused variables.

### While Loop Equivalent

```go
for <condition> {
    // Code
}
```

Or

```go
for {
    // Code
}
```

### Package Level Variables

In Go, package-level variables are defined at the top outside all functions. They can only be declared using the `var` keyword. It is a best practice to define variables as locally as possible.

## Go Packages

Go allows the creation of packages to organize code better. To import a package, use `<app module name in go.mod>/<module name>`. Exported functions must have the first letter capitalized.

## Variable Scopes in Go

In Go, variables have three levels of scope:

- Global: Accessible throughout the entire application.
- Package: Accessible within a package.
- Local: Accessible within a specific scope, such as a function.

## Maps

Maps in Go aggregate data with keys and values of the same data type for type safety, performance optimization, and simplicity.

## Structs

Structs are useful when dealing with different data types across values or nested data. They are particularly suitable for handling JSON.

To create a struct in Go:

```go
type StructName struct {
    Var1 string
    Var2 bool
    Var3 uint8
}
```

## Concurrency

Concurrency in Go allows code to run without blocking the main thread. Goroutines can be used to execute functions concurrently.

To make a function concurrent:

```go
go myFunction()
```

By default, the main goroutine does not wait for other goroutines to finish. To address this, use a `sync.WaitGroup` to make the main goroutine wait for others to finish.

### Concurrency in Go vs Other Languages

- Writing concurrent code in other languages is more complex and has more overhead.

### Threads vs Goroutines

- Creating threads is more expensive and has a slow startup time.
- Goroutines in Go are lightweight and efficient, allowing the creation of thousands rapidly.

Go uses green threads, abstractions of actual threads, making communication between them easy through channels.

- GO

Go also know for Golang, was created by Google in 2007, its open source and kinda new.

- Go was designed to run on multiple cores and build support concurrency
- Concurrency in Go is cheap and easy

The main use case of Go is for peromrant apps.

- Getting started

```bash
go mod init my-app
```

it describes the module: name/module path and go version used in the program

All the code must belong to a package

package main // in the first line of the code.

In go, if you declare a variable and dont use it youll get an error

* Slices in Go

Slice is an abstraction of an Array, which has a dynamic size, you declare it with an array withou a size , like this:

var bookings []string //declaration
bookings = append(bookings, firstName+" "+lastName) // to append

- Loops

In Go, there is only for loop, normal for, for each..., game is game.

for each loop in Go

for \_, booking := range bookings {
names := strings.Fields(booking) // splits the string with white spaces as separator
firstNames = append(firstNames, names[0])
}

in Go, _ (underscore) are used to identify unused variables.

We can also do for <condition> {

} //this is like a do-while writtin in for loop lmao

a while equivalent would be for { //code } or for true{ //code }

* Package level variables

In Go, we can declare variables that are accessible through files whitin a package to make the code cleaner.
- They are defined at the top outside all functions
- They can only be declared using the var keyword

PS: the best practice is to define a var locally as possible!!!

* Go Packages

In Go we can create packages to organize better the code. After creating a file we must tell to which package the file belongs to.

We can create multiple packages in Go, if we have multiple packages, its ideal to create a folders for each packages, to import a package, we should use <app module name in go.mod>/<module name>. Also the functions which must be accessible outside the package must exported, to do so, we only need to capitalize the first letter on the function.

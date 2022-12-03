// Compile and execute one or two file
go run 
go run main.go

// Just compile a bunch of go source code files
go build 
go build main.go

// format all the code in each file in current directory
go fmt

// compile and "install" a package
go install

// download the raw source code of someone else's package 
go get

// run any test associated with the current project
go test

# What does 'package main' mean?
### types of package
// Executable and Reusable
// Executable generates a file that we can run
// Reusable: Code used as 'helpers'. Good place to put reusable logic

// Package main -> run go build -> main.exe (Executable package)
// Package blahblah -> run go build -> nothing! (Reusable package)

// saying import, we can import standard package like fmt or reusable package make by other engineer
Compile and execute one or two file

```
go run
go run main.go
```

Just compile a bunch of go source code files

```
go build
go build main.go
```

Format all the code in each file in current directory

```
go fmt
```

Compile and "install" a package

```
go install
```

Download the raw source code of someone else's package

```
go get
```

run any test associated with the current project

```
go test
```

## What does 'package main' mean?

### types of package

// Executable and Reusable
// Executable generates a file that we can run
// Reusable: Code used as 'helpers'. Good place to put reusable logic

// package main -> run go build -> main.exe (Executable package)
// must define a function named main, which is run auto when the program runs
// package blahblah -> run go build -> nothing! (Reusable package)
// saying package blahblah means this file is a part of package blahblah

// saying import, we can import standard package like fmt or reusable package make by other engineer

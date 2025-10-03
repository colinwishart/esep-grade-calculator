# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

Runtime requirements:
Golang 1.24+

Check Golang version using:
```
go version
```

To install Golang follow the instructions at:
https://go.dev/doc/install

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```

Summary of Files:
 - go.mod: This file is necessary to use Go’s dependency management. It defines the module path as esep/grade-calculator and which version of Go the module should use.
 - grade_calculator.go: This go file implements the grade calculation functionality, including initializing a new calculator, getting a final grade, adding a new grade to the calculator, and computing the numerical grade.
 - grade_calculator_test.go: This go file defines unit tests for the functions and structures implemented in grade_calculator.go.
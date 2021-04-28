package main

import "fmt"

type T struct {
	noCopy
	A, B int
}

func print(t T) {
	fmt.Printf("%d%d", t.A, t.B)

}
func main() {
	a := T{A: 1, B: 9}
	b := a
	print(a)
	b.A = 2
	b.B = 6
	print(b)
}

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

/*
$ go run no-copy-struct/main.go
1926

$ go vet no-copy-struct/main.go
# command-line-arguments
no-copy-struct/main.go:10:14: print passes lock by value: command-line-arguments.T
no-copy-struct/main.go:16:7: assignment copies lock value to b: command-line-arguments.T
no-copy-struct/main.go:17:8: call of print copies lock value: command-line-arguments.T
no-copy-struct/main.go:20:8: call of print copies lock value: command-line-arguments.T
*/

package main

import (
	"basic/awspkg"
	"fmt"
)

func main() {
	// 👇 Print "hello"
	fmt.Println("hello")

	// 👇 Print output of <package-name-check-sts.go>.PrintCallerIdentity()
	fmt.Println(awspkg.PrintCallerIdentity())
}

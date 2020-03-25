package main

import (
	"fmt"
	"os"

	"github.com/yjiang-dev/simplemath/client/rpc"
)

func main() {
	if len(os.Args) == 1 {
		usage()
		os.Exit(1)
	}
	method := os.Args[1]
	switch method {
	case "gcd":
		if len(os.Args) < 4 {
			usage()
			os.Exit(1)
		}
		rpc.GreatCommonDivisor(os.Args[2], os.Args[3])
	default:
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Println("Welcome to Simple Math Client")
	fmt.Println("Usage:")
	fmt.Println("gcd num1 num2")
	fmt.Println("Enjoy")
}

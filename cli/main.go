package main

import (
	"github.com/csgrimes1/gidiom/gen"
	"fmt"
	"os"
)

func main() {
	fileName := os.Args[1]

	fmt.Println(fileName)
	gen.Compile(fileName)
}

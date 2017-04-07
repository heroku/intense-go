package main

import (
	"fmt"
	"github.com/jesperfj/intense-go/pkg/iplookup"
	"os"
)

func main() {
	fmt.Println(os.Args[1])
	fmt.Println(iplookup.Lookup(os.Args[1]))
	d := iplookup.Lookup2(os.Args[1])
	fmt.Println(d.Lat)
}

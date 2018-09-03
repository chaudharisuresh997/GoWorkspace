package main

import (
	"fmt"
	"onebigproject/commonpackages/stringutil"
)

func main() {
	s := "Hello youtube !!"
	s = stringutil.Reverse(s)

	fmt.Println(s)
}

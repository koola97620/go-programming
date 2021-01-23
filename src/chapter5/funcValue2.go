package main

import (
	"fmt"
	"strings"
)

func add2(r rune) rune { return r + 1 }
func main() {
	fmt.Println(strings.Map(add2, "VMS"))
}

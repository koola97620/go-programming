package main

import (
	"fmt"
	"strings"
)

func square(n int) int { return n * n }
func add1(r rune) rune { return r + 1 }
func main() {
	f1 := square
	fmt.Println(f1(3)) // 9 출력

	var f func(int) int // 함수 타입의 제로 값은 nil 이다.
	fmt.Println(f(3))   // panic: runtime error: invalid memory address or nil pointer dereference

	fmt.Println(strings.Map(add1, "VMS"))
}

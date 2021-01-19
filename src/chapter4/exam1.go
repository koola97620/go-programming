package main

import "fmt"

func main() {
	var a [3]int = [3]int{1, 2, 3}
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	// 인덱스와 원소 출력
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// 원소만 출력
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	m := make(map[string]int)
	m["alice"] = 31
	m["charlie"] = 34

	m2 := map[string]int{
		"alice":   31,
		"charlie": 33,
	}

	fmt.Print(m2["alice"])

}

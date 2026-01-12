package main

import (
	"fmt"
	"strings"
)

func main() {
	//문자열 합치기
	strs := []string{"a", "b", "c"}
	fmt.Println(strings.Join(strs, ":"))

	//문자열 대치
	str := "a.b.c"
	r := strings.Replace(str, ".", "_", -1)
	fmt.Println(r)
}
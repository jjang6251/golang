package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	strInt := "100"
	intStr := string(strInt)
	fmt.Println(intStr, reflect.TypeOf(intStr))

	i, err := strconv.Atoi(strInt)
	fmt.Println(intStr, reflect.TypeOf(i))

	strInt = "987654321"
	//strconv.ParseInt(문자열, 진법, 비트수)
	i8, err := strconv.ParseInt(strInt, 0, 8)
	i16, err := strconv.ParseInt(strInt, 0, 16)
	i32, err := strconv.ParseInt(strInt, 0, 32)
	i64, err := strconv.ParseInt(strInt, 16, 64)

	fmt.Println(i8, err, reflect.TypeOf(i8))
	fmt.Println(i16, err, reflect.TypeOf(i16))
	fmt.Println(i32, err, reflect.TypeOf(i32))
	fmt.Println(i64, err, reflect.TypeOf(i64))
}
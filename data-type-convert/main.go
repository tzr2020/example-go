package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 整数字符串->整型
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int, err:", err)
	} else {
		fmt.Printf("type: %T, value: %#v\n", i1, i1)
	}

	// 整型->整数字符串
	i2 := 100
	s2 := strconv.Itoa(i2)
	fmt.Printf("type: %T, value: %#v\n", s2, s2)

	// 布尔字符串->布尔型
	s3 := "t" // 1、0、t、f、T、F、true、false、True、False、TRUE、FALSE
	b1, err := strconv.ParseBool(s3)
	if err != nil {
		fmt.Println("can't convert to bool, err:", err)
	} else {
		fmt.Printf("type: %T, value: %#v\n", b1, b1)
	}

	// 整数字符串->整型-10进制64位
	i3, err := strconv.ParseInt("-100", 10, 64)
	if err != nil {
		fmt.Println("can't convert to int64, err:", err)
	} else {
		fmt.Printf("type: %T, value: %#v\n", i3, i3)
	}

	// 小数字符串->浮点型-64位
	f1, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		fmt.Println("can't convert to float64, err:", err)
	} else {
		fmt.Printf("type: %T, value: %#v\n", f1, f1)
	}
}

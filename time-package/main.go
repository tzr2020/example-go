package main

import (
	"fmt"
	"time"
)

func main() {
	// 时间类型
	now := time.Now()
	fmt.Printf("now of type: %T\n", now)
	fmt.Printf("now = %v\n", now)

	year := now.Year()
	fmt.Printf("year of type: %T\n", year)
	fmt.Printf("year = %v\n", year)

	// 时间戳类型
	timeStamp := now.Unix()
	// timeStamp := now.UnixNano()
	fmt.Printf("timeStamp of type: %T\n", timeStamp)
	fmt.Printf("timeStamp = %v\n", timeStamp)

	now = time.Unix(timeStamp, 0)
	fmt.Printf("now of type: %T\n", now)
	fmt.Printf("now = %v\n", now)

	// 时间间隔类型
	nsec := time.Nanosecond
	fmt.Printf("nsec of type: %T\n", nsec)
	fmt.Printf("nsec = %#v\n", nsec)
	hour := time.Hour
	fmt.Printf("hour of type: %T\n", hour)
	fmt.Printf("hour = %#v\n", hour)

	// 时间操作
	later := now.Add(time.Hour)
	fmt.Printf("later of type: %T\n", later)
	fmt.Printf("later = %v\n", later)
	sub := later.Sub(now)
	fmt.Printf("sub of type: %T\n", sub)
	fmt.Printf("sub = %v\n", sub)

	fmt.Println(later.Equal(now))
	fmt.Println(later.Before(now))
	fmt.Println(later.After(now))

	// 定时器
	ticker := time.Tick(time.Second)
	afterSec := time.Now().Add(3 * time.Second).Second()
	for t := range ticker {
		fmt.Println(t) //每秒都会执行的任务
		if t.Second() == afterSec {
			break // 3秒后结束
		}
	}

	// 时间格式化-200612345
	fmt.Println(now.Format("2006年01月02日 15时04分05秒")) // 24小时制
	fmt.Println(now.Format("2006年01月02日 03时04分05秒")) // 12小时制
}

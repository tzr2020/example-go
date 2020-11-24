package main

import (
	"encoding/json"
	"fmt"
)

// Server Struct
type Server struct {
	ServerName string
	ServerIP   string
}

// Servers Struct
type Servers struct {
	Servers []Server
}

// 解析JSON
func main() {

	// 解析到结构体
	var s Servers
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)

	// 解析到接口
	// 通过inteface{}和type assert的配合，解析未知结构的JSON
	b := []byte(`{"name":"Wed","Age":6,"Parents":["Gom","Mor"]}`)
	var f interface{}

	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)

	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string:", vv)
		case float64:
			fmt.Println(k, "is float64:", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}

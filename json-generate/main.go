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

// 生成JSON
func main() {

	var s Servers
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}

package main

import (
	"fmt"
	"net"
)
func main() {
	conn, err :=net.Dail("tcp","localhost:1234")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("successfully connected")
	b:=make([]byte,10)
	if _ ,err =conn.Write([]byte("G'olibbek")); err !=nil {
		panic(err)
	}
	if _, err = 
	
	
	conn.Read(b); err !=nil {
		panic (err)
	}
	fmt.Println(string(b))
}
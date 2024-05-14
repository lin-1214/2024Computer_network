package main
import (
	"net"
)

// func check(e error) { 
//     if e != nil { 
//         panic(e) 
//     } 
// } 
 
func main() { 
    ln, _ := net.Listen("tcp", ":31301")
	conn, _ := ln.Accept() 
    defer ln.Close() 
    defer conn.Close()
}
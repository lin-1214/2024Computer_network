package main 
 
import (
	"fmt" 
	"bufio" 
	"net" 
	"net/http"
	"os"
)
 
 
func check(e error) { 
    if e != nil { 
        panic(e) 
    } 
} 
 
func main() {
	
	fmt.Println("Launching server...") 
    ln, _ := net.Listen("tcp", ":12013") 
    defer ln.Close() 

	for {
		conn, _ := ln.Accept()
	
		reader := bufio.NewReader(conn) 
		req, err := http.ReadRequest(reader) 
		check(err)
		
		request_file := fmt.Sprintf(".%v", req.URL.String())
		f, err := os.Stat(request_file)

		if os.IsNotExist(err) {
			fmt.Println("File not found")
			
		} else {
			fmt.Printf("File size = %v\n", f.Size())
		}
		
		conn.Close()
	} 
}
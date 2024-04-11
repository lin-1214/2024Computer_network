package main
import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strconv"
	"strings"
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
		defer conn.Close()
		f_out, err := os.Create("whatever.newfile")
		check(err)

		cnt := 1

		reader := bufio.NewReader(conn)
		writer := bufio.NewWriter(f_out)
		socket_writer := bufio.NewWriter(conn)
		size, errr := reader.ReadString('\n')
		size = strings.Replace(size, "\n", "", -1)
		fmt.Printf("Upload file size: %v\n", size)
		n, err := strconv.Atoi(size)
		check(err)

		for {
			line, err := reader.ReadString('\n')
			check(err)
			
			content := strconv.Itoa(cnt) + " " + line
			writer.WriteString(content)
			writer.Flush()

			n = n-len(line) 
			cnt++
			if (n <= 0) {
				break
			}
		}

		stat, err := f_out.Stat()
		fmt.Printf("Output file size: %v\n", stat.Size())
		socket_writer.WriteString(fmt.Sprintf("%v bytes received, %v bytes file generated\n", size, stat.Size()))
		socket_writer.Flush()

		check(errr)
	}
}
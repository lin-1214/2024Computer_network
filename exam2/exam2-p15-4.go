package main
import (
	"net"
    "bufio"
    "fmt"
)

func check(e error) { 
    if e != nil { 
        panic(e) 
    } 
} 
 
func main() {

    // open server connection 
    ln, _ := net.Listen("tcp", ":31301")
	conn, _ := ln.Accept() 
    defer ln.Close() 
    defer conn.Close()

    // write message to client when receive message
    reader := bufio.NewReader(conn)
    writer := bufio.NewWriter(conn)
    message, err := reader.ReadString('\n')
    check(err)

    if (message == "PLAY\n") {
        writer.WriteString("Welcome to Hangman Not Quite! I have an English word for you to figure out. Give me a character. I will show where the character appears in the word.\n")
        writer.Flush()
    }

    // receive character
    ch := ""
    ch , _ = reader.ReadString('\n')

    fmt.Printf("%v", ch)
}
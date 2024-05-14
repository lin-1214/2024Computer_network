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
    

    guess := ""
    n := len("Internet")
    i := 0
    for {
        if (ch == fmt.Sprintf("%v\n", string("Internet"[i])) || ch == fmt.Sprintf("%v\n", string("internet"[i]))) {
            guess = guess + string("Internet"[i])
        } else {
            guess = guess + "_"
        }
        i = i + 1
        if (i >= n) {
            break
        }
    }

    writer.WriteString(guess)
    writer.Flush()
}
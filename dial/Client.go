package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	portAddress := os.Args[1]
	conn, err := net.Dial("tcp", ":"+portAddress)
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("In Line Memory Data Structure!!!!!!!!!")
	for {
		fmt.Print(">>>>>>>:")
		input, _ := inputReader.ReadString('\n')
		if len(input) != 2 {
			_, err = conn.Write([]byte(input))
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("<<<<<<<<<:" + message)
			continue
		}
		fmt.Println("Enter something")
	}
}

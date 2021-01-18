package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

var dataMap = make(map[string]string)

func main() {
	portAddress := os.Args[1]
	fmt.Println("starting the server ...")
	ln, err := net.Listen("tcp", ":"+portAddress)
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error listening", err.Error())
			return
		}
		go HandleClient(conn)
	}
}

//HandleClient handles the client connections and helps to work with In-memory Data Structure
func HandleClient(conn net.Conn) {
	defer conn.Close()
	Scanner := bufio.NewScanner(conn)
	for Scanner.Scan() {
		commands := Scanner.Text()
		cmd := strings.Fields(strings.ToUpper(commands))
		if len(cmd) < 2 {
			fmt.Fprintln(conn, "Enter Proper Values")
			continue
		}
		key := cmd[1]
		switch cmd[0] {
		case "SET":
			if len(cmd) != 3 {
				fmt.Fprintln(conn, "Enter Proper Values")
				continue
			}
			value := cmd[2]
			fmt.Fprintln(conn, SetData(key, value))
		case "GET":
			fmt.Fprintln(conn, GetData(key))
		case "DEL":
			fmt.Fprintln(conn, RemoveData(key))
		case "INC":
			fmt.Fprintln(conn, UpdateData(key, "INC"))
		case "DEC":
			fmt.Fprintln(conn, UpdateData(key, "DEC"))
		default:
			fmt.Fprintln(conn, "INVALID COMMAND")
		}
	}
}

//SetData sets the key value pairs into the dataMap
func SetData(key, value string) string {
	err := IsDataExists(key)
	if err != nil {
		dataMap[key] = value
		return "Variable Succesfully Added"
	}
	return "Already Exists"
}

//GetData gets the value of a given key data from the datamap
func GetData(key string) string {
	err := IsDataExists(key)
	if err == nil {
		return dataMap[key]
	}
	return "Variable Not Exists"
}

//IsDataExists checks whether the data exists or not in datamap
func IsDataExists(key string) error {
	if _, ok := dataMap[key]; ok {
		return nil
	}
	return errors.New("Data Not Exists")
}

//RemoveData removes the data from the datamap
func RemoveData(key string) string {
	if _, ok := dataMap[key]; ok {
		delete(dataMap, key)
		return "Data Deleted"
	}
	return "Variable Not Exists"
}

//UpdateData updates the value of a key data based on command i.e, increment and decrement
func UpdateData(key string, cmd string) string {
	if value, ok := dataMap[key]; ok {
		numberValue, _ := strconv.Atoi(value)
		if cmd == "INC" {
			numberValue++
		} else if cmd == "DEC" {
			numberValue--
		}
		dataMap[key] = strconv.Itoa(numberValue)
		return "Variable Updated"
	}
	return "Variable Not Exists"
}

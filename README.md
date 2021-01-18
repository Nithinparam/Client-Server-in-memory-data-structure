# Client-Server In Memory Data Structure
### Simple Redis like In Memory Data Structure using socket programming

##  Introduction
    Redis is a in-memory data structure store likewise here using socket programming we used to store data through client and server. Here client set data which is stored in server side and gets the data whenever clients needed

## Specifications
---
## SET 
    Set command is used to set the data or variable to the memory

        SET KEY_NAME VAlUE
## GET
    Get Command is used to get the data or variable that already stored in memory

        GET KEY_NAME
##  INC
    INC Command is used to increment the data value of a key data

        INC KEY_NAME
## DEC
    DEC Command is used to decrement the data value of a key data

        DEC KEY_NAME
##  DEL
    DEL Command is used to delete the data from memory

        DEL KEY_NAME
    

## Packages
    net/http
    bufio
	errors
	fmt
	os
	strconv
	strings

### [bufio.NewScanner](https://golang.org/pkg/bufio/#NewScanner)
    // NewScanner returns a new Scanner to read from r.
### [Scanner.scan()](https://golang.org/pkg/bufio/#Scanner.Scan)
    //Scan advances the Scanner to the next token, which will then be available through the Bytes or Text method. It returns false when the scan stops.
### [Scanner.Text()](https://golang.org/pkg/bufio/#Scanner.Text)
    //Text returns the most recent token generated by a call to Scan as a newly allocated string holding its bytes.
### [Strings.fields()](https://golang.org/pkg/strings/#Fields)
    //Fields splits the string around each instance of one or more consecutive white space characters, and returns slice of strings

## Set Up
    After having project files open two command prompt one for server another for client (you can have many number of command prompt for client side) 
    
    while running the project port number should be provide for both client and server side as a argument

***SERVER SIDE***
```GO
$cd listen/
$ go run sever.go 9000 <-- here 9000 is port number
starting the server ...
```

***CLIENT SIDE***
```Go
$cd dial/
$ go run client.go 9000 <-- here 9000 is port number
$ In Line Memory Data Structure!!!!!!!!!
>>>>>>>:SET A 1
<<<<<<<<<:Variable Succesfully Added
>>>>>>>:GET A
<<<<<<<<<:1
>>>>>>>:INC A
<<<<<<<<<:Variable Updated
>>>>>>>:DEC A
<<<<<<<<<:Variable Updated
>>>>>>>:DEL A
<<<<<<<<<:Data Deleted
>>>>>>>:GET A
<<<<<<<<<:Variable Not Exists
```



package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	// Open the file in write-only mode. If the file doesn't exist, it will be created.
	file, err := os.Create("sum.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Redirect the standard output to the file.
	// This will make all subsequent print statements write to the file.
	// To reset the output to the terminal, you can use os.Stdout again.
	// But here, we keep it redirected to the file for the entire program.
	// If you want to redirect only specific output, you can use file.Write() instead.
	os.Stdout = file

	// Your Go code here. For example, calculate and print the sum from 1 to 10.
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += i
	}
	fmt.Println("Sum is ", sum)

	// Flushing the buffer ensures that all the data is written to the file before closing it.
	// This step is not always necessary, but it ensures data integrity in some cases.

	processOnArray()
	// struct
	server_1 := Server{
		ServerName: "Gaugau",
		Counter:    1,
	}

	server_2 := Server{
		ServerName: "MeoMeo",
		Counter:    2,
	}

	server_3 := Server{
		ServerName: "Angang",
		Counter:    3,
	}
	fmt.Println("Server 1: ", server_1.ServerName, ", ", server_1.Counter)
	fmt.Println("Server 2: ", server_2.ServerName, ", ", server_2.Counter)
	fmt.Println("Server 3: ", server_3.ServerName, ", ", server_3.Counter)
	server_3.__Print__()

	file.Sync()
	time.Sleep(2 * time.Second)

	x := 'A'
	y := 'B'
	z := x + y
	fmt.Println(z)
}

func printHelloWorld() {
	fmt.Println("Hello World")
}

func processOnArray() {
	var mang = [5]int{10, 20, 30, 40, 50}
	for _, so := range mang {
		fmt.Println(so)
	}

	for idx := range mang {
		fmt.Println(idx)
		fmt.Println(mang[idx])
	}
}

func processOnMapOrHashtable() {

}

// Khai báo một cấu trúc với tên là "Server"
type Server struct {
	ServerName string
	Counter    int
}

func (s Server) __Print__() string {
	result := "Method of struct Server"
	fmt.Println(result)
	return result
}

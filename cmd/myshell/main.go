package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	// Uncomment this block to pass the first stage
	for(true){
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err!=nil{
			fmt.Println("Error reading input")
		}
		if command=="exit 0\n"{
			break
		}
		if "echo " == command[:5]{
			fmt.Println(command[5:len(command)-1])
			continue
		}
		fmt.Println(command[:len(command)-1]+": command not found ")
	}
	os.Exit(0)
}

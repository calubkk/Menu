package main

import "fmt"

func main() {
	for {
		fmt.Print("Please input a command: ")
		var cmd string
		fmt.Scanln(&cmd)
		switch cmd {
		case "help":
			fmt.Println("This is help command.")
			fmt.Println("We only have command 'hello' and 'quit'.")
		case "hello":
			fmt.Println("Hello.")
		case "quit":
			fmt.Println("Program ended.")
			return
		default:
			fmt.Println("Wrong!")
		}
	}
}

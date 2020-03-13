package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func commandList() {
	fmt.Println("Welcome to the Library!")
	fmt.Println("=============================================================================")
	fmt.Println("Here are the commands that you can input:")
	fmt.Println("1. list                              -> to display all book and its status")
	fmt.Println("2. get [code_of_book]                -> to show name of book by code")
	fmt.Println("3. add [code_of_book] [name_of_book] -> to add new book")
	fmt.Println("4. rent [code_of_book]               -> to update status of book rented")
	fmt.Println("5. return [code_of_book]             -> to update status of book returned ")
	fmt.Println("6. rented                            -> to display all rented books")
	fmt.Println("7. exit                              -> to exit the program")
	fmt.Println("=============================================================================")
}

func getCommand() string {
	fmt.Print("Command : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	command := scanner.Text()
	return command
}

func cmdList(commands []string) {
	if len(commands) > 1 && commands[1] != "" {
		fmt.Println("Illegal parameter")
	} else {
		fmt.Println("Process")
	}
}

func cmdGet(commands []string) {
	if len(commands) > 2 && commands[1] != "" {
		fmt.Println("Illegal parameter")
	} else {
		fmt.Println("Process")
	}
}

func cmdAdd(commands []string) {
	if len(commands) > 3 && commands[1] != "" {
		fmt.Println("Illegal parameter")
	} else {
		fmt.Println("Process")
	}
}

func cmdRent(commands []string) {
	if len(commands) > 2 && commands[1] != "" {
		fmt.Println("Illegal parameter")
	} else {
		fmt.Println("Process")
	}
}

func cmdReturn(commands []string) {
	if len(commands) > 2 && commands[1] != "" {
		fmt.Println("Illegal parameter")
	} else {
		fmt.Println("Process")
	}
}

func cmdRented(commands []string) {
	if len(commands) > 2 && commands[1] != "" {
		fmt.Println("Illegal parameter")
	} else {
		fmt.Println("Process")
	}
}

func main() {
	for {
		clear()
		commandList()
		command := getCommand()
		leadCloseWhitespace := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`) //whitespace in the beginning and end of string
		insideWhitespace := regexp.MustCompile(`[\s\p{Zs}]{2,}`)               //more than 1 whitespace
		cleanCommand := leadCloseWhitespace.ReplaceAllString(command, "")
		cleanCommand = insideWhitespace.ReplaceAllString(cleanCommand, " ")

		commands := strings.Split(cleanCommand, " ")

		switch commands[0] {
		case "list":
			cmdList(commands)
		case "get":
			cmdGet(commands)
		case "add":
			cmdAdd(commands)
		case "rent":
			cmdRent(commands)
		case "return":
			cmdReturn(commands)
		case "rented":
			cmdRented(commands)
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Command not found!")
		}

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
	fmt.Println("=============================================================================")
}

func getCommand() string {
	fmt.Print("Command : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	command := scanner.Text()
	return command
}

func main() {
	for {
		clear()
		commandList()
		command := getCommand()
		commands := strings.Split(command, " ")

		fmt.Println(commands)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
	}
}

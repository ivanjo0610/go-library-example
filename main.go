package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type Book struct {
	Name   string
	Status int
}

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

func cmdList(books map[string]Book, commands []string) {
	if len(commands) > 1 && commands[1] != "" {
		fmt.Println("Illegal parameter")
	} else {
		fmt.Println("List of Books ([code] - [name] [status]) : ")
		ctr := 0
		for code, book := range books {
			ctr++
			status := ""
			if book.Status == 1 {
				status = "(Rented)"
			}
			fmt.Printf("%d. %s - %s %s\n", ctr, code, book.Name, status)
		}
		if ctr == 0 { //if no books
			fmt.Println("This library currently doesn't have any books")
		}
	}
}

func cmdGet(books map[string]Book, commands []string) {
	if (len(commands) > 2 && commands[2] != "") || len(commands) < 2 {
		fmt.Println("Illegal parameter")
	} else {
		if book, ok := books[commands[1]]; ok {
			fmt.Printf("Book name : %s\n", book.Name)
		} else {
			fmt.Println("Book not found!")
		}
	}
}

func cmdAdd(books map[string]Book, commands []string) map[string]Book {
	if len(commands) < 3 {
		fmt.Println("Illegal parameter")
	} else {
		// join commands index 2 .. len - 1 to form the name of the book
		name := strings.Join(commands[2:len(commands)], " ")

		var book Book
		book.Name = name
		book.Status = 0

		books[commands[1]] = book
		fmt.Println("Book added!")
	}

	return books
}

func cmdRent(books map[string]Book, commands []string) map[string]Book {
	if (len(commands) > 2 && commands[2] != "") || len(commands) < 2 {
		fmt.Println("Illegal parameter")
	} else {
		if book, ok := books[commands[1]]; ok {
			if book.Status == 1 {
				fmt.Printf("%s is already rented\n", book.Name)
			} else {
				book.Status = 1
				books[commands[1]] = book
				fmt.Printf("%s rented\n", book.Name)
			}
		} else {
			fmt.Println("Book not found!")
		}
	}

	return books
}

func cmdReturn(books map[string]Book, commands []string) map[string]Book {
	if (len(commands) > 2 && commands[2] != "") || len(commands) < 2 {
		fmt.Println("Illegal parameter")
	} else {
		if book, ok := books[commands[1]]; ok {
			if book.Status == 0 {
				fmt.Printf("%s is not rented\n", book.Name)
			} else {
				book.Status = 0
				books[commands[1]] = book
				fmt.Printf("%s returned\n", book.Name)
			}
		} else {
			fmt.Println("Book not found!")
		}
	}

	return books
}

func cmdRented(books map[string]Book, commands []string) {
	if len(commands) > 1 && commands[1] != "" {
		fmt.Println("Illegal parameter")
	} else {
		fmt.Println("List of Rented Books ([code] - [name]) : ")
		ctr := 0
		for code, book := range books {
			if book.Status == 1 {
				ctr++
				fmt.Printf("%d. %s - %s\n", ctr, code, book.Name)
			}
		}
		if ctr == 0 { //if no books rented
			fmt.Println("Currently no books is rented!")
		}
	}
}

func main() {
	books := make(map[string]Book)

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
			cmdList(books, commands)
		case "get":
			cmdGet(books, commands)
		case "add":
			books = cmdAdd(books, commands)
		case "rent":
			books = cmdRent(books, commands)
		case "return":
			books = cmdReturn(books, commands)
		case "rented":
			cmdRented(books, commands)
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Command not found!")
		}

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
	}
}

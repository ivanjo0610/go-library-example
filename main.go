package main

import (
	"bufio"
	"errors"
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

func cmdList(books map[string]Book, commands []string) error {
	if len(commands) > 1 && commands[1] != "" {
		return errors.New("Illegal parameter")
	}
	if len(books) == 0 { //if no books
		return errors.New("This library currently doesn't have any books")
	}
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
	return nil
}

func cmdGet(books map[string]Book, commands []string) (string, error) {
	if (len(commands) > 2 && commands[2] != "") || len(commands) < 2 {
		return "", errors.New("Illegal parameter")
	}
	if book, ok := books[commands[1]]; ok {
		return book.Name, nil
	}
	return "", errors.New("Book not found")
}

func cmdAdd(books map[string]Book, commands []string) (map[string]Book, error) {
	if len(commands) < 3 {
		return books, errors.New("Illegal parameter")
	}
	// join commands index 2 .. len - 1 to form the name of the book
	name := strings.Join(commands[2:len(commands)], " ")

	// check if code already exist
	if _, ok := books[commands[1]]; ok {
		return books, errors.New("Failed to add book! Book code already exist")
	}

	var book Book
	book.Name = name
	book.Status = 0

	books[commands[1]] = book
	return books, nil
}

func cmdRent(books map[string]Book, commands []string) (map[string]Book, error) {
	if (len(commands) > 2 && commands[2] != "") || len(commands) < 2 {
		return books, errors.New("Illegal parameter")
	}

	book, ok := books[commands[1]]
	if !ok {
		return books, errors.New("Book not found")
	}
	if book.Status == 1 {
		errMessage := fmt.Sprintf("Failed to rent book! %s is already rented", book.Name)
		return books, errors.New(errMessage)
	}
	book.Status = 1
	books[commands[1]] = book

	return books, nil
}

func cmdReturn(books map[string]Book, commands []string) (map[string]Book, error) {
	if (len(commands) > 2 && commands[2] != "") || len(commands) < 2 {
		return books, errors.New("Illegal parameter")
	}
	if book, ok := books[commands[1]]; ok {
		if book.Status == 0 {
			errMessage := fmt.Sprintf("Failed to return book! %s is not rented", book.Name)
			return books, errors.New(errMessage)
		}
		book.Status = 0
		books[commands[1]] = book

		return books, nil
	}
	return books, errors.New("Book not found")
}

func cmdRented(books map[string]Book, commands []string) error {
	if len(commands) > 1 && commands[1] != "" {
		return errors.New("Illegal parameter")
	}
	ctr := 0
	for code, book := range books {
		if book.Status == 1 {
			if ctr == 1 { //header for the first (in if so it is not written if no books)
				fmt.Println("List of Rented Books ([code] - [name]) : ")
			}
			ctr++
			fmt.Printf("%d. %s - %s\n", ctr, code, book.Name)
		}
	}
	if ctr == 0 { //if no books rented
		return errors.New("Currently no books is rented")
	}
	return nil
}

func main() {
	books := make(map[string]Book)
	var err error
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
			err = cmdList(books, commands)
			if err != nil {
				fmt.Println(err.Error())
			}
		case "get":
			name, err := cmdGet(books, commands)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("Book name : %s\n", name)
			}
		case "add":
			books, err = cmdAdd(books, commands)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Book added!")
			}
		case "rent":
			books, err = cmdRent(books, commands)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("%s rented\n", books[commands[1]].Name)
			}
		case "return":
			books, err = cmdReturn(books, commands)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("%s returned\n", books[commands[1]].Name)
			}
		case "rented":
			err = cmdRented(books, commands)
			if err != nil {
				fmt.Println(err.Error())
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Command not found!")
		}

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
	}
}

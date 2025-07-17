package controllers

import (
	"library_management/models"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var reader = bufio.NewReader(os.Stdin)


func validateInput(input string) bool {
	return strings.TrimSpace(input) != ""
	
}

func getInput(msg string) string {
    printBlueInput(msg)
    for {
        input, err := readLine()
        if err != nil {
            PrintError("Error reading input, please try again")
            continue
        }
        if validateInput(input) {
            return input
        } else {
            PrintError("Invalid input, please enter again")
        }
    }
}

func getint(msg string) int{
	for {
		input := getInput(msg)
		value, error := strconv.Atoi(input) // asci to integer
		if error == nil{
			return value
	
		}else if value <= 0 {
			PrintError("invalid input, please enter positive number")
			 
		}
		PrintError("invalid input, please enter integer number again")
	}

	}

func printbookList(books []models.Book) {
    if len(books) == 0 {
        PrintError("No books found.")
        return
    }
    fmt.Printf("\033[32m%-5s | %-20s | %-20s | %-10s\033[0m\n", "ID", "Title", "Author", "Status")
    fmt.Println(strings.Repeat("-", 65))
    for _, book := range books {
        fmt.Printf("\033[34m%-5d | %-20s | %-20s | %-10s\033[0m\n", book.ID, book.Title, book.Author, book.Status)
    }
}


func readLine() (string, error){
	input, error:= reader.ReadString('\n')
	return strings.TrimSpace(input), error
}




func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
func PrintError(msg string) {
	fmt.Printf("\033[31m%s\033[0m\n", msg) // Red
}

func PrintSuccess(msg string) {
	fmt.Printf("\033[32m%s\033[0m\n", msg) // Green
}

func PrintBlue(msg string) {
	fmt.Printf("\033[36m%s\033[0m\n", msg) // Blue
}
func printBlueInput(msg string){
	fmt.Printf("\033[36m%s:\033[0m ", msg) // Blue without new line

}

func PrintBold(msg string) {
	fmt.Printf("\033[1m%s\033[0m\n", msg) 
}

func PrintUnderline(msg string) {
	fmt.Printf("\033[4m%s\033[0m\n", msg) // undeline
}
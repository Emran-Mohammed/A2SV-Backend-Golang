package controllers

import(
	"fmt"
	"library_management/models"
    "library_management/services"
)


var currentLibrary = NewLibrary()

func NewLibrary() *services.Library{ 
	return &services.Library{
		Books: make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}


func App(){
		for {
        ClearScreen()
        PrintBold("welcome to library managment system")
        fmt.Println("1. Add Book")
        fmt.Println("2. Add Member")
        fmt.Println("3. Remove Book")
        fmt.Println("4. Borrow Book")
        fmt.Println("5. Return Book")
        fmt.Println("6. List Available Books")
        fmt.Println("7. List Borrowed Books")
        fmt.Println("0. Quit")
        printBlueInput("Select an option")
        choice, _ := readLine()
        switch choice {
        case "1":
            AddBook()
        case "2":
            AddMember()
        case "3":
            RemoveBook()
        case "4":
            BorrowBook()
        case "5":
            ReturnBook()
        case "6":
            ListAvailableBooks()
        case "7":
            ListBorrowedBooks()
        case "0":
			PrintSuccess("exiting ...")
            return
        default:
            PrintError("Invalid option. Please try again.")
        }
        fmt.Println("Press Enter to continue...")
        readLine()
    }
}


func AddBook(){
	ClearScreen()
	PrintBold("Add a book")
	title := getInput("Please Enter the title")
	author := getInput("Please Enter the author")
	currentLibrary.AddBook(models.Book{
		Title: title, Author: author, Status: models.Available,
	})
	PrintSuccess("you add the book succesfully")

}
func AddMember(){
	ClearScreen()
	PrintBold("Add a Member")
	name:= getInput("please enter your name")
	currentLibrary.AddMember(models.Member{
		Name: name,
	})

	PrintSuccess("you add the member succesfully")

}
func RemoveBook(){
	ClearScreen()
	PrintBold("Remove a book")
	bookID := getint("please enter the id of the book you want to remove")
	currentLibrary.RemoveBook(bookID)


	PrintSuccess("you remove the book succesfully")

}
func BorrowBook(){
	ClearScreen()
	PrintBold("Borrow a book")
	bookID := getint("please enter the id of the book you want to borrow")
	memberID := getint("please enter member id")
	err := currentLibrary.BorrowBook(bookID, memberID)
	if err != nil{
		PrintError(err.Error()) // .Error() change the error to a stirng
	}else{
		PrintSuccess("you return the book succesfully")
	}

}
func ReturnBook(){
	ClearScreen()
	PrintBold("Return a book")
	bookID := getint("please enter the id of the book you want to return")
	memberID := getint("please enter member id")
	err := currentLibrary.ReturnBook(bookID, memberID)
	if err != nil{
		PrintError(err.Error()) // .Error() change the error to a stirng
	}else{
		PrintSuccess("you return the book succesfully")
	}


}
func ListAvailableBooks(){
	ClearScreen()
	PrintBold("List of Availabe books")
	newList := currentLibrary.ListAvailableBooks()
	printbookList(newList)

}
func ListBorrowedBooks(){
	ClearScreen()
	PrintBold("List of borrowed books")
	memberID := getint("please enter member id")
	newList, err := currentLibrary.ListBorrowedBooks(memberID)
	if err!= nil{
		PrintError(err.Error())
	}else{
		printbookList(newList)

	}

}
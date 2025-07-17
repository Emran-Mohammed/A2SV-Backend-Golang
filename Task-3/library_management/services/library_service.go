package services
import(
	"library_management/models"
	"fmt"
)

type LibraryManager interface{
	AddBook(book models.Book)
	RemoveBook (bookID int) error
	BorrowBook (bookID int, memberID int) error
	ReturnBook (bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book


	AddMember(member models.Member)
}

type Library struct{
	Books map[int]models.Book
	Members map[int]models.Member
	nextBookID int
	nextMemberID int
}

func (l *Library) AddBook(book models.Book){
	l.nextBookID ++
	book.ID= l.nextBookID
	l.Books[book.ID] = book
}
func (l *Library) AddMember(member models.Member){
	l.nextMemberID ++
	member.ID = l.nextMemberID
	l.Members[member.ID] = member
}

func (l *Library)RemoveBook(bookID int) {
	// we do not check wheater the book is found or not since our goal is to remove it
	delete(l.Books, bookID)
	for id , member := range l.Members{
		for i, book := range member.BorrowedBooks{
			if book.ID == bookID{
				member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
				break
			}

		}
		l.Members[id] = member

	}

}

func ( l *Library) BorrowBook(bookID int, memberID int) error{
	_ , exists := l.Members[memberID]
	if !exists {
		return fmt.Errorf("error, sorry you are not the member of the Library")
	}
	book , exists := l.Books[bookID]
	if !exists {
		return fmt.Errorf("error,sorry the booK with %v id is not found in the library", bookID)
	} else if book.Status == models.Available{
		
		book.Status = models.Borrowed
		l.Books[bookID] = book

		member := l.Members[memberID]
		member.BorrowedBooks = append(member.BorrowedBooks , book)
		l.Members[memberID] = member
		return nil
	}else{
		return fmt.Errorf("error, sorry the booK with %v id is not Avialable", bookID)

	}
	
}

func (l *Library) ReturnBook(bookID int, memberID int) error{
	member , exists := l.Members[memberID]
	if !exists {
		return fmt.Errorf("error, sorry you are not the member of the Library")
	}
	found := false
	var bookIndex int
	for index, book := range member.BorrowedBooks{
		if book.ID == bookID{
			bookIndex = index
			found = true
			break
		}
	}
	if !found{
		return fmt.Errorf("you did not borrow a book with %v id", bookID)
	}
	book , exists := l.Books[bookID]
	if !exists {
		return fmt.Errorf("error,sorry the booK with %v id is not found in the library", bookID)
	}else if  book.Status == models.Available{
		return fmt.Errorf("you did not borrow a book with %v id", bookID)

	}else if book.Status == models.Borrowed{
		
		book.Status = models.Available
		l.Books[bookID] = book
		newBorrowList := member.BorrowedBooks

		newBorrowList= append(newBorrowList[:bookIndex], newBorrowList[bookIndex+1:]...)
		member.BorrowedBooks = newBorrowList
		l.Members[memberID] = member

		return nil
	
	}
	return fmt.Errorf("somthing gone wrong")
}

func (l * Library)ListAvailableBooks() []models.Book{
	var availableBooks []models.Book
	for _, book := range l.Books{
		if book.Status == models.Available{
			availableBooks = append(availableBooks, book)

		}
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) ([]models.Book, error){
	var borrowedBooks []models.Book
	member , exists := l.Members[memberID]
	if !exists {
		return borrowedBooks, fmt.Errorf("error, sorry you are not the member of the Library")
	} else{
		return member.BorrowedBooks, nil
	}


}



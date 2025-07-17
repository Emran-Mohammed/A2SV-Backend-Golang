# Library Management System (Go)

## Overview

This project is a simple **Library Management System** written in Go. It demonstrates basic CRUD operations, user input handling, and struct-based design using Go’s standard library. The system allows you to add, remove, borrow, and return books, as well as manage library members.

---

## Project Structure

```
library_management/
│
├── controllers/
│   ├── helpers.go
│   └── library_controller.go
│
├── docs/
│   └── documentation.md
│
├── models/
│   ├── book.go
│   └── member.go
│
├── services/
│   └── library_service.go
│
├── go.mod
└── main.go
```

---

## How to Clone and Run

1. **Clone the repository:**
   ```bash
   git clone https://github.com/Emran-Mohammed/A2SV-Backend-Golang.git
   cd A2SV-Backend-Golang/Task-3/library_management
   ```

2. **(Optional) Initialize Go modules (if not already):**
   ```bash
   go mod tidy
   ```

3. **Run the application:**
   ```bash
   go run main.go
   ```

4. **Follow the on-screen menu to interact with the system.**

---

## Main Components

### 1. **models**

- **book.go**  
  Defines the `Book` struct and the `Status` type (`Available`, `Borrowed`).

- **member.go**  
  Defines the `Member` struct, which includes an ID, Name, and a list of borrowed books.

---

### 2. **services**

- **library_service.go**  
  Contains the `Library` struct and methods for:
  - Adding/removing books and members
  - Borrowing and returning books
  - Managing the internal state of the library

---

### 3. **controllers**

- **library_controller.go**  
  Handles user interaction and calls service methods. Functions include:
  - `AddBook`, `RemoveBook`, `AddMember`, `BorrowBook`, `ReturnBook`, `ListAvailableBooks`, etc.
  - Uses helper functions for input/output and validation.

- **helpers.go**  
  Utility functions for:
  - Reading and validating user input
  - Printing colored and formatted output
  - Clearing the terminal screen

---

### 4. **main.go**

- Entry point of the application.
- Calls `controllers.App()` to start the interactive menu.

---

## Features

- **Add Book:**  
  Add a new book with title and author. Status defaults to `Available`.

- **Remove Book:**  
  Remove a book by its ID from the library and from all members’ borrowed lists.

- **Add Member:**  
  Register a new member by name.

- **Borrow Book:**  
  A member can borrow a book if it is available.

- **Return Book:**  
  A member can return a borrowed book.

- **List Available Books:**  
  Display all books with their status in a formatted, colored table.

- **Input Validation:**  
  All user input is validated for correctness (e.g., integer IDs, non-empty strings).

---

## Example Usage

- Add a book:
  ```
  Please Enter the title: Go Programming
  Please Enter the author: Alan A. A. Donovan
  you add the book successfully
  ```

- Borrow a book:
  ```
  please enter the id of the book you want to borrow: 1
  please enter member id: 1
  you return the book successfully
  ```

- List available books:
  ```
  ID    | Title                | Author               | Status    
  -------------------------------------------------------------
  1     | Go Programming       | Alan A. A. Donovan   | borrowed
  ```

---

## Notes

- The system uses in-memory data structures (maps and slices); data is not persisted.
- The terminal output uses ANSI color codes for better readability.
- The code is modular and easy to extend for more features (e.g., book search, due dates).

---

## Author

- Emran Seid Mohammed
- github.com/Emran-Mohammed
---

## License

This project is for educational
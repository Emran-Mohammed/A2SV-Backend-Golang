# Task Manager API Documentation

---

## Overview

The **Task Manager API** is a RESTful service built with Go and Gin for managing tasks.  
It supports creating, reading, updating, and deleting tasks, each with a title, description, due date, and status.

---

## Folder Structure

```
task_manager/
├── main.go
├── controllers/
│   └── task_controller.go
├── models/
│   └── task.go
├── data/
│   └── task_service.go
├── router/
│   └── router.go
├── docs/
│   └── api_documentation.md
└── go.mod
```

---

## How to Run

1. **Clone the repository:**
   ```bash
   git clone https://github.com/Emran-Mohammed/A2SV-Backend-Golang.git
   cd A2SV-Backend-Golang/Task-4/task_manager
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the server:**
   ```bash
   go run main.go
   ```
   The API will be available at `http://localhost:8080`.

---

## Models

### Task

```go
type Task struct {
    ID          int       `json:"id"`
    Title       string    `json:"title" binding:"required"`
    Description string    `json:"description"`
    DueDate     time.Time `json:"duedate" binding:"required"`
    Status      Status    `json:"status" binding:"required"`
}
```

- **Status** must be one of: `"progress"`, `"completed"`, `"pending"`.

---

## API Endpoints

### 1. Get All Tasks

- **GET** `/tasks`
- **Response:**  
  `200 OK`  
  Returns a list of all tasks.

### 2. Get Task by ID

- **GET** `/tasks/:id`
- **Response:**  
  `200 OK`  
  Returns the task with the specified ID.  
  `400 Bad Request` if ID is invalid or not found.

### 3. Create Task

- **POST** `/tasks`
- **Request Body Example:**
  ```json
  {
    "title": "Write documentation",
    "description": "Complete the API documentation for the project.",
    "duedate": "2024-07-25T00:00:00Z",
    "status": "progress"
  }
  ```
- **Response:**  
  `201 Created`  
  Returns the created task.  
  `400 Bad Request` if required fields are missing or status is invalid.

### 4. Update Task

- **PUT** `/tasks/:id`
- **Request Body Example:**
  ```json
  {
    "title": "Update documentation",
    "description": "Revise the API documentation.",
    "duedate": "2024-08-01T00:00:00Z",
    "status": "completed"
  }
  ```
- **Response:**  
  `200 OK`  
  Returns the updated task.  
  `400 Bad Request` if ID is invalid, required fields are missing, or status is invalid.

### 5. Delete Task

- **DELETE** `/tasks/:id`
- **Response:**  
  `204 No Content`  
  Task is deleted.  
  `400 Bad Request` if ID is invalid.

---

## Status Field

- Allowed values: `"progress"`, `"completed"`, `"pending"`
- Validation is enforced in the controller using the `IsValid()` method.

---

## Date Format

- The `duedate` field must be in ISO 8601 format, e.g. `"2024-07-25T00:00:00Z"`
- This is automatically parsed to Go's `time.Time` type.

---

## Error Handling

- All endpoints return clear error messages and appropriate HTTP status codes for invalid input or missing resources.

---

## Example Task JSON

```json
{
  "title": "Review code",
  "description": "Go through the codebase and ensure all functions are documented.",
  "duedate": "2024-08-01T00:00:00Z",
  "status": "progress"
}
```

---

## Author

- Emran Seid Mohammed  
- [GitHub: Emran-Mohammed](https://github.com/Emran-Mohammed)

---

## License

This project is for educational
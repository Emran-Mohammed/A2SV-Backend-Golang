# Task Manager API Documentation

---

## Overview

The **Task Manager API** is a RESTful service built with Go, Gin, and MongoDB for managing tasks and users with authentication and role-based authorization.
It supports creating, reading, updating, and deleting tasks, as well as user registration and login.  
All data is persisted in a MongoDB database.  
**JWT authentication** is required for most endpoints, and some actions are restricted to admin users.

---

## Folder Structure

```
task_manager/
├── main.go
├── controllers/
│   ├── auth_controller.go
│   └── task_controller.go
├── models/
│   ├── task.go
│   └── user.go
├── data/
│   ├── mongo.go
│   ├── task_service.go
│   └── user_service.go
├── middleware/
│   └── auth_middleware.go
├── router/
│   └── router.go
├── config/
│   └── config.go
├── docs/
│   └── api_documentation.md
└── go.mod
```

---

## How to Run

1. **Start MongoDB:**  
   Make sure MongoDB is running locally on `mongodb://localhost:27017`.

2. **Clone the repository:**
   ```bash
   git clone https://github.com/Emran-Mohammed/A2SV-Backend-Golang.git
   cd A2SV-Backend-Golang/Task-6/task_manager
   ```

3. **Install dependencies:**
   ```bash
   go mod tidy
   ```

4. **Run the server:**
   ```bash
   go run main.go
   ```
   The API will be available at `http://localhost:8080`.

---

## Authentication & Authorization

- **JWT authentication** is required for all `/api` endpoints.
- Obtain a JWT by registering and logging in via `/auth/register` and `/auth/login`.
- **Admin-only endpoints** are under `/api/admin` and require the user to have the `"admin"` role.
- Regular users can only access general `/api/tasks` endpoints.

---

## Models

### User

```go
type User struct {
    ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Username string             `json:"username" bson:"username" binding:"required"`
    Password string             `json:"password" bson:"password" binding:"required"`
    Role     Role               `json:"role,omitempty" bson:"role"`
}
type Role string
const (
    RoleAdmin Role = "admin"
    RoleUser  Role = "user"
)
```
- **Role** defaults to `"user"` if not provided.
- **Username** must be unique.

### Task

```go
type Task struct {
    ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Title       string             `json:"title" binding:"required" bson:"title"`
    Description string             `json:"description" bson:"description"`
    DueDate     time.Time          `json:"duedate" binding:"required" bson:"duedate"`
    Status      Status             `json:"status" binding:"required,oneof=progress completed pending" bson:"status"`
}
```
- **Status** must be one of: `"progress"`, `"completed"`, `"pending"`.

---

## API Endpoints

### Authentication

#### Register
- **POST** `/auth/register`
- **Request Body:**
  ```json
  {
    "username": "yourname",
    "password": "yourpassword",
    "role": "admin" // or "user", optional (defaults to "user")
  }
  ```
- **Responses:**
  - `201 Created` – User registered successfully
  - `400 Bad Request` – Username already taken or invalid input

#### Login
- **POST** `/auth/login`
- **Request Body:**
  ```json
  {
    "username": "yourname",
    "password": "yourpassword"
  }
  ```
- **Responses:**
  - `200 OK` – Returns JWT token
  - `401 Unauthorized` – Invalid username or password

---

### Task Management

#### All `/api` endpoints require JWT in the `Authorization` header:
```
Authorization: Bearer <your_token>
```

#### Get All Tasks
- **GET** `/api/tasks`
- **Response:**  
  `200 OK` – List of all tasks

#### Get Task by ID
- **GET** `/api/tasks/:id`
- **Response:**  
  `200 OK` – Task details  
  `400 Bad Request` – Invalid ID

---

### Admin-only Task Management

**These endpoints require the user to have the `"admin"` role.**

#### Create Task
- **POST** `/api/admin/tasks`
- **Request Body:**
  ```json
  {
    "title": "Task Title",
    "description": "Task details",
    "duedate": "2024-08-01T12:00:00Z",
    "status": "pending"
  }
  ```
- **Response:**  
  `201 Created` – Task created  
  `400 Bad Request` – Invalid input

#### Update Task
- **PUT** `/api/admin/tasks/:id`
- **Request Body:** (same as Create)
- **Response:**  
  `200 OK` – Task updated  
  `400 Bad Request` – Invalid input or ID

#### Delete Task
- **DELETE** `/api/admin/tasks/:id`
- **Response:**  
  `204 No Content` – Task deleted  
  `400 Bad Request` – Invalid ID

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

## Error Handling

- All endpoints return clear error messages and appropriate HTTP status codes for invalid input, authentication, or authorization failures.

---

## Using the API with Postman

1. **Register** a user via `/auth/register`.
2. **Login** via `/auth/login` and copy the returned JWT token.
3. For protected endpoints, set the `Authorization` header:
   ```
   Authorization: Bearer <your_token>
   ```
4. Use `/api/admin/...` endpoints only with an admin token.

---

## Author

- Emran Seid Mohammed  
- [GitHub: Emran-Mohammed](https://github.com/Emran-Mohammed)

---

## License

This project is for Education purpose
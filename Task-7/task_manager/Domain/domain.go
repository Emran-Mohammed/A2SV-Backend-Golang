package domain
import (
	"time"
	"context"
)


type Task struct{
	ID			string
	Title		string 		 
	Description string 		
	DueDate		time.Time  	
	Status 		Status

}

type Status string

const(
	Progress Status = "progress"
	Completed Status = "completed"
	Pending Status = "pending"
)


type User struct{
	ID	string
	Username string 
	Password string 
	Role Role 
	
}

type Role string 
const(
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)




type ITaskRepository interface{
	CreateTask(ctx context.Context, task *Task)(*Task, error)
	GetTasks(ctx context.Context) ([]Task, error)
	GetTaskByID(ctx context.Context, id string)(*Task, error)
	UpdateTask(ctx context.Context, id string, updatedTask *Task)(*Task, error)
	DeleteTask (ctx context.Context, id string) error
}

type IUserRepository interface{
	Register(ctx context.Context, user *User) error
	Login(ctx context.Context, user *User) (*User, error)

}

type IJWTService interface {
	GenerateToken(user *User) (string, error)
}

type IPasswordService interface{
	HashPassword(password string) (string, error)
	VerifyPassword(hashedPassword, password string) bool
}

type ITaskUsecase interface {
	CreateTask(ctx context.Context, task *Task) (*Task, error)
	GetTasks(ctx context.Context) ([]Task, error)
	GetTaskByID(ctx context.Context, id string) (*Task, error)
	UpdateTask(ctx context.Context, id string, task *Task) (*Task, error)
	DeleteTask(ctx context.Context, id string) error
}

type IUserUsercase interface{
	Register(ctx context.Context, user *User) error
	Login(ctx context.Context, user *User) (string, error)
}
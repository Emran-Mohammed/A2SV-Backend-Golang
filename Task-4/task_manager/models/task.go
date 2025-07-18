package models
import "time"


type Task struct{
	ID			int			`json:"id"`
	Title		string 		`json:"title" binding:"required"`
	Description string 		`json:"description"`
	DueDate		time.Time  	`json:"duedate" binding:"required"`
	Status 		Status		`json:"status" binding:"required"`

}

type Status string

const(
	Progress Status = "progress"
	Completed Status = "completed"
	Pending Status = "pending"
)
func (s Status) IsValid() bool {
    return s == Progress || s == Completed || s == Pending
}
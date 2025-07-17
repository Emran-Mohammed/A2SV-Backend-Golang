package models


type Book struct{
	ID int
	Title string
	Author string
	Status Status
}

type Status string

const(
	Available Status = "avialable" 
	Borrowed Status = "borrowed"

)
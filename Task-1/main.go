package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Student struct{
	name string
	subjectsMap map[string]float64
	avgGrade float64
}


func main(){

	
	var choice int;
	
	
	for{
		fmt.Println("Welcome to Student's grade calculater")
		fmt.Println("1. start")
		fmt.Println("2. Exit")
		fmt.Print("Enter 1 or 2: ")
		fmt.Scan(&choice)
		var discard string
		fmt.Scanln(&discard)
		
		switch choice{
			
		case 1:
			student := Student{
			subjectsMap: make(map[string]float64),
			}
			getStudentInfo(&student)
			getgrade(&student)
			showResult(&student)
		case 2:
			fmt.Println("Exiting...")
        	return
		default:
			fmt.Println("invalid choice. please the choice again.")

		}


		
	}

	
	
	
}

func getStudentInfo(s *Student){
	fmt.Print("What is your name: ")
	reader := bufio.NewReader(os.Stdin)
	nameInput, _ := reader.ReadString('\n')
	s.name = strings.TrimSpace(nameInput)

	
	
}



func getgrade(s *Student){

	var numberOfSubjects int;
	for {
		fmt.Print("how many subjects: ")
		fmt.Scanln(&numberOfSubjects)
		if numberOfSubjects > 0 {
			break
		}
		fmt.Println("Number must be greater than 0.")
	}


	var subjectName string
	var subjectGrade float64
	var sum float64 = 0
	for i:=0 ; i < numberOfSubjects; i++ {
		fmt.Printf("what is the name of the %d subject: ", i+1)
		fmt.Scanln(&subjectName)

		fmt.Printf("what is your %s grade: ", subjectName)
		fmt.Scanln(&subjectGrade)
		for {

			if validateGrade(subjectGrade){
				break
			}
			fmt.Printf("invalide grade. Please Enter your correct %s grade again: ", subjectName)
			fmt.Scanln(&subjectGrade)

		}

		sum += subjectGrade
		s.subjectsMap[subjectName] = subjectGrade

	}
	s.avgGrade= sum / float64(numberOfSubjects)
}

func showResult(s *Student){

	fmt.Println("----------------------------------------")
	fmt.Printf("Dear %s here it is the name and grade of subjects you have given\n", s.name)
	fmt.Println("----------------------------------------")
	fmt.Printf("%-20s | % -10s\n", "Subject name", "Grade")
	fmt.Println("----------------------------------------")
	for subject, grade := range s.subjectsMap{

		fmt.Printf("%-20s | %-10.2f\n", subject, grade)
	}
	fmt.Println("----------------------------------------")
	fmt.Printf("The average grade is %.2f\n", s.avgGrade)
	fmt.Println()
}

func validateGrade(grade float64) bool{
	return grade >= 0 && grade <= 100

}


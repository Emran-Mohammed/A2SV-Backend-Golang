package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


var name string;
var numberOfSubjects int;
var subjectsMap = make(map[string]float64)
var resultGrade float64

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
			subjectsMap = make(map[string]float64)
        	resultGrade = 0
			getStudentInfo()
			getgrade()
			showResult()
		case 2:
			fmt.Println("Exiting...")
        	return
		default:
			fmt.Println("invalid choice. please the choice again.")

		}


		
	}

	
	
	
}

func getStudentInfo(){
	fmt.Print("What is your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)


	for {
		fmt.Print("how many subjects: ")
		fmt.Scanln(&numberOfSubjects)
		if numberOfSubjects > 0 {
			break
		}
		fmt.Println("Number must be greater than 0.")
	}
	
	
	

}


func getgrade(){
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
		subjectsMap[subjectName] = subjectGrade

	}
	resultGrade = sum / float64(numberOfSubjects)
}

func showResult(){

	fmt.Printf("Dear %s you given name and grade of %d subjects \n", name, numberOfSubjects)
	fmt.Println("----------------------------------------")
	fmt.Printf("%-20s | % -10s\n", "Subject name", "Grade")
	fmt.Println("----------------------------------------")
	for subject, grade := range subjectsMap{

		fmt.Printf("%-20s | %-10.2f\n", subject, grade)
	}
	fmt.Println("----------------------------------------")
	fmt.Printf("The result grade is %.2f\n", resultGrade)
}

func validateGrade(grade float64) bool{
	return grade >= 0 && grade <= 100

}


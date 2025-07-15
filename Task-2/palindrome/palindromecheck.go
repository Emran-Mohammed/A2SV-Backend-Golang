package main 

import (
	"fmt"
	"unicode"
	"strings"
)

func main(){
		// Test cases
	fmt.Println(palindromeCheck("Madam"))                 // true
	fmt.Println(palindromeCheck("A man, a plan, a canal, Panama")) // true
	fmt.Println(palindromeCheck("Hello"))                 // false
	fmt.Println(palindromeCheck("No lemon, no melon"))    // true

}

func palindromeCheck(sentence string) bool{
	sentence = clearSentence(sentence)
	left := 0
	right := len(sentence) - 1
	
	for left < right{
		if sentence[left] != sentence[right]{
			return false
		}

		left ++
		right --

	}
	return true


}

func clearSentence( s string) string{
	var builder strings.Builder

	for _ , ch := range s{

		if unicode.IsPunct(ch) || ch == ' '{
			continue
		}else{
			builder.WriteRune(unicode.ToLower(ch))
		}

	}

	return builder.String()

}
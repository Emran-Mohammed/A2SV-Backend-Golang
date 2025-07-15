package main
import(
	"fmt"
	"strings"
	"unicode"
)

func main(){
	word := "the hello! there. the here is Hello."

	fmt.Println(wordFreqCount(word))

}

func wordFreqCount(sentence string) map[string]int{
	countMap := make(map[string]int)
	words:= strings.Fields(sentence)
	for _, word := range words{
		word = strings.ToLower(removePunc(word))
		_ , exist := countMap[word]
		if exist {
			countMap[word] += 1

		}else{
			countMap[word] = 1
		}
	}

	return countMap
}

func removePunc(s string) string{
	var builder strings.Builder
	for _, ch := range s{
		if unicode.IsPunct(ch){
			continue
		} else {
			builder.WriteRune(ch)
		}

	}
	return builder.String()
}
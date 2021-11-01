package main

import(
	"fmt"
	"strings"
	"practice/strings/reverse"
)

func main(){
	s := "hello world"
	rev := stringOps.Reverse(s)
	noVowels := stringOps.RemoveVowels(s)
	fmt.Println(rev)
	fmt.Println(noVowels)
}


//rotationalCypher practice problem
func rotationalCypher(input string, rotationalFactor int) string {
	var newS strings.Builder
	for i:= 0; i <len(input); i++{
		curChar := input[i]
		if curChar >= 62 && curChar <= 122 {	
		//letter char
			//65=A....122=z
			//curChar + rotationalFactor must be >=65 and <=122
			n := curChar + byte(rotationalFactor)
			if n <= 122{
				newS.WriteByte(n)
			} else{
				remain := n % 122
				char := remain + 65
				newS.WriteByte(char)
			}
		} else if curChar >= 48 && curChar <= 57 {
			n := curChar + byte(rotationalFactor)
			if n <= 57{
				newS.WriteByte(n)
			} else {
				remain := n % 57
				char := remain + 48
				newS.WriteByte(char)
			}
		} else {
			newS.WriteByte(curChar)
		}
	}
	return newS.String()
}

package stringOps

import(
	"strings"
)


//simple Reverse a String O(n) 
func Reverse(input string) string {
	var b strings.Builder
	for i := len(input)-1; i >= 0; i-- {
		b.WriteByte(input[i])
	}
	return b.String()
}

func RemoveVowels(input string) string {
	var b strings.Builder
	for z := 0; z < len(input); z++{
		curCh := input[z]
		if curCh == 'a' || curCh == 'e' || curCh == 'i' || curCh == 'o' || curCh == 'u' {

		} else {
			b.WriteByte(curCh)
		}
	}
	return b.String()
}

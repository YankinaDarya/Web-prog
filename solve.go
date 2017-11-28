package main

import (
	"unicode"
	"fmt"
)

func RemoveEven(input []int) []int {
	var ans []int;
	for i := 0; i < len(input); i++ {
		if input[i] % 2 == 1{
			ans = append(ans, input[i]);
		}
	}
	return ans;
}
func PowerGenerator(a int) func() int {
	x := 1
	return func() int {
		x *= a
		return x
	}
}

func DifferentWordsCount(x string) int {
	str := ""
	set := make(map[string]bool)
	ans := 0
	for _, c := range (x + " ") {
		if unicode.IsLetter(c) {
			str += string(unicode.ToLower(c))
		} else if str != "" {
			if !set[str] {
				ans += 1
			}
			set[str] = true
			str = ""
		}
	}
	return ans
}

/*func main() {
	input := []int{0, 3, 2, 5}

	var output = RemoveEven(input);

	for _, elem := range output {
		if elem % 2 == 1 {
			fmt.Print(elem);
		}
	}

	fmt.Println("Hello, playground")
}*/


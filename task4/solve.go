package main

import (
	"unicode"
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

func PowerGenerator(gen int) func() int{
	i := gen;
	return func() (ret int){
		ret = i;
		i *= gen;
		return;
	}
}

func DifferentWordsCount(str string) int {
	word := "";
	allWords := make(map[string]int);
	for i := 0; i < len(str); i++{
		c := str[i];
		if unicode.IsLetter(rune(c)){
			word = word + string(unicode.ToLower(rune(c)));
		} else {
			if len(word) != 0{
				allWords[word]++;
			}
			word = "";
		}
	}
	if len(word) != 0{
		allWords[word]++
	}
	return len(allWords);
}

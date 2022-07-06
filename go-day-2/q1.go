package main

import (
	"fmt"
)

func findFrequencyofCharacter(word []string, ch chan map[string]int) {
	frequency := map[string]int{}
	for _, wordValue := range word {
		for i := 0; i < len(wordValue); i++ {
			frequency[string(wordValue[i])]++
		}
	}
	ch <- frequency
}
func main() {
	word := []string{"quick", "brown", "fox", "lazy", "dog"}

	word1 := word[0:2]
	word2 := word[2:4]
	word3 := word[4:]

	channel := make(chan map[string]int, 3)
	go findFrequencyofCharacter(word1, channel)
	go findFrequencyofCharacter(word2, channel)
	go findFrequencyofCharacter(word3, channel)

	finalFrequencyWord := map[string]int{}

	for i := 0; i < 3; i++ {
		tempFreqWord := <-channel
		for key, value := range tempFreqWord {
			finalFrequencyWord[key] += value
		}
	}

	fmt.Println(finalFrequencyWord)
}

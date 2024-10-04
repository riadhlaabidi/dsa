package main

import "fmt"

func dividePlayers(skill []int) int64 {
	freq := make(map[int]int)

	totalSkill := 0

	for _, s := range skill {
		totalSkill += s
		freq[s]++
	}

	targetSkill := totalSkill / (len(skill) / 2)

	if totalSkill%targetSkill != 0 {
		return -1
	}

	totalChemistry := int64(0)

	for currSkill, currCount := range freq {
		pairSkill := targetSkill - currSkill
		pairCount, exist := freq[pairSkill]

		if !exist || pairCount != currCount {
			return -1
		}

		totalChemistry += int64(currSkill * pairSkill * currCount)
	}

	return totalChemistry / 2
}

func main() {
	players := []int{3, 2, 5, 1, 3, 4}

	expected := int64(22)
	actual := dividePlayers(players)

	if actual != expected {
		fmt.Printf("Expected %d, but got %d", expected, actual)
	}

	fmt.Printf("Correct %v", actual)
}

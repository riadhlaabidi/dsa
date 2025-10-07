package avoidflood

import "sort"

func avoidFlood(rains []int) []int {
	noRainIndices := []int{}
	indexByLake := make(map[int]int)
	res := make([]int, len(rains))

	for i, rain := range rains {
		if rain == 0 {
			noRainIndices = append(noRainIndices, i)
			continue
		}

		if prevSameLakeIndex, exists := indexByLake[rain]; exists {
			n := len(noRainIndices)

			if len(noRainIndices) == 0 || noRainIndices[n-1] < prevSameLakeIndex {
				return []int{}
			}

			j := sort.SearchInts(noRainIndices, prevSameLakeIndex)
			res[noRainIndices[j]] = rain
			noRainIndices = append(noRainIndices[:j], noRainIndices[j+1:]...)
		}

		res[i] = -1
		indexByLake[rain] = i
	}

	for _, i := range noRainIndices {
		res[i] = 1
	}

	return res
}

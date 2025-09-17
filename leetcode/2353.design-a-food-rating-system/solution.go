package solution

import "container/heap"

type FoodRating struct {
	name    string
	cuisine string
	rating  int
	index   int
}

type HighestRatingHeap []*FoodRating

func (h HighestRatingHeap) Len() int {
	return len(h)
}

func (h HighestRatingHeap) Less(i, j int) bool {
	if h[i].rating == h[j].rating {
		return h[i].name < h[j].name
	}
	return h[i].rating > h[j].rating
}

func (h HighestRatingHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *HighestRatingHeap) Push(x any) {
	*h = append(*h, x.(*FoodRating))
}

func (h *HighestRatingHeap) Pop() any {
	n := len(*h)
	old := *h
	popped := old[n-1]
	*h = old[:n-1]
	return popped
}

type FoodRatings struct {
	foodRatingByFood       map[string]*FoodRating
	highestRatingByCuisine map[string]*HighestRatingHeap
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	highestRatingByCuisine := make(map[string]*HighestRatingHeap)
	foodRatingByFood := make(map[string]*FoodRating)

	for i, c := range cuisines {
		if _, exists := highestRatingByCuisine[c]; !exists {
			hrh := make(HighestRatingHeap, 0, len(cuisines))
			highestRatingByCuisine[c] = &hrh
		}

		index := highestRatingByCuisine[c].Len()
		foodRating := &FoodRating{foods[i], c, ratings[i], index}
		foodRatingByFood[foods[i]] = foodRating
		heap.Push(highestRatingByCuisine[c], foodRating)
	}

	return FoodRatings{foodRatingByFood, highestRatingByCuisine}
}

func (frs *FoodRatings) changeRating(food string, rating int) {
	foodRating := frs.foodRatingByFood[food]
	foodRating.rating = rating
	heap.Fix(frs.highestRatingByCuisine[foodRating.cuisine], foodRating.index)
}

func (frs *FoodRatings) highestRated(cuisine string) string {
	highestRated := (*(frs.highestRatingByCuisine[cuisine]))[0]
	return highestRated.name
}

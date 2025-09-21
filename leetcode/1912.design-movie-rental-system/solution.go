package movierental

import "container/heap"

type movieEntry struct {
	shop  int
	movie int
	price int
	index int
}

func newMovie(shop, movie, price, index int) *movieEntry {
	return &movieEntry{shop, movie, price, index}
}

// generic swap for movieEntry queues
func swap[T ~[]*movieEntry](c T, i, j int) {
	c[i], c[j] = c[j], c[i]
	c[i].index = i
	c[j].index = j
}

// generic push for movieEntry queues
func push[T ~[]*movieEntry](c *T, x any) {
	*c = append(*c, x.(*movieEntry))
}

// generic pop for movieEntry queues
func pop[T ~[]*movieEntry](c *T) any {
	old := *c
	n := len(old)
	popped := old[n-1]
	*c = old[0 : n-1]
	old[n-1] = nil
	return popped
}

// Unrented movies priority queue, sorted by
// - smaller price
// - smallest shop in case of a tie
type unrented []*movieEntry

func (c unrented) Len() int { return len(c) }
func (c unrented) Less(i, j int) bool {
	if c[i].price == c[j].price {
		return c[i].shop < c[j].shop
	}
	return c[i].price < c[j].price
}
func (c unrented) Swap(i, j int) { swap(c, i, j) }
func (c *unrented) Push(x any)   { push(c, x) }
func (c *unrented) Pop() any     { return pop(c) }

// Rented movies priority queue, sorted by
// - smaller price
// - smaller shop in case of a tie
// - smaller movie in case of double tie
type rented []*movieEntry

func (c rented) Len() int { return len(c) }
func (c rented) Less(i, j int) bool {
	if c[i].price == c[j].price {
		if c[i].shop == c[j].shop {
			return c[i].movie < c[j].movie
		}
		return c[i].shop < c[j].shop
	}
	return c[i].price < c[j].price
}
func (c rented) Swap(i, j int) { swap(c, i, j) }
func (c *rented) Push(x any)   { push(c, x) }
func (c *rented) Pop() any     { return pop(c) }

type MovieRentingSystem struct {
	unrentedByMovie     map[int]*unrented
	movieByMovieAndShop map[int]map[int]*movieEntry
	rentedMovies        *rented
}

func Constructor(n int, entries [][]int) MovieRentingSystem {
	unrentedByMovie := make(map[int]*unrented)
	movieByMovieAndShop := make(map[int]map[int]*movieEntry)
	rentedMovies := make(rented, 0)

	for _, entry := range entries {
		shop := entry[0]
		movie := entry[1]
		price := entry[2]

		if _, exists := unrentedByMovie[movie]; !exists {
			unrentedQueue := make(unrented, 0)
			unrentedByMovie[movie] = &unrentedQueue
		}
		newMovie := newMovie(shop, movie, price, len(*unrentedByMovie[movie]))
		heap.Push(unrentedByMovie[movie], newMovie)

		if _, exists := movieByMovieAndShop[movie]; !exists {
			movieByMovieAndShop[movie] = make(map[int]*movieEntry)
		}
		movieByMovieAndShop[movie][shop] = newMovie
	}

	return MovieRentingSystem{
		unrentedByMovie,
		movieByMovieAndShop,
		&rentedMovies,
	}
}

func (this *MovieRentingSystem) Search(movie int) []int {
	shops := []int{}

	if c, exists := this.unrentedByMovie[movie]; exists {
		size := min(5, c.Len())
		popped := make([]*movieEntry, size)

		for i := range size {
			popped[i] = heap.Pop(c).(*movieEntry)
			shops = append(shops, popped[i].shop)
		}

		for i := range size {
			popped[i].index = c.Len()
			heap.Push(c, popped[i])
		}
	}

	return shops
}

func (this *MovieRentingSystem) Rent(shop int, movie int) {
	rentedMovie := this.movieByMovieAndShop[movie][shop]
	heap.Remove(this.unrentedByMovie[movie], rentedMovie.index)

	rentedMovie.index = this.rentedMovies.Len()
	heap.Push(this.rentedMovies, rentedMovie)
}

func (this *MovieRentingSystem) Drop(shop int, movie int) {
	droppedMovie := this.movieByMovieAndShop[movie][shop]
	heap.Remove(this.rentedMovies, droppedMovie.index)

	droppedMovie.index = this.unrentedByMovie[movie].Len()
	heap.Push(this.unrentedByMovie[movie], droppedMovie)
}

func (this *MovieRentingSystem) Report() [][]int {
	movies := [][]int{}
	size := min(5, this.rentedMovies.Len())
	popped := make([]*movieEntry, size)

	for i := range size {
		popped[i] = heap.Pop(this.rentedMovies).(*movieEntry)
		movies = append(movies, []int{popped[i].shop, popped[i].movie})
	}

	for i := range size {
		popped[i].index = this.rentedMovies.Len()
		heap.Push(this.rentedMovies, popped[i])
	}

	return movies
}

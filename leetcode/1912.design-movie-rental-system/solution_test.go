package movierental

import (
	"dsa/internal/testutil"
	"testing"
)

func TestMovieRentalCreate(t *testing.T) {
	entries := [][]int{
		{0, 1, 5},
		{0, 2, 6},
		{0, 3, 7},
		{1, 1, 4},
		{1, 2, 7},
		{2, 1, 5},
	}
	system := Constructor(3, entries)
	testutil.AssertNotNil(t, system.rentedMovies)
	testutil.AssertEquals(t, 0, system.rentedMovies.Len())

	testutil.AssertNotNil(t, system.movieByMovieAndShop)
	testutil.AssertEquals(t, 3, len(system.movieByMovieAndShop))
	testutil.AssertEquals(t, 3, len(system.movieByMovieAndShop[1]))
	testutil.AssertEquals(t, 2, len(system.movieByMovieAndShop[2]))
	testutil.AssertEquals(t, 1, len(system.movieByMovieAndShop[3]))

	testutil.AssertNotNil(t, system.unrentedByMovie)
	testutil.AssertEquals(t, 3, len(system.unrentedByMovie))
	testutil.AssertNotNil(t, system.unrentedByMovie[1])
	testutil.AssertEquals(t, 3, system.unrentedByMovie[1].Len())
	testutil.AssertEquals(t, 4, (*system.unrentedByMovie[1])[0].price)
	testutil.AssertNotNil(t, system.unrentedByMovie[2])
	testutil.AssertEquals(t, 2, system.unrentedByMovie[2].Len())
	testutil.AssertEquals(t, 6, (*system.unrentedByMovie[2])[0].price)
	testutil.AssertNotNil(t, system.unrentedByMovie[3])
	testutil.AssertEquals(t, 1, system.unrentedByMovie[3].Len())
	testutil.AssertEquals(t, 7, (*system.unrentedByMovie[3])[0].price)
}

func TestMovieRental(t *testing.T) {
	entries := [][]int{
		{0, 1, 5},
		{0, 2, 6},
		{0, 3, 7},
		{1, 1, 4},
		{1, 2, 7},
		{2, 1, 5},
	}

	system := Constructor(3, entries)
	testutil.AssertSliceEquals(t, []int{1, 0, 2}, system.Search(1))
	testutil.AssertEquals(t, 0, len(system.Search(4)))

	system.Rent(0, 1)
	system.Rent(1, 2)

	got := system.Report()
	want := [][]int{{0, 1}, {1, 2}}

	testutil.AssertEquals(t, len(got), len(want))
	for i := range got {
		testutil.AssertSliceEquals(t, want[i], got[i])
	}

	system.Drop(1, 2)
	testutil.AssertSliceEquals(t, []int{0, 1}, system.Search(2))
}

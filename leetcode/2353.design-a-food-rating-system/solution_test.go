package solution

import "testing"

func assertEquals(t *testing.T, want, got string) {
	if got != want {
		t.Errorf("want = %s, got = %s", want, got)
	}
}

func TestFoodRatings(t *testing.T) {
	foods := []string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"}
	cuisines := []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"}
	ratings := []int{9, 12, 8, 15, 14, 7}

	foodRatings := Constructor(foods, cuisines, ratings)

	assertEquals(t, "kimchi", foodRatings.highestRated("korean"))
	assertEquals(t, "ramen", foodRatings.highestRated("japanese"))

	foodRatings.changeRating("sushi", 16)

	assertEquals(t, "sushi", foodRatings.highestRated("japanese"))

	foodRatings.changeRating("ramen", 16)

	assertEquals(t, "ramen", foodRatings.highestRated("japanese"))
}

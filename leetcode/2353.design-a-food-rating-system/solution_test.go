package foodrating

import "dsa/internal/testutil"
import "testing"

func TestFoodRatings(t *testing.T) {
	foods := []string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"}
	cuisines := []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"}
	ratings := []int{9, 12, 8, 15, 14, 7}

	foodRatings := Constructor(foods, cuisines, ratings)

	testutil.AssertEquals(t, "kimchi", foodRatings.highestRated("korean"))
	testutil.AssertEquals(t, "ramen", foodRatings.highestRated("japanese"))

	foodRatings.changeRating("sushi", 16)

	testutil.AssertEquals(t, "sushi", foodRatings.highestRated("japanese"))

	foodRatings.changeRating("ramen", 16)

	testutil.AssertEquals(t, "ramen", foodRatings.highestRated("japanese"))
}

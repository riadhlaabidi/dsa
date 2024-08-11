
package me.riadh.algorithms.search;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class LinearSearchTest {

    @Test
    void testNeedleInArray() {
        int idx = LinearSearch.search(new int[] { 60, 50, 2, 9, 6, 36, 44, 2 }, 2);
        assertEquals(2, idx);
    }

    @Test
    void testNeedleNotInArray() {
        int idx = LinearSearch.search(new int[] { 60, 50, 2, 9, 6, 36, 44, 2 }, 23);
        assertEquals(-1, idx);
    }

}

package me.riadh.algorithms.sorting;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

import org.junit.jupiter.api.Test;

class InsertionSortTest {

    @Test
    void testOneValueArraySort() {
        int[] oneValue = new int[] { 0 };
        int[] sortedOneValue = new int[] { 0 };
        InsertionSort.sort(oneValue);
        assertArrayEquals(sortedOneValue, oneValue);
        InsertionSort.reverseOrderSort(oneValue);
        assertArrayEquals(sortedOneValue, oneValue);
    }

    @Test
    void testPositiveValueArraySort() {
        int[] positiveValues = new int[] { 19, 15, 6, 7, 8, 8, 0, 50, 60, 1, 100, 99 };
        int[] sortedPositiveValues = new int[] { 0, 1, 6, 7, 8, 8, 15, 19, 50, 60, 99, 100 };
        int[] reverseOrderSortedPositiveValues = new int[] { 100, 99, 60, 50, 19, 15, 8, 8, 7, 6, 1, 0 };
        InsertionSort.sort(positiveValues);
        assertArrayEquals(sortedPositiveValues, positiveValues);
        InsertionSort.reverseOrderSort(positiveValues);
        assertArrayEquals(reverseOrderSortedPositiveValues, positiveValues);
    }

    @Test
    void testTwoValuesArraySort() {
        int[] twoValues = new int[] { 3, 2 };
        int[] sortedTwoValues = new int[] { 2, 3 };
        int[] reverseOrderSortedTwoValues = new int[] { 3, 2 };
        InsertionSort.sort(twoValues);
        assertArrayEquals(sortedTwoValues, twoValues);
        InsertionSort.reverseOrderSort(twoValues);
        assertArrayEquals(reverseOrderSortedTwoValues, twoValues);
    }

    @Test
    void testNegativeValuesArraySort() {
        int[] negativeValues = new int[] { -2, -2, -3, -4, -5, -6, -7 };
        int[] sortedNegativeValues = new int[] { -7, -6, -5, -4, -3, -2, -2 };
        int[] reverseOrderSortedNegativeValues = new int[] { -2, -2, -3, -4, -5, -6, -7 };
        InsertionSort.sort(negativeValues);
        assertArrayEquals(sortedNegativeValues, negativeValues);
        InsertionSort.reverseOrderSort(negativeValues);
        assertArrayEquals(reverseOrderSortedNegativeValues, negativeValues);
    }

    @Test
    void testMixedNegativeAndPositiveValuesArraySort() {
        int[] mixedNegativeAndPositiveValues = new int[] { 49, 2, -6, 8, 1, 5, 20, -90, 300, 2 };
        int[] sortedMixedNegativeAndPositiveValues = new int[] { -90, -6, 1, 2, 2, 5, 8, 20, 49, 300 };
        int[] reverseOrderSortedMixedNegativeAndPositiveValues = new int[] { 300, 49, 20, 8, 5, 2, 2, 1, -6, -90 };
        InsertionSort.sort(mixedNegativeAndPositiveValues);
        assertArrayEquals(sortedMixedNegativeAndPositiveValues, mixedNegativeAndPositiveValues);
        InsertionSort.reverseOrderSort(mixedNegativeAndPositiveValues);
        assertArrayEquals(reverseOrderSortedMixedNegativeAndPositiveValues, mixedNegativeAndPositiveValues);
    }

    @Test
    void testEmptyArray() {
        int[] emptyArray = new int[] {};
        InsertionSort.sort(emptyArray);
        assertArrayEquals(new int[] {}, emptyArray);
        InsertionSort.reverseOrderSort(emptyArray);
        assertArrayEquals(new int[] {}, emptyArray);
    }
}

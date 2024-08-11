package me.riadh.algorithms.sorting;

/**
 * Insertion sort algorithms
 */
public final class InsertionSort {

    private InsertionSort() {
    }

    /**
     * Sorts an array of integers in monotonically increasing order.
     * 
     * @param array Array of integers
     */
    public static void sort(int[] array) {
        for (int i = 1; i < array.length; i++) {
            int key = array[i];
            int j = i - 1;
            while (j >= 0 && key < array[j]) {
                array[j + 1] = array[j];
                j--;
            }
            array[j + 1] = key;
        }
    }

    /**
     * Sorts an array of integers in monotonically decreasing order.
     * 
     * @param array Array of integers
     */
    public static void reverseOrderSort(int[] array) {
        for (int i = 1; i < array.length; i++) {
            int key = array[i];
            int j = i - 1;
            while (j >= 0 && key > array[j]) {
                array[j + 1] = array[j];
                j--;
            }
            array[j + 1] = key;
        }

    }
}

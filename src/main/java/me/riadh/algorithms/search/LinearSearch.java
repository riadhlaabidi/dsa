package me.riadh.algorithms.search;

public final class LinearSearch {

    private LinearSearch() {
    }

    /**
     * Searches for a needle inside an array of values.
     * 
     * @param input  input array
     * @param needle the value to search for
     * @return index of the first occurence of the needle if found, -1 otherwise.
     */
    public static int search(int[] input, int needle) {
        for (int i = 0; i < input.length; i++) {
            if (input[i] == needle) {
                // Needle found, return it's index in the input array
                return i;
            }
        }
        // Needle not found
        return -1;
    }
}

package me.riadh.exercices;

public final class Binary {

    private Binary() {
    }

    /**
     * Adds n-bit binary integers a and b
     * 
     * @param a n-bit array representation of number a
     * @param b n-bit array representation of number b
     * @param n number of bits
     * @return array of (n+1)-bits representing the sum of a and b
     */
    public static int[] add(int[] a, int[] b, int n) {
        if (n <= 0) {
            throw new RuntimeException("Invalid binary number, n should be greater than 0");
        }
        if (a.length < n || b.length < n) {
            throw new RuntimeException("Binary number (a and b) input length should be equal to n = " + n);
        }
        int[] result = new int[n + 1];
        int carried = 0;

        for (int i = 0; i < n; i++) {
            result[i] = (a[i] + b[i] + carried) % 2;
            carried = (a[i] + b[i]) / 2;
        }

        result[n + 1] = carried;
        return result;
    }
}

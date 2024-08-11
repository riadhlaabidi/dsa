package me.riadh.algorithms.recursion;

public class RecursiveFunctions {

    /**
     * Adds numbers from 0 to N recursively
     * 
     * @param n number
     * @return sum of numbers from 0 to N
     */
    public static int sum(int n) {
        if (n == 0)
            return 0;
        return sum(n - 1) + n;
    }

    /**
     * Multiplies two integers a & b
     * 
     * @param a first integer
     * @param b second integer
     * @return Long result of the multiplication
     */
    public static long multiply(int a, int b) {
        if (a == 0)
            return 0;
        if (a < 0)
            return multiply(a + 1, b) - b;
        return multiply(a - 1, b) + b;
    }
}

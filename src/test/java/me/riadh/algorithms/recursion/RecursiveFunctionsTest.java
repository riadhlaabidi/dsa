package me.riadh.algorithms.recursion;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class RecursiveFunctionsTest {

    @Test
    void testSum() {
        int sum = RecursiveFunctions.sum(5);
        assertEquals(15, sum);
    }

    @Test
    void testMultiplyPositiveValues() {
        int a = 695;
        int b = 323;
        long result = RecursiveFunctions.multiply(a, b);
        assertEquals(a * b, result);
    }

    @Test
    void testMultiplyNegativeValues() {
        int a = -45;
        int b = -32;
        long result = RecursiveFunctions.multiply(a, b);
        assertEquals(a * b, result);
    }

    @Test
    void testMultiplyPositiveAndNegativeValues() {
        int a = 5555;
        int b = -3;
        long result = RecursiveFunctions.multiply(a, b);
        assertEquals(a * b, result);
    }

    @Test
    void testMultiplyNegativeAndPositiveValues() {
        int a = -621;
        int b = 3266;
        long result = RecursiveFunctions.multiply(a, b);
        assertEquals(a * b, result);
    }

    @Test
    void testMultiplyZeroAndNegativeValues() {
        int a = 0;
        int b = -5266;
        long result = RecursiveFunctions.multiply(a, b);
        assertEquals(a * b, result);
    }

    @Test
    void testMultiplyNegativeAndZeroValues() {
        int a = -8690;
        int b = 0;
        long result = RecursiveFunctions.multiply(a, b);
        assertEquals(a * b, result);
    }
}

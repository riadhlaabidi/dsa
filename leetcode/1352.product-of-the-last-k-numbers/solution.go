package main

import (
	"fmt"
)

type ProductOfNumbers struct {
	products []int
	size     int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		products: []int{1},
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.products = []int{1}
		this.size = 0
	} else {
		prev := this.products[this.size]
		this.products = append(this.products, prev*num)
		this.size++
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if k > this.size {
		return 0
	}

	return this.products[this.size] / this.products[this.size-k]
}

func expectEquals(actual, expected int) {
	if actual != expected {
		panic(fmt.Sprintf("Expected %d, but got %d\n", expected, actual))
	}
	fmt.Printf("Correct %d\n", actual)
}

func main() {
	pn := Constructor()
	pn.Add(3)
	pn.Add(0)
	pn.Add(2)
	pn.Add(5)
	pn.Add(4)
	expectEquals(pn.GetProduct(2), 20)
	expectEquals(pn.GetProduct(3), 40)
	expectEquals(pn.GetProduct(4), 0)
	pn.Add(8)
	expectEquals(pn.GetProduct(2), 32)
}

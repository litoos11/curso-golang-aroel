package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func PrintList(list ...any) {
	for _, item := range list {
		fmt.Println(item)
	}
}

type integer int
type Numbers interface {
	~int | ~float64 | ~float32 | ~uint
}

func Sum[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

func Includes[T comparable](list []T, item T) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}

func Filter[T constraints.Ordered](list []T, callback func(T) bool) []T {
	result := make([]T, 0, len(list))

	for _, item := range list {
		if callback(item) {
			result = append(result, item)
		}
	}
	return result
}

type Product[T uint | string] struct {
	Id    T
	Name  string
	Price float64
}

func main() {
	// PrintList(42, "hello", 3.14, true, []int{1, 2, 3})
	// var num1 integer = 100
	// var num2 integer = 200
	// fmt.Println("Sum:", Sum(num1, num2, 50))
	// fmt.Println("Sum:", Sum[float32](1.36, 2, 3, 4, 5))

	// stringsList := []string{"apple", "banana", "cherry"}
	// numbersList := []int{10, 20, 30, 40, 50}

	// fmt.Println("Includes 'banana':", Includes(stringsList, "banana"))
	// fmt.Println("Includes 25:", Includes(numbersList, 25))

	// fmt.Println("Filtered strings (length > 5):", Filter(stringsList, func(s string) bool { return len(s) > 5 }))
	// fmt.Println("Filtered numbers (> 25):", Filter(numbersList, func(n int) bool { return n > 25 }))

	product1 := Product[uint]{Id: 1, Name: "Laptop", Price: 999.99}
	product2 := Product[string]{Id: "A100", Name: "Smartphone", Price: 499.99}

	fmt.Println("Product Details:", product1)
	fmt.Println("Product Details:", product2)
}

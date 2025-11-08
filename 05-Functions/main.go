package main

import (
	"fmt"
	"strconv"
)

type transformFunction func(int) int

func double(a int) int {
	return a * 2
}
func triple(a int) int {
	return a * 3
}

// taking a function as a parameter
func transformArray(numbers *[]int, transform func(int) int) []int {
	dnumbers := []int{}
	for _, val := range *numbers {
		dnumbers = append(dnumbers, transform(val))
	}
	return dnumbers
}

// taking a function as a parameter with shorthand enabeled
func trnasformArray2(numbers *[]int, transform transformFunction) []int {
	dnumbers := []int{}
	for _, val := range *numbers {
		dnumbers = append(dnumbers, transform(val))
	}
	return dnumbers
}

// returning a function based on value
func returnAFunctionIllustration() func(int) int {
	return double
}

// func transfromNumber (numbers *[]int,)
func main() {
	fmt.Print("Hello there")
	var a any = 12.44
	typed_val, ok := a.(float64)
	fmt.Println(typed_val)
	fmt.Println(strconv.FormatBool(ok))
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	result := transformArray(&numbers, triple)
	trnasformArray2(&numbers, double)
	for idx, val := range result {
		prs := fmt.Sprintf("elt %d,->%d", idx, val)
		fmt.Println(prs)
	}
	returnAFunctionIllustration()
	//testing out lambda function/anonymous function
	transformed := transformArray(&numbers, func(number int) int {
		return number * 6
	})
	for idx, val := range transformed {
		fmt.Printf("%d,%d\n", idx, val)
	}
	//closures
	factor_five_func := createTransformer(5)
	fmt.Println(factor_five_func(60))
	//create and test variadic functions:
	sum(4, 10)
	sum(2, 3, 4, 5, 6)
	//reverse variadic: passign slices as kwargs
	sum(4, numbers...)

}

// custom function creators that can be used to construct other functions
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

// variadic functions
func sum(startingValue int, numbers ...int) {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	fmt.Printf("%d, %d\n", startingValue, sum)

}

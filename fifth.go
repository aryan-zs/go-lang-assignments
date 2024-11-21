package main

import (
	"fmt"
)

// Customer struct definition
type Customer struct {
	Name   string
	Age    int
	Salary int
}

// Function to compare two customers based on the given attribute and direction
func compare(c1, c2 Customer, attribute string, direction string) bool {

	switch attribute {
	case "age":
		if direction == "asc" {
			return c1.Age < c2.Age
		}
		return c1.Age > c2.Age
	case "name":
		if direction == "asc" {
			return c1.Name < c2.Name
		}
		return c1.Name > c2.Name
	case "salary":
		if direction == "asc" {
			return c1.Salary < c2.Salary
		}
		return c1.Salary > c2.Salary
	}
	return false
}

// SelectionSort function to sort the slice of customers
func SelectionSort(customers []Customer, attribute string, direction string) {
	n := len(customers)
	for i := 0; i < n-1; i++ {

		minIndex := i

		for j := i + 1; j < n; j++ {
			if compare(customers[j], customers[minIndex], attribute, direction) {
				minIndex = j
			}
		}
		customers[i], customers[minIndex] = customers[minIndex], customers[i]
	}
}

// SortTester function to sort the customers and return their names in sorted order
func SortTester(customers []Customer, attribute string, direction string) (names []string) {

	SelectionSort(customers, attribute, direction)
	names = make([]string, len(customers))
	for i, customer := range customers {
		names[i] = customer.Name
	}
	return names
}

func fifth() {

	customers := []Customer{
		{"Aryan", 45, 10000},
		{"Aman", 99, 5000},
		{"Abhishek", 32, 5000},
	}

	names := SortTester(customers, "age", "asc")
	fmt.Println("the sorted customer  on their age in ascending order is ", names)

	names = SortTester(customers, "name", "desc")
	fmt.Println("the sorted customer on their name in descending order is ", names)

	names = SortTester(customers, "salary", "asc")
	fmt.Println("the sorted customer on their salary  in ascending order is ", names)
}

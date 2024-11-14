package main

import "fmt"

func double_val(n int) int {
	return n * 2
}

func greet(name string) {
	fmt.Println("hello  " + name)
}

func cir_perimeter(n float64) float64 {
	return 2 * n * pi
}

func sq_perimeter(n float64) float64 {
	return 4 * n
}

const pi = 3.1415926

func main() {
	var a int
	fmt.Println("Enter the number")
	fmt.Scanf("%d", &a)
	fmt.Println(double_val(a))
	var name string
	fmt.Println("Enter the name")
	fmt.Scanf("%s", &name)

	greet(name)

	var side float64
	fmt.Println("Enter the side:")
	fmt.Scanf("%g", &side)

	fmt.Printf("the perimeter of the square of side %v  is %v \n ", side, sq_perimeter(side))

	var radius float64
	fmt.Println("enter the radus:")
	fmt.Scanf("%g", &radius)
	fmt.Printf("the perimeter of the circle of radius %v  is %v \n", radius, cir_perimeter(radius))

}

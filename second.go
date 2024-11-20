package main

import (
	"errors"
	"fmt"
)

func what_to_wear() {

	var season string
	fmt.Println("Enter the season :")
	fmt.Scanf("%s", &season)

	if season == "summer" {
		var is_humid bool
		fmt.Println("summer season  can you plese tell me is it humid or not")
		fmt.Scanf("%t", &is_humid)
		if is_humid {
			fmt.Println("you can wear some cotton shirt and shorts")
		} else {
			fmt.Println("you can wear some synthetic shirt  with shorts")
		}

	} else if season == "winter" {

		var is_snowing bool
		fmt.Println("winter season  can you plese tell me is it snowing or not")
		fmt.Scanf("%t", &is_snowing)
		if is_snowing {
			fmt.Println("you can wear some waterproof jackets with waterproof pants ")
		} else {
			fmt.Println("you can wear some  jackets with  pants ")
		}

	} else if season == "autum" {
		var is_raining bool
		fmt.Println("autum season  can you plese tell me is it raining or not")
		fmt.Scanf("%t", &is_raining)
		if is_raining {
			fmt.Println("you can wear anyting but dont forget to have raincoat ")

		} else {
			fmt.Println("you can wear full sleves  and pants")
		}
	}
}

func calc(num1 float64, num2 float64, op string) (float64, error) {

	var res float64

	if num2 == 0 {
		return res, errors.New("division by zero")
	} else {
		switch op {
		case "*":
			res = num1 * num2
		case "+":
			res = num1 + num2
		case "-":
			res = num1 - num2
		case "/":
			res = num1 / num2

		}
	}
	return res, nil
}

type vehicle struct {
	name          string
	fuel_capacity int
}

func voting_right() {
	var a int
	fmt.Println("enter you age:")
	fmt.Scanf("%d", &a)

	if a < 18 {
		fmt.Println("Sorrry not eligible to vote")
	} else if a == 18 {
		fmt.Println(" eligible apply for voter id !!")
	} else {
		fmt.Println("yes you can vote ")
	}
}

func vehicle_typ() {

	var v1 vehicle = vehicle{name: "truck", fuel_capacity: 500}

	switch v1.name {
	case "truck":
		fmt.Printf("the current vehicle is truck with tank capacity of  %d litres\n", v1.fuel_capacity)
	case "bike":
		fmt.Printf("the current vehicle is bike with tank capacity of %d litres\n", v1.fuel_capacity)
	case "car":
		fmt.Printf("the current vehicle is car with tank capacity of %d litres\n", v1.fuel_capacity)
	default:
		fmt.Println("the vehicle is not known")
	}

}

func isEven(a int) {

	if a%2 == 0 {
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}
}

func cal() {
	var num1 float64
	var num2 float64
	num1 = 20
	num2 = 10

	op := "*"

	result, error := calc(num1, num2, op)

	if error == nil {
		fmt.Printf("%v %v %v =%v \n", num1, op, num2, result)
	} else {
		fmt.Println(error)
	}

}

func sumNatural(n int) {

	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("the sum of number from 1 to %v is %v \n", n, sum)

}

func swap_pointers(a, b *int) {

	temp := *a
	*a = *b
	*b = temp

}

func swap_number_by_using_pointer() {

	a := 20

	b := 30

	a_p := &a
	b_p := &b

	fmt.Printf("%v is a  and %v is b \n", a, b)

	swap_pointers(a_p, b_p)

	fmt.Printf("%v is a  and %v is b \n", a, b)

}

type bank_acc struct {
	account_holder string
	age            int
	balance        float32
	account_no     string
	IFSC           string
}

func (b *bank_acc) check_bal() {
	fmt.Printf("the current balance  of %v  is %v", b.account_holder, b.balance)
}
func (b *bank_acc) deposite(a float32) {
	fmt.Println("added money :", a)
	b.balance += a
}
func (b *bank_acc) withdraw(a float32) {
	fmt.Println("deducted money :", a)
	b.balance -= a
}

func account() {

	var acc1 bank_acc
	acc1.account_holder = "aryan joshi"
	acc1.age = 22
	acc1.balance = 1000
	acc1.account_no = "1022321"
	acc1.IFSC = "SBIN00S23"

	acc1.check_bal()
	acc1.deposite(500)
	acc1.check_bal()
	acc1.withdraw(1000)
	acc1.check_bal()

}

func second() {
	voting_right()
	what_to_wear()
	vehicle_typ()
	isEven(87)
	cal()
	sumNatural(50)
	swap_number_by_using_pointer()
	account()
}

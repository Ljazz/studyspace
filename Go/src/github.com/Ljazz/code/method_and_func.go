package main

import (
	"fmt"
)

type Employee struct {
	name string
	salary int
	currency string
}

/*
displaySalary() method converted to function with Employee as parameter
*/

func displaySalary(e Employee){
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main(){
	emp1 := Employee{
		name: "Sam Adolf",
		salary: 5000,
		currency: "$",
	}
	displaySalary(emp1)
}
package main

import "oop/employee"

func main() {
	// e := employee.Employee{
	// 	FirstName:   "Sam",
	// 	LastName:    "Adolf",
	// 	TotalLeaves: 30,
	// 	LeavesTaken: 20,
	// }
	var e employee.Employee
	e.LeavesRemaining()
}

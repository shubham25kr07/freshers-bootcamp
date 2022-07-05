package main

import (
	"fmt"
)

type CalculateSalaryofEmployee interface {
	Calulate() int
}

type FullTime struct {
	empId         int
	basicPay      int
	totalDaysWork int
}
type Contractor struct {
	empId         int
	basicPay      int
	totalDaysWork int
}

type FreeLancer struct {
	empId          int
	basicPay       int
	totalHoursWork int
}

func (f FullTime) Calulate() int {
	return f.basicPay * f.totalDaysWork
}
func (c Contractor) Calulate() int {
	return c.basicPay * c.totalDaysWork
}
func (f FreeLancer) Calulate() int {
	return f.basicPay * f.totalHoursWork
}
func measure(c CalculateSalaryofEmployee) {
	fmt.Println(c.Calulate())
}

func main() {

	fullTime := FullTime{empId: 1, basicPay: 500, totalDaysWork: 28}
	contract := Contractor{empId: 2, basicPay: 100, totalDaysWork: 28}
	freeLancer := FreeLancer{empId: 3, basicPay: 10, totalHoursWork: 20}

	fmt.Println("FullTime salary")
	measure(fullTime)

	fmt.Println("Contractor salary")
	measure(contract)

	fmt.Println("FreeLancer salary")
	measure(freeLancer)
}

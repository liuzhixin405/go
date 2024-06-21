package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	totalAmount := 100.0
	numbers := 10
	//amounts := doubeleAverage(totalAmount, numbers)
	amounts := remainingAverage(totalAmount, numbers)
	//amounts := fixedAmount(totalAmount, numbers)
	//amounts := linearDecrement(totalAmount, numbers)
	for i, amount := range amounts {
		fmt.Printf("Person %d: %.2f\n", i+1, amount)
	}
}

// 双重平均
func doubeleAverage(totalAmount float64, numbers int) []float64 {

	amounts := make([]float64, numbers)

	for i := numbers; i > 0; i-- {
		if i == 1 {
			amounts[numbers-1] = totalAmount
		} else {
			avg := totalAmount / float64(i)
			amount := r.Float64() * (avg * 2)
			if amount < 0.01 {
				amount = 0.01
			}
			amounts[numbers-i] = amount
			totalAmount -= amount
		}
	}
	return amounts
}

// 剩余平均
func remainingAverage(totalAmount float64, numPeople int) []float64 {

	amounts := make([]float64, numPeople)
	for i := 0; i < numPeople; i++ {
		avg := totalAmount / float64(numPeople-i)
		amount := r.Float64() * avg * 2
		if amount < 0.01 {
			amount = 0.01
		}
		if amount > 20 {
			log.Println("金额大于20")
		}
		amounts[i] = amount
		totalAmount -= amount
	}
	return amounts
}

// 固定金额
func fixedAmount(totalAmount float64, numPeople int) []float64 {
	amountPerPerson := totalAmount / float64(numPeople)
	amounts := make([]float64, numPeople)
	for i := range amounts {
		amounts[i] = amountPerPerson
	}
	return amounts
}

// 线性递减
func linearDecrement(totalAmount float64, numPeople int) []float64 {
	amounts := make([]float64, numPeople)
	decrement := (totalAmount / float64(numPeople)) / float64(numPeople)
	currentAmount := (totalAmount / float64(numPeople)) + (decrement * float64(numPeople-1) / 2)
	for i := 0; i < numPeople; i++ {
		amounts[i] = currentAmount
		currentAmount -= decrement
	}
	return amounts
}

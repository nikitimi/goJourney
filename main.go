package main

import (
	"fmt"
	"time"
)

type Consumption struct {
	ConsumptionType string
	DateCreated     time.Time
}

type Management interface {
	PrintConsumptionType()
	PrintConsumptionDate()
	// ChangeConsumptionType()
}

func (p Consumption) PrintConsumptionType() {
	fmt.Printf("Your consumption type is %v\n", p.ConsumptionType)
}

func (p Consumption) PrintConsumptionDate() {
	y, d, m := p.DateCreated.Year(), p.DateCreated.Day(), p.DateCreated.Month()
	fmt.Printf("Your consumption is created on %v %v, %v\n", m, d, y)
}

func (p *Consumption) ChangeConsumptionType(newConsumptionType string) {
	fmt.Println("Pointer?")
	p.ConsumptionType = newConsumptionType
}

func createObjectFromStruct() {
	foodConsumption := Consumption{
		ConsumptionType: "food",
		DateCreated:     time.Now(),
	}
	travelConsumption := Consumption{
		ConsumptionType: "travel",
		DateCreated:     time.Now(),
	}
	foodConsumption.ChangeConsumptionType("feed")
	m := []Management{foodConsumption, travelConsumption}

	for _, v := range m {
		v.PrintConsumptionDate()
		v.PrintConsumptionType()
	}
}

func doPanic() {
	defer func() { // Anonymous function.
		if e := recover(); e != nil {
			fmt.Printf("Recovering with %s\n", e)
		}
	}()
	panic("I'm panicking AHHH!")
	fmt.Println("This will be unreachable code.")
}

func flexCollections() []int {
	// Array, assigning 10 in item 5.
	x := [5]int{4: 10}

	// Slice with 4 length and 10 capacity.
	y := make([]int, 4, 10)

	// Slice proper.
	arr := []int{}
	length := 5
	base := 10

	fmt.Println(x, y)

	for x := 0; x < length; x += 1 {
		arr = append(arr, base*(x+1))
	}

	return arr
}

func main() {
	fmt.Println("Initializing panicking hormone.")
	doPanic()
	fmt.Println("Deym! recovered.")

	fmt.Printf("\nI've received slice => %v\n", flexCollections())
	createObjectFromStruct()
}

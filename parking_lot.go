package main

import (
	"fmt"
	"sort"
)

type Car struct {
	Number string
}

type ParkingLot struct {
	Capacity int
	Slots    map[int]*Car
	Free     []int
}

func NewParkingLot(capacity int) *ParkingLot {
	free := make([]int, capacity)
	for i := 0; i < capacity; i++ {
		free[i] = i + 1
	}
	return &ParkingLot{
		Capacity: capacity,
		Slots:    make(map[int]*Car),
		Free:     free,
	}
}

func (p *ParkingLot) Park(carNumber string) {
	if len(p.Free) == 0 {
		fmt.Println("Sorry, parking lot is full")
		return
	}
	sort.Ints(p.Free)
	slot := p.Free[0]
	p.Free = p.Free[1:]
	p.Slots[slot] = &Car{Number: carNumber}
	fmt.Printf("Allocated slot number: %d\n", slot)
}

func (p *ParkingLot) Leave(carNumber string, hours int) {
	var slotToFree int = -1
	for slot, car := range p.Slots {
		if car.Number == carNumber {
			slotToFree = slot
			break
		}
	}
	if slotToFree == -1 {
		fmt.Printf("Registration number %s not found\n", carNumber)
		return
	}

	delete(p.Slots, slotToFree)
	p.Free = append(p.Free, slotToFree)

	charge := 10
	if hours > 2 {
		charge += (hours - 2) * 10
	}
	fmt.Printf("Registration number %s with Slot Number %d is free with Charge $%d\n",
		carNumber, slotToFree, charge)
}

func (p *ParkingLot) Status() {
	fmt.Println("Slot No. Registration No.")
	keys := make([]int, 0, len(p.Slots))
	for k := range p.Slots {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, slot := range keys {
		fmt.Printf("%d %s\n", slot, p.Slots[slot].Number)
	}
}

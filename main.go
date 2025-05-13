package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: parking-app <input_file>")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}
	defer file.Close()

	var parkingLot *ParkingLot

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		tokens := strings.Fields(line)
		switch tokens[0] {
		case "create_parking_lot":
			if len(tokens) != 2 {
				fmt.Println("Invalid command: create_parking_lot {capacity}")
				continue
			}
			capacity, err := strconv.Atoi(tokens[1])
			if err != nil || capacity <= 0 {
				fmt.Println("Invalid capacity")
				continue
			}
			parkingLot = NewParkingLot(capacity)
			fmt.Printf("Created parking lot with %d slots\n", capacity)
		case "park":
			if len(tokens) != 2 {
				fmt.Println("Invalid command: park {car_number}")
				continue
			}
			if parkingLot == nil {
				fmt.Println("Parking lot not created")
				continue
			}
			parkingLot.Park(tokens[1])
		case "leave":
			if len(tokens) != 3 {
				fmt.Println("Invalid command: leave {car_number} {hours}")
				continue
			}
			if parkingLot == nil {
				fmt.Println("Parking lot not created")
				continue
			}
			hours, err := strconv.Atoi(tokens[2])
			if err != nil || hours <= 0 {
				fmt.Println("Invalid hours")
				continue
			}
			parkingLot.Leave(tokens[1], hours)
		case "status", "Status":
			if parkingLot == nil {
				fmt.Println("Parking lot not created")
				continue
			}
			parkingLot.Status()
		default:
			fmt.Printf("Unknown command: %s\n", tokens[0])
		}
	}
}

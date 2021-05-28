package main

import (
	"fmt"
	"lab3/approximation"
	"lab3/reading"
	"log"
)

func main() {
	x, err := reading.FromInput()
	if err != nil {
		log.Fatal(err)
	}

	degree := int(x)
	fmt.Printf("provided degree value: %d\n", degree)

	data, err := reading.FileToMap("./data.csv")
	if err != nil {
		log.Fatal(err)
	}

	approximation.FindApproximationFunc(data[1], data[2], degree)
}

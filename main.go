package main

import (
	"fmt"
)

var radius float32

var x int

var y int

var l int

func main() {
	// var p *int
	// x := 10
	// fmt.Println(x)

	// increment(p)

	// fmt.Println(p)

	// calculatecirculeArea(2)

	// printcirculeArea(2)

	calculateSquareArea(5, 6)

}

// func increment(p int) {
// 	p += 1
// 	return

// }

// func calculatecirculeArea(radius int) (float32, error) {
// 	if radius <= 0 {
// 		return float32(0), errors.New("Радиус должен быть отрицательным")
// 	}
// 	return float32(radius) * float32(radius) * math.Pi, nil

// }

// func printcirculeArea(radius int) {

// 	fmt.Printf("Радиус круга: %d см. \n", radius)
// 	fmt.Println("Формула для расчета площади круга: S = π × r2")

// 	circuleArea, err := calculatecirculeArea(radius)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return

// 	}

// 	fmt.Printf("Площадь круга: %f32 см. кв. \n", circuleArea)

// }

func calculateSquareArea(x int, y int) {

	var l = x * y

	fmt.Println(l)
	fmt.Println()

}

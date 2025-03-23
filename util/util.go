package util

import (
	"fmt"
	"math"
	"math/rand"
)

// Return a random number between min and max
func RandBetween(min, max float64) float64 {
	//fmt.Println("util.RandBetween, min:", min, "max:", max)
	return min + (max-min)*rand.Float64()
}

// Return a boolean with a given probability of being true
func RandBool(prob float64) bool {
	return rand.Float64() < prob
}

// Calculate the distance between two points, each with direction and distance from the origin
func Distance(distance1, angle1, distance2, angle2 float64) float64 {
	// Convert polar coordinates to Cartesian coordinates
	x1 := distance1 * math.Cos(angle1)
	y1 := distance1 * math.Sin(angle1)
	x2 := distance2 * math.Cos(angle2)
	y2 := distance2 * math.Sin(angle2)
	distance := math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
	if distance > 2.0 {
		// Exit the program if the distance is greater than 1.0
		panic(fmt.Sprintf("Distance > 2.0, d1: %f a1: %f d2: %f a2: %f", distance1, angle1, distance2, angle2))

	}
	return distance
}

// Sigmoid activation function for neural networks
func Sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

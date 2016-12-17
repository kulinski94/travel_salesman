package main

import (
	"fmt"
	"math"

	"github.com/bradhe/stopwatch"
)

type costs [][]float64

// RunBruteForce - calculate all paths costa and select the best
func RunBruteForce(originalCities Cities) (Cities, float64, uint32) {
	start := stopwatch.Start()

	var minRoute Cities

	var minCosts float64

	initDistanceByCoordinates(originalCities)

	fmt.Println("Results:", minRoute, minCosts)
	watch := stopwatch.Stop(start)

	fmt.Printf("Milliseconds elapsed: %v\n", watch.Milliseconds())

	return originalCities, minCosts, watch.Milliseconds()
}

func initDistanceByCoordinates(cities Cities) costs {

	costs := initCostsMatrix(cities)

	for i := 0; i < len(cities); i++ {
		fmt.Println("Row")
		for j := 0; j < len(cities); j++ {
			costs[i][j] = calculateTravelCostsBetweenCities(cities, i, j)
			fmt.Print(" ", costs[i][j])
		}
	}
	return costs
}

func initCostsMatrix(cities Cities) costs {
	var costs costs
	for i := 0; i < len(cities); i++ {
		var row []float64
		for j := 0; j < len(cities); j++ {
			row = append(row, 0)
		}
		costs = append(costs, row)
	}
	return costs
}

func calculateTravelCostsBetweenCities(cities Cities, i, j int) float64 {
	dx := cities[i].XCord - cities[j].XCord
	dy := cities[i].YCord - cities[j].YCord
	return math.Sqrt(float64(dx*dx) + float64(dy*dy))
}

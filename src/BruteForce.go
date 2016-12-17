package main

import (
	"fmt"

	"github.com/bradhe/stopwatch"
)

type costs [][]float64

var travelCosts costs

var count int

var minCosts float64

var minRoute []int

// RunBruteForce - calculate all paths costa and select the best
func RunBruteForce(originalCities Cities) (Cities, float64, uint32) {
	start := stopwatch.Start()

	var route []int

	minCosts = float64(-1)

	count = 0

	//first city always zer
	route = append(route, 0)

	travelCosts = initDistanceByCoordinates(originalCities)

	for i := 1; i < len(originalCities); i++ {
		fmt.Println("Route", route)
		if len(route) == 1 {
			route = append(route, i)
		} else {
			route[1] = i
		}
		checkRoute(route, 2, len(originalCities))
	}

	watch := stopwatch.Stop(start)

	var optimalPath Cities
	for i := 0; i < len(minRoute); i++ {
		optimalPath = append(optimalPath, originalCities[minRoute[i]])
	}
	fmt.Println("Min Route", minRoute)
	return optimalPath, minCosts, watch.Milliseconds()
}

func checkRoute(route []int, offset int, count int) {
	if offset == count {
		count++
		cost := calculateCost(route)
		if minCosts < 0 || cost < minCosts {
			minCosts = cost
			minRoute = route
		}

		return
	}

loop:
	for i := 1; i < count; i++ {
		for j := 0; j < offset; j++ {
			if route[j] == i {
				continue loop
			}
		}

		if len(route)-1 < offset {
			route = append(route, i)
		} else {
			route[offset] = i
		}
		checkRoute(route, offset+1, count)
	}
}

func calculateCost(route []int) float64 {
	cost := float64(0)

	for i := 1; i < len(route); i++ {
		cost += travelCosts[route[i-1]][route[i]]
	}
	//return to starting city
	cost += travelCosts[route[len(route)-1]][route[0]]

	return cost
}

func initDistanceByCoordinates(cities Cities) costs {

	costs := initCostsMatrix(cities)

	for i := 0; i < len(cities); i++ {
		for j := 0; j < len(cities); j++ {
			costs[i][j] = CalculateTravelCostsBetweenCities(cities[i], cities[j])
		}
	}
	return costs
}

func initCostsMatrix(cities Cities) costs {
	var newCosts costs
	for i := 0; i < len(cities); i++ {
		var row []float64
		for j := 0; j < len(cities); j++ {
			row = append(row, 0)
		}
		newCosts = append(newCosts, row)
	}
	return newCosts
}

package main

import (
	"fmt"

	"github.com/bradhe/stopwatch"
)

// RunNearestNeighbor - start from one city and go to closest one
func RunNearestNeighbor(originalCities Cities) (Cities, float64, uint32) {
	start := stopwatch.Start()

	leftCities := originalCities

	var bestOrder Cities

	for i := 0; i < len(originalCities); i++ {

		if i == 0 {
			bestOrder = append(bestOrder, leftCities[0])
			leftCities = removeIndex(leftCities, 0)
		} else {
			currentCity := bestOrder[len(bestOrder)-1]

			closestCityIndex := 0
			closestDistance := float64(1000000)

			for j, secondCity := range leftCities {
				dist := CalculateTravelCostsBetweenCities(currentCity, secondCity)
				if dist < closestDistance {
					closestDistance = dist
					closestCityIndex = j
				}
			}
			bestOrder = append(bestOrder, leftCities[closestCityIndex])
			leftCities = removeIndex(leftCities, closestCityIndex)
		}
	}

	distance := calcPath(bestOrder)

	watch := stopwatch.Stop(start)
	return bestOrder, distance, watch.Milliseconds()
}

func calcPath(cities Cities) float64 {
	distance := float64(0)
	for i := 0; i < len(cities); i++ {
		if i == len(cities)-1 {
			distance += CalculateTravelCostsBetweenCities(cities[i], cities[0])
		} else {
			distance += CalculateTravelCostsBetweenCities(cities[i], cities[i+1])
		}
	}
	fmt.Println("distance: ", distance)
	return distance
}

func removeIndex(s Cities, index int) Cities {
	return append(s[:index], s[index+1:]...)
}

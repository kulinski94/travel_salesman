package main

import (
	"fmt"
	"math"

	"github.com/bradhe/stopwatch"
)

// ClosestCityAlgorithm - start from one city and go to closest one
func ClosestCityAlgorithm(originalCities Cities) (Cities, float64, uint32) {
	start := stopwatch.Start()

	leftCities := originalCities

	var bestOrder Cities

	for i := 0; i < len(originalCities); i++ {

		fmt.Println("left", leftCities)
		if i == 0 {
			bestOrder = append(bestOrder, leftCities[0])
			leftCities = removeIndex(leftCities, 0)
		} else {
			currentCity := bestOrder[len(bestOrder)-1]

			closestCityIndex := 0
			closestDistance := float64(1000000)

			for j, secondCity := range leftCities {
				dist := calcDistanceBeetweenDots(currentCity, secondCity)
				fmt.Println("distance closest", dist, closestDistance)
				if dist < closestDistance {
					closestDistance = dist
					closestCityIndex = j
					fmt.Println("distance closest", dist, closestDistance, closestCityIndex)
				}
			}
			bestOrder = append(bestOrder, leftCities[closestCityIndex])
			leftCities = removeIndex(leftCities, closestCityIndex)
		}
	}

	distance := calcPath(bestOrder)

	watch := stopwatch.Stop(start)
	fmt.Printf("Milliseconds elapsed: %v\n", watch.Milliseconds())

	return bestOrder, distance, watch.Milliseconds()
}

func calcPath(cities Cities) float64 {
	distance := float64(0)
	for i := 0; i < len(cities); i++ {
		if i == len(cities)-1 {
			distance += calcDistanceBeetweenDots(cities[i], cities[0])
		} else {
			distance += calcDistanceBeetweenDots(cities[i], cities[i+1])
		}
	}
	fmt.Println("distance: ", distance)
	return distance
}

func calcDistanceBeetweenDots(a, b City) float64 {
	first := math.Pow(float64(b.XCord)-float64(a.XCord), 2)
	second := math.Pow(float64(b.YCord)-float64(a.YCord), 2)
	distance := math.Sqrt(first + second)
	return distance
}

func removeIndex(s Cities, index int) Cities {
	return append(s[:index], s[index+1:]...)
}

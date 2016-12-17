package main

import (
	"fmt"
	"math"
	"math/rand"
)

func ClosestCityAlgorithm(originalCities Cities) (Cities, float64) {
	leftCities := originalCities

	var bestOrder Cities

	currentCityIndex := 0

	for i := 0; i < len(originalCities); i++ {
		if currentCityIndex == 0 {
			bestOrder = append(bestOrder, leftCities[0])
			leftCities = removeIndex(leftCities, 0)
		} else {
			currentCity := bestOrder[len(bestOrder)-1]

			closestCityIndex := 0
			closestDistance := float64(10000)

			for j, secondCity := range leftCities {
				dist := calcDistanceBeetweenDots(currentCity, secondCity)
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
	return bestOrder, distance
}

func removeIndex(s Cities, index int) Cities {
	return append(s[:index], s[index+1:]...)
}

func generateCities(count int) Cities {
	var cities Cities
	weight := 540
	height := 480
	i := 0
	for i < count {
		cities = append(cities, City{XCord: rand.Intn(weight), YCord: rand.Intn(height)})
		i++
	}
	return cities
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

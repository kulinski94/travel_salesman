package main

import "math"

// City - keep coordinates of city
type City struct {
	Name  string `json:"name"`
	XCord int    `json:"xCord"`
	YCord int    `json:"yCord"`
}

// RestResponse - path and distance of the path
type RestResponse struct {
	Paths     Cities  `json:"paths"`
	Distance  float64 `json:"distance"`
	TimeTaken uint32  `json:"timeTaken"`
}

//RestRequest - cities and algorithm to find best path
type RestRequest struct {
	Cities    Cities `json:"cities"`
	Algorithm string `json:"algorithm"`
}

//Cities - array of City object
type Cities []City

//CalculateTravelCostsBetweenCities - distance between two cities
func CalculateTravelCostsBetweenCities(a, b City) float64 {
	dx := a.XCord - b.XCord
	dy := a.YCord - b.YCord
	return math.Sqrt(float64(dx*dx) + float64(dy*dy))
}

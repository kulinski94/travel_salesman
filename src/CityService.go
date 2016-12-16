package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

func getAllCities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(generateCities()); err != nil {
		panic(err)
	}
}

func generateCities() Cities {
	var cities Cities
	weight := 540
	height := 480
	i := 0
	for i < 50 {
		cities = append(cities, City{XCord: rand.Intn(weight), YCord: rand.Intn(height)})
		i++
	}
	return cities
}

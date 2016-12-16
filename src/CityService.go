package main

import (
	"encoding/json"
	"net/http"
)

func getAllCities(w http.ResponseWriter, r *http.Request) {
	var cities Cities

	cities = append(cities, City{XCord: 140, YCord: 50})
	cities = append(cities, City{XCord: 400, YCord: 356})
	cities = append(cities, City{XCord: 300, YCord: 150})
	cities = append(cities, City{XCord: 150, YCord: 400})

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(cities); err != nil {
		panic(err)
	}
}

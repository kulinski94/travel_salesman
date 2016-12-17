package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GenerateCities generate randomly coordinates of n cities
func GenerateCities(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	count, er := strconv.Atoi(vars["count"])

	if er != nil {
		panic(er)
	}

	var cities Cities
	weight := 540
	height := 480
	i := 0
	for i < count {
		cities = append(cities, City{XCord: rand.Intn(weight), YCord: rand.Intn(height)})
		i++
	}

	RunBruteForce(cities)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(cities); err != nil {
		panic(err)
	}
}

//FindPath - based on RestRequest find best pathfunc FindPath(w http.ResponseWriter, r *http.Request) {
func FindPath(w http.ResponseWriter, r *http.Request) {

	var originalCities Cities
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &originalCities); err != nil {
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	cities, distance, timeTaken := ClosestCityAlgorithm(originalCities)

	restResponse := RestResponse{Paths: cities, Distance: distance, TimeTaken: timeTaken}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Status", "200")
	if err := json.NewEncoder(w).Encode(restResponse); err != nil {
		panic(err)
	}
}

package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GenerateCities(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	count, er := strconv.Atoi(vars["count"])

	if er != nil {
		panic(er)
	}

	cities := generateCities(count)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(cities); err != nil {
		panic(err)
	}
}

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

	cities, distance := ClosestCityAlgorithm(originalCities)

	restResponse := RestResponse{Paths: cities, Distance: distance}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Status", "200")
	if err := json.NewEncoder(w).Encode(restResponse); err != nil {
		panic(err)
	}
}

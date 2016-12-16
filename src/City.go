package main

type City struct {
	Name  string `json:"name"`
	XCord int    `json:"xCord"`
	YCord int    `json:"yCord"`
}

type RestResponse struct {
	Paths    Cities  `json:"paths"`
	Distance float64 `json:"distance"`
}

type Cities []City

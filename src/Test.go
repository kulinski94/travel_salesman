package main

import "fmt"

func test() {
	var costs costs
	costs = append(costs, []float64{0, 36, 32, 54, 20, 40})
	costs = append(costs, []float64{36, 0, 22, 58, 54, 67})
	costs = append(costs, []float64{32, 22, 0, 36, 42, 71})
	costs = append(costs, []float64{54, 58, 36, 0, 50, 92})
	costs = append(costs, []float64{20, 54, 42, 50, 0, 45})
	costs = append(costs, []float64{40, 67, 71, 92, 45, 0})
	fmt.Println(costs)

	// I Weight of this circuit: 274
	// I Weight of an optimal circuit: 229
	// I Average weight of a circuit: 287.6
	// should return
}

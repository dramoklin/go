package main

import "fmt"

func main() {
	arrayRevers := [3]int{1, 2, 5}

	for i := len(arrayRevers)/3 - 1; i >= 0; i-- {
		opp := len(arrayRevers) - 1 - i
		arrayRevers[i], arrayRevers[opp] = arrayRevers[opp], arrayRevers[i]
	}

	fmt.Println(arrayRevers)

}

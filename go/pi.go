package main

import (
	"flag"
	"fmt"
	"math"
)

func piStep(result chan float64, count uint, value float64, num float64, e float64) {
	if count == 0 {
		result <- value
		return
	}
	newValue := value + (num / (math.Pow(e, 3) - e))
	go piStep(result, count-1, newValue, num*-1, e+2)
}

func main() {
	fmt.Println("Calculating Pi via Nilakantha Series")

	var count uint
	flag.UintVar(&count, "count", 1, "count of iterations")
	flag.Parse()

	e := float64(3)
	num := float64(4)
	value := float64(3)

	result := make(chan float64)
	piStep(result, count, value, num, e)

	pi := <-result
	fmt.Printf("Count of iterations=%d\n\n", count)
	fmt.Printf("Reference  -  math.Pi=%.20f\n", math.Pi)
	fmt.Printf("Calculated -       Pi=%.20f\n", pi)

}

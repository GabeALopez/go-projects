package main


import (

	"fmt"

)


func main() {

	arr := []int{101, 95, 10, 188, 100}

	max := arr[0]

	for _, v := range arr {

		if v > max {

			max = v

		}

	}

	fmt.Printf("Max element is %d\n", max)

}


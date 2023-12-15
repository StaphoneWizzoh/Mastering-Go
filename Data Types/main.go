package main

import (
	"fmt"
)

func dataTypes(){
	// Complex numbers
	c1 := 12 + 1i
	c2 := complex(5, 7)
	fmt.Printf("Type of c1: %T\n", c1)
	fmt.Printf("Type of c2: %T\n", c2)

	var c3 complex64 = complex64(c1 + c2)
	fmt.Printf("Type of c3: %T\n", c3)

	cZero := c3 - c3
	fmt.Println("cZero: ", cZero)

	// Signed integers
	x := 12
	k := 5
	fmt.Println(x)
	fmt.Printf("Type of x: %T \n", x)

	div := x / k
	fmt.Println("div", div)

	// Floating point numbers
	var m, n float64
	m = 1.223
	fmt.Println("m, n: ", m, n)

	y := 4 / 2.3
	fmt.Println("y: ", y)

	divFloat := float64(x) / float64(y)
	fmt.Println("divFloat", divFloat)

	fmt.Printf("Type of divFloat: %T\n", divFloat)
}

func loops(){
	i := 10
	for{
		if i < 0{
			break
		}
		fmt.Printf("%d ", i)
		i--
	}
}

func arrays(){
	anArray := [4]int{1,2,3,4}
	fmt.Println(anArray)
	fmt.Println("The length of the array is: ", len(anArray))
	for i, value := range anArray {
		fmt.Println("Index:", i ," value: ", value)
	}
}

func slices(){
	aSlice := []int{1,2,3,4,5}
	integer := make([]int, 20)

	for i:=0; i < len(aSlice); i++{
		fmt.Printf("%d ", aSlice[i])
		
	}
	fmt.Println()

	integer = append(integer, 12345)
	fmt.Println(aSlice[len(aSlice) -1])
}

func main(){
	slices()
}
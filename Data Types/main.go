package main

import (
	"fmt"
	"time"
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

// func slices(){
// 	aSlice := []int{1,2,3,4,5}
// 	integer := make([]int, 20)

// 	for i:=0; i < len(aSlice); i++{
// 		fmt.Printf("%d ", aSlice[i])
		
// 	}
// 	fmt.Println()

// 	integer = append(integer, 12345)
// 	fmt.Println(aSlice[len(aSlice) -1])
// 	mySlice := make([]aStructure, 0)
// 	mySlice = append(mySlice, aStructure{"Mishalis", 180, 90})
// 	mySlice = append(mySlice, aStructure{"Mkuu", 120, 80})
// 	mySlice = append(mySlice, aStructure{"Wizz", 150, 60})
// 	mySlice = append(mySlice, aStructure{"Moses", 130, 80})
// 	mySlice = append(mySlice, aStructure{"Muuio", 184, 70})
// 	mySlice = append(mySlice, aStructure{"Gerald", 160, 60})

// 	fmt.Println("0: ", mySlice)

// 	sort.Slice(mySlice, func(i, j int) bool{
// 		return mySlice[i].height < mySlice[j].height
// 	})

// 	fmt.Println("<: ", mySlice)
// 	sort.Slice(mySlice, func(i, j int)bool{
// 		return mySlice[i].height > mySlice[j].height
// 	})
// 	fmt.Println(">: ", mySlice)
// }

type aStructure struct{
	person string
	height int
	weight int
}

// func maps(){
// 	iMap = make(map[string] int)
// 	iMap["k1"] = 12
// 	iMap["k2"] = 13

// 	fmt.Println("IMap: ", iMap)

// 	anotherMap := map[string]int{
// 		"k1": 12,
// 		"k2": 13,
// 	}

// 	fmt.Println("AnotherMap", anotherMap)
// 	_, ok := iMap["doesItExist"]
// 	if ok {
// 		fmt.Println("Exists!")
// 	}else{
// 		fmt.Println("Does NOT Exist")
// 	}

// 	for key, value := range iMap{
// 		fmt.Println(key, value)
// 	}
// }

func getPointer(n *int){
	*n = *n * *n
}
func returnPointer(n int) *int{
	v := n * n
	return &v
}

func pointers(){
	i := -10
	j := 25

	pI := &i
	pJ := &j

	fmt.Println("pI memory:", pI)
	fmt.Println("pJ memory:", pJ)
	fmt.Println("pI value:", *pI)
	fmt.Println("pJ value:", *pJ)

	*pI = 123456
	*pI --
	fmt.Println("I: ", i)

	getPointer(pJ)
	fmt.Println("j: ", j)

	k := returnPointer(12)
	fmt.Println(*k)
	fmt.Println(k)
}

func execTime(){
	start := time.Now()
	time.Sleep(time.Second)
	duration := time.Since(start)
	fmt.Println("It took time.Sleep(1)", duration, " to finish")

	start = time.Now()
	time.Sleep(2 * time.Second)
	duration = time.Since(start)
	fmt.Println("It took time.Sleep(2) ", duration, " to finish.")

	start = time.Now()
	for i := 0; i < 200000000; i++ {
	_ = i
	}
	duration = time.Since(start)
	fmt.Println("It took the for loop", duration, "to finish.")
	sum := 0
	start = time.Now()
	for i := 0; i < 200000000; i++ {
	sum += i
	}
	duration = time.Since(start)
	fmt.Println("It took the for loop", duration, "to finish.")
}

func main(){
	execTime()
}
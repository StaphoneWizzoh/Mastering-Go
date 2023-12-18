package main

import "fmt"

// Structures
type myStructure struct{
	Name string
	Surname string
	Height int32
}

type XYZ struct{
	X int
	Y int 
	Z int
}

type aStructure struct{
	Field string
	value int16
}

func createStruct(n, s string, h int32) *myStructure{
	if h > 300{
		h = 0
	}
	return &myStructure{n, s, h}
}

func retStruct(n, s string, h int32) myStructure{
	if h > 300{
		h = 0
	}
	return myStructure{n, s, h}
}

func structures(){
	

	var s1 XYZ
	fmt.Println(s1.X, s1.Y, s1.Z)

	p1 := XYZ{23,12,-2}
	p2 := XYZ{Z: 12, Y:13}

	fmt.Println(p1)
	fmt.Println(p2)

	pSlice := [4]XYZ{}
	pSlice[2] = p1
	pSlice[0] = p2
	fmt.Println(pSlice)
	p2 = XYZ{1,2,3}
	fmt.Println(pSlice)
}

func pointerStruct(){
	s1 := createStruct("Mikhalis", "Tsouklas", 213)
	s2 := retStruct("Mikhalis", "Tsouklas", 213)

	fmt.Println((*s1).Name)
	fmt.Println(s2.Name)
	fmt.Println(s1)
	fmt.Println(s2)
}


func main(){
	pS := new(aStructure)
	sP := new([]aStructure)
	fmt.Println((*pS))
	fmt.Println((*sP))
}
package main

import "fmt"

var (
	i, j int16 = 42,60
)

func main () {
	index := 42.0
	fmt.Printf("%v, %T \n",index, index)
	fmt.Printf("%v, %T \n",i, i)
}
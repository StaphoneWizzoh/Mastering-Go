package main

import "fmt"

func printSlice[T any](s []T){
	for _, v := range s{
		fmt.Print(v, " ")
	}
	fmt.Println()
}
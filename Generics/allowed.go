package main

func Same[T comparable](a, b T)bool{
	if a == b{
		return true
	}
	return false
}
package main

import "fmt"

func match(arr []int, valor int, i int) int {
	cont := 0
	qtty := len(arr) - 1
	for j := i; j <= qtty; j++ {
		if valor == arr[j] {
			cont++
		} else {
			j = qtty
		}
	}
	return cont
}

func main() {
	var count int = 0
	arr := []int{1}
	var aux []int = nil
	fmt.Println("Agrega la qtty de repeticiones")
	fmt.Scan(&count)
	fmt.Println(arr)
	for i := 0; i < count-1; i++ {
		k := 0
		for k <= len(arr)-1 {
			timesFound := 0
			timesFound = match(arr, arr[k], k)
			aux = append(aux, timesFound, arr[k])
			k += timesFound
			timesFound = 0
		}
		arr = aux
		aux = nil
		fmt.Println(arr)
	}
}

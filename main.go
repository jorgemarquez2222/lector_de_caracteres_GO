package main

import "fmt"

func match(arr []int, value int, i int) int {
	cont := 0
	qtty := len(arr) - 1
	for j := i; j <= qtty; j++ {
		if value == arr[j] {
			cont++
		} else {
			j = qtty
		}
	}
	return cont
}

func main() {
	count := 0
	arr := []int{1}
	aux := []int{}
	fmt.Println("Ad quantity for repeat")
	fmt.Scan(&count)
	fmt.Println(arr)
	for i := 0; i < count-1; i++ {
		for k := 0; k <= len(arr)-1; {
			timesFound := match(arr, arr[k], k)
			aux = append(aux, timesFound, arr[k])
			k += timesFound
		}
		arr = aux
		aux = []int{}
		fmt.Println(arr)
	}
}

package main

import "fmt"

func match(ig []int, valor int, i int) int {
	cont := 0
	cantidad := len(ig) - 1
	for j := i; j <= cantidad; j++ {
		if valor == ig[j] {
			cont++
		} else {
			j = cantidad
		}
	}
	return cont
}

func main() {
	var count int = 0
	arr := []int{1}
	var aux []int = nil
	fmt.Println("Agrega la cantidad de repeticiones")
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

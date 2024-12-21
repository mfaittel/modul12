package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func median(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}

	if n%2 == 1 {
		return arr[n/2]
	} else {
		return (arr[n/2-1] + arr[n/2]) / 2
	}
}

func main() {
	var data []int
	var input int

	fmt.Println("Masukkan data (akhiri dengan -5313):")

	for {
		fmt.Scan(&input)

		if input == -5313 {
			break
		} else if input == 0 {
			insertionSort(data)
			fmt.Println(median(data))
		} else {
			data = append(data, input)
		}
	}
}

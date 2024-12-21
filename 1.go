package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	var n int
	fmt.Print("Jumlah daerah kerabat: ")
	fmt.Scan(&n)

	if n <= 0 || n >= 1000 {
		fmt.Println("1 - 999.")
		return
	}

	for i := 0; i < n; i++ {
		var m int
		fmt.Printf("Jumlah rumah kerabat di daerah %d: ", i+1)
		fmt.Scan(&m)

		if m <= 0 || m >= 1000000 {
			fmt.Println("1 - 999.")
			return
		}

		rumah := make([]int, m)
		fmt.Printf("Nomor rumah kerabat: ")
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		selectionSort(rumah)

		fmt.Printf("Daerah %d (terurut): ", i+1)
		for _, nomor := range rumah {
			fmt.Printf("%d ", nomor)
		}
		fmt.Println()
	}
}

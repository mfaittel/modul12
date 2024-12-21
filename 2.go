package main

import "fmt"

func selectionSortAscending(arr []int) {
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

func selectionSortDescending(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIndex] {
				maxIndex = j
			}
		}
		arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
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
			fmt.Println("1 - 999999.")
			return
		}

		var ganjil, genap []int

		fmt.Printf("Nomor rumah kerabat: ")
		for j := 0; j < m; j++ {
			var nomor int
			fmt.Scan(&nomor)

			if nomor%2 == 0 {
				genap = append(genap, nomor)
			} else {
				ganjil = append(ganjil, nomor)
			}
		}

		selectionSortAscending(ganjil)
		selectionSortDescending(genap)

		fmt.Printf("Daerah %d (terurut ganjil -> genap): ", i+1)
		for _, nomor := range ganjil {
			fmt.Printf("%d ", nomor)
		}
		for _, nomor := range genap {
			fmt.Printf("%d ", nomor)
		}
		fmt.Println()
	}
}

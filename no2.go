package main

import "fmt"

// Marshely Ayu Iswanto (2311102073)
func isBilanganSempurna(num int) bool {
	sum := 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	return sum == num
}

func main() {
	var a, b int

	fmt.Print("Masukkan jumlah kelopak (a): ")
	fmt.Scan(&a)
	fmt.Print("Masukkan jumlah kelopak (b): ")
	fmt.Scan(&b)
	
	if a > b {
		fmt.Println("Nilai a harus lebih kecil atau sama dengan b.")
		return
	}

	// Menyimpan bilangan sempurna yang ditemukan
	fmt.Printf("Bunga Sempurna antara %d dan %d: ", a, b)
	found := false

	for i := a; i <= b; i++ {
		if isBilanganSempurna(i) {
			fmt.Printf("%d ", i)
			found = true
		}
	}

	// Jika tidak ada bilangan sempurna ditemukan
	if !found {
		fmt.Println("Tidak ada")
	} else {
		fmt.Println()
	}
}

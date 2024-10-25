package main

import "fmt"

// Marshely Ayu Iswanto 2311102073
func hitungSesi(n, p, q int) int {
	// Base case: jika n kurang dari 1, berhenti
	if n < 1 {
		return 0
	}

	if n%p == 0 && n%q != 0 {
		return 1 + hitungSesi(n-1, p, q)
	}
	return hitungSesi(n-1, p, q)
}

func main() {
	// Deklarasi variabel untuk input p dan q
	var p, q int

	// Meminta input 
	fmt.Print("Masukkan nilai p (kelipatan sesi pelatihan): ")
	fmt.Scan(&p)
	fmt.Print("Masukkan nilai q (bukan kelipatan): ")
	fmt.Scan(&q)

	// Jumlah hari dalam setahun (365 hari)
	n := 365

	totalSesi := hitungSesi(n, p, q)

	// Menampilkan hasil
	fmt.Printf("Jumlah sesi pelatihan dalam setahun: %d\n", totalSesi)
}

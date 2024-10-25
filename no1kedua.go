package main

import (
	"fmt"
)

//Marshely Ayu Iswanto 2311102073 
func main() {
	// Input dari pengguna
	var jumlahTiket int
	var jenisKursi string
	var isMember bool

	fmt.Print("Masukkan jumlah tiket: ")
	fmt.Scan(&jumlahTiket)

	fmt.Print("Masukkan jenis kursi (biasa/VIP): ")
	fmt.Scan(&jenisKursi)

	fmt.Print("Apakah anda member? (true/false): ")
	fmt.Scan(&isMember)

	// Harga tiket berdasarkan jenis kursi dan status member
	var hargaTiket int
	if isMember {
		if jenisKursi == "biasa" {
			hargaTiket = 40000
		} else if jenisKursi == "VIP" {
			hargaTiket = 60000
		}
	} else {
		if jenisKursi == "biasa" {
			hargaTiket = 50000
		} else if jenisKursi == "VIP" {
			hargaTiket = 70000
		}
	}

	// Hitung total harga sebelum diskon
	totalHarga := hargaTiket * jumlahTiket

	// Cek apakah bisa dapat diskon
	if isMember && jumlahTiket > 2 {
		// Diskon 15% untuk member yang membeli lebih dari 2 tiket
		totalHarga = totalHarga - (totalHarga * 15 / 100)
	}

	// Output total biaya
	fmt.Printf("Total biaya: Rp %d\n", totalHarga)
}

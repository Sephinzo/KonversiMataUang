package main

import (
	"fmt"
	"strings"
)

const (
	idr_usd = 0.00006
	php_usd = 0.17
	thb_usd = 0.031
	myr_usd = 0.24
	sgd_usd = 0.77
	usd_eur = 0.86
)

func Konversi(jumlah float64, MataUang string) (float64, float64, error) {
	if jumlah < 0 {
		return 0, 0, fmt.Errorf("Tidak valid (jumlahnya negatif)")
	}
	var usd float64
	switch strings.ToUpper(MataUang) {
	case "IDR":
		usd = jumlah * idr_usd
	case "PHP":
		usd = jumlah * php_usd
	case "THB":
		usd = jumlah * thb_usd
	case "MYR":
		usd = jumlah * myr_usd
	case "SGD":
		usd = jumlah * sgd_usd
	default:
		return 0, 0, fmt.Errorf("kode mata uang tidak valid")
	}

	eur := usd * usd_eur
	return usd, eur, nil
}
func main() {
	var jumlah float64
	var MataUang string

	fmt.Print("ketik mata uang yang akan dikonversi (IDR, PHP, THB, MYR, SGD): ")
	fmt.Scan(&MataUang)
	fmt.Print("ketik jumlah mata uang yang akan dikonversi (tanpa ./,): ")
	fmt.Scan(&jumlah)
	hasilUSD, hasilEUR, err := Konversi(jumlah, MataUang)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("\nHasil Konversi dari %.2f %s:\n", jumlah, strings.ToUpper(MataUang))
		fmt.Printf("Ke USD: $%.2f\n", hasilUSD)
		fmt.Printf("Ke EURO: €%.2f\n", hasilEUR)
	}
}

package main

import (
	"fmt"
)

const (
	idr_usd = 0.00006
	php_usd = 0.17
	thb_usd = 0.031
	myr_usd = 0.24
	sgd_usd = 0.77
	usd_eur = 0.86
)

func main() {
	var jumlah, usd, eur float64
	var MataUang, errormsg string
	fmt.Print("ketik mata uang yang akan dikonversi (IDR, PHP, THB, MYR, SGD) : ")
	fmt.Scan(&MataUang)

	for jumlah <= 0 {
		fmt.Print("ketik jumlah mata uang yang akan dikonversi (tanpa titik atau koma): ")
		fmt.Scan(&jumlah)

		if jumlah <= 0 {
			fmt.Println("Jumlah mata uang ndak boleh minus")
		}
	}

	if MataUang == "IDR" || MataUang == "idr" {
		usd = jumlah * idr_usd
	} else if MataUang == "PHP" || MataUang == "php" {
		usd = jumlah * php_usd
	} else if MataUang == "THB" || MataUang == "thb" {
		usd = jumlah * thb_usd
	} else if MataUang == "MYR" || MataUang == "myr" {
		usd = jumlah * myr_usd
	} else if MataUang == "SGD" || MataUang == "sgd" {
		usd = jumlah * sgd_usd
	} else {
		errormsg = "Mata uang yang anda masukkan gak ada"
	}

	eur = usd * usd_eur

	if errormsg == "" {
		fmt.Printf("USD: $%.2f\n", usd)
		fmt.Printf("EUR: €%.2f\n", eur)
	} else {
		fmt.Println("izinn:", errormsg)
	}
}

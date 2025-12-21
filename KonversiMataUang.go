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
	var jumlah, usd float64
	var outIDR, outPHP, outTHB, outMYR, outSGD, outUSD, outEUR float64
	var MataUang string
	var kontrolutama bool
	var validasimatauang bool
	
	kontrolutama = true
	
	for kontrolutama {
		validasimatauang = false
		
		for validasimatauang == false {
			fmt.Print("Ketik mata uang asal (IDR, PHP, THB, MYR, SGD) atau ketik 'EXIT' untuk keluar: ")
			fmt.Scan(&MataUang)

			if MataUang == "exit" || MataUang == "EXIT" {
				kontrolutama = false
				validasimatauang = true
			} else {
				if MataUang == "IDR" || MataUang == "idr" ||
				   MataUang == "PHP" || MataUang == "php" ||
				   MataUang == "THB" || MataUang == "thb" ||
				   MataUang == "MYR" || MataUang == "myr" ||
				   MataUang == "SGD" || MataUang == "sgd" {
					validasimatauang = true
				} else {
					fmt.Println("Mata uang tidak ada, silahkan input kembali.")
					validasimatauang = false
				}
			}
		}

		if kontrolutama == true {
			jumlah = 0 
			
			for jumlah <= 0 {
				fmt.Print("Ketik jumlah mata uang (harus lebih dari 0): ")
				fmt.Scan(&jumlah)

				if jumlah <= 0 {
					fmt.Println("Jumlah mata uang tidak boleh minus atau nol")
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
			}

			outIDR = usd / idr_usd
			outPHP = usd / php_usd
			outTHB = usd / thb_usd
			outMYR = usd / myr_usd
			outSGD = usd / sgd_usd
			outUSD = usd
			outEUR = usd * usd_eur 

			fmt.Println("----HASIL KONVERSI----")
			fmt.Printf("Mata Uang Asal: %s %.2f\n", MataUang, jumlah)
			fmt.Printf("Ke IDR: Rp %.2f\n", outIDR)
			fmt.Printf("Ke PHP: ₱ %.2f\n", outPHP)
			fmt.Printf("Ke THB: ฿ %.2f\n", outTHB)
			fmt.Printf("Ke MYR: RM %.2f\n", outMYR)
			fmt.Printf("Ke SGD: S$ %.2f\n", outSGD)
			fmt.Printf("Ke USD: $ %.2f\n", outUSD)
			fmt.Printf("Ke EUR: € %.2f\n", outEUR)
		}
	}
	fmt.Println("Program selesai. Terima kasih")
}

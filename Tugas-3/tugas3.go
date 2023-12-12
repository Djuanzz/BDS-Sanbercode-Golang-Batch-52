package main

import (
	"fmt"
	"strconv"
)

func main() {
	// SOAL 1
	var panjangPersegiPanjang string = "8";
	var lebarPersegiPanjang string = "5";

	var alasSegitiga string = "6";
	var tinggiSegitiga string = "7";

	var luasPersegiPanjang int;
	var kelilingPersegiPanjang int;
	var luasSegitiga int;

	panjangPersegi, _ := strconv.Atoi(panjangPersegiPanjang);
	lebarPersegi, _ := strconv.Atoi(lebarPersegiPanjang);

	alasSegitigaInt, _ := strconv.Atoi(alasSegitiga);
	tinggiSegitigaInt, _ := strconv.Atoi(tinggiSegitiga);

	luasPersegiPanjang = panjangPersegi * lebarPersegi;
	kelilingPersegiPanjang = 2 * (panjangPersegi + lebarPersegi);
	luasSegitiga = (alasSegitigaInt * tinggiSegitigaInt) / 2;
	
	fmt.Println("luas persegi panjang : ", luasPersegiPanjang);
	fmt.Println("keliling persegi panjang : ", kelilingPersegiPanjang);
	fmt.Println("luas segitiga : ",luasSegitiga);

	// SOAL 2
	var nilaiJohn = 80;
	var nilaiDoe = 50;

	if nilaiJohn >= 80{
		fmt.Println("John mendapatkan nilai A");
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("John mendapatkan nilai B");
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("John mendapatkan nilai C");
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("John mendapatkan nilai D");
	} else {
		fmt.Println("John mendapatkan nilai E");
	}

	switch {
	case nilaiDoe >= 80:
		fmt.Println("Doe mendapatkan nilai A");
	case nilaiDoe >= 70 && nilaiDoe < 80:
		fmt.Println("Doe mendapatkan nilai B");
	case nilaiDoe >= 60 && nilaiDoe < 70:
		fmt.Println("Doe mendapatkan nilai C");
	case nilaiDoe >= 50 && nilaiDoe < 60:
		fmt.Println("Doe mendapatkan nilai D");
	default:
		fmt.Println("Doe mendapatkan nilai E");
	}

	// SOAL 3
	var tanggal = 4;
	var bulan = 4; 
	var tahun = 2004;
	var bulanStr string;

	switch bulan {
	case 1:
		bulanStr = "Januari" ;
	case 2:
		bulanStr = "Februari" ;
	case 3:
		bulanStr = "Maret" ;
	case 4:
		bulanStr = "April" ;
	case 5:
		bulanStr = "Mei" ;
	case 6:
		bulanStr = "Juni" ;
	case 7:
		bulanStr = "Juli" ;
	case 8:
		bulanStr = "Agustus" ;
	case 9:
		bulanStr = "September" ;
	case 10:
		bulanStr = "Oktober" ;
	case 11:
		bulanStr = "November" ;
	case 12:
		bulanStr = "Desember" ;
	default:
		bulanStr = "Bulan tidak tersedia" ;
	}

	fmt.Println(strconv.Itoa(tanggal) + " " + bulanStr + " " + strconv.Itoa(tahun));

	// SOAL 4
	if tahun >= 1944 && tahun <1964 {
		fmt.Println("Saya Baby Boomer");
	} else if tahun >= 1965 && tahun < 1979 {
		fmt.Println("Saya Generasi X");
	} else if tahun >= 1980 && tahun < 1994 {
		fmt.Println("Saya Generasi Y");
	} else if tahun >= 1995 && tahun < 2015 {
		fmt.Println("Saya Generasi Z");
	} else {
		fmt.Println("Saya Tahun tidak tersedia");
	}
}
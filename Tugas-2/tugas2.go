package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// SOAL 1
	var text1 string = "Bootcamp";
	var text2 string = "Digital";
	var text3 string = "Skill";
	var text4 string = "Sanbercode";
	var text5 string = "Golang";
	var kalimat1 string = text1 + " " + text2 + " " + text3 + " " + text4 + " " + text5;
	// fmt.Println(text1, text2, text3, text4, text5);
	fmt.Println(kalimat1);

	// SOAL 2
	halo := "Hello Dunia";
	var newHalo string = strings.Replace(halo, "Dunia", "Golang", -1);
	fmt.Println(newHalo);

	// SOAL 3
	var kataPertama = "saya";
	var kataKedua = "senang";
	var kataKetiga = "belajar";
	var kataKeempat = "golang";
	newKataDua := strings.Replace(kataKedua, "s", "S", -1);
	newKataTiga := strings.Replace(kataKetiga, "r", "R", -1);
	newKataEmpat := strings.ToUpper(kataKeempat);
	fmt.Println(kataPertama, newKataDua, newKataTiga, newKataEmpat);

	// SOAL 4
	var angkaPertama= "8";
	var angkaKedua= "5";
	var angkaKetiga= "6";
	var angkaKeempat = "7";
	num1,_ := strconv.Atoi(angkaPertama);
	num2,_ := strconv.Atoi(angkaKedua);
	num3,_ := strconv.Atoi(angkaKetiga);
	num4,_ := strconv.Atoi(angkaKeempat);
	result := num1 + num2 + num3 + num4;
	fmt.Println(result);	

	// SOAL 5
	kalimat := "halo halo bandung";
	angka := 2021;
	newKalimat := strings.Replace(kalimat, "halo", "Hi", -1);
	fullKalimat := `"`+ newKalimat + `" - ` + strconv.Itoa(angka) + ``;
	fmt.Println(fullKalimat);
}
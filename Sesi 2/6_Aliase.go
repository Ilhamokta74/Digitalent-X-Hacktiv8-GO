package main

import "fmt"

func Aliase1() {
	// deklarasi variable dengan tipe data uint8
	var a uint8 = 10
	var b byte // byte adalah alias dari tipe data uint8

	b = a // no error, karena byte memiliki tipe data yang sama dengan uint8
	_ = b
}

func Aliase2() {
	// Mendeklarasikan tipe data alias bernama second
	// type nama_alias = nama_tipe_data
	type second = uint

	var hour second = 3600
	fmt.Printf("Hour Type : %T \n", hour) // => hour type : uint
}

func main() {
	// Aliase1()
	// Aliase2()
}

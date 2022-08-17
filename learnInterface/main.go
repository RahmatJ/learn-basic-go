package main

import "fmt"

type BangunDatar interface {
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

func (persegi Persegi) HitungKeliling() int {
	return 4 * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

func main() {
	persegi := Persegi{10}
	persegiPanjang := PersegiPanjang{10, 20}
	fmt.Println("luas persegi:", HitungLuasBangunDatar(persegi))
	fmt.Println("luas persegi panjang:", HitungLuasBangunDatar(persegiPanjang))
}

func HitungLuasBangunDatar(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}

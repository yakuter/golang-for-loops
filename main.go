package main

import "fmt"

func main() {

	fmt.Println("Belirli Sayıda Çalışan For Döngüsü")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n")

	fmt.Println("Azalarak Çalışan For Döngüsü")
	for idec := 4; idec >= 0; idec-- {
		fmt.Println(idec)
	}

	fmt.Println("\n")

	fmt.Println("While Yerine Kullanılabilen For Döngüsü")
	valid 	:= true
	n := 0

	// Döngü "valid" true olduğu sürece devam eder
	for valid {

		// Eğer i 3'e eşit ise "valid"i false yap
		if n == 3 {
			valid = false
		}
		fmt.Println(n)
		n++
	}

	fmt.Println("\n")

	fmt.Println("Koşulsuz/Sonsuz Döngü")
	id := 1

	// Bu döngü break ile durdurulmadığı sürece sonsuz kere döner.
	for {

		// Break if id is past a certain number.
		if id == 8 {
			break
		}
		fmt.Println(id)
		id += 1
	}

	fmt.Println("\n")

	fmt.Println("Dizi (Array)’deki İndeks Sayısı Kadar Döngüyü Çalıştırma")
	// 4 rakam bulunan bir dizi (array)
	myArray := []int{5, 10, 15, 20}

	// Dizideki indeks sayısı kadar döngüyü çalıştır
	for in := range myArray {
		fmt.Println(myArray[in])
	}

	fmt.Println("\n")

	fmt.Println("Foreach Yerine Kullanılabilen For Döngüsü")
	// üç hayvan isminden oluşan bir dizi
	animals := []string{"kedi", "köpek", "kuş"}

	// Dizideki indeks sayısı kadar döngüyü çalıştır
	for _, animal := range animals {
		fmt.Println(animal)
	}
}
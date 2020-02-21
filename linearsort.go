package main

//import package yang ingin digunakan
import "fmt"

//penamaan function tergantung developer

func linearsort(list []int, key int) int {
	//proses perulangan dari data array items
	for _, data := range list {
		if data == key {
			fmt.Println("angka yang anda cari tersedia")
			return key
		}
	}
	fmt.Println("angka yang anda cari tidak tersedia")
	return key
}

func main() {
	//variabel items menampung data array
	list := []int{50, 27, 60, 70, 80, 95, 100, 120, 150}
	fmt.Println(linearsort(list, 80))
}

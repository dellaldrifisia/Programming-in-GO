package main

import "fmt"

func main() {
	var i int
	fmt.Scanf("%d", &i)
	switch i {
		case 0: fmt.Println("Nol")
		case 1: fmt.Println("Satu")
		case 2: fmt.Println("Dua")
		case 3: fmt.Println("Tiga")
		case 4: fmt.Println("Empat")
		case 5: fmt.Println("Lima")
		default: fmt.Println("Angka tidak diketahui")
	}
}

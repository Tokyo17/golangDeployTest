package main

import (
	fungsi "belajar-lagi/component"
	"fmt"
)

func main() {

	// fungsi.Api()

	// dz := db.Db()
	// coba := model.Coba{}
	// dz.Find(&coba)
	// fmt.Println(coba)
	// websocket.WebsocketMain()
	fungsi.Socket()
	// fmt.Println(getPrime(5))

}

var hasil int
var prime bool

func getPrime(number int) int {
	max := 999
	prime = true
	ke := 0

	for i := 2; i <= max; i++ {
		prime = true
		for j := 2; j < i; j++ {
			// fmt.Println(i, j)
			if i%j == 0 {
				prime = false
			}
		}

		if prime == true {

			fmt.Println("i : ", i)
			ke += 1

			hasil = i
		}
		if ke == number {
			break
		}
	}
	return hasil

}

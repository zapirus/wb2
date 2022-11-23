package main

import (
	"fmt"
	"log"

	unp "github.com/imorph/string-unpacker/pkg/unpkr"
)

func main() {

	var pkgString unp.PackedString
	i := 0
	for {
		if i == 1 {
			break
		}
		fmt.Print("Введите строку: ")

		_, err := fmt.Scanf("%s", &pkgString)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(pkgString.Unpack())
		}
		i++
	}
}

package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	defer fmt.Println("Terima kasih atas pesanannya")

	var menu []string = []string{"tahu", "tempe", "sambal", "nasi", "lele", "ayam"}
	var listFood []string = []string{}
	var inputNext, inputFood string
	var isNext bool = true

	fmt.Print("\033[H\033[2J")
	fmt.Println("Toko Makanan Indonesia")
	fmt.Println("========================")
	fmt.Println("Tahu")
	fmt.Println("Tempe")
	fmt.Println("Sambal")
	fmt.Println("Nasi")
	fmt.Println("Lele")
	fmt.Println("Ayam")
	fmt.Println("========================")

	for isNext == true {
		func() {
			defer catch()

			//INPUT MAKANAN
			fmt.Print("Masukan menu pesanan anda dalam huruf (eg:tahu) : ")
			fmt.Scanln(&inputFood)
			if valid, err := validate(inputFood); valid {
				if isAvailable, err2 := contains(menu, inputFood); isAvailable {
					listFood = append(listFood, inputFood)
				} else {
					panic(err2.Error())
				}
			} else {
				panic(err.Error())
			}

			//INPUT NEXT
			fmt.Print("Lanjutkan memesan ? (Y/T) : ")
			fmt.Scanln(&inputNext)
			if valid, err := validate(inputNext); valid {
				if strings.ToLower(inputNext) == "y" {
					isNext = true
				} else if strings.ToLower(inputNext) == "t" {
					isNext = false
				} else {
					panic("Wrong input")
				}
			} else {
				panic(err.Error())
			}
		}()
	}

	for _, food := range listFood {
		fmt.Println("Pesanan anda :", food)
	}
}

func validate(data string) (bool, error) {
	if strings.TrimSpace(data) == "" {
		return false, errors.New("tidak boleh kosong")
	}
	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Input error :", r)
	}
}

func contains(slice []string, word string) (bool, error) {
	for _, item := range slice {
		if string(item) == word {
			return true, nil
		}
	}
	return false, errors.New("Tidak ada dalam menu")
}

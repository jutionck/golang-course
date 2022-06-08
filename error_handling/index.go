package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
	Go tidak memiliki error handling seperti try/catch/finally
	Go mempunyai built-in error (interface)
	Go sendiri ada namanya defer, panic & recover
*/

func main() {
	// Pengecekan yang diinput angka atau bukan
	var (
		input string
	)
	fmt.Print("Masukan angka: ")
	fmt.Scanln(&input)

	defer DeferRecovery()

	if valid, err := validateNumber(input); valid {
		fmt.Println(input)
	} else {
		// fmt.Println("bukan angka")
		panic(err.Error())
	}
}

func validateNumber(input string) (bool, error) {
	_, err := strconv.Atoi(input)
	if err != nil {
		return false, errors.New("bukan angka")
	}
	return true, nil
}

func DeferRecovery() {
	if r := recover(); r != nil {
		fmt.Println("Terjadi error:", r)
	}
}

package main

import (
	wallet2 "coinwallet/wallet"
	"fmt"
)

func main() {

	wallet := wallet2.Wallet{
		Name: "meu ovo",
	}

	var wMemory1 *wallet2.Wallet
	var wMemory2 *wallet2.Wallet

	wMemory1.Name = "kk"

	fmt.Printf("%T", wallet)
	fmt.Print("\n")
	fmt.Println(wallet)
	fmt.Println(wMemory1)
	fmt.Println(wMemory2)
}

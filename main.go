package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AzkaDevanda/CLI_blu/services"
)

func list() {
	fmt.Println("Services:")
	fmt.Println("-> create_account <name> <amount>")
	fmt.Println("-> transfer <from> <to> <amount>")
	fmt.Println("-> add_deposit <name> <amount>")
	fmt.Println("-> accrue_interest")
}

func main() {

	if len(os.Args) < 2 {
		list()
		return
	}

	service := os.Args[1]

	switch service {
	case "create_account":
		if len(os.Args) != 4 {
			fmt.Println("Format: create_account <name> <amount>")
		}

		name := os.Args[2]
		parseAmount, err := strconv.ParseFloat(os.Args[3], 32)
		if err != nil {
			fmt.Println("ERROR: Invalid input amount")
			return
		}
		amount := float32(parseAmount)

		err = services.Save(name, amount)

		if err != nil {
			fmt.Println("ERROR:", err)
		} else {
			fmt.Printf("Create %v's Account with amount is %v", name, amount)

		}

	case "transfer":
		if len(os.Args) != 5 {
			fmt.Println("Format: transfer <from> <to> <amount>")
		}

		from := os.Args[2]
		to := os.Args[3]
		parseAmount, err := strconv.ParseFloat(os.Args[4], 32)
		if err != nil {
			panic(err)
		}
		amount := float32(parseAmount)

		// LOGIC SERVICE TRANSFER
		err = services.Transfer(from, to, amount)
		if err != nil {
			fmt.Println("ERROR:", err)
		} else {
			fmt.Printf("Transfer amount from %v's to %v's account by %v", from, to, amount)

		}

	case "add_deposit":
		if len(os.Args) != 4 {
			fmt.Println("Format : add_deposit <name> <amount>")
		}

		name := os.Args[2]
		parseAmount, err := strconv.ParseFloat(os.Args[3], 32)
		if err != nil {
			panic(err)
		}
		amount := float32(parseAmount)

		// LOGIC SERVICE ADD DEPOSIT
		err = services.Deposit(name, amount)
		if err != nil {
			fmt.Println("ERROR:", err)

		} else {
			fmt.Printf("add deposit %v's with amount %v success", name, amount)

		}

	case "accure_interest":
		if len(os.Args) != 2 {
			fmt.Println("Format : accrue_interest")
		}

		// LOGIC ACCURE INTEREST
		services.AccureInterest()

	default:
		fmt.Println("Unknow Service", service)
		list()
	}

}

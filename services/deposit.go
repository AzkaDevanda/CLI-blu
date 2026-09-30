package services

import "github.com/AzkaDevanda/CLI_blu/repository/deposit"

func Deposit(name string, amount float32) error {
	return deposit.Save(name, amount)
}

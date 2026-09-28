package services

import (
	"github.com/AzkaDevanda/CLI_blu/repository/account"
)

func Save(name string, amount float32) error {
	return account.Save(name, amount)
}

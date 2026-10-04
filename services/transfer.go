package services

import (
	"errors"

	"github.com/AzkaDevanda/CLI_blu/repository/account"
	"github.com/AzkaDevanda/CLI_blu/repository/payment"
)

func Transfer(from string, to string, amount float32) error {

	accounts, err := validateTransfer(from, to, amount)
	if err != nil {
		return err
	}

	accounts[from] -= amount
	accounts[to] += amount

	// UPDATE ACCOUNT BALANCE

	err = account.UpdateAll(accounts)
	if err != nil {
		return err
	}

	return payment.Save(from, to, amount)
}

func validateTransfer(from string, to string, amount float32) (map[string]float32, error) {
	account, err := account.FindAll()

	if err != nil {
		return nil, err
	}

	if from == to {
		return nil, errors.New("Tidak bisa transfer ke akun sendiri")
	}

	fromBalance, fromExist := account[from]
	_, toExist := account[to]

	if !fromExist || !toExist {
		return nil, errors.New("Akun pengirim atau penerima tidak ditemukan")
	}

	if fromBalance < amount {
		return nil, errors.New("Saldo tidak cukup")
	}

	return account, nil

}

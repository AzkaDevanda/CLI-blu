package account

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const filebae = "filebase/account.txt"

func Save(name string, amount float32) error {

	account, err := FindAll()

	if err != nil {
		return err
	}

	if _, exist := account[name]; exist {
		return errors.New("Account sudah ada")
	}

	row := fmt.Sprintf("%s,%.2f\n", name, amount)

	file, err := os.OpenFile(filebae, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(row)
	if err != nil {
		return err
	}

	return nil
}

func DeleteByName(name string) error {
	account, err := FindAll()

	if err != nil {
		return err
	}

	if _, exist := account[name]; exist {
		delete(account, name)
	} else {
		return errors.New("account tidak ada")
	}

	return nil
}

func FindByName(name string) (float32, error) {
	accounts, err := FindAll()
	if err != nil {
		return 0, err
	}

	balance, exists := accounts[name]
	if !exists {
		return 0, fmt.Errorf("account '%s' tidak ditemukan", name)
	}

	return balance, nil

}

func FindAll() (map[string]float32, error) {
	accounts := make(map[string]float32)

	file, err := os.Open(filebae)

	if err != nil {
		if os.IsNotExist(err) {
			return accounts, nil
		}
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) == 2 {
			name := strings.TrimSpace(parts[0])
			val, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 32)
			if err == nil {
				accounts[name] = float32(val)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return accounts, nil

}

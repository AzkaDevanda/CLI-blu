package account

import (
	"fmt"
	"os"
)

const filebae = "filebase/account.txt"

func Add_account(name string, amount float32) error {
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

package services

import (
	"fmt"
	"time"

	"github.com/AzkaDevanda/CLI_blu/repository/deposit"
)

func Deposit(name string, amount float32) error {
	return deposit.Save(name, amount)
}

func AccureInterest() error {
	deposits, err := deposit.FindAll()

	if err != nil {
		return err
	}

	hasAccured := false

	for i := range deposits {
		depositTime, err := time.ParseInLocation("2006-01-02 15:04:05", deposits[i].CreatedAt, time.Local)
		if err != nil {
			return nil
		}

		minutes := int(time.Since(depositTime).Minutes())
		if minutes >= 1 {
			interest := deposits[i].Amount * 0.001 * float32(minutes)

			deposits[i].Amount += interest

			hasAccured = true

			fmt.Printf("Bunga untuk %s bertambah: %.2f durasi %d menit", deposits[i].Name, interest, int32(minutes))
			return deposit.UpdateAll(deposits, interest)

		}
	}

	if !hasAccured {
		fmt.Println("Belum ada deposit yang mencapai 1 menit")
	}
	return nil
}

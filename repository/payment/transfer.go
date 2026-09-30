package payment

import (
	"fmt"
	"os"
	"time"
)

// const accountFilebasse = "filebase/account.txt"
const paymentFilebase = "filebase/payment.txt"

func Save(from string, to string, amount float32) error {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	row := fmt.Sprintf("%s,%s,%.2f,%s\n", to, from, amount, currentTime)

	file, err := os.OpenFile(paymentFilebase, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

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

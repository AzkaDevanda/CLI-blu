package deposit

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Deposit struct {
	Name      string
	Amount    float32
	CreatedAt string
}

const filebase = "filebase/deposit.txt"

func Save(name string, amount float32) error {

	deposits, err := FindAll()

	if err != nil {
		return err
	}

	for _, d := range deposits {
		if d.Name == name {
			return errors.New("Account sudah ada")
		}
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	row := fmt.Sprintf("%s,%.2f,%s\n", name, amount, currentTime)

	file, err := os.OpenFile(filebase, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

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

func UpdateAll(deposit []Deposit, amount float32) error {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	file, err := os.Create(filebase)
	if err != nil {
		return err
	}

	defer file.Close()

	for _, d := range deposit {
		row := fmt.Sprintf("%s,%.2f,%s\n", d.Name, d.Amount, currentTime)
		_, err := file.WriteString(row)
		if err != nil {
			return err
		}
	}
	return nil
}

func FindAll() ([]Deposit, error) {
	var deposits []Deposit

	file, err := os.Open(filebase)

	if err != nil {
		if os.IsNotExist(err) {
			return deposits, nil
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
		if len(parts) == 3 {
			name := strings.TrimSpace(parts[0])
			val, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 32)
			if err != nil {
				continue
			}
			createdAt := strings.TrimSpace(parts[2])
			deposits = append(deposits, Deposit{
				Name:      name,
				Amount:    float32(val),
				CreatedAt: createdAt,
			})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return deposits, nil

}

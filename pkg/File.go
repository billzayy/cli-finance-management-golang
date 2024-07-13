package pkg

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"os"
	"strings"
)

func CreateFile(filename string, item []Financial) error {
	data, err := json.Marshal(item)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644) // Create File with this os.WriteFile()
}

func LoadFile(filename string, item *[]Financial) error {
	file, err := os.ReadFile(filename)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, item) //? Convert data from file into item *[]string

	if err != nil {
		return err
	}

	return nil
}

func GetInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil //? Read arguments of flag and change into string
	}

	scanner := bufio.NewScanner(r) //? Scan the input from Client
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	if len(scanner.Text()) == 0 {
		return "", errors.New("empty is not allowed")
	}

	return scanner.Text(), nil
}

package easyenv

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Errors returned
var (
	ErrBadFormat = errors.New("Incorrect format of environment variables")
)

// GetEnv returns the map filled with env vars from the .env file
func GetEnv(path string) (map[string]string, error) {
	wd, _ := os.Getwd()

	file, err := os.Open(wd + "/" + path)
	check(err)
	defer file.Close()

	variables := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		varSlice := strings.FieldsFunc(scanner.Text(), func(c rune) bool {
			if c == '=' || c == ' ' {
				return true
			}
			return false
		})

		if len(varSlice) != 2 {
			return nil, ErrBadFormat
		}

		variables[varSlice[0]] = varSlice[1]
	}

	// checks if Scan() returned false because of EOF
	check(scanner.Err())

	return variables, nil
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	env "github.com/Netflix/go-env"
)

type environmentStruct struct {
	Home string `env:"HOME"`

	Extras env.EnvSet // not explicitly invoked environment variables go there
}

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
func GetEnv() (map[string]string, error) {
	file, err := os.Open("./.env")
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

		fmt.Println(varSlice)
	}

	// checks if Scan() returned false because of EOF
	check(scanner.Err())

	return variables, nil
}

func main() {
	var environment environmentStruct

	extras, err := env.UnmarshalFromEnviron(&environment)

	if err != nil {
		log.Fatal(err)
	}

	environment.Extras = extras

	fmt.Println("Home =", environment.Home)

	fmt.Println("Number of extra environment variables:", len(environment.Extras))

	_, err = GetEnv()

	check(err)
}

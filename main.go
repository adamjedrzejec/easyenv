package main

import (
	"fmt"
	"log"
	"os"

	env "github.com/Netflix/go-env"
	easyenv "github.com/adamjedrzejec/go-env-vars/EasyEnv"
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

func main() {
	var environment environmentStruct

	extras, err := env.UnmarshalFromEnviron(&environment)

	if err != nil {
		log.Fatal(err)
	}

	environment.Extras = extras

	fmt.Println("Home =", environment.Home)

	fmt.Println("Number of extra environment variables:", len(environment.Extras))

	wd, _ := os.Getwd()

	envVars, err := easyenv.GetEnv(wd + "/.env")
	check(err)

	for key, val := range envVars {
		fmt.Printf("%s -> %s\n", key, val)
	}
}

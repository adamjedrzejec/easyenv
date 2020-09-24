package main

import (
	"fmt"
	"log"

	env "github.com/Netflix/go-env"
)

type environmentStruct struct {
	Home string `env:"HOME"`

	Extras env.EnvSet // not explicitly invoked environment variables go there
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
	for key, val := range environment.Extras {
		fmt.Printf("%s=%s\n", key, val)
	}
}

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	secret, found := os.LookupEnv("GHAS_ACCEPTANCE_TEST_ONE")
	if !found || secret == "" {
		log.Fatalln("Failed to find GHAS_ACCEPTANCE_TEST_ONE value")
	}

	if secret != "I am secret one!" {
		log.Fatalln("Received GHAS_ACCEPTANCE_TEST_ONE value but did not match expected value")
	} else {
		fmt.Printf("The value of GHAS_ACCEPTANCE_TEST_ONE matched the expected value! Yay!")
	}
}

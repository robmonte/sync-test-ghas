package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	testType, exists := os.LookupEnv("TYPE")
	if !exists {
		log.Fatalf("missing required \"TYPE\" field to determine test type")
	}

	fmt.Printf("Running %q test type\n", testType)

	switch testType {
	case "LOAD_TEST":
		loadTypeTestCheck()

	case "UPDATE_TEST":
		updateTypeTestCheck()

	default:
		log.Fatalf("received unsupported test type")
	}
}

func loadTypeTestCheck() {
	envVar1 := "STOREGHAS_TEST_LOAD_KEY_1_BAD_SYMBOLS________"
	secret, found := os.LookupEnv(envVar1)
	if !found || secret == "" {
		log.Fatalf("Failed to find %q value\n", envVar1)
	}

	if secret != "I am secret one!" {
		log.Fatalf("received %q value but did not match expected value\n", envVar1)
	} else {
		fmt.Printf("the value of %q matched the expected value! Yay!\n", envVar1)
	}

	envVar2 := "STOREGHAS_TEST_LOAD_KEY_2"
	secret, found = os.LookupEnv(envVar2)
	if !found || secret == "" {
		log.Fatalf("failed to find %q value\n", envVar2)
	}

	if secret != "I am secret two!" {
		log.Fatalf("received %q value but did not match expected value\n", envVar2)
	} else {
		fmt.Printf("the value of %q matched the expected value! Yay!\n", envVar2)
	}

	envVar3 := "STOREGHAS_TEST_LOAD_KEY_3"
	secret, found = os.LookupEnv(envVar3)
	if !found || secret == "" {
		log.Fatalf("failed to find %q value\n", envVar3)
	}

	if secret != "1234" {
		log.Fatalf("received %q value but did not match expected value\n", envVar3)
	} else {
		fmt.Printf("the value of %q matched the expected value! Yay!\n", envVar3)
	}

	envVar4 := "STOREGHAS_TEST_LOAD_KEY_4"
	secret, found = os.LookupEnv(envVar4)
	if !found || secret == "" {
		log.Fatalf("failed to find %q value\n", envVar4)
	}

	for i := 0; i < len(secret); i++ {
		fmt.Printf("%v ", secret[i])
	}
	fmt.Println()

	if secret != "{\"0\":\"my\",\"1\":\"map\",\"2\":\"value\"}" {
		log.Fatalf("received %q value but did not match expected value\n", envVar4)
	} else {
		fmt.Printf("the value of %q matched the expected value! Yay!\n", envVar4)
	}
}

func updateTypeTestCheck() {
	envVar1 := "STOREGHAS_TEST_UPDATE_KEY_1"
	secret, found := os.LookupEnv(envVar1)
	if !found || secret == "" {
		log.Fatalf("Failed to find %q value\n", envVar1)
	}

	if secret != "StoreGHAS-Test-Update-Value-1-Changed" {
		log.Fatalf("received %q value but did not match expected value\n", envVar1)
	} else {
		fmt.Printf("the value of %q matched the expected value! Yay!\n", envVar1)
	}

	envVar2 := "STOREGHAS_TEST_UPDATE_KEY_2"
	secret, found = os.LookupEnv(envVar2)
	if !found || secret == "" {
		log.Fatalf("Failed to find %q value\n", envVar2)
	}

	if secret != "{0:\"StoreGHAS-Test-Update-Value-2-Changed\"}" {
		log.Fatalf("received %q value but did not match expected value\n", envVar2)
	} else {
		fmt.Printf("the value of %q matched the expected value! Yay!\n", envVar2)
	}
}

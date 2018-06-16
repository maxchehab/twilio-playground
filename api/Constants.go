package main

import (
	"bufio"
	"log"
	"os"
)

// ConstantType represents the constants used in the application
type ConstantType struct {
	AccountSid string
}

// Constants is a singleton representing the constants used in the application
var Constants = ConstantType{
	readLine("../secrets/AccountSid"),
}

func readLine(path string) (line string) {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	_ = scanner.Scan()
	line = scanner.Text()
	err = scanner.Err()
	if err != nil {
		log.Panic(err)
	}
	return
}

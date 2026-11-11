package main

import (
	"errors"
	"fmt"
	"net"
	"net/mail"
	"strings"
)

var invalidAddress = errors.New("Invalid Email address")

// var invalidInput = errors.New("Invalid input")
var domain string

func validate(s string) ([]string, error) {
	s = strings.TrimSpace(s)
	parsedDomain, err := mail.ParseAddress(s)
	if err != nil {
		return []string{}, invalidAddress
	}

	var ind int
	if ind = strings.LastIndex(parsedDomain.Address, "@"); ind != 1 {
		fmt.Printf("Indexed men,, %d", ind)
		return []string{}, invalidAddress
	}
	splitList := strings.Split(parsedDomain.Address, "@")
	if len(splitList[0]) > 64 {
		return []string{}, invalidAddress
	}
	if strings.Contains(splitList[0], ".") || !strings.Contains(splitList[1], ".") {

		return []string{}, invalidAddress
	}
	return splitList, nil
}
func mx(domain string) error {
	domain = strings.TrimSpace(domain)

	result, err := net.LookupMX(domain)
	if err != nil {
		return invalidAddress
	}
	for index, data := range result {
		fmt.Printf("%v -- %v -- %v\n", index+1, data.Pref, data.Host)
	}
	return nil
}

func userInput() error {
	for {
		fmt.Print("Enter domain: ")
		_, err := fmt.Scan(&domain)
		if err != nil {
			continue
		} else {
			break
		}
	}
	return nil
}

func main() {
	if err := userInput(); err != nil {
		fmt.Printf("Had an error: %v", err)
		return
	}
	passedList, err := validate(domain)
	if err != nil {
		fmt.Printf("Had an error: %v", err)
		return
	}
	if err = mx(passedList[1]); err != nil {
		fmt.Printf("Had an error: %v", err)
		return
	}

}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"os"
	"strings"
)

var invalidAddress = errors.New("invalid email address")

// var invalidInput = errors.New("Invalid input")
// var domain string

func validate(s string) ([]string, error) {
	s = strings.TrimSpace(s)
	parsedAddress, err := mail.ParseAddress(s)
	if err != nil {
		return []string{}, invalidAddress
	}

	ind := strings.LastIndex(parsedAddress.Address, "@")
	splitList := []string{
		parsedAddress.Address[:ind],
		parsedAddress.Address[ind+1:],
	}
	// fmt.Println(splitList[1])

	if !strings.Contains(splitList[1], ".") {

		return []string{}, invalidAddress
	}
	return splitList, nil
}
func mx(domain string) ([]*net.MX, error) {
	domain = strings.TrimSpace(domain)
	return net.LookupMX(domain)
}
func checkAArecords(s string) (addrs []string, err error) {
	return net.LookupHost(s)
}

func userInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter domain: ")
	domain, err := reader.ReadString('\n')
	domain = strings.TrimSpace(domain)

	if err != nil {
		return "", err
	}

	return domain, nil
}

func main() {
	email, err := userInput()
	// fmt.Printf("This is the email %v %v", email, err)
	if err != nil {
		fmt.Printf("Had an error: %v", err)
		return
	}
	passedList, err := validate(email)
	if err != nil {
		fmt.Printf("Had an error: %v", err)
		return
	}
	result, err := mx(passedList[1])
	if err != nil {
		fmt.Printf("Had an error: %v", err)
		aaResults, err := checkAArecords(passedList[1])
		if err != nil {
			fmt.Printf("Had an error: %v", err)
			return
		}
		for index, data := range aaResults {
			fmt.Printf("%d -- %v", index, data)
		}

	} else {
		for i, data := range result {
			var summary string
			if data.Host == "." {
				summary = "No host names"
				break
			}
			summary += fmt.Sprintf("%v -- %d -- %s", i, data.Pref, data.Host)
			fmt.Println(summary)
		}
	}
}

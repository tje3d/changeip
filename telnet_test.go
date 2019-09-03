package main

import (
	"log"
	"strings"
	"testing"

	"changeip/telnet"
)

func TestTelnetConnect(t *testing.T) {
	_, err := telnet.New()

	if err != nil {
		log.Fatalln(err)
	}
}

func TestTelnetChangeIP(t *testing.T) {
	currentIP, err := getMyIP(-1)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("CurrentIP: " + currentIP)
	client, err := telnet.New()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Changing IP")
	client.ChangeIP()

	secondIP, err := getMyIP(-1)

	if err != nil {
		log.Fatalln(err)
	}

	if strings.Compare(secondIP, currentIP) == 0 {
		log.Fatalln("IP Didnt Changed")
	}

	log.Println("SecondIP: " + secondIP)
}

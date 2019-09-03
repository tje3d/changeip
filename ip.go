package main

import (
	"changeip/telnet"
	"log"
	"strings"
)

func changeIPTo(tcClient *telnet.MyTelnet, currentIP string, changeTO string) string {
	if strings.HasPrefix(currentIP, changeTO) {
		log.Fatalln("IP Succesfully changed to " + currentIP)
		return currentIP
	}

	tcClient.ChangeIP()

	secondIP, _ := getMyIP(-1)
	spin.Stop()
	log.Println(secondIP)
	spin.Start()

	if !strings.HasPrefix(secondIP, changeTO) {
		return changeIPTo(tcClient, secondIP, changeTO)
	}

	return secondIP
}

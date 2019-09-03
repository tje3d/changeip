package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/manifoldco/promptui"
)

func initializeConfig() {
	var hasDotEnv bool

	godotenv.Load()

	if len(os.Getenv("IP")) == 0 {
		promptIP := promptui.Prompt{
			Label:   "Please enter modem ip",
			Default: "192.168.1.1",
		}
		ModemIP, _ := promptIP.Run()
		os.Setenv("IP", ModemIP)
		hasDotEnv = false
	}

	if len(os.Getenv("PASSWORD")) == 0 {
		promptIP := promptui.Prompt{
			Label: "Please enter modem password",
			Mask:  '*',
		}
		ModemIP, _ := promptIP.Run()
		os.Setenv("PASSWORD", ModemIP)
		hasDotEnv = false
	}

	if !hasDotEnv {
		godotenv.Write(map[string]string{
			"IP":       os.Getenv("IP"),
			"PASSWORD": os.Getenv("PASSWORD"),
		}, ".env")
	}
}

func askForIPPrefix() (string, error) {
	prompt := promptui.Select{
		Label: "Select IP Prefix",
		Items: []string{"2.", "5.", "85."},
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}

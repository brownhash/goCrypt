package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os/exec"
	"strings"
)

// throws error if vim is not installed on the system
func CheckVim() error {
	out, err := exec.Command("command", "-v", "vim").Output()
	if err != nil {
		return err
	}
	check := string(out)
	if strings.Contains(check, "/usr/bin/vim") {
		return nil
	} else {
		return errors.New("vim is not installed, please install it first ")
	}
}

// writes data to the fileName provided
func WriteToFile(fileName string, data []byte) error {
	err := ioutil.WriteFile(fileName, data, 0644)
	return err
}

// reads data from the fileName provided
func ReadFromFile(fileName string) ([]byte, error) {
	data, err := ioutil.ReadFile(fileName)
	return data, err
}

// parse the arguments supplied
func ArgParse(args []string) ([]string, error) {
	if len(args) < 2 || len(args) > 5 { // throws an error if no user arguments are received
		return []string{"", "", ""}, errors.New("no arguments were given ")
	}
	option := args[1]

	// if options are recorded then checks for supporting options
	// throws error if required number of support arguments are not received
	if option == "-w" || option == "-r" {
		if len(args) == 1 { // throws an error when -w or -r is used but no filename is supplied
			return []string{"", "", ""}, errors.New("filename not provided ")
		}
		filename := args[2]
		if len(args) == 3{
			return []string{option, filename, ""}, nil
		} else {
			keyOpt := args[3]
			if keyOpt == "-k" { // throws an error when -k is used but no key is supplied
				if len(args) == 4 {
					return []string{option, filename, ""}, errors.New("no key supplied ")
				} else {
					key := args[4]
					return []string{option, filename, key}, nil
				}
			} else if keyOpt == "-f" {
				if len(args) == 4 {
					return []string{option, filename, ""}, errors.New("no file name supplied ")
				} else {
					confFile := args[4]
					keyConfig, err := ReadFromFile(confFile)

					if err != nil {
						return []string{"", "", ""}, errors.New(err.Error())
					}

					keyModal := make(map[string]string, 0)
					err = json.Unmarshal(keyConfig, &keyModal)

					if err != nil {
						return []string{"", "", ""}, errors.New(err.Error())
					}

					return []string{option, filename, keyModal["key"]}, nil
				}
			} else { // throws error if supplied options are not recognised
				return []string{"", "", ""}, errors.New("invalid option, use -h to see available options ")
			}
		}

	} else if option == "-h" {
		return []string{option, "", ""}, nil
	} else { // throws error if supplied options are not recognised
		return []string{"", "", ""}, errors.New("invalid option, use -h to see available options ")
	}
}

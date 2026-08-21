package main

import (
	"fetch/distros"
	"bufio"
	"fmt"
	"strings"
	"os"
)

func detectDistro() (string, error) {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return "", err
 	}

  	defer file.Close()

   	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	line := scanner.Text()

     	if strings.HasPrefix(line, "ID=") {
      		id := strings.TrimPrefix(line, "ID=")
        	id = strings.Trim(id, "\"")
         	return id, nil
      	}
    }

    if err := scanner.Err(); err != nil {
    	return "", err
    }

    return "", fmt.Errorf("ID не найден в /etc/os-release")
}

func main() {
	distro, err := detectDistro()
	if err != nil {
		fmt.Printf("Ошибка при определении дистрибутива: %v\n", err)
		os.Exit(1)
	}

	switch distro {
		case "linuxmint":
			distros.MintFetch()
		case "debian":
			distros.DebianFetch()
		case "nixos":
			distros.NixFetch()
		case "artix":
			distros.ArtixFetch()
		case "android":
			distros.AndroidFetch()
		default:
			fmt.Printf("Дистрибутив %s не поддерживается dfetch\n", distro)
	}
}

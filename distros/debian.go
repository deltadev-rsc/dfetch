package distros

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const _RESET_COLOR = "\x1b[0m"
const _RED_COLOR = "\x1b[31m"
const _GREEN_COLOR = "\x1b[32m"
const _YELLOW_COLOR = "\x1b[33m"
const _BLUE_COLOR = "\x1b[34m"
const _MAGENTA_COLOR = "\x1b[35m"
const _CYAN_COLOR = "\x1b[36m"

func debGetUserName() (string, error) {
	cmd := exec.Command("whoami")
	output, err := cmd.Output()

	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении команды: %v", err)
	}

	username := string(output[:len(output) - 1])
	return username, nil
}

func debGetHostName() (string, error) {
	cmd := exec.Command("uname", "-n")
	output, err := cmd.Output()

	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении команды: %v", err)
	}

	hostname := strings.TrimSpace(string(output))
	return hostname, nil
}

func debGetOperatingSystemName() (string, error) {
	cmd := exec.Command("uname", "-o")
	output, err := cmd.Output()

	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении команды: %v", err)
	}

	osname := strings.TrimSpace(string(output))
	return osname, nil
}

func debGetKernel() (string, error) {
	cmd := exec.Command("uname", "-sr")
	output, err := cmd.Output()

	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении команды: %v", err)
	}

	kernel := strings.TrimSpace(string(output))
	return kernel, nil
}

func debGetShell() (string) {
	shell := "SHELL"

	if val := os.Getenv(shell); val != "" {
		return val
	}

	if _, err := os.Stat("/usr/bin/fish"); err == nil {
		return "fish"
	}

	if _, err := os.Stat("/usr/bin/bash"); err == nil {
		return "bash"
	}

	if _, err := os.Stat("/usr/bin/sh"); err == nil {
		return "sh"
	}

	return "Unknown"
}

func debGetDesktopEnvironment() (string) {
	vars := []string { "XDG_CURRENT_DESKTOP", "DESKTOP_SESSION", "GDMSESSION" }

	for _, v := range vars {
		if val := os.Getenv(v); val != "" {
			return val
		}
	}

	return "Unknown"
}

func debGetInit() (string) {
	cmd := exec.Command("ps", "-p", "1", "-o", "comm=")
	output, err := cmd.Output()

	if err != nil {
		return "Unknown"
	}

	init := strings.TrimSpace(string(output))

	if init == "init" {
		return debRefineInit()
	}

	return init
}

func debRefineInit() (string) {
	if _, err := os.Stat("/run/openrc"); err == nil {
		return "OpenRC"
	}

	if _, err := os.Stat("/var/service"); err == nil {
		return "Runit"
	}

	if _, err := os.Stat("/etc/init.d"); err == nil {
		return "SysVinit"
	}

	return "init (Unknown)"
}

func DebianFetch() {
	username, err := debGetUserName()
	hostname, err := debGetHostName()
	osname, err := debGetOperatingSystemName()
	kernel, err := debGetKernel()
	shell := debGetShell()
	de_wm := debGetDesktopEnvironment()
	init := debGetInit()
	if err != nil {
		fmt.Println("Ошибащка: ", err)
		return
	}

	fmt.Println("")
	fmt.Println(_RED_COLOR, " ", _RESET_COLOR)
	fmt.Println(_RED_COLOR, "   ▟██████████▙   ", _RESET_COLOR, _RED_COLOR,  " User:", _RESET_COLOR, username, _RESET_COLOR)
	fmt.Println(_RED_COLOR, "  ▟█▛        ▜█▙  ", _RESET_COLOR, _MAGENTA_COLOR,  "───────────────────",  _RESET_COLOR)
	fmt.Println(_RED_COLOR, " ██▛   ▟██    ▜█▙ ", _RESET_COLOR, _RED_COLOR,  "󰌢 Host:", _RESET_COLOR, hostname, _RESET_COLOR)
	fmt.Println(_RED_COLOR, "██    ▟█▛      ██", _RESET_COLOR, _RED_COLOR,  " OS:",  _RESET_COLOR, osname,  _RESET_COLOR)
	fmt.Println(_RED_COLOR, "██   ▟█▛      ▟█▛  ", _RESET_COLOR, _RED_COLOR,  " Kernel:", _RESET_COLOR, kernel,  _RESET_COLOR)
	fmt.Println(_RED_COLOR, "▜█▙  ▜█████████▛          ", _RESET_COLOR, _RED_COLOR,  " Shell:", _RESET_COLOR, shell,   _RESET_COLOR)
	fmt.Println(_RED_COLOR, " ▜█▙            ", _RESET_COLOR, _RED_COLOR,  " DE/WM:", _RESET_COLOR, de_wm,   _RESET_COLOR)
	fmt.Println(_RED_COLOR, "  ▜█▙           ", _RESET_COLOR, _RED_COLOR,  " Init system:", _RESET_COLOR, init, _RESET_COLOR)
	fmt.Println(_RED_COLOR, "   ▜█▙          ", _RESET_COLOR, _RED_COLOR,  "", _RESET_COLOR)
	fmt.Println(_RED_COLOR, "    ▜█▙         ", _RESET_COLOR, _RED_COLOR,  "█",
			_GREEN_COLOR, "█", _YELLOW_COLOR, "█",
			_BLUE_COLOR,  "█", _MAGENTA_COLOR, "█", _CYAN_COLOR, "█", _RESET_COLOR)
	fmt.Println("")
}

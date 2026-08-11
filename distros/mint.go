package distros

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const _RESET_COL = "\x1b[0m"
const _RED_COL = "\x1b[31m"
const _GREEN_COL = "\x1b[32m"
const _YELLOW_COL = "\x1b[33m"
const _BLUE_COL = "\x1b[34m"
const _MAGENTA_COL = "\x1b[35m"
const _CYAN_COL = "\x1b[36m"

func mintGetUserName() (string, error) {
	cmd := exec.Command("whoami")
	output, err := cmd.Output()

	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении команды: %v", err)
	}

	username := string(output[:len(output) - 1])
	return username, nil
}

func mintGetHostName() (string, error) {
	cmd := exec.Command("uname", "-n")
	output, err := cmd.Output()

	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении команды: %v", err)
	}

	hostname := strings.TrimSpace(string(output))
	return hostname, nil
}

func mintGetOperatingSystemName() (string, error) {
	cmd := exec.Command("uname", "-o")
	output, err := cmd.Output()

	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении команды: %v", err)
	}

	osname := strings.TrimSpace(string(output))
	return osname, nil
}

func mintGetKernel() (string, error) {
	cmd := exec.Command("uname", "-sr")
	output, err := cmd.Output()

	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении команды: %v", err)
	}

	kernel := strings.TrimSpace(string(output))
	return kernel, nil
}

func mintGetShell() (string) {
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

func mintGetDesktopEnvironment() (string) {
	vars := []string { "XDG_CURRENT_DESKTOP", "DESKTOP_SESSION", "GDMSESSION" }

	for _, v := range vars {
		if val := os.Getenv(v); val != "" {
			return val
		}
	}

	return "Unknown"
}

func mintGetInit() (string) {
	cmd := exec.Command("ps", "-p", "1", "-o", "comm=")
	output, err := cmd.Output()

	if err != nil {
		return "Unknown"
	}

	init := strings.TrimSpace(string(output))

	if init == "init" {
		return mintRefineInit()
	}

	return init
}

func mintRefineInit() (string) {
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

func MintFetch() {
	username, err := mintGetUserName()
	hostname, err := mintGetHostName()
	osname, err := mintGetOperatingSystemName()
	kernel, err := mintGetKernel()
	shell := mintGetShell()
	de_wm := mintGetDesktopEnvironment()
	init := mintGetInit()
	if err != nil {
		fmt.Println("Ошибащка: ", err)
		return
	}

	fmt.Println("")
	fmt.Println(_GREEN_COL, " ", _RESET_COL)
	fmt.Println(_GREEN_COL, "██  ▟██████▙ ▟███████▙ ", _RESET_COL, _GREEN_COL,  " User:", username,    RESET_COL)
	fmt.Println(_GREEN_COL, "██  ██    █████     ██ ", _RESET_COL, MAGENTA_COL,  "───────────────────",  RESET_COL)
	fmt.Println(_GREEN_COL, "██  ██     ███      ██ ", _RESET_COL, _GREEN_COL,  "󰌢 Host:",    hostname, RESET_COL)
	fmt.Println(_GREEN_COL, "██  ██     ███      ██ ", _RESET_COL, _GREEN_COL,  " OS:",      osname,   RESET_COL)
	fmt.Println(_GREEN_COL, "██  ██     ███      ██ ", _RESET_COL, _GREEN_COL,  " Kernel:",  kernel,   RESET_COL)
	fmt.Println(_GREEN_COL, "██  ██     ███      ██ ", _RESET_COL, _GREEN_COL,  " Shell:",   shell,    RESET_COL)
	fmt.Println(_GREEN_COL, "██  ██     ███      ██ ", _RESET_COL, _GREEN_COL,  " DE/WM:",   de_wm,    RESET_COL)
	fmt.Println(_GREEN_COL, "██  ██     ███      ██ ", _RESET_COL, _GREEN_COL,  " Init system:", init, RESET_COL)
	fmt.Println(_GREEN_COL, "▜█                  █▛  ", _RESET_COL, _GREEN_COL,  "", _RESET_COL)
	fmt.Println(_GREEN_COL, " ▜██████████████████▛   ", _RESET_COL, _RED_COL,  "█", _GREEN_COL, "█", _YELLOW_COL, "█", _BLUE_COL, "█", _MAGENTA_COL, "█", _CYAN_COL, "█", _RESET_COL)
	fmt.Println("")
}

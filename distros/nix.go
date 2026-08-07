package distros

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const RESET_COL = "\x1b[0m"
const RED_COL = "\x1b[31m"
const GREEN_COL = "\x1b[32m"
const YELLOW_COL = "\x1b[33m"
const BLUE_COL = "\x1b[34m"
const MAGENTA_COL = "\x1b[35m"
const CYAN_COL = "\x1b[36m"

func nixGetUserName() (string, error) {
	cmd := exec.Command("whoami")
	output, err := cmd.Output()

	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении команды: %v", err)
	}

	username := string(output[:len(output) - 1])
	return username, nil
}

func nixGetHostName() (string, error) {
	cmd := exec.Command("uname", "-n")
	output, err := cmd.Output()

	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении команды: %v", err)
	}

	hostname := strings.TrimSpace(string(output))
	return hostname, nil
}

func nixGetOperatingSystemName() (string, error) {
	cmd := exec.Command("uname", "-o")
	output, err := cmd.Output()

	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении команды: %v", err)
	}

	osname := strings.TrimSpace(string(output))
	return osname, nil
}

func nixGetKernel() (string, error) {
	cmd := exec.Command("uname", "-sr")
	output, err := cmd.Output()

	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении команды: %v", err)
	}

	kernel := strings.TrimSpace(string(output))
	return kernel, nil
}

func nixGetShell() (string) {
	shell := "SHELL"

	if val := os.Getenv(shell); val != "" {
		return val
	}

	if _, err := os.Stat("/run/current-system/sw/bin/fish"); err == nil {
		return "fish"
	}

	if _, err := os.Stat("/run/current-system/sw/bin/bash"); err == nil {
		return "bash"
	}

	if _, err := os.Stat("/run/current-system/sw/bin/sh"); err == nil {
		return "sh"
	}

	return "Unknown"
}

func nixGetDesktopEnvironment() (string) {
	vars := []string { "XDG_CURRENT_DESKTOP", "DESKTOP_SESSION", "GDMSESSION" }

	for _, v := range vars {
		if val := os.Getenv(v); val != "" {
			return val
		}
	}

	return "Unknown"
}

func nixGetInit() (string) {
	cmd := exec.Command("ps", "-p", "1", "-o", "comm=")
	output, err := cmd.Output()

	if err != nil {
		return "Unknown"
	}

	init := strings.TrimSpace(string(output))

	if init == "init" {
		return nixRefineInit()
	}

	return init
}

func nixRefineInit() (string) {
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

func NixFetch() {
	username, err := nixGetUserName()
	hostname, err := nixGetHostName()
	osname, err := nixGetOperatingSystemName()
	kernel, err := nixGetKernel()
	shell := nixGetShell()
	de_wm := nixGetDesktopEnvironment()
	init := nixGetInit()
	if err != nil {
		fmt.Println("Ошибащка: ", err)
		return
	}

	fmt.Println("")
	fmt.Println(CYAN_COL, "          ▜█▙   ▜█▙  ▟█▀             ", RESET_COL)
	fmt.Println(CYAN_COL, "           ▜█▙   ██ ▟█▀              ", RESET_COL, CYAN_COL,  " User:", username,    RESET_COL)
	fmt.Println(CYAN_COL, "       ███████████▀ ██    ▟█▀        ", RESET_COL, MAGENTA_COL,  "───────────────────",  RESET_COL)
	fmt.Println(CYAN_COL, "           ▟█▀      ▜█▙ ▟██▀         ", RESET_COL, CYAN_COL,  "󰌢 Host:",    hostname, RESET_COL)
	fmt.Println(CYAN_COL, "          ▟█▀        ▜█▟██▀          ", RESET_COL, CYAN_COL,  " OS:",      osname,   RESET_COL)
	fmt.Println(CYAN_COL, "     ███████          ███████        ", RESET_COL, CYAN_COL,  " Kernel:",  kernel,   RESET_COL)
	fmt.Println(CYAN_COL, "       ▟██▜█▙        ▟██▀            ", RESET_COL, CYAN_COL,  " Shell:",   shell,    RESET_COL)
	fmt.Println(CYAN_COL, "      ▟█▀  ▜█▙      ▟█▀              ", RESET_COL, CYAN_COL,  " DE/WM:",   de_wm,    RESET_COL)
	fmt.Println(CYAN_COL, "            ▜██▙ ▜████████           ", RESET_COL, CYAN_COL,  " Init system:", init, RESET_COL)
	fmt.Println(CYAN_COL, "            ▟████▙   ▜█▙             ", RESET_COL, CYAN_COL,  "", RESET_COL)
	fmt.Println(CYAN_COL, "           ▟█▀  ▜█▙   ▜█▙            ", RESET_COL, RED_COL,  "█", GREEN_COL, "█", YELLOW_COL, "█", BLUE_COL, "█", MAGENTA_COL, "█", CYAN_COL, "█", RESET_COL)
	fmt.Println("")
}

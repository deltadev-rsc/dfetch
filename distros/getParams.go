package distros

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetUserName() (string, error) {
	cmd := exec.Command("whoami")
	output, err := cmd.Output()

	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении команды: %v", err)
	}

	username := string(output[:len(output) - 1])
	return username, nil
}

func GetHostName() (string, error) {
	cmd := exec.Command("uname", "-n")
	output, err := cmd.Output()

	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении команды: %v", err)
	}

	hostname := strings.TrimSpace(string(output))
	return hostname, nil
}

func GetOperatingSystemName() (string, error) {
	cmd := exec.Command("uname", "-o")
	output, err := cmd.Output()

	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении команды: %v", err)
	}

	osname := strings.TrimSpace(string(output))
	return osname, nil
}

func GetKernel() (string, error) {
	cmd := exec.Command("uname", "-sr")
	output, err := cmd.Output()

	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении команды: %v", err)
	}

	kernel := strings.TrimSpace(string(output))
	return kernel, nil
}

func GetShell() (string) {
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

func GetDesktopEnvironment() (string) {
	vars := []string { "XDG_CURRENT_DESKTOP", "DESKTOP_SESSION", "GDMSESSION" }

	for _, v := range vars {
		if val := os.Getenv(v); val != "" {
			return val
		}
	}

	return "Unknown"
}

func GetInit() (string) {
	cmd := exec.Command("ps", "-p", "1", "-o", "comm=")
	output, err := cmd.Output()

	if err != nil {
		return "Unknown"
	}

	init := strings.TrimSpace(string(output))

	if init == "init" {
		return RefineInit()
	}

	return init
}

func RefineInit() (string) {
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

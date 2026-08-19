package distros

import (
	"fmt"
)

func DebianFetch() {
	username, err := GetUserName()
	hostname, err := GetHostName()
	osname, err :=   GetOperatingSystemName()
	kernel, err :=   GetKernel()
	shell := GetShell()
	de_wm := GetDesktopEnvironment()
	init  := GetInit()
	if err != nil {
		fmt.Println("Ошибащка: ", err)
		return
	}

	fmt.Println("")
	fmt.Println(RED_COLOR, " ", RESET_COLOR)
	fmt.Println(RED_COLOR, "   ▟██████████▙    ", RESET_COLOR, RED_COLOR,  " User:", RESET_COLOR, username, RESET_COLOR)
	fmt.Println(RED_COLOR, "  ▟█▛        ▜█▙   ", RESET_COLOR, MAGENTA_COLOR,  "───────────────────",        RESET_COLOR)
	fmt.Println(RED_COLOR, " ██▛   ▟██    ▜█▙  ", RESET_COLOR, RED_COLOR,  "󰌢 Host:",   RESET_COLOR, hostname, RESET_COLOR)
	fmt.Println(RED_COLOR, "██    ▟█▛      ██  ", RESET_COLOR, RED_COLOR,  " OS:",     RESET_COLOR, osname,  RESET_COLOR)
	fmt.Println(RED_COLOR, "██   ▟█▛      ▟█▛  ", RESET_COLOR, RED_COLOR,  " Kernel:", RESET_COLOR, kernel,  RESET_COLOR)
	fmt.Println(RED_COLOR, "▜█▙  ▜█████████▛   ", RESET_COLOR, RED_COLOR,  " Shell:",  RESET_COLOR, shell,   RESET_COLOR)
	fmt.Println(RED_COLOR, " ▜█▙               ", RESET_COLOR, RED_COLOR,  " DE/WM:",  RESET_COLOR, de_wm,   RESET_COLOR)
	fmt.Println(RED_COLOR, "  ▜█▙              ", RESET_COLOR, RED_COLOR,  " Init system:", RESET_COLOR, init, RESET_COLOR)
	fmt.Println(RED_COLOR, "   ▜█▙             ", RESET_COLOR, RED_COLOR,  "", RESET_COLOR)
	fmt.Println(RED_COLOR, "    ▜█▙            ", RESET_COLOR, RED_COLOR,  "█",
			GREEN_COLOR, "█", YELLOW_COLOR, "█",
			BLUE_COLOR,  "█", MAGENTA_COLOR, "█", CYAN_COLOR, "█", RESET_COLOR)
	fmt.Println("")
}

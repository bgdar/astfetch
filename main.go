package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func displayHorizontal(lines []string, info []string) {
	maxLines := len(lines)
	if len(info) > maxLines {
		maxLines = len(info)
	}
	// jika maish ada logo atau info maka tampilkan jika tidak kosongkan ("")
	// tujuanya untuk menyamakan
	for i := 0; i < maxLines; i++ {
		logoPart := ""
		infoPart := ""

		// 		maxLen := 0
		// for _, l := range labels {
		// 	if len(l) > maxLen {
		// 		maxLen = len(l)
		// 	}
		// }
		if i < len(lines) {
			logoPart = lines[i]
		} else {
			logoPart = ""
		}
		if i < len(info) {
			infoPart = info[i]
		} else {
			infoPart = ""
		}
		if !strings.Contains(infoPart, "none") {
			// Format: logo di kiri, info di kanan
			fmt.Printf("%-19s   %s\n", logoPart, infoPart)
		}
	}

}
func displayVertical(lines []string, info []string) {
	// logo
	for _, line := range lines {
		if line != "none" {
			fmt.Println(line)
		}
	}
	fmt.Print("")
	// info
	for _, in := range info {
		fmt.Println(in)
	}
}

func main() {

	// simpan array string untuk di tampilkan
	info := []string{
		fmt.Sprintf("%s on %s", GetUser(), GetOSName()),
		fmt.Sprint("---------------------------"),
		fmt.Sprintf("kernel      : %s", GetKernel()),
		fmt.Sprintf("uptime      : %s", GetUptime()),
		fmt.Sprintf("shell       : %s", GetShell()),
		fmt.Sprintf("resolution  : %s", GetResolution()),
		fmt.Sprintf("window      : %s", GetWindows()),
		fmt.Sprintf("theme       : %s", GetTheme()),
		fmt.Sprintf("terminal    : %s", GetTerminal()),
		fmt.Sprintf("cpu         : %s",GetCpu()),
		fmt.Sprintf("gpu         : %s",GetGpu()),
		fmt.Sprintf("memory     : %s",GetMemory()),
	}
	// ambil sesui os yg jalan
	lines := MainLogoDistroOS[runtime.GOOS]

	// ambil argumets setelah nama file atau nama entri program jika di build
	var args = os.Args[1:]
	var arg string

	if len(args) > 0 {
		arg = args[0]
	}
	if arg == "d" {
		displayVertical(lines, info)
	} else {
		//DEFAULT
		displayHorizontal(lines, info)
	}

}

package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

/// tampilakn dalam bentuk lurus ke samping 
func displayHorizontal(lines []string, info []string) {



	maxLines := len(lines)
	maxInfos := len(info)

	// cari baris yang paling besar 
	if maxInfos > maxLines {
		maxLines = maxInfos
	}

	fmt.Print("")

	// lop untuk menampilkan 
	// jika maish ada logo atau info maka tampilkan jika tidak kosongkan ("")
	// tujuanya untuk menyamakan
	for i := 0; i < maxLines; i++ {
		logoPart := ""
		infoPart := ""

		// Tampilkan row logo hanya jika masih dalam range lines
		if i < len(lines){
				logoPart = lines[i]
			}

		// Tampilkan info jika masih ada
		if i < maxInfos {
			infoPart = info[i]
		}

		// Jika info tidak mengandung "none", cetak keduanya
		if !strings.Contains(infoPart, "none") {
			fmt.Printf("%-19s  %s\n", logoPart, infoPart)
		}
	}

}

/// tampilakn dalam bentuk lurus ke bawah 
func displayVertical(lines []string, info []string) {
	fmt.Print("")
	// logo
	for _, line := range lines {
		if line != "none" {
			fmt.Println(line)
		}
	}
	fmt.Print("")
	// info
	for _, in := range info {
		if !strings.Contains(in,"none"){
			fmt.Println(in)
		}
	}
}

func main() {

	// simpan array string untuk di tampilkan
	info := []string{
		fmt.Sprintf("\033[38;5;47;5;255m%s\033[0m on %s", GetUser(), GetOSName()),
		fmt.Sprint("---------------------------"),
		fmt.Sprintf("\033[38;5;40m \033[0m %-12s \033[38;5;38m %-20s  \033[0m", "Kernel", GetKernel()),
		fmt.Sprintf("\033[38;5;38m󱑎 \033[0m %-12s \033[38;5;45m %-20s  \033[0m", "Uptime", GetUptime()),
		fmt.Sprintf("\033[38;5;38m \033[0m %-12s \033[1;5;40m %-20s  \033[0m", "Shell", GetShell()),
		fmt.Sprintf("\033[38;5;56m \033[0m %-12s \033[38;5;277m %-20s  \033[0m", "Resolution", GetResolution()),
		fmt.Sprintf("\033[38;5;38m󰨡 \033[0m  %-12s \033[38;5;226m %-20s  \033[0m", "Window", GetWindows()),
		fmt.Sprintf("\033[38;5;38m \033[0m  %-12s \033[38;5;201m %-20s \033[0m", "Theme", GetTheme()),
		fmt.Sprintf("\033[1;245m \033[0m  %-12s \033[38;5;15m %-20s  \033[0m", "Terminal", GetTerminal()),
		fmt.Sprintf("\033[38;5;38m \033[0m  %-12s \033[1;5;96m %-20s   \033[0m", "CPU", GetCpu()),
		fmt.Sprintf("\033[38;5;196m \033[0m  %-12s \033[1;5;91m %-20s   \033[0m", "GPU", GetGpu()),
		fmt.Sprintf("\033[1;5;221m \033[0m  %-12s \033[1;5;166m %-20s  \033[0m", "memory", GetMemory()),
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

package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/bgdar/astfetch/ansi"
)

// / tampilakn dalam bentuk lurus ke samping
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
		if i < len(lines) {
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

// / tampilakn dalam bentuk lurus ke bawah
func displayVertical(lines []string, info []string) {
	fmt.Print("")
	// logo
	for _, line := range lines {
		if line != "none" {
			fmt.Println(line)
		}
	}
	fmt.Println("                                 ")
	// info
	for _, in := range info {
		if !strings.Contains(in, "none") {
			fmt.Println(in)
		}
	}
}

func main() {

	// simpan array string untuk di tampilkan
	info := []string{
		fmt.Sprintf(ansi.FgGreen+"%s"+ansi.Reset+ansi.TbTurquoise+" on "+ansi.Reset+"%s", GetUser(), GetOSName()),
		fmt.Sprint(ansi.TbWhite + "---------------------------" + ansi.Reset),
		fmt.Sprintf("  %-12s "+ansi.FgRed+" %-20s  "+ansi.Reset, "Kernel", GetKernel()),
		fmt.Sprintf(ansi.TbYellow+"󱑎 "+ansi.Reset+" %-12s "+ansi.Reset+ansi.FgGreen+" %-20s  "+ansi.Reset, "Uptime", GetUptime()),
		fmt.Sprintf(ansi.TbGray+" "+ansi.Reset+" %-12s "+ansi.FgGray+" %-20s  "+ansi.Reset, "Shell", GetShell()),
		fmt.Sprintf(ansi.TbTeal+" "+ansi.Reset+" %-12s "+ansi.FgWhite+" %-20s  "+ansi.Reset, "Resolution", GetResolution()),
		fmt.Sprintf(ansi.TbMagenta+"󰨡 "+ansi.Reset+" %-12s "+ansi.FgBlue+" %-20s  "+ansi.Reset, "Desktop", GetWindows()),
		fmt.Sprintf(ansi.TbCyan+" "+ansi.Reset+" %-12s "+ansi.FgPurple+" %-20s "+ansi.Reset, "Theme", GetTheme()),
		fmt.Sprintf(ansi.TbGray+" "+ansi.Reset+" %-12s "+ansi.FgDarkGray+" %-20s  "+ansi.Reset, "Terminal", GetTerminal()),
		fmt.Sprintf(ansi.TbBlue+" "+ansi.Reset+" %-12s "+ansi.FgWhite+" %-20s   "+ansi.Reset, "CPU", GetCpu()),
		fmt.Sprintf(ansi.TbRed+" "+ansi.Reset+" %-12s "+ansi.FgWhite+" %-20s   "+ansi.Reset, "GPU", GetGpu()),
		fmt.Sprintf(ansi.TbBrown+" "+ansi.Reset+" %-12s "+ansi.FgWhite+" %-20s  "+ansi.Reset, "memory", GetMemory()),
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

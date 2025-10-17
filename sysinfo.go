package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)


// / return string : kembalikan nama user saat ini ($USER)
func GetUser() string {
	return os.Getenv("USER")

}

// / return string  | "none" : hosname
func GetHostname() string {
	name, err := os.Hostname()
	if err != nil {
		fmt.Println("terjadi error saat mendapatkan Hostname :", err)
		return "none"
	}
	return name
}

// / rerurn string : nama os yang sedang di gunaakn
func GetOSName() string {
	return runtime.GOOS
}

// kembalikan type os /
// example return : ubuntu || arch
func GetDistroLinux() string {
	content, err := os.ReadFile("/etc/os-release")
	if err != nil {
		fmt.Println("error distro linux :", err.Error())
		return "none"
	}
	// Memecah konten berdasarkan baris
	lines := strings.Split(string(content), "\n")
	// Mencari baris "ID=" untuk mendapatkan ID distro
	result := ""

	for _, line := range lines {
		if strings.HasPrefix(line, "ID=") {
			// Mengambil nilai setelah "ID=" yakni nama nya misal ( ID=arch)
			distroID := strings.Split(line, "=")[1]
			// Menghilangkan tanda kutip jika ada
			distroID = strings.Trim(distroID, "\"")
			result = string(distroID)
		}
	}
	return result
}

// / return string | "noen" : kernel
func GetKernel() string {
	out, err := exec.Command("uname", "-r").CombinedOutput()
	if err != nil {
		fmt.Println("terjaid error saat mendapatkan kernel :", err)
		return "none"
	}
	return strings.TrimSpace(string(out))
}
func GetUptime() string {
	out, err := exec.Command("uptime", "-p").Output()
	if err != nil {
		println("uptime error :", err)
		return "none"
	}
	return strings.TrimSpace(string(out))
}

// / dapatkan informasi shell yang di gunakan saat ini
func GetShell() string {
	switch runtime.GOOS {
	case "windows":
		// Ambil dari ComSpec (CMD bawaan)
		comSpec := os.Getenv("ComSpec")
		if comSpec != "" {
			idx := strings.LastIndex(comSpec, "\\") // pakai backslash di Windows
			return string(comSpec[idx+1:])
		}
		// Jika tidak ada, fallback ke parent process (misal PowerShell)
		return "powershell.exe" // asumsi umum
	default:
		// Linux/macOS
		path := os.Getenv("SHELL")
		idx := strings.LastIndex(path, "/")
		if idx >= 0 {
			return strings.TrimSpace(string(path[idx+1:]))
		}
	}
	return "none"

}

// / mendapatkan resolusi   lebar X tinggi
func GetResolution() string {
	switch runtime.GOOS {
	case "linux":
		out, _ := exec.Command("xrandr").Output()
		re := regexp.MustCompile(`current\s+(\d+)\s+x\s+(\d+)`)
		match := re.FindStringSubmatch(string(out))
		if len(match) == 3 {
			return fmt.Sprintf("%sx%s", match[1], match[2])
		} else {
			return "none"
		}
	case "windows":
		cmd := exec.Command("powershell", "-Command", "Get-CimInstance Win32_VideoController | Select-Object -First 1 CurrentHorizontalResolution,CurrentVerticalResolution")
		out, _ := cmd.Output()
		return string(out)
	case "darwin":
		cmd := exec.Command("system_profiler", "SPDisplaysDataType")
		out, _ := cmd.Output()
		re := regexp.MustCompile(`Resolution:\s+(\d+)\s+x\s+(\d+)`)
		match := re.FindStringSubmatch(string(out))
		if len(match) == 3 {
			return fmt.Sprintf("%sx%s", match[1], match[2])
		}
	}
	return "none"
}

// / return : string : kembalikan informasi Destop envariotmeny
func GetWindows() string {
	val, ok := os.LookupEnv("XDG_CURRENT_DESKTOP")
	if ok {
		return val
	}
	return "none"
}

// / return string  | "noen": informasi theme yang di gunakan
func GetTheme() string {
	// Deteksi KDE Plasma
	if err := exec.Command("pgrep", "-x", "plasmashell").Run(); err == nil {
		out, _ := exec.Command("grep", "widgetStyle=", fmt.Sprintf("%s/.config/kdeglobals", getHome())).Output()
		parts := strings.SplitN(string(out), "=", 2)
		if len(parts) == 2 {
			return strings.TrimSpace(parts[1])
		}
		return "none"
	}
	// Deteksi GNOME
	if err := exec.Command("pgrep", "-x", "gnome-shell").Run(); err == nil {
		cmd := exec.Command("gsettings", "get", "org.gnome.desktop.interface", "gtk-theme")
		var out bytes.Buffer
		cmd.Stdout = &out
		if err := cmd.Run(); err == nil {
			theme := strings.TrimSpace(out.String())
			theme = strings.Trim(theme, "'") // hapus tanda kutip
			return theme
		}

		return "none"
	}
	// Deteksi XFCE4
	if err := exec.Command("pgrep", "-x", "xfce4-session").Run(); err == nil {
		cmd := exec.Command("xfconf-query", "-c", "xsettings", "-p", "/Net/ThemeName")
		var out bytes.Buffer
		cmd.Stdout = &out
		if err := cmd.Run(); err == nil {
			theme := strings.TrimSpace(out.String())
			return theme
		}
		return "none"
	}

	return "none"
}

func GetTerminal() string {
	switch runtime.GOOS {
	case "linux", "mac":
		return os.Getenv("TERM")
	case "windows":
		ppid := os.Getppid()
		cmd := exec.Command("wmic", "process", "where", fmt.Sprintf("ProcessId=%d", ppid), "get", "Name")
		var out bytes.Buffer
		cmd.Stdout = &out
		_ = cmd.Run()

		lines := strings.Split(strings.TrimSpace(out.String()), "\n")
		if len(lines) > 1 {
			return strings.TrimSpace(lines[1])
		}
		return "none"
	}
	return "none"
}

func GetCpu()string { 

switch runtime.GOOS {
	case "linux":
		out, err := exec.Command("bash", "-c", "grep 'model name' /proc/cpuinfo | head -1").Output() 
		if err != nil {
			fmt.Println("cpu error -> out : ",err.Error())
			return "none"
		}
		name := strings.TrimSpace(strings.SplitN(string(out), ":", 2)[1])

		cores, _ := exec.Command("nproc").Output()
		coreCount := strings.TrimSpace(string(cores))

		freq, _ := exec.Command("bash", "-c", "lscpu | grep 'MHz' | head -1 | awk '{print $3/1000}'").Output()
		freqGHz := strings.TrimSpace(string(freq))

		return fmt.Sprintf("%s (%s) @ %sGHz", name, coreCount, freqGHz)

	case "darwin":
		out, _ := exec.Command("sysctl", "-n", "machdep.cpu.brand_string").Output()
		return strings.TrimSpace(string(out))

	case "windows":
		out, _ := exec.Command("wmic", "cpu", "get", "Name,NumberOfLogicalProcessors,MaxClockSpeed").Output()
		lines := strings.Split(strings.TrimSpace(string(out)), "\n")
		if len(lines) > 1 {
			fields := strings.Fields(lines[1])
			return fmt.Sprintf("%s (%s) @ %.2fGHz", strings.Join(fields[:len(fields)-2], " "), fields[len(fields)-2], parseGHz(fields[len(fields)-1]))
		}
	}

	return "none"
}
func parseGHz(mhz string) float64 {
	var m float64
	fmt.Sscanf(mhz, "%f", &m)
	return m / 1000
}

func GetGpu() string {
	var out []byte
	var err error

	switch runtime.GOOS {
	case "linux":
		out, err = exec.Command("bash", "-c", "lspci | grep -E 'VGA|3D' | head -1 | sed 's/.*controller: //'").Output()
	case "windows":
		out, err = exec.Command("wmic", "path", "win32_VideoController", "get", "Name").Output()
	case "darwin":
		out, err = exec.Command("bash", "-c", "system_profiler SPDisplaysDataType | grep 'Chipset Model' | head -1 | sed 's/.*: //'").Output()
	default:
		return "none"
	}

	if err != nil || len(out) == 0 {
		return "none"
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	gpu := strings.TrimSpace(lines[len(lines)-1])

	// Hilangkan kata "Corporation", "Intel", atau simbol kotak berlebih
	gpu = strings.ReplaceAll(gpu, "Corporation", "")
	gpu = strings.TrimSpace(gpu)
	return gpu
}

/// mendapatkan ukuran memory 
func GetMemory()string {

	switch runtime.GOOS {
	case "linux","mac" :
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return "none"
	}
	var total, available uint64
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "MemTotal:") {
			fmt.Sscanf(line, "MemTotal: %d kB", &total)
		}
		if strings.HasPrefix(line, "MemAvailable:") {
			fmt.Sscanf(line, "MemAvailable: %d kB", &available)
		}
	}
	used := total - available
	return fmt.Sprintf("Memory: %dMiB / %dMiB (%.1f%%)",
		used/1024, total/1024, float64(used)/float64(total)*100)
	case "windows" :
	out, err := exec.Command("wmic", "OS", "get", "TotalVisibleMemorySize,FreePhysicalMemory", "/Value").Output()
	if err != nil {
		return "none"
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	var total, free uint64
	for _, line := range lines {
		if strings.HasPrefix(line, "TotalVisibleMemorySize=") {
			val := strings.TrimPrefix(line, "TotalVisibleMemorySize=")
			total, _ = strconv.ParseUint(val, 10, 64)
		}
		if strings.HasPrefix(line, "FreePhysicalMemory=") {
			val := strings.TrimPrefix(line, "FreePhysicalMemory=")
			free, _ = strconv.ParseUint(val, 10, 64)
		}
	}

	used := total - free
	return fmt.Sprintf("%dMiB / %dMiB (%.1f%%)", used/1024, total/1024, float64(used)/float64(total)*100)
	}

	return "none"
}

func getHome() string {
	cmd := exec.Command("bash", "-c", "echo $HOME")
	out, _ := cmd.Output()
	return strings.TrimSpace(string(out))
}

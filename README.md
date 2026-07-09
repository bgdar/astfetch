<h1 align="center">
Astfetch
</h1>
tools astetik-fetch , informasi system mirip sperti neofetch

<br>

### Parameter `Display`

default **horizontal**

- `d` => **down** tampil secara **vertical**
<br>

### Function Info
* GetUser(): Mengambil nama pengguna aktif dari env variable USER pada semua OS.
* GetHostname(): Mengambil nama host/jaringan perangkat menggunakan API internal bawaan Go pada semua OS.
* GetOSName(): Mengambil nama dasar sistem operasi lewat konstanta runtime.GOOS pada semua OS.
* GetDistroLinux(): Membaca nama distro Linux dari file /etc/os-release (`Linux`).
* GetKernel(): Mengambil versi kernel sistem lewat CLI uname -r ( `Linux` dan `macOS`).
* GetUptime(): Mengambil durasi aktif komputer lewat CLI uptime -p ( `Linux` dan `macOS`).
* GetShell(): Mengambil jenis shell aktif lewat env SHELL (Linux/Mac) atau env ComSpec (`Windows`).
* GetResolution(): Mengambil resolusi layar via xrandr (`Linux`), PowerShell (`Windows`), atau system_profiler (Mac).
* GetWindows(): Mengambil nama Desktop Environment Linux lewat env XDG_CURRENT_DESKTOP ( `Linux` ).
* GetTheme(): Mengambil tema aktif lewat pencarian proses UI seperti GNOME, KDE, atau XFCE (`Linux`).
* GetTerminal(): Mengambil nama aplikasi terminal via env TERM (`Linux`` dan `Mac`) atau query wmic (`Windows`).
* GetCpu(): Mengambil model dan spek CPU via /proc/cpuinfo (`Linux`), sysctl (`Mac`), atau wmic (`Windows`).
* GetGpu(): Mengambil nama kartu grafis via lspci (`Linux`), system_profiler (`Mac`), atau wmic (`Windows`).
* GetMemory(): Mengambil info RAM via file /proc/meminfo (`Linux`) atau query wmic OS (`Windows`).

<br>

### Demo
![contoh hasil](./img/astfetch.png) 

<br>

### how to use

1. install excutable langsung <https://github.com/bgdar/astfetch/releases/tag/v1.0> 

2. install ke directori default `go/bin`
```bash 
# cloning repo 
https://github.com/bgdar/astfetch.git  

# install 
go install
```

<br>

### other

> sysinfo.sh => tidak saya gunakan , tapi full golang
> jika ingin gunakan cukup panggil

```bash
bash sysinfo.sh
# atau
./sysinfo.sh
```


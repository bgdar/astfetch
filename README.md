# astfetch

tools astetik-fetch , informasi system mirip sperti neofetch tapi dengan konfigurasi manual sesaui kode asci yg di berikan

## detail info

OS dan Kernel → runtime.GOOS, runtime.GOARCH, gopsutil/host

Hostname → os.Hostname()

Uptime → gopsutil/host

CPU → model, core count → gopsutil/cpu

Memory → total & used → gopsutil/mem

Disk → kapasitas storage → gopsutil/disk

GPU (opsional, lebih ribet, bisa lewat lspci atau library tambahan)

Shell yang digunakan (ambil dari environment variable $SHELL)

Terminal (env $TERM)

Resolution (kalau di Linux, bisa panggil xrandr atau fbset)

Theme / WM / DE (biasanya lewat environment variable, misal

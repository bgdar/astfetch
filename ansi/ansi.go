package ansi

const (
	// Reset
	Reset = "\033[0m"

	// Foreground colors (38;5)
	FgBlack     = "\033[38;5;16m"
	FgRed       = "\033[38;5;196m" // merah terang
	FgGreen     = "\033[38;5;46m"  // hijau terang
	FgYellow    = "\033[38;5;226m"
	FgBlue      = "\033[38;5;39m"
	FgMagenta   = "\033[38;5;201m" // magenta / ungu muda
	FgCyan      = "\033[38;5;51m"  // cyan / biru muda
	FgWhite     = "\033[38;5;255m"
	FgGray      = "\033[38;5;252m"
	FgDarkGray  = "\033[38;5;128m"
	FgGold      = "\033[38;5;214m"
	FgOrange    = "\033[38;5;208m" // oranye
	FgPink      = "\033[38;5;213m" // Pink lembut
	FgPurple    = "\033[38;5;93m"  // Ungu tua
	FgTeal      = "\033[38;5;37m"  // Biru kehijauan
	FgBrown     = "\033[38;5;94m"  // Coklat
	FgLightBlue = "\033[38;5;153m" // Biru muda pastel
	FgLime      = "\033[38;5;118m" // Hijau limau
	FgTurquoise = "\033[38;5;80m"  // Turquoise

	// Text tebal (bold)

	TbRed       = "\033[1;38;5;196m" // merah terang (bold)
	TbGreen     = "\033[1;38;5;46m"  // hijau terang (bold)
	TbYellow    = "\033[1;38;5;226m"
	TbBlue      = "\033[1;38;5;39m"
	TbMagenta   = "\033[1;38;5;201m"
	TbCyan      = "\033[1;38;5;51m"
	TbWhite     = "\033[1;38;5;255m"
	TbGray      = "\033[1;38;5;252m"
	TbDarkGray  = "\033[1;38;5;235m"
	TbGold      = "\033[1;38;5;214m"
	TbOrange    = "\033[1;38;5;208m"
	TbPink      = "\033[1;38;5;213m"
	TbPurple    = "\033[1;38;5;93m"
	TbTeal      = "\033[1;38;5;37m"
	TbBrown     = "\033[1;38;5;94m"
	TbLightBlue = "\033[1;38;5;153m"
	TbLime      = "\033[1;38;5;118m"
	TbTurquoise = "\033[1;38;5;80m"

	// Background colors (48;5)
	BgBlack    = "\033[48;5;16m"
	BgWhite    = "\033[48;5;15m" // Putih terang
	BgGold     = "\033[48;5;214m"
	BgRed      = "\033[48;5;196m"
	BgGreen    = "\033[48;5;46m"
	BgBlue     = "\033[48;5;39m"
	BgYellow   = "\033[48;5;226m"
	BgMagenta  = "\033[48;5;201m"
	BgCyan     = "\033[48;5;51m"
	BgOrange   = "\033[48;5;208m"
	BgGray     = "\033[48;5;252m"
	BgDarkGray = "\033[48;5;235m"
)

// example :
// ini bisa ketikan di gabung : "\033[38;5;39m" + "\033[48;5;226m" -> ESC[38;5;39mESC[48;5;226m , ini valid

// fmt.Println(
// 	ansi.FgBlue + ansi.BgYellow + " BLUE TEXT ON YELLOW BG " + ansi.Reset,
// )
//
// fmt.Println(
// 	ansi.FgYellow + ansi.BgBlue + " YELLOW TEXT ON BLUE BG " + ansi.Reset,
// )

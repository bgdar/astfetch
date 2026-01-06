package main

import "github.com/bgdar/astfetch/ansi"

var MainLogoDistroOS = map[string][]string{

	// windows APRov
	"windows": {
		"                                         ",
		ansi.Reset + ansi.FgBlue + " ┌───────────┬───────────┐" + ansi.Reset + ansi.Reset,
		ansi.Reset + ansi.FgBlue + " │ " + ansi.Reset + ansi.BgBlue + " ████████" + ansi.Reset + ansi.FgBlue + " │ " + ansi.Reset + ansi.BgBlue + "████████ " + ansi.Reset + ansi.FgBlue + " │ " + ansi.Reset,
		ansi.Reset + ansi.FgBlue + " │ " + ansi.Reset + ansi.BgBlue + " ████████" + ansi.Reset + ansi.FgBlue + " │ " + ansi.Reset + ansi.BgBlue + "████████ " + ansi.Reset + ansi.FgBlue + " │ " + ansi.Reset,

		ansi.Reset + ansi.FgBlue + " │ " + ansi.Reset + ansi.BgBlue + " ████████" + ansi.Reset + ansi.FgBlue + " │ " + ansi.Reset + ansi.BgBlue + "████████ " + ansi.Reset + ansi.FgBlue + " │ " + ansi.Reset,

		ansi.Reset + ansi.FgBlue + " ├───────────┼───────────┤" + ansi.Reset + ansi.Reset,
		ansi.Reset + ansi.FgBlue + " │ " + ansi.Reset + ansi.BgBlue + " ████████" + ansi.Reset + ansi.FgBlue + " │ " + ansi.Reset + ansi.BgBlue + "████████ " + ansi.Reset + ansi.FgBlue + " │ " + ansi.Reset,
		ansi.Reset + ansi.FgBlue + " │ " + ansi.Reset + ansi.BgBlue + " ████████" + ansi.Reset + ansi.FgBlue + " │ " + ansi.Reset + ansi.BgBlue + "████████ " + ansi.Reset + ansi.FgBlue + " │ " + ansi.Reset,
		ansi.Reset + ansi.FgBlue + " │ " + ansi.Reset + ansi.BgBlue + " ████████" + ansi.Reset + ansi.FgBlue + " │ " + ansi.Reset + ansi.BgBlue + "████████ " + ansi.Reset + ansi.FgBlue + " │ " + ansi.Reset,
		ansi.Reset + ansi.FgBlue + " └───────────┴───────────┘" + ansi.Reset + ansi.BgBlack + " " + ansi.Reset,
	},
	// urutan linux sudah fixed
	// Linux sudah APRov
	"linux": {
		"                    ",
		"                    ",
		"                    ",

		"        " + ansi.FgWhite + ansi.BgDarkGray + ".--." + ansi.Reset + "        ",
		"       " + ansi.FgWhite + ansi.BgDarkGray + "|" + ansi.FgRed + ansi.BgBlack + "o" + ansi.Reset + ansi.FgBlack + ansi.BgDarkGray + " " + ansi.Reset + ansi.FgWhite + ansi.BgBlack + "-" + ansi.Reset + ansi.FgWhite + ansi.BgDarkGray + " |" + ansi.Reset + "       ",
		"       " + ansi.FgWhite + ansi.BgDarkGray + "|" + ansi.Reset + ansi.FgYellow + "\\=/" + ansi.Reset + ansi.FgWhite + ansi.BgDarkGray + " |" + ansi.Reset + "       ",
		"      " + ansi.FgWhite + ansi.BgDarkGray + "//" + ansi.Reset + ansi.FgBlack + ansi.BgWhite + "   \\" + ansi.Reset + ansi.FgWhite + ansi.BgDarkGray + " \\" + ansi.Reset + "     ",
		"    " + ansi.FgWhite + ansi.BgDarkGray + "('|" + ansi.Reset + ansi.FgBlack + ansi.BgWhite + "     " + ansi.Reset + ansi.FgWhite + ansi.BgDarkGray + "|')" + ansi.Reset + "     ",
		"    " + ansi.FgWhite + ansi.BgDarkGray + "/'\\" + ansi.Reset + ansi.FgBlack + ansi.BgWhite + "\\_  _/" + ansi.Reset + ansi.FgWhite + ansi.BgDarkGray + "`\\" + ansi.Reset + "     ",
		"    " + ansi.FgBlack + ansi.BgGold + "\\___)" + ansi.Reset + ansi.FgWhite + ansi.BgDarkGray + "-" + ansi.Reset + ansi.FgBlack + ansi.BgGold + "(___/" + ansi.Reset + "     ",
	},
	"mac": {
		"                                         ",
		"                                         ",
		"            " + ansi.FgBlack + ansi.BgGray + " .::`" + ansi.Reset,
		"           " + ansi.FgBlack + ansi.BgGray + ".::'" + ansi.Reset + "        ",
		"     __    " + ansi.FgBlack + ansi.BgGray + ":'" + ansi.Reset + "__          ",
		"   " + ansi.FgBlack + ansi.BgGray + " .'`'`'-__-' `''" + ansi.Reset + "      ",
		"  " + ansi.FgBlack + ansi.BgGray + " :            ." + ansi.Reset + "    ",
		" " + ansi.FgBlack + ansi.BgGray + ".'             :" + ansi.Reset + "    ",
		"  " + ansi.FgBlack + ansi.BgGray + ",      " + ansi.Reset + ansi.FgCyan + ansi.BgRed + "a" + ansi.Reset + ansi.FgGreen + ansi.BgDarkGray + "s" + ansi.Reset + ansi.FgGold + ansi.BgBlue + "t" + ansi.Reset + ansi.BgGray + "    `.." + ansi.Reset + "    ",
		"  " + ansi.FgBlack + ansi.BgGray + " :_             ;" + ansi.Reset + "    ",
		"    " + ansi.FgBlack + ansi.BgGray + "\"-...___-.--\"'" + ansi.Reset + "       ",
	},
}

// coleksi distro linux
// var LinuxDistroLogo = map[string][]string{
//
// 	"ubuntu": {
// 	"",
// },
// }

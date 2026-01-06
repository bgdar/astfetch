package ansi


/// gunakan untuk menggabungkan langsung 
func Colorize(fg, bg, text string) string {
	return fg + bg + text + Reset
}
// example :
// for _, line := range lines {
// 		fmt.Println(
// 			ansi.Colorize(ansi.FgBlue, ansi.BgBlue, line),
// 		)
// 	}

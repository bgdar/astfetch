package main

import "runtime"

var MainLogoDistroOS = map[string]string{
	"linux": `
        .--.            
       |o_o |           
       |:_/ |           
      //   \ \          
     (|     | )         
    /'\_   _/'\         
    \___)=(___/         
	`,
	"windows": `
       +-----------+
       |     |     |
       |-----+-----|
       |     |     |
       +-----------+
	`,
	"macos": `
       .:'          
   __ :'__          
 .'  ` + "`" + `         
:            :      
:            :      
 :          :       
   ` + "`" + `                 
	`,
}

// kembikan logo os
func GetLogoOS() string {
	var typeOs = runtime.GOOS
	switch typeOs {
	case "linux":
		// di validasi lagi untuk distro , jika tidak ada , maka tapikna logo linux
		return MainLogoDistroOS["linux"]
	case "windows":
		return MainLogoDistroOS["windows"]
	case "mac":
		return MainLogoDistroOS["mac"]
	default:
		return "NOOOO OS"
	}
}

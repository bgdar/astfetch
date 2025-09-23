package main

import (
	"fmt"
	"os/exec"
)

// astfetch != asteticek fetch
func main() {

	var ouput, errSysinfo = exec.Command("bash", "./sysinfo.sh").Output()
	if errSysinfo != nil {
		panic(errSysinfo)
	}

	// var systemType, errSystemType = exec.Command("echo", "$OSTYPE").Output()
	// if errSystemType != nil {
	// 	panic(errSystemType)
	//
	// }

	// fmt.Print(GetLogoOS())
	// info system nya
	// fmt.Print(string(ouput))
	fmt.Printf("%-30s %-30s", GetLogoOS(), string(ouput))

}

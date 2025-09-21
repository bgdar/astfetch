package main

import (
	"fmt"
	"os/exec"
)

func main() {

	fmt.Print("astfetch")

	var ouput, errSysinfo = exec.Command("bash", "./sysinfo.sh").Output()
	if errSysinfo != nil {
		panic(errSysinfo)
	}

	fmt.Print(string(ouput))

}

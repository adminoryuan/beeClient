package main

import (
	"fmt"
	"regexp"

	"github.com/adminoryuan/beeclient/common"
	"github.com/adminoryuan/beeclient/untils"
)

type MemoryInfo struct {
	memused  string
	swapused string
}

func getMemoryInfo() common.Memoryinfo {
	rex := `MiB Mem : .+(\d+.\d+) used,.+\n.+(\d+.\d+) used`
	res, err := untils.ReadChangeCmdCount("/usr/bin/top", 6, "-bn", "1")
	if err != nil {
		fmt.Println(res)
	}
	re, _ := regexp.Compile(rex)

	all := re.FindStringSubmatch(res)
	if len(all) < 3 {
		fmt.Println("len", len(all))
		return common.Memoryinfo{}
	}
	return common.Memoryinfo{Memused: all[1], Swapused: all[2]}
}

package main

import (
	"fmt"
	"regexp"

	"github.com/adminoryuan/beeclient/common"
	"github.com/adminoryuan/beeclient/untils"
)

func getCpuinfo() common.Cpuinfo {
	res, err := untils.ReadChangeCmdCount("/usr/bin/top", 6, "-bn", "1")
	if err != nil {
		fmt.Println(res)
	}
	regex := `Cpu\(s\): (\d+.\d+) us,  (\d+.\d+) sy.+(\d+.\d+) wa`

	re, _ := regexp.Compile(regex)

	all := re.FindStringSubmatch(res)
	if len(all) < 2 {
		return common.Cpuinfo{}
	}
	return common.Cpuinfo{

		Us: all[1],
		Sy: all[2],
		Io: all[3],
	}
}

package main

import (
	"fmt"

	"github.com/adminoryuan/beeclient/common"

	"github.com/adminoryuan/beeclient/untils"
)

//获取磁盘情况
func getDiskInfo() common.Diskinfo {
	res := untils.GetProcessOutSteam(untils.BinBash, "-c", "df -h /")
	fmt.Println(string(res))
	return common.Diskinfo{}
}

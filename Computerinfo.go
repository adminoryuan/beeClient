package main

import (
	"github.com/adminoryuan/beeclient/common"
	"github.com/adminoryuan/beeclient/untils"
)

type ComputeInfo struct {
}

//获取服务器性能
func (c *ComputeInfo) getComputerinfo() (common.Computerinfo, string) {
	cpuCount := untils.GetProcessOutSteam(untils.BinBash, "-c", "cat /proc/cpuinfo |grep 'physical id'|sort|uniq|wc -l ")
	cpuCore := untils.GetProcessOutSteam(untils.BinBash, "-c", "cat /proc/cpuinfo |grep 'core id'|wc -l")
	boardMan := untils.GetProcessOutSteam(untils.BinBash, "-c", "sudo dmidecode -t 2")
	addr := untils.GetProcessOutSteam(untils.BinBash, "-c", "curl ifconfig.me")
	return common.Computerinfo{
		MemCount:  getMemorySum(),
		CpuCount:  string(cpuCount),
		MainBoard: string(boardMan),
		CpuCore:   string(cpuCore),
		DiskCount: "1000000kb",
	}, string(addr)
}

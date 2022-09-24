package main

import "github.com/adminoryuan/beeclient/common"

type BeeCol struct{}

//收集服务器所有性能指标
func (c *BeeCol) Collection() common.CollectInfo {
	cpuinfo := getCpuinfo()
	diskinfo := getDiskInfo()
	meminfo := getMemoryInfo()
	proinfo := getProcessinfo()
	return common.CollectInfo{
		Cpu:     &cpuinfo,
		Disk:    &diskinfo,
		Meminfo: &meminfo,
		Proinfo: &proinfo,
	}
}

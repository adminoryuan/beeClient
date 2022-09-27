package main

import (
	"fmt"
	"testing"
)

func TestCompute(t *testing.T) {
	info := ComputeInfo{}
	cinfo, addr := info.getComputerinfo()
	fmt.Println(cinfo.CpuCore)
	fmt.Println(cinfo.CpuCount)
	fmt.Println(cinfo.MainBoard)
	fmt.Println(cinfo.MemCount)

	fmt.Println(addr)
}

package main

import (
	"fmt"
	"testing"
)

func TestCpu(t *testing.T) {

	info := getCpuinfo()
	fmt.Println(info.Us)
	fmt.Println(info.Sy)
	fmt.Println(info.Io)
}

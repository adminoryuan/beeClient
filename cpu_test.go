package main

import (
	"fmt"
	"testing"
)

func TestCpu(t *testing.T) {
	for i := 0; i < 100; i++ {

		info := getCpuinfo()
		fmt.Println(info.Us)
		fmt.Println(info.Sy)
		fmt.Println(info.Io)
	}
}

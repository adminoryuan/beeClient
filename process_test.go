package main

import (
	"fmt"
	"testing"
)

func TestProcess(t *testing.T) {
	pro := getProcessinfo()
	for i := 0; i < len(pro.Processinfo); i++ {
		fmt.Println(pro.Processinfo[i])
	}
}

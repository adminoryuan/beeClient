package main

import (
	"fmt"
	"testing"
)

func TestMemeory(t *testing.T) {
	info := getMemoryInfo()
	fmt.Println(info.Memused)

	fmt.Println(info.Swapused)
}

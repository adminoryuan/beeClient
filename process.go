package main

import (
	"reflect"
	"strings"

	"github.com/adminoryuan/beeclient/common"

	"github.com/adminoryuan/beeclient/untils"
)

type process struct {
}
type ProcessInfo struct {
	pid     string
	user    string
	command string
	mem     string
	cpu     string
}

func getProcessinfo() common.Processinfo {
	res, _ := untils.ReadChangeCmdCount(untils.TopPath, 17, "-bn", "1")
	headlen := 7
	resSplits := strings.Split(res, "\n")
	processinfo := make([]*common.Process, 0)

	if len(resSplits) < headlen {
		return common.Processinfo{}
	}
	for i := headlen; i < len(resSplits); i++ {
		row := resSplits[i]
		thenProInfo := paserProinfo(row)
		if !reflect.DeepEqual(&thenProInfo, ProcessInfo{}) {

			processinfo = append(processinfo, &thenProInfo)
		}
	}
	return common.Processinfo{
		Processinfo: processinfo,
	}
}

func paserProinfo(item string) common.Process {
	row := strings.Split(item, " ")
	//newRow := [12]string{}
	newRow := make([]string, 0)

	index := 0
	for i := 0; i < len(row); i++ {
		if strings.EqualFold(row[i], "") {
			continue
		}
		newRow = append(newRow, row[i])
		index++
	}
	if len(newRow) < 9 {
		return common.Process{}
	}
	return common.Process{
		Pid:     newRow[0],
		User:    newRow[1],
		Cpu:     newRow[8],
		Mem:     newRow[9],
		Command: newRow[11],
	}

}

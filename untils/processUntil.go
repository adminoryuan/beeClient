package untils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

var (
	TopPath string = "/usr/bin/top"
	IpPath  string = "/usr/sbin/ip"
	BinBash string = "/bin/bash"
)

func GetProcessOutSteam(pro string, args ...string) []byte {
	cmd := exec.Command(pro, args...)

	var out bytes.Buffer
	cmd.Stdout = &out

	cmd.Start()

	cmd.Wait()
	return out.Bytes()
}

//返回实时变化的命令行 读取有限行
func ReadChangeCmdCount(commandName string, len int, params ...string) (string, error) {
	cmd := exec.Command(commandName, params...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Print("cmd.StdoutPipe: ", err)
		return "", err
	}
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		log.Print("cmd start  error")

		return "", err
	}
	res := ""
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)
	for i := 0; i < len; i++ {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		res += line
	}
	cmd.Process.Kill()
	err = cmd.Wait()
	return res, err
}

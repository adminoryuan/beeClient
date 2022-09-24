package untils

import (
	"fmt"
	"testing"
)

func TestUntils(t *testing.T) {
	body := GetProcessOutSteam(TopPath, "-w")

	fmt.Println(string(body))

}
func TestXxx(t *testing.T) {

	res, _ := ReadChangeCmdCount(TopPath, 20, "-bn", "1")
	fmt.Println(res)
}

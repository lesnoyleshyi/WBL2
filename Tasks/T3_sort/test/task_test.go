package test

import (
	"fmt"
	"os/exec"
	"reflect"
	"testing"
)

var TestCases = [][]string{
	{"./f1.txt"},
	{"./f1.txt", "-r"},
	{"./f1.txt", "-k", "1"},
	{"./f1.txt", "-r", "-u"},
	{"./f3.txt", "-n", "-r"},
	{"./f3.txt", "-n", "-b"},
	{"./f3.txt", "-n", "-M"},
	{"./f3.txt", "-h"},
	{"./f3.txt", "-h", "-r"},
	{"./f3.txt", "-n"},
}

func Test(t *testing.T) { //nolint:paralleltest
	for _, args := range TestCases {
		out1, _ := exec.Command("../sort", args...).Output()
		out2, _ := exec.Command("sort", args...).Output()

		if reflect.DeepEqual(out1, out2) == false {
			t.Error("args failed")
			fmt.Println(string(out1)) //nolint:forbidigo
			fmt.Println(string(out2)) //nolint:forbidigo
		}
	}
}

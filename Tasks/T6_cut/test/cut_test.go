package test

import (
	"fmt"
	"os/exec"
	"reflect"
	"testing"
)

var TestCases = [][]string{
	{"-f", "1-2,4-5", "./f1.txt"},
	{"-f", "1,2,3,4", "-d", "b", "./f2.txt"},
	{"-f", "1,-1,3,4", "-d", " ", "./f2.txt"},
	{"-f", "1,2,3,4", "-d", ",", "-s", "./f2.txt"},
	{"-f", "2,1-4", "-d", ",", "-s", "./f2.txt"},
	{"-f", "1,2,3,4", "-d", "z", "./f3.txt"},
	{"-f", "2,1-4", "-d", ".", "-s", "./f3.txt"},
}

func Test(t *testing.T) {
	for _, args := range TestCases {
		out1, _ := exec.Command("../cut_go", args...).Output()
		out2, _ := exec.Command("cut", args...).Output()

		if reflect.DeepEqual(out1, out2) == false {
			t.Error("test failed")
			fmt.Println(string(out1)) //nolint:forbidigo
			fmt.Println("----------")
			fmt.Println(string(out2)) //nolint:forbidigo
		}
	}
}

package test

import (
	"fmt"
	"os/exec"
	"reflect"
	"testing"
)

var TestCases = [][]string{
	{"1", "./f1.txt", "-A", "2", "-B", "2", "-c", "-i"},
	{"1", "./f1.txt", "-A", "2", "-B", "5", "-i"},
	{"1", "./f1.txt", "-A", "2", "-B", "2", "-i", "-F"},
	{"1", "./f1.txt", "-A", "4", "-B", "2", "-C", "3", "-i", "-F"},
	{"A", "./f3.txt", "-A", "2", "-B", "2", "-c", "-i"},
	{"A", "./f3.txt", "-A", "2", "-B", "2", "-i"},
	{"A", "./f3.txt", "-A", "2", "-B", "2", "-i", "-c"},
	{"A", "./f3.txt", "-C", "1", "-B", "2", "-c", "-i"},
	{"A", "./f3.txt", "-C", "3", "-B", "5", "-i"},
	{"A", "./f3.txt", "-C", "4", "-B", "2", "-i", "-c"},
}

func Test(t *testing.T) {
	for _, args := range TestCases {
		out1, _ := exec.Command("../grep", args...).Output()
		out2, _ := exec.Command("grep", args...).Output()

		if reflect.DeepEqual(out1, out2) == false {
			t.Error("test failed")
			fmt.Println(string(out1)) //nolint:forbidigo
			fmt.Println("------")
			fmt.Println(string(out2)) //nolint:forbidigo
		}
	}
}

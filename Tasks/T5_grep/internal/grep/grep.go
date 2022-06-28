package grep

import (
	packkey "WBL2/Tasks/T5_grep/internal/key"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// GetLines return lines slice
func GetLines(key *packkey.Key) ([]string, error) {
	var fileName string

	if len(os.Args) >= 3 { //nolint:gomnd
		fileName = os.Args[2]
	} else {
		fileName = os.Stdin.Name()
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error wheh open file: %w", err)
	}

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil

}

// Compare line with pattern according to flags
func Compare(key *packkey.Key, line, pattern string) bool {
	if key.IgnoreCase {
		line, pattern = strings.ToLower(line), strings.ToLower(pattern)
	}

	if key.Fixed {
		return line == pattern
	}

	ret, _ := regexp.MatchString(pattern, line)
	return ret
}

// Contains check whether item is in slice or not
func Contains(lines []int, idx int) bool {
	for _, a := range lines {
		if a == idx {
			return true
		}
	}
	return false
}

// Group for grep to control (non-)overlapping lines
type Group struct {
	index []int
	left  int
	right int
}

// GetBorders return left and right border of current group
func GetBorders(key *packkey.Key, idx int, len int) (int, int) {
	left := idx - key.Before
	right := idx + key.After

	if left < 0 {
		left = 0
	}
	if right >= len {
		right = len - 1
	}

	return left, right
}

// UniqueOutput control removes overlaps in output
func UniqueOutput(prevGroup, group *Group, idx int, out []string, lines []string) []string {
	if len(group.index) == 0 && idx != 0 {
		if Contains(prevGroup.index, group.left) == false && idx != group.left {
			out = append(out, "--")
		}
	}
	if Contains(prevGroup.index, idx) == false {
		out = append(out, lines[idx])
		group.index = append(group.index, idx)
	}
	return out
}

// Grep set flag options and executes filtration of input
func Grep(key *packkey.Key) ([]string, error) {

	if key.Context > 0 {
		if key.Before == 0 {
			key.Before = key.Context
		}
		if key.After == 0 {
			key.After = key.Context
		}
	}

	lines, err := GetLines(key)
	if err != nil {
		return nil, fmt.Errorf("Error when get lines: %w", err)
	}

	pattern := os.Args[1]
	out := make([]string, 0)
	prevGroup, group := &Group{}, &Group{}
	for i, line := range lines {
		if Compare(key, line, pattern) != key.Invert {
			left, right := GetBorders(key, i, len(lines))
			group = &Group{left: left, right: right}
			for j := left; j <= right; j++ {
				out = UniqueOutput(prevGroup, group, j, out, lines)
			}
		}
		prevGroup.index = append(prevGroup.index, group.index...)
		prevGroup.left, prevGroup.right = group.left, group.right
	}

	if key.LineNum {
		for i, l := range out {
			out[i] = fmt.Sprintf("%d:%s", i+1, l)
		}
	}

	return out, nil
}

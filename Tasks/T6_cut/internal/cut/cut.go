package cut

import (
	packkey "WBL2/Tasks/T6_cut/internal/key"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Cut structure
type Cut struct {
	key *packkey.Key
}

//New returns instance of type cut
func New(key *packkey.Key) *Cut {
	return &Cut{key: key}
}

//GetLines return slice of lines from file
func (*Cut) GetLines() ([]string, error) {
	fileName := os.Args[len(os.Args)-1]

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

//GetFields return list of specified fields
func (c *Cut) GetFields() []int {
	index := make([]string, 0)

	for _, val := range c.key.Fields {
		index = append(index, strings.Split(val, "-")...)
	}

	indexInt := make([]int, 0)
	for _, val := range index {
		d, _ := strconv.Atoi(val)
		indexInt = append(indexInt, d)
	}

	return indexInt
}

//contains check whether item is in slice or not
func contains(lines []int, idx int) bool {
	for _, a := range lines {
		if a == idx {
			return true
		}
	}
	return false
}

//Cut split input lines and chose needed columns
func (c *Cut) Cut() ([]string, error) {
	lines, err := c.GetLines()
	if err != nil {
		return nil, fmt.Errorf("error when get lines: %w", err)
	}

	result := make([]string, 0)
	str := strings.Builder{}
	for _, line := range lines {
		str.Reset()
		if c.key.Separated == false || strings.Contains(line, c.key.Delimiter) {
			for i, s := range strings.Split(line, c.key.Delimiter) {
				if contains(c.GetFields(), i+1) {
					str.WriteString(s + c.key.Delimiter)
				}
			}
			result = append(result, str.String()[:len(str.String())-1])
		}
	}

	return result, nil
}

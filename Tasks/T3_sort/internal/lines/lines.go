package lines

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Lines is redefenition of [][]string
type Lines [][]string

// GetLines return file lines
func GetLines(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error while open file: %w", err)
	}
	defer file.Close()

	l := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l = append(l, []string{"", scanner.Text()})
	}

	return l, nil
}

// source https://github.com/skarademir/naturalsort/blob/master/naturalsort.go
var r = regexp.MustCompile(`[^0-9]+|[0-9]+`)

// SortHumanNumeric compares input in human readable format
func (l Lines) SortHumanNumeric(i, j int) bool {
	spliti := r.FindAllString(strings.Replace(l[i][0], " ", "", -1), -1)
	splitj := r.FindAllString(strings.Replace(l[j][0], " ", "", -1), -1)

	for index := 0; index < len(spliti) && index < len(splitj); index++ {
		if spliti[index] != splitj[index] {
			// Both slices are numbers
			if isNumber(spliti[index][0]) && isNumber(splitj[index][0]) {
				// Remove Leading Zeroes
				stringi := strings.TrimLeft(spliti[index], "0")
				stringj := strings.TrimLeft(splitj[index], "0")
				if len(stringi) == len(stringj) {
					for indexchar := 0; indexchar < len(stringi); indexchar++ {
						if stringi[indexchar] != stringj[indexchar] {
							return stringi[indexchar] < stringj[indexchar]
						}
					}
					return len(spliti[index]) < len(splitj[index])
				}
				return len(stringi) < len(stringj)
			}
			// One of the slices is a number (we give precedence to numbers regardless of ASCII table position)
			if isNumber(spliti[index][0]) || isNumber(splitj[index][0]) {
				return isNumber(spliti[index][0])
			}
			// Both slices are not numbers
			return spliti[index] < splitj[index]
		}

	}
	// Fall back for cases where space characters have been annihilated by the replacement call
	// Here we iterate over the unmolsested string and prioritize numbers over
	for index := 0; index < len(l[i][0]) && index < len(l[j][0]); index++ {
		if isNumber(l[i][0][index]) || isNumber(l[j][0][index]) {
			return isNumber(l[i][0][index])
		}
	}
	return l[i][0] < l[j][0]
}

func isNumber(input uint8) bool {
	return input >= '0' && input <= '9'
}

// StandardSort compares input in lexicographical format
func (l Lines) StandardSort(i, j int) bool {
	return l[i][0] <= l[j][0]
}

// SortNumeric compares input in numeric format
func (l Lines) SortNumeric(i, j int) bool {
	d1, err1 := strconv.ParseFloat(l[i][0], 64) //nolint:gomnd
	d2, err2 := strconv.ParseFloat(l[j][0], 64) //nolint:gomnd

	if err1 != nil && err2 != nil {
		return l[i][0] < l[j][0]
	} else if err1 != nil {
		return true
	} else if err2 != nil {
		return false
	}

	return d1 < d2
}

var month = []string{"jan", "feb", "mar", "apr", "may", "june", "july", "aug", "sep", "oct", "nov", "dec"}

// MonthIndex returns index of month
func MonthIndex(str string) int {
	i := 0
	for _, m := range month {
		if strings.HasPrefix(strings.ToLower(str), m) {
			return i
		}
		i++
	}
	return -1
}

// SortMonth compares input in month format
func (l Lines) SortMonth(i, j int) bool {
	i1 := MonthIndex(l[i][0])
	i2 := MonthIndex(l[j][0])
	if i1 != -1 && i2 != -1 {
		return i1 < i2
	} else if i1 != -1 {
		return true
	} else if i2 != -1 {
		return false
	}

	return l[i][0] < l[j][0]
}

// Unique removes repeating lines from output
func (l Lines) Unique() Lines {
	inResult := make(map[string]bool)
	var result Lines
	for _, line := range l {
		str := line[0]
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, line)
		}
	}
	return result
}

func (l Lines) Reverse() {
	for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
		l[i], l[j] = l[j], l[i]
	}
}

// IgnoreTailingSpaces removes tailing spaces
func (l Lines) IgnoreTailingSpaces() {
	for i := range l {
		l[i][1] = strings.TrimRight(l[i][1], " ")
	}
}

// SetColumn sets column to compare with
func (l Lines) SetColumn(lines Lines, column int) {
	column--
	for i := range l {
		strArr := strings.Split(l[i][1], " ")
		if column > len(strArr)-1 {
			l[i][0] = strArr[len(strArr)-1]
		} else {
			l[i][0] = strArr[column]
		}
	}
}

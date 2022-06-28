package sort

import (
	packkey "WBL2/Tasks/T3_sort/internal/key"
	paklines "WBL2/Tasks/T3_sort/internal/lines"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
)

var errMisFilename = errors.New("missing filename")
var errMutExflags = errors.New("mutually exclusive flags")

// CheckArguments validates received program flags
func CheckArguments(lines paklines.Lines, key *packkey.Key) error {
	if len(os.Args) < 2 { //nolint:gomnd
		return errMisFilename
	}

	a := []bool{key.NumericSort, key.MonthSort, key.HumanNumericSort}
	k := 0
	for _, flag := range a {
		if flag {
			k++
		}
	}

	if k > 1 {
		return errMutExflags
	}

	return nil
}

// Sort select sorting mode and additional options
func Sort(key *packkey.Key) paklines.Lines {
	var lines paklines.Lines

	lines, err := paklines.GetLines(os.Args[1])
	if err != nil {
		log.Fatal(fmt.Errorf("error when getting lines: %v", err))
	}

	if err := CheckArguments(lines, key); err != nil {
		log.Fatal(fmt.Errorf("invalid arguments: %v", err))
	}

	lines.SetColumn(lines, key.K)

	switch {
	case key.NumericSort:
		sort.Slice(lines, lines.SortNumeric)
	case key.HumanNumericSort:
		sort.Slice(lines, lines.SortHumanNumeric)
	case key.MonthSort:
		sort.Slice(lines, lines.SortMonth)
	default:
		sort.Slice(lines, lines.StandardSort)
	}

	if key.Unique {
		lines = lines.Unique()
	}
	if key.Reverse {
		lines.Reverse()
	}
	if key.IgnoreTailingBlanks {
		lines.IgnoreTailingSpaces()
	}

	return lines
}

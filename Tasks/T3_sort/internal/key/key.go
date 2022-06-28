package key

import (
	"github.com/spf13/cobra"
)

// New return instance of type Key
func New() *Key {
	return &Key{}
}

// Key is a struct for program flags received from cobra
type Key struct {
	K                   int
	NumericSort         bool
	Reverse             bool
	Unique              bool
	MonthSort           bool
	IgnoreTailingBlanks bool
	Check               bool
	HumanNumericSort    bool
}

// SetKeys receives flags and set them to fields of struct Key
func (key *Key) SetKeys(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().BoolP("help", "", false, "help for this command")
	rootCmd.Flags().IntVarP(&key.K, "key", "k", 1, "field")
	rootCmd.Flags().BoolVarP(&key.NumericSort, "numeric-sort", "n", false, "numeric-sort")
	rootCmd.Flags().BoolVarP(&key.Reverse, "reverse", "r", false, "reverse")
	rootCmd.Flags().BoolVarP(&key.Unique, "unique", "u", false, "unique")
	rootCmd.Flags().BoolVarP(&key.MonthSort, "month-sort", "M", false, "month-sort")
	rootCmd.Flags().BoolVarP(&key.IgnoreTailingBlanks, "ignore-tailing-blanks", "b", false, "ignore-tailing-blanks")
	rootCmd.Flags().BoolVarP(&key.Check, "check", "c", false, "check")
	rootCmd.Flags().BoolVarP(&key.HumanNumericSort, "human-numeric-sort", "h", false, "human-numeric-sort")
}

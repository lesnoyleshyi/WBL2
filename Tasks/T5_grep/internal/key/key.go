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
	After      int
	Before     int
	Context    int
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
	Pattern    string
	Filename   string
}

// SetKeys receives flags and set them to fields of struct Key
func (key *Key) SetKeys(rootCmd *cobra.Command) {
	rootCmd.Flags().IntVarP(&key.After, "after-context", "A", 0, "after-context")
	rootCmd.Flags().IntVarP(&key.Before, "before-context", "B", 0, "before-context")
	rootCmd.Flags().IntVarP(&key.Context, "context", "C", 0, "context")
	rootCmd.Flags().BoolVarP(&key.Count, "count", "c", false, "count")
	rootCmd.Flags().BoolVarP(&key.IgnoreCase, "ignore-case", "i", false, "ignore-case")
	rootCmd.Flags().BoolVarP(&key.Invert, "invert-match", "v", false, "invert-match")
	rootCmd.Flags().BoolVarP(&key.Fixed, "fixed-strings", "F", false, "fixed-strings")
	rootCmd.Flags().BoolVarP(&key.LineNum, "line-number", "n", false, "line-number")
}

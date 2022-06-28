package key

import (
	"time"

	"github.com/spf13/cobra"
)

// New return instance of type Key
func New() *Key {
	return &Key{}
}

// Key is a struct for program flags received from cobra
type Key struct {
	Timeout time.Duration
	Host    string
	Port    string
}

// SetKeys receives flags and set them to fields of struct Key
func (key *Key) SetKeys(rootCmd *cobra.Command) {
	rootCmd.Flags().DurationVarP(&key.Timeout, "timeout", "t", 10*time.Second, "fields")
}

package key

import (
	"github.com/spf13/cobra"
)

//New return instance of type Key
func New() *Key {
	return &Key{}
}

//Key is a struct for program flags received from cobra
type Key struct {
	Fields    []string
	Delimiter string
	Separated bool
}

//SetKeys receives flags and set them to fields of struct Key
func (key *Key) SetKeys(rootCmd *cobra.Command) {
	rootCmd.Flags().StringSliceVarP(&key.Fields, "fields", "f", nil, "fields")
	rootCmd.Flags().StringVarP(&key.Delimiter, "delimeter", "d", "\t", "delimiter")
	rootCmd.Flags().BoolVarP(&key.Separated, "separated", "s", false, "separated")
}

package main

import (
	"WBL2/Tasks/T5_grep/internal/grep"
	packkey "WBL2/Tasks/T5_grep/internal/key"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func main() {
	cmd := &cobra.Command{} //nolint:exhaustruct
	key := packkey.New()

	key.SetKeys(cmd)

	if err := cmd.Execute(); err != nil || len(os.Args) < 2 {
		log.Fatal(fmt.Errorf("required argument missing: %w", err))
	}

	out, _ := grep.Grep(key)
	for _, l := range out {
		fmt.Println(l) //nolint:forbidigo
	}
}

package main

import (
	packcut "WBL2/Tasks/T6_cut/internal/cut"
	packkey "WBL2/Tasks/T6_cut/internal/key"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func main() {
	cmd := &cobra.Command{}
	key := packkey.New()
	cut := packcut.New(key)
	key.SetKeys(cmd)

	if err := cmd.Execute(); err != nil || len(os.Args) < 3 {
		log.Fatal(fmt.Errorf("required argument missing: %v", err))
	}

	lines, _ := cut.Cut()
	for _, line := range lines {
		fmt.Println(line)
	}
}

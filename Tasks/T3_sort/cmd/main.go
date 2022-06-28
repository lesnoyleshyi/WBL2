package main

import (
	packkey "WBL2/Tasks/T3_sort/internal/key"
	"WBL2/Tasks/T3_sort/internal/sort"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func main() {
	cmd := &cobra.Command{}
	key := packkey.New()
	key.SetKeys(cmd)

	if err := cmd.Execute(); err != nil || len(os.Args) < 2 {
		log.Fatal(fmt.Errorf("required argument missing: %v", err))
	}

	lines := sort.Sort(key)
	for _, line := range lines {
		fmt.Println(line[1])
	}
}

package hexclient

import (
	"fmt"
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/graph/formats/dot"
)

var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Root command for graph operations",
	Long:  `Usage: nb graph add [filename]`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			fmt.Println("Not enough arguments. Usage: nb graph add [filename]")
		}
	},
}

var graphAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Root command for graph operations",
	Long:  `Usage: nb graph add [filename]`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			fmt.Println("Not enough arguments. Usage: nb graph add [filename]")
		}

		astFile, err := dot.ParseFile(args[0])

		if err != nil {
			fmt.Printf("Error opening file: %v", err)
			return
		}

		fmt.Println("graph: %s", astFile.String())

	},
}

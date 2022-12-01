/*
Copyright © 2022 Robot Motel Limited
*/
package hexgen

import (
	"fmt"
	"github.com/3vilM33pl3/hexrepo/src/hexgen/internal/pkg/hexgraph"
	"strconv"

	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen [s]",
	Short: "generate map of s size",
	Long: `Generate a map of s size. For example:
	hexgen gen 10 
will generate a circular map with a radius of 10 hexagons.

`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen called")

		size, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Size argument not a number: %s", err)
			return
		}

		m := hexgraph.NewMap("test")
		m.Generate(size)
		m.DOT()

	},
}

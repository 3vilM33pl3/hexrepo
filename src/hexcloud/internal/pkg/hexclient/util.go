package hexclient

import (
	"fmt"
	"github.com/3vilM33pl3/hexrepo/src/hexcloud/pkg/hexgrid"
	"github.com/spf13/cobra"
	"strconv"
)

var pairCmd = &cobra.Command{
	Use:   "pair [x] [y]",
	Short: "Pairs two signed integers into a single unsigned integer",
	Long: `Pairs two signed integers into a single unsigned integer. It uses Szudzik's algorithm to pair 
and maps all negative outcomes to positive even numbers and positive to positive odd numbers.
This is to avoid the need for a sign bit in the resulting integer.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Not enough arguments")
		}

		x, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("First argument not a number: %s", err)
			return
		}

		y, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Second argument not a number: %s", err)
			return
		}

		fmt.Printf("%d", hexgrid.Pair(int64(x), int64(y)))

	},
}

var unPairCmd = &cobra.Command{
	Use:   "unpair [s]",
	Short: "Unpairs a single unsigned integer into two signed integers",
	Long:  `A longer description that spans multiple lines and likely contains examples.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Not enough arguments")
		}

		s, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Argument not a number: %s", err)
			return
		}

		x, y := hexgrid.Unpair(int64(s))
		fmt.Printf("%d %d", x, y)

	},
}

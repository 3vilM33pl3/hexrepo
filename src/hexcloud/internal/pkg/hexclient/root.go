/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package hexclient

import (
	"flag"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "nb",
	Short: "Command line interface for game content manipulation",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")

	rootCmd.PersistentFlags().BoolP("secure", "s", true, "secure connection")
	rootCmd.PersistentFlags().StringP("addr", "a", "localhost:8080", "server address")

	rootCmd.AddCommand(repoCmd)
	repoCmd.AddCommand(repoAddCmd)
	repoAddCmd.AddCommand(repoAddFileCmd)
	repoAddCmd.AddCommand(repoAddDataCmd)
	repoAddDataCmd.AddCommand(repoAddDataFileCmd)
	repoCmd.AddCommand(repoGetCmd)
	repoGetCmd.AddCommand(repoGetDataCmd)
	repoCmd.AddCommand(repoDelCmd)
	repoDelCmd.AddCommand(repoDelDataCmd)

	rootCmd.AddCommand(mapCmd)
	mapCmd.AddCommand(mapAddCmd)
	mapAddCmd.AddCommand(mapAddData)
	mapCmd.AddCommand(mapGetCmd)
	mapCmd.AddCommand(mapUpdateCmd)
	mapUpdateCmd.AddCommand(mapUpdateDataCmd)
	mapCmd.AddCommand(mapRemoveCmd)
	mapRemoveCmd.AddCommand(mapRemoveDataCmd)

	mapCmd.PersistentFlags().Int64P("radius", "r", 0, "radius of hexagon circle")

	rootCmd.AddCommand(hexStatusCmd)
	hexStatusCmd.AddCommand(hexStatusServerCmd)
	hexStatusCmd.AddCommand(hexStatusStorageCmd)
	hexStatusCmd.AddCommand(hexStatusClientCmd)

}

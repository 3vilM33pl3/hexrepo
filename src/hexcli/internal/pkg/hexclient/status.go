package hexclient

import (
	"fmt"
	"github.com/spf13/cobra"
)

var hexStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "status of hexagon network",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside hex status command: %v\n", args)

	},
}

var hexStatusServerCmd = &cobra.Command{
	Use:   "server",
	Short: "status of hexagon server",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)
		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		}

		status, err := client.StatusServer()
		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		} else {
			fmt.Printf("Status server: %s\n", status)
		}

	},
}

var hexStatusStorageCmd = &cobra.Command{
	Use:   "storage",
	Short: "status of hexagon storage",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr := "localhost:8080"
		if len(args) > 0 {
			serverAddr = args[0]
		}

		secure, _ := cmd.Flags().GetBool("secure")
		client, err := NewClient(serverAddr, secure)

		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		}

		status, err := client.StatusStorage()
		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		} else {
			fmt.Printf("Status storage: %s\n", status)
		}
	},
}

var hexStatusClientCmd = &cobra.Command{
	Use:   "client",
	Short: "status of hexagon network clients",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr := "localhost:8080"
		if len(args) > 0 {
			serverAddr = args[0]
		}

		secure, _ := cmd.Flags().GetBool("secure")
		client, err := NewClient(serverAddr, secure)

		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		}

		status, err := client.StatusClients()
		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		} else {
			fmt.Printf("Status clients: %s\n", status)
		}

	},
}

package hexclient

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"hexcloud/internal/pkg/hexcloud"
	"log"
	"os"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Manipulate hexagons",
	Long:  `A longer description that spans multiple lines and likely contains examples.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("repo called")
	},
}

var repoAddCmd = &cobra.Command{
	Use:   "add [ref]",
	Short: "add hexagon to repository with reference [ref] and [exits]",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)
		var infoList hexcloud.HexInfoList

		ref := args[0]
		if err != nil {
			fmt.Printf("Error parsing exits %s", err)
		}
		infoList.HexInfo = append(infoList.HexInfo, &hexcloud.HexInfo{ID: ref})

		err = client.RepoAddHexagonInfo(&infoList)
		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		}

	},
}

var repoAddFileCmd = &cobra.Command{
	Use:   "file [filename]",
	Short: "add hexagon from file [filename] to repository",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)
		var infoList hexcloud.HexInfoList

		f, err := os.Open(args[0])
		if err != nil {
			fmt.Printf("Error opening file %s", args[0])
			return
		}
		rc := bufio.NewReader(f)
		csvLines, err := csv.NewReader(rc).ReadAll()
		if err != nil {
			log.Printf("Error reading hexdata file: %v", err)
			return
		}

		for _, line := range csvLines {
			hexInfo := &hexcloud.HexInfo{
				ID: line[0],
			}
			infoList.HexInfo = append(infoList.HexInfo, hexInfo)
		}

		err = client.RepoAddHexagonInfo(&infoList)
		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		}

	},
}

var repoAddDataFileCmd = &cobra.Command{
	Use:   "file [filename]",
	Short: "add hexagon from file [filename] to repository",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)
		var infoList hexcloud.HexInfoList

		f, err := os.Open(args[0])
		if err != nil {
			fmt.Printf("Error opening file %s", args[0])
			return
		}
		rc := bufio.NewReader(f)
		csvLines, err := csv.NewReader(rc).ReadAll()
		if err != nil {
			log.Printf("Error reading hexdata file: %v", err)
			return
		}

		for _, line := range csvLines {
			data := make(map[string]string)
			data[line[1]] = line[2]

			hexInfo := &hexcloud.HexInfo{
				ID:   line[0],
				Data: data,
			}
			infoList.HexInfo = append(infoList.HexInfo, hexInfo)
		}

		err = client.RepoAddHexagonInfo(&infoList)
		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		}

	},
}

var repoAddDataCmd = &cobra.Command{
	Use:   "data [ref] [key] [value]",
	Short: "add data hexagon",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)
		var infoList hexcloud.HexInfoList

		if err != nil {
			fmt.Printf("Error parsing exits %s", err)
			return
		}

		if len(args) < 3 {
			fmt.Println("Not enough arguments")
		}

		ref := args[0]
		key := args[1]
		value := args[2]
		hexInfo := &hexcloud.HexInfo{ID: ref}
		hexInfo.Data = make(map[string]string)
		hexInfo.Data[key] = value
		infoList.HexInfo = append(infoList.HexInfo, hexInfo)

		err = client.RepoAddHexagonInfo(&infoList)
		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		}

	},
}

var repoDelCmd = &cobra.Command{
	Use:   "del [ref]",
	Short: "delete hexagon from repository with reference [ref] ",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)
		var infoList hexcloud.HexIDList

		for _, id := range args {
			infoList.HexID = append(infoList.HexID, id)
		}

		err = client.RepoDeleteHexagonInfo(&infoList)

		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		}
	},
}

var repoDelDataCmd = &cobra.Command{
	Use:   "data [ref]",
	Short: "delete hexagon data from repository with reference [ref] and [key] ",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)
		var hexIDData hexcloud.HexIDData

		if len(args) < 2 {
			fmt.Printf("Not enough arguments\n")
		}
		hexIDData.HexID = args[0]
		hexIDData.DataKey = args[1]

		err = client.RepoDeleteHexagonInfoData(&hexIDData)

		if err != nil {
			fmt.Printf("Error deleting %s", err)
			return
		}
	},
}

var repoGetCmd = &cobra.Command{
	Use:   "get [ref]",
	Short: "get hexagon info from repository with reference [ref] ",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)
		var infoList hexcloud.HexIDList

		infoList.HexID = append(infoList.HexID, args[0])

		result, err := client.RepoGetHexagonInfo(&infoList)

		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		}

		for _, hex := range result.HexInfo {
			fmt.Printf("%s \n", hex.ID)
			for key, value := range hex.Data {
				fmt.Printf("    %s %s\n", key, value)
			}
		}
	},
}

var repoGetDataCmd = &cobra.Command{
	Use:   "data [ref] [key]",
	Short: "get hexagon data value from repository with reference [ref] and [key] ",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)
		var hexIDData hexcloud.HexIDData

		if len(args) < 2 {
			fmt.Printf("Not enough arguments\n")
		}
		hexIDData.HexID = args[0]
		hexIDData.DataKey = args[1]

		result, err := client.RepoAddHexagonInfoData(&hexIDData)

		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		}

		fmt.Printf("    %s %s %s\n", result.HexID, result.DataKey, result.Value)
	},
}

var repoGetAllCmd = &cobra.Command{
	Use:   "all",
	Short: "get all hexagon info from repository",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)

		result, err := client.RepoGetAllHexagonInfo()

		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		}

		for _, hex := range result.HexInfo {
			fmt.Printf("%s \n", hex.ID)
		}
	},
}

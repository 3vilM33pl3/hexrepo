package hexclient

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"hexcloud/internal/pkg/hexcloud"
	"strconv"
)

var mapCmd = &cobra.Command{
	Use:   "map",
	Short: "Manipulate map",
	Long:  `Functionality for adding, updating, removing everything map related`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hex called")
	},
}

var mapAddCmd = &cobra.Command{
	Use:   "add [x,y,z] [content ref]",
	Short: "add hexagon with coordinate [x,y,z] and content reference [content ref]",
	Long: "Example: nb hex place --secure=false -- 0 -5 5 N 0000-0000-0000-0000" +
		"The double dash is needed to indicate no more flags are coming and everything is interpreted as an normal argument. " +
		"This is needed so that negative numbers don't get interpreted as flags",
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)

		if len(args) < 4 {
			fmt.Printf("Not enough arguments")
			return
		}
		x, y, z, err := extractHexCoord(args)
		if err != nil {
			return
		}

		hex := &hexcloud.HexLocation{
			X:     x,
			Y:     y,
			Z:     z,
			HexID: args[3],
		}

		hexLocList := &hexcloud.HexLocationList{}
		hexLocList.HexLoc = append(hexLocList.HexLoc, hex)

		err = client.MapAdd(hexLocList)

		if err != nil {
			fmt.Printf("Error placing hexagon on map: %s", err)
			return
		}

	},
}

var mapAddData = &cobra.Command{
	Use:   "data [0,0,0] [key] [value]",
	Short: "add data to map",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)

		x, y, z, err := extractHexCoord(args)

		hexLocData := hexcloud.HexLocData{
			X:       x,
			Y:       y,
			Z:       z,
			DataKey: args[3],
			Value:   args[4],
		}

		err = client.MapAddData(&hexLocData)

		if err != nil {
			glog.Errorf("%s\n", err)
		}

	},
}

var mapGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get hexagon at `coord` [x,y,z]",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")
		radius, _ := cmd.Flags().GetInt64("radius")

		client, err := NewClient(serverAddr, secure)

		x, y, z, err := extractHexCoord(args)
		if err != nil {
			return
		}

		result, err := client.MapGet(&hexcloud.HexLocation{
			X: x,
			Y: y,
			Z: z,
		}, radius)

		if err != nil {
			fmt.Printf("Error retrieving hexagon information on map: %s", err)
			return
		}

		for _, hex := range result.HexLoc {
			fmt.Println("--Coordinates & ID--")
			fmt.Printf("%d %d %d %s\n", hex.X, hex.Y, hex.Z, hex.HexID)
			fmt.Println("--global--")
			for key, value := range hex.GlobalData {
				fmt.Printf("'%s' - '%s'\n", key, value)
			}
			fmt.Println("--local--")
			for key, value := range hex.LocalData {
				fmt.Printf("'%s' - '%s'\n", key, value)
			}
		}

	},
}

var mapRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove hexagon at `coord` [x,y,z]",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)
		var hexList hexcloud.HexLocationList

		x, y, z, err := extractHexCoord(args)
		if err != nil {
			return
		}

		hex := &hexcloud.HexLocation{
			X:     x,
			Y:     y,
			Z:     z,
			HexID: "",
		}

		hexList.HexLoc = append(hexList.HexLoc, hex)
		err = client.MapRemove(&hexList)

	},
}

var mapRemoveDataCmd = &cobra.Command{
	Use:   "data",
	Short: "remove hexagon data at `coord` [x,y,z] and [key]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			fmt.Println("Not enough arguments")
			return
		}

		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)

		x, y, z, err := extractHexCoord(args)
		if err != nil {
			return
		}

		data := make(map[string]string)
		data[args[3]] = ""
		hex := &hexcloud.HexLocation{
			X:         x,
			Y:         y,
			Z:         z,
			HexID:     "",
			LocalData: data,
		}

		err = client.MapRemoveData(hex)

	},
}

var mapUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "update hexagon from map at `coord` [x,y,z]",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)
		var hexList hexcloud.HexLocationList

		x, y, z, err := extractHexCoord(args)
		if err != nil {
			return
		}

		hexId := args[3]

		hex := &hexcloud.HexLocation{
			X:     x,
			Y:     y,
			Z:     z,
			HexID: hexId,
		}

		hexList.HexLoc = append(hexList.HexLoc, hex)
		_, err = client.MapUpdate(hex)

	},
}

var mapUpdateDataCmd = &cobra.Command{
	Use:   "data",
	Short: "update hexagon from map at `coord` [x,y,z]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 4 {
			fmt.Println("Not enough arguments")
			return
		}

		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)
		var hexList hexcloud.HexLocationList

		x, y, z, err := extractHexCoord(args)
		if err != nil {
			return
		}

		data := make(map[string]string)
		data[args[3]] = args[4]
		hex := &hexcloud.HexLocation{
			X:         x,
			Y:         y,
			Z:         z,
			HexID:     "",
			LocalData: data,
		}

		hexList.HexLoc = append(hexList.HexLoc, hex)
		_, err = client.MapUpdate(hex)

	},
}

func extractHexCoord(args []string) (x int64, y int64, z int64, err error) {

	x, err = strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		fmt.Printf("x value is not a number %s : %e", args[0], err)
		return x, y, z, err
	}

	y, err = strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		fmt.Printf("y value is not a number %s : %e", args[1], err)
		return x, y, z, err
	}

	z, err = strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		fmt.Printf("z value is not a number %s : %e", args[2], err)
		return x, y, z, err
	}

	return x, y, z, nil
}

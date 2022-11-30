package hexclient

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"hexcloud/internal/pkg/hexcloud"
	"hexcloud/pkg/hexgrid"
	"log"
	"os"
	"strconv"
	"strings"
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

var mapAddFileCmd = &cobra.Command{
	Use:   "file",
	Short: "add hexagons included in the file",
	Long: "Example: nb hex add file map.csv" +
		"csv format:" +
		"-2,2,0,0000-0000-0000-0001" +
		"-1,2,-1,0000-0000-0000-0001",
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)

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

		hexLocList := &hexcloud.HexLocationList{}

		for _, line := range csvLines {
			x, err := strconv.ParseInt(line[0], 10, 64)
			if err != nil {
				glog.Errorf("x: %s\n", err)
				return
			}

			y, err := strconv.ParseInt(line[1], 10, 64)
			if err != nil {
				glog.Errorf("y: %s\n", err)
				return
			}

			z := -x - y

			hex := &hexcloud.HexLocation{
				X:     x,
				Y:     y,
				Z:     z,
				HexID: line[3],
			}

			hexLocList.HexLoc = append(hexLocList.HexLoc, hex)

		}

		err = client.MapAdd(hexLocList)

		if err != nil {
			fmt.Printf("Error placing hexagon on map: %s", err)
			return
		}

	},
}

var mapAddDataCmd = &cobra.Command{
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

		if len(args) > 4 {
			xi, err := strconv.ParseInt(args[4], 10, 64)
			if err == nil {
				yi, err := strconv.ParseInt(args[4], 10, 64)
				if err == nil {
					hexLocData.Value = strconv.FormatInt(hexgrid.Pair(xi, yi), 10)
				}
			}
		}

		hexLocDataList := &hexcloud.HexLocDataList{}
		hexLocDataList.HexLocData = append(hexLocDataList.HexLocData, &hexLocData)

		err = client.MapAddData(hexLocDataList)

		if err != nil {
			glog.Errorf("%s\n", err)
		}

	},
}

var mapAddDataFileCmd = &cobra.Command{
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

		hexLocDataList := &hexcloud.HexLocDataList{}

		for _, line := range csvLines {
			x, err := strconv.ParseInt(line[0], 10, 64)
			if err != nil {
				glog.Errorf("x: %s\n", err)
				return
			}

			y, err := strconv.ParseInt(line[1], 10, 64)
			if err != nil {
				glog.Errorf("y: %s\n", err)
				return
			}

			z := -x - y

			hexLocData := hexcloud.HexLocData{
				X:       x,
				Y:       y,
				Z:       z,
				DataKey: line[2],
				Value:   line[3],
			}

			if hexLocData.DataKey == "link" {
				coords := strings.Split(hexLocData.Value, ":")
				xi, err := strconv.ParseInt(coords[0], 10, 64)
				if err == nil {
					yi, err := strconv.ParseInt(coords[1], 10, 64)
					if err == nil {
						hexLocData.Value = strconv.FormatInt(hexgrid.Pair(xi, yi), 10)
					}
				}
			}

			hexLocDataList.HexLocData = append(hexLocDataList.HexLocData, &hexLocData)

		}

		err = client.MapAddData(hexLocDataList)

		if err != nil {
			glog.Errorf("%s\n", err)
			return
		}

		err = client.RepoAddHexagonInfo(&infoList)
		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
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

var mapRemoveAllCmd = &cobra.Command{
	Use:   "all",
	Short: "remove all hexagon from map",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)
		if err != nil {
			return
		}

		err = client.MapRemoveAll()

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

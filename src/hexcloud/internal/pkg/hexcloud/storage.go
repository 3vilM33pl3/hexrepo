package hexcloud

import (
	"context"
	"github.com/golang/glog"
)

type HexStorage interface {
	AddHexagonToRepo(hexInfo *HexInfo)
	GetHexagonInfoAll() (hexInfoList *HexInfoList)
	GetHexagonInfo(hexID string) (hexInfo *HexInfo)
	AddHexagonToMap(hexLocation *HexLocation)
	GetHexagonFromMap(x int64, y int64) (hexLocation *HexLocation)
	DeleteHexagonFromMap(hexLocation *HexLocation)
	DeleteHexagonFromRepo(hexID string)
	SizeMap() (count int)
	SizeRepo() (count int)
	DeleteHexagonDataFromRepo(id string, key string)
	GetHexagonInfoData(id string, key string) (hexIDData *HexIDData)
	AddDataToMap(data *HexLocData) (err error)
	MapUpdate(data *HexLocation) (err error)
	MapRemove(data *HexLocation) (err error)
	MapRemoveData(data *HexLocation) (err error)
	DeleteAllHexagonsFromMap() (err error)
	DeleteAllHexagonsFromRepo() (err error)
}

func NewHexStorage(newdb bool, dbName string, example bool, remote bool) (storage HexStorage, err error) {
	var hs HexStorage

	if remote {
		ctx := context.Background()
		hs, err = NewFirestoreClient(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		hst := &HexStorageSQLite{}
		hst.Local = true
		hst.Database, err = InitialiseDatabase(newdb, dbName, example)
		if err != nil {
			glog.Fatalf("Error initializing database: %s", err)
		}
		hs = hst
	}

	return hs, nil
}

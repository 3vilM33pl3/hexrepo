package hexcloud

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/golang/glog"
	"log"
	"runtime"
)

func UnimplementedStorageControllerFunction() {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	fmt.Printf("Unimplemented function: %s:%d %s\n", frame.File, frame.Line, frame.Function)
}

type HexStorageFirestore struct {
	local   bool
	client  *firestore.Client
	hexMap  *firestore.CollectionRef
	hexRepo *firestore.CollectionRef
}

func (h *HexStorageFirestore) DeleteAllHexagonsFromMap() (err error) {
	iter := h.hexMap.Documents(context.Background())
	numDeleted := 0

	for {
		doc, err := iter.Next()
		if err != nil {
			return err
		}
		_, err = doc.Ref.Delete(context.Background())
		if err != nil {
			return err
		}
		numDeleted++
	}

	glog.Infof("Deleted %d documents from map", numDeleted)

	return nil
}

func (h *HexStorageFirestore) DeleteAllHexagonsFromRepo() (err error) {
	iter := h.hexRepo.Documents(context.Background())
	numDeleted := 0

	for {
		doc, err := iter.Next()
		if err != nil {
			return err
		}
		_, err = doc.Ref.Delete(context.Background())
		if err != nil {
			return err
		}
		numDeleted++
	}

	glog.Infof("Deleted %d documents from hexagon repo", numDeleted)

	return nil
}

func NewFirestoreClient(ctx context.Context) (storage *HexStorageFirestore, err error) {
	projectID := "robot-motel"

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil, err
	}

	hexMap := client.Collection("map")
	hexRepo := client.Collection("hexrepo")

	storage = &HexStorageFirestore{
		local:   false,
		client:  client,
		hexMap:  hexMap,
		hexRepo: hexRepo,
	}

	return storage, nil
}

func castToInterface(in map[string]string) (out map[string]interface{}) {
	out = make(map[string]interface{}, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

func castFromInterface(in map[string]interface{}) (out map[string]string) {
	out = make(map[string]string, len(in))
	for k, v := range in {
		out[k] = v.(string)
	}
	return out
}

func (h *HexStorageFirestore) AddHexagonToRepo(hexInfo *HexInfo) {
	ctx := context.Background()

	if hexInfo.ID != "" {
		result, err := h.hexRepo.Doc(hexInfo.ID).Set(ctx, castToInterface(hexInfo.GetData()))
		if err != nil {
			glog.Errorf("%s\n%v", result, err)
		}
	}
}

func (h *HexStorageFirestore) GetHexagonInfoAll() (hexInfoList *HexInfoList) {
	UnimplementedStorageControllerFunction()

	return
}

func (h *HexStorageFirestore) GetHexagonInfo(hexID string) (hexInfo *HexInfo) {
	ctx := context.Background()

	docRef := h.hexRepo.Doc(hexID)
	doc, err := docRef.Get(ctx)

	if err != nil {
		glog.Errorf("%v", err)
		return nil
	}

	hexInfo = newHexInfo(hexID, doc)

	return
}

func newHexInfo(id string, doc *firestore.DocumentSnapshot) *HexInfo {
	hexInfo := &HexInfo{
		ID: id,
	}
	hexInfo.Data = castFromInterface(doc.Data())

	return hexInfo
}

func (h *HexStorageFirestore) AddHexagonToMap(hexLocation *HexLocation) {
	ctx := context.Background()

	id := Pair(hexLocation.X, hexLocation.Y)
	result, err := h.hexMap.Doc(fmt.Sprintf("%d", id)).Set(ctx, map[string]interface{}{
		"x":          hexLocation.X,
		"y":          hexLocation.Y,
		"z":          hexLocation.Z,
		"hexID":      hexLocation.HexID,
		"localData":  castToInterface(hexLocation.LocalData),
		"globalData": castToInterface(hexLocation.GlobalData),
	})
	if err != nil {
		glog.Errorf("%s\n%v", result, err)
	}
}

func (h *HexStorageFirestore) GetHexagonFromMap(x int64, y int64) (hexLocation *HexLocation) {
	ctx := context.Background()

	id := Pair(x, y)
	docRef := h.hexMap.Doc(fmt.Sprintf("%d", id))
	doc, err := docRef.Get(ctx)

	if err != nil {
		glog.Errorf("%v", err)
		return nil
	}

	hexLocation = newHexLocation(id, doc)

	return hexLocation
}

func newHexLocation(id int64, doc *firestore.DocumentSnapshot) *HexLocation {
	x, y := Unpair(id)
	hexLocation := &HexLocation{
		X:         x,
		Y:         y,
		Z:         x - y,
		HexID:     doc.Data()["hexID"].(string),
		LocalData: castFromInterface(doc.Data()["localData"].(map[string]interface{})),
	}

	return hexLocation

}

func (h *HexStorageFirestore) DeleteHexagonFromMap(hexLocation *HexLocation) {
	ctx := context.Background()
	id := Pair(hexLocation.X, hexLocation.Y)

	result, err := h.hexMap.Doc(fmt.Sprintf("%d", id)).Delete(ctx)
	if err != nil {
		glog.Errorf("%s\n%v", result, err)
	}

	return

}

func (h *HexStorageFirestore) DeleteHexagonFromRepo(hexID string) {
	ctx := context.Background()

	_, err := h.hexRepo.Doc(hexID).Delete(ctx)
	if err != nil {
		glog.Errorf("%v", err)
	}

}

func (h *HexStorageFirestore) SizeMap() (count int) {

	return
}

func (h *HexStorageFirestore) SizeRepo() (count int) {
	UnimplementedStorageControllerFunction()

	return
}

func (h *HexStorageFirestore) DeleteHexagonDataFromRepo(hexID string, key string) {
	ctx := context.Background()

	result, err := h.hexRepo.Doc(hexID).Update(ctx, []firestore.Update{
		{Path: key, Value: firestore.Delete},
	})

	if err != nil {
		glog.Errorf("%s\n%v", result, err)
	}

}

func (h *HexStorageFirestore) GetHexagonInfoData(hexID string, key string) (hexIDData *HexIDData) {
	ctx := context.Background()

	docRef := h.hexRepo.Doc(hexID)
	doc, err := docRef.Get(ctx)

	if err != nil {
		glog.Errorf("%v", err)
		return nil
	}

	hexInfo := newHexInfo(hexID, doc)
	hexIDData = &HexIDData{
		HexID:   hexID,
		DataKey: key,
		Value:   hexInfo.Data[key],
	}

	return hexIDData
}

func (h *HexStorageFirestore) AddDataToMap(data *HexLocData) (err error) {
	ctx := context.Background()

	id := Pair(data.X, data.Y)

	result, err := h.hexMap.Doc(fmt.Sprintf("%d", id)).Update(ctx, []firestore.Update{
		{
			Path:  "localData." + data.DataKey,
			Value: data.Value,
		},
	})

	if err != nil {
		glog.Errorf("%s\n%v", result, err)
	}

	return
}

func (h *HexStorageFirestore) MapUpdate(hexLocation *HexLocation) (err error) {
	ctx := context.Background()

	id := Pair(hexLocation.X, hexLocation.Y)
	result, err := h.hexMap.Doc(fmt.Sprintf("%d", id)).Set(ctx, map[string]interface{}{
		"x":          hexLocation.X,
		"y":          hexLocation.Y,
		"hexID":      hexLocation.HexID,
		"localData":  castToInterface(hexLocation.LocalData),
		"globalData": castToInterface(hexLocation.GlobalData),
	}, firestore.MergeAll)

	if err != nil {
		glog.Errorf("%s\n%v", result, err)
	}

	return
}

func (h *HexStorageFirestore) MapRemove(data *HexLocation) (err error) {
	ctx := context.Background()
	id := Pair(data.X, data.Y)

	result, err := h.hexMap.Doc(fmt.Sprintf("%d", id)).Delete(ctx)
	if err != nil {
		glog.Errorf("%s\n%v", result, err)
	}

	return
}

func (h *HexStorageFirestore) MapRemoveData(hexLocation *HexLocation) (err error) {
	ctx := context.Background()

	hexID := Pair(hexLocation.X, hexLocation.Y)

	for key, _ := range hexLocation.LocalData {
		result, err := h.hexMap.Doc(fmt.Sprintf("%d", hexID)).Update(ctx, []firestore.Update{
			{Path: "localData." + key, Value: firestore.Delete},
		})

		if err != nil {
			glog.Errorf("%s\n%v", result, err)
			return err
		}
	}

	return err
}

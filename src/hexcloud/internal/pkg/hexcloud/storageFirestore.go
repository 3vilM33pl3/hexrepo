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

func biject(a int64) (A uint64) {
	if a >= 0 {
		A = uint64(2 * a)
	} else {
		A = uint64(-2*a - 1)
	}
	return
}

func elegantPair(x, y int64) (result uint64) {
	a := biject(x)
	b := biject(y)

	if a >= b {
		result = a*a + a + b
	} else {
		result = a + b*b
	}
	return
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
	UnimplementedStorageControllerFunction()

}

func (h *HexStorageFirestore) GetHexagonFromMap(x int64, y int64) (hexLocation *HexLocation) {
	UnimplementedStorageControllerFunction()

	return
}

func (h *HexStorageFirestore) DeleteHexagonFromMap(hexLocation *HexLocation) {
	UnimplementedStorageControllerFunction()

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

func (h *HexStorageFirestore) DeleteHexagonDataFromRepo(id string, key string) {
	UnimplementedStorageControllerFunction()

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
	UnimplementedStorageControllerFunction()

	return
}

func (h *HexStorageFirestore) MapUpdate(data *HexLocation) (err error) {
	UnimplementedStorageControllerFunction()

	return
}

func (h *HexStorageFirestore) MapRemove(data *HexLocation) (err error) {
	UnimplementedStorageControllerFunction()

	return
}

func (h *HexStorageFirestore) MapRemoveData(data *HexLocation) (err error) {
	UnimplementedStorageControllerFunction()

	return
}

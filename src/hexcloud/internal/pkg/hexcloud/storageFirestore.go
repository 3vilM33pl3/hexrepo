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

func cast(in map[string]string) (out map[string]interface{}) {
	out = make(map[string]interface{}, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

func (h *HexStorageFirestore) AddHexagonToRepo(hexInfo *HexInfo) {
	ctx := context.Background()

	if hexInfo.ID != "" {
		result, err := h.hexRepo.Doc(hexInfo.ID).Set(ctx, cast(hexInfo.GetData()))
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
	UnimplementedStorageControllerFunction()

	return
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
	UnimplementedStorageControllerFunction()
}

func (h *HexStorageFirestore) SizeMap() (count int) {
	UnimplementedStorageControllerFunction()

	return
}

func (h *HexStorageFirestore) SizeRepo() (count int) {
	UnimplementedStorageControllerFunction()

	return
}

func (h *HexStorageFirestore) DeleteHexagonDataFromRepo(id string, key string) {
	UnimplementedStorageControllerFunction()

}

func (h *HexStorageFirestore) GetHexagonInfoData(id string, key string) (hexIDData *HexIDData) {
	UnimplementedStorageControllerFunction()

	return
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

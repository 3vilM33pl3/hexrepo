
…
api/hexagon.proto"
Empty"Ã
HexLocation
X (RX
Y (RY
Z (RZ
HexID (	RHexID9
	LocalData (2.HexLocation.LocalDataEntryR	LocalData<

GlobalData (2.HexLocation.GlobalDataEntryR
GlobalData<
LocalDataEntry
key (	Rkey
value (	Rvalue:8=
GlobalDataEntry
key (	Rkey
value (	Rvalue:8"z
HexInfo
ID (	RID&
Data (2.HexInfo.DataEntryRData7
	DataEntry
key (	Rkey
value (	Rvalue:8"1
HexInfoList"
hexInfo (2.HexInfoRhexInfo"7
HexLocationList$
hexLoc (2.HexLocationRhexLoc"e
HexagonGetRequest$
hexLoc (2.HexLocationRhexLoc
radius (Rradius
fill (Rfill"!
	HexIDList
HexID (	RHexID"Q
	HexIDData
HexID (	RHexID
dataKey (	RdataKey
value (	Rvalue"f

HexLocData
X (RX
Y (RY
Z (RZ
dataKey (	RdataKey
value (	Rvalue"
Status
msg (	Rmsg""
Result
success (Rsuccess2«
HexagonService+
RepoAddHexagonInfo.HexInfoList.Result.
RepoGetHexagonInfo
.HexIDList.HexInfoList0
RepoGetHexagonInfoData
.HexIDData
.HexIDData-
RepoGetAllHexagonInfo.Empty.HexInfoList)
RepoDelHexagonInfo
.HexIDList.Result-
RepoDelHexagonInfoData
.HexIDData.Result#
MapAdd.HexLocationList.Result"

MapAddData.HexLocData.Result.
MapGet.HexagonGetRequest.HexLocationList"
	MapUpdate.HexLocation.Result&
MapUpdateData.HexLocation.Result&
	MapRemove.HexLocationList.Result&
MapRemoveData.HexLocation.Result"
GetStatusServer.Empty.Status#
GetStatusStorage.Empty.Status#
GetStatusClients.Empty.Statusbproto3
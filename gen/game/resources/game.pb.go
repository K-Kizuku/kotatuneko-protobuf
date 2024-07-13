// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: game/resources/game.proto

package resources

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GameState int32

const (
	GameState_GAME_STATE_UNKNOWN GameState = 0
	GameState_GAME_STATE_WAIT    GameState = 1
	GameState_GAME_STATE_PLAY    GameState = 2
	GameState_GAME_STATE_EXIT    GameState = 3
	GameState_GAME_STATE_FINISH  GameState = 4
	GameState_GAME_STATE_RESULT  GameState = 5
)

// Enum value maps for GameState.
var (
	GameState_name = map[int32]string{
		0: "GAME_STATE_UNKNOWN",
		1: "GAME_STATE_WAIT",
		2: "GAME_STATE_PLAY",
		3: "GAME_STATE_EXIT",
		4: "GAME_STATE_FINISH",
		5: "GAME_STATE_RESULT",
	}
	GameState_value = map[string]int32{
		"GAME_STATE_UNKNOWN": 0,
		"GAME_STATE_WAIT":    1,
		"GAME_STATE_PLAY":    2,
		"GAME_STATE_EXIT":    3,
		"GAME_STATE_FINISH":  4,
		"GAME_STATE_RESULT":  5,
	}
)

func (x GameState) Enum() *GameState {
	p := new(GameState)
	*p = x
	return p
}

func (x GameState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameState) Descriptor() protoreflect.EnumDescriptor {
	return file_game_resources_game_proto_enumTypes[0].Descriptor()
}

func (GameState) Type() protoreflect.EnumType {
	return &file_game_resources_game_proto_enumTypes[0]
}

func (x GameState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameState.Descriptor instead.
func (GameState) EnumDescriptor() ([]byte, []int) {
	return file_game_resources_game_proto_rawDescGZIP(), []int{0}
}

type HandState int32

const (
	HandState_HAND_STATE_UNKNOWN  HandState = 0
	HandState_HAND_STATE_HOLDING  HandState = 1
	HandState_HAND_STATE_OPENING  HandState = 2
	HandState_HAND_STATE_THROWING HandState = 3
)

// Enum value maps for HandState.
var (
	HandState_name = map[int32]string{
		0: "HAND_STATE_UNKNOWN",
		1: "HAND_STATE_HOLDING",
		2: "HAND_STATE_OPENING",
		3: "HAND_STATE_THROWING",
	}
	HandState_value = map[string]int32{
		"HAND_STATE_UNKNOWN":  0,
		"HAND_STATE_HOLDING":  1,
		"HAND_STATE_OPENING":  2,
		"HAND_STATE_THROWING": 3,
	}
)

func (x HandState) Enum() *HandState {
	p := new(HandState)
	*p = x
	return p
}

func (x HandState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HandState) Descriptor() protoreflect.EnumDescriptor {
	return file_game_resources_game_proto_enumTypes[1].Descriptor()
}

func (HandState) Type() protoreflect.EnumType {
	return &file_game_resources_game_proto_enumTypes[1]
}

func (x HandState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HandState.Descriptor instead.
func (HandState) EnumDescriptor() ([]byte, []int) {
	return file_game_resources_game_proto_rawDescGZIP(), []int{1}
}

type ObjectState int32

const (
	ObjectState_OBJECT_STATE_UNKNOWN   ObjectState = 0
	ObjectState_OBJECT_STATE_QUIESCENT ObjectState = 1
	ObjectState_OBJECT_STATE_MOVING    ObjectState = 2
	ObjectState_OBJECT_STATE_FORCING   ObjectState = 3
)

// Enum value maps for ObjectState.
var (
	ObjectState_name = map[int32]string{
		0: "OBJECT_STATE_UNKNOWN",
		1: "OBJECT_STATE_QUIESCENT",
		2: "OBJECT_STATE_MOVING",
		3: "OBJECT_STATE_FORCING",
	}
	ObjectState_value = map[string]int32{
		"OBJECT_STATE_UNKNOWN":   0,
		"OBJECT_STATE_QUIESCENT": 1,
		"OBJECT_STATE_MOVING":    2,
		"OBJECT_STATE_FORCING":   3,
	}
)

func (x ObjectState) Enum() *ObjectState {
	p := new(ObjectState)
	*p = x
	return p
}

func (x ObjectState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ObjectState) Descriptor() protoreflect.EnumDescriptor {
	return file_game_resources_game_proto_enumTypes[2].Descriptor()
}

func (ObjectState) Type() protoreflect.EnumType {
	return &file_game_resources_game_proto_enumTypes[2]
}

func (x ObjectState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ObjectState.Descriptor instead.
func (ObjectState) EnumDescriptor() ([]byte, []int) {
	return file_game_resources_game_proto_rawDescGZIP(), []int{2}
}

type ObjectKind int32

const (
	ObjectKind_OBJECT_KIND_UNKNOWN     ObjectKind = 0
	ObjectKind_OBJECT_KIND_RECTANGULAR ObjectKind = 1
)

// Enum value maps for ObjectKind.
var (
	ObjectKind_name = map[int32]string{
		0: "OBJECT_KIND_UNKNOWN",
		1: "OBJECT_KIND_RECTANGULAR",
	}
	ObjectKind_value = map[string]int32{
		"OBJECT_KIND_UNKNOWN":     0,
		"OBJECT_KIND_RECTANGULAR": 1,
	}
)

func (x ObjectKind) Enum() *ObjectKind {
	p := new(ObjectKind)
	*p = x
	return p
}

func (x ObjectKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ObjectKind) Descriptor() protoreflect.EnumDescriptor {
	return file_game_resources_game_proto_enumTypes[3].Descriptor()
}

func (ObjectKind) Type() protoreflect.EnumType {
	return &file_game_resources_game_proto_enumTypes[3]
}

func (x ObjectKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ObjectKind.Descriptor instead.
func (ObjectKind) EnumDescriptor() ([]byte, []int) {
	return file_game_resources_game_proto_rawDescGZIP(), []int{3}
}

type Vector3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z float32 `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *Vector3) Reset() {
	*x = Vector3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_resources_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vector3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vector3) ProtoMessage() {}

func (x *Vector3) ProtoReflect() protoreflect.Message {
	mi := &file_game_resources_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vector3.ProtoReflect.Descriptor instead.
func (*Vector3) Descriptor() ([]byte, []int) {
	return file_game_resources_game_proto_rawDescGZIP(), []int{0}
}

func (x *Vector3) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vector3) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Vector3) GetZ() float32 {
	if x != nil {
		return x.Z
	}
	return 0
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Color    string `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_resources_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_game_resources_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_game_resources_game_proto_rawDescGZIP(), []int{1}
}

func (x *Player) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Player) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId string      `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	Layer    int32       `protobuf:"varint,9,opt,name=layer,proto3" json:"layer,omitempty"`
	Kinds    ObjectKind  `protobuf:"varint,2,opt,name=kinds,proto3,enum=game.resources.ObjectKind" json:"kinds,omitempty"`
	State    ObjectState `protobuf:"varint,3,opt,name=state,proto3,enum=game.resources.ObjectState" json:"state,omitempty"`
	Position *Vector3    `protobuf:"bytes,4,opt,name=position,proto3" json:"position,omitempty"`
	Size     *Vector3    `protobuf:"bytes,8,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_resources_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_game_resources_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_game_resources_game_proto_rawDescGZIP(), []int{2}
}

func (x *Object) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *Object) GetLayer() int32 {
	if x != nil {
		return x.Layer
	}
	return 0
}

func (x *Object) GetKinds() ObjectKind {
	if x != nil {
		return x.Kinds
	}
	return ObjectKind_OBJECT_KIND_UNKNOWN
}

func (x *Object) GetState() ObjectState {
	if x != nil {
		return x.State
	}
	return ObjectState_OBJECT_STATE_UNKNOWN
}

func (x *Object) GetPosition() *Vector3 {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *Object) GetSize() *Vector3 {
	if x != nil {
		return x.Size
	}
	return nil
}

type Stat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId              string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Score                 int32  `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	Rank                  int32  `protobuf:"varint,3,opt,name=rank,proto3" json:"rank,omitempty"`
	RankScore             int32  `protobuf:"varint,4,opt,name=rank_score,json=rankScore,proto3" json:"rank_score,omitempty"`
	RankScoreDiff         int32  `protobuf:"varint,5,opt,name=rank_score_diff,json=rankScoreDiff,proto3" json:"rank_score_diff,omitempty"`
	RankScoreDiffRate     int32  `protobuf:"varint,6,opt,name=rank_score_diff_rate,json=rankScoreDiffRate,proto3" json:"rank_score_diff_rate,omitempty"`
	RankScoreDiffRateRank int32  `protobuf:"varint,7,opt,name=rank_score_diff_rate_rank,json=rankScoreDiffRateRank,proto3" json:"rank_score_diff_rate_rank,omitempty"`
}

func (x *Stat) Reset() {
	*x = Stat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_resources_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stat) ProtoMessage() {}

func (x *Stat) ProtoReflect() protoreflect.Message {
	mi := &file_game_resources_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stat.ProtoReflect.Descriptor instead.
func (*Stat) Descriptor() ([]byte, []int) {
	return file_game_resources_game_proto_rawDescGZIP(), []int{3}
}

func (x *Stat) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *Stat) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Stat) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *Stat) GetRankScore() int32 {
	if x != nil {
		return x.RankScore
	}
	return 0
}

func (x *Stat) GetRankScoreDiff() int32 {
	if x != nil {
		return x.RankScoreDiff
	}
	return 0
}

func (x *Stat) GetRankScoreDiffRate() int32 {
	if x != nil {
		return x.RankScoreDiffRate
	}
	return 0
}

func (x *Stat) GetRankScoreDiffRateRank() int32 {
	if x != nil {
		return x.RankScoreDiffRateRank
	}
	return 0
}

type Hand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         string    `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	State          HandState `protobuf:"varint,2,opt,name=state,proto3,enum=game.resources.HandState" json:"state,omitempty"`
	CenterPosition *Vector3  `protobuf:"bytes,3,opt,name=center_position,json=centerPosition,proto3" json:"center_position,omitempty"`
	ActionPosition *Vector3  `protobuf:"bytes,4,opt,name=action_position,json=actionPosition,proto3" json:"action_position,omitempty"`
	ForceDirection *Vector3  `protobuf:"bytes,5,opt,name=force_direction,json=forceDirection,proto3" json:"force_direction,omitempty"`
}

func (x *Hand) Reset() {
	*x = Hand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_resources_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hand) ProtoMessage() {}

func (x *Hand) ProtoReflect() protoreflect.Message {
	mi := &file_game_resources_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hand.ProtoReflect.Descriptor instead.
func (*Hand) Descriptor() ([]byte, []int) {
	return file_game_resources_game_proto_rawDescGZIP(), []int{4}
}

func (x *Hand) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Hand) GetState() HandState {
	if x != nil {
		return x.State
	}
	return HandState_HAND_STATE_UNKNOWN
}

func (x *Hand) GetCenterPosition() *Vector3 {
	if x != nil {
		return x.CenterPosition
	}
	return nil
}

func (x *Hand) GetActionPosition() *Vector3 {
	if x != nil {
		return x.ActionPosition
	}
	return nil
}

func (x *Hand) GetForceDirection() *Vector3 {
	if x != nil {
		return x.ForceDirection
	}
	return nil
}

var File_game_resources_game_proto protoreflect.FileDescriptor

var file_game_resources_game_proto_rawDesc = []byte{
	0x0a, 0x19, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x33, 0x0a, 0x07, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x7a,
	0x22, 0x4f, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x22, 0x82, 0x02, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12,
	0x30, 0x0a, 0x05, 0x6b, 0x69, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x05, 0x6b, 0x69, 0x6e, 0x64,
	0x73, 0x12, 0x31, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1b, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x52,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0xff, 0x01, 0x0a, 0x04, 0x53, 0x74, 0x61, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x61, 0x6e, 0x6b,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d,
	0x72, 0x61, 0x6e, 0x6b, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x44, 0x69, 0x66, 0x66, 0x12, 0x2f, 0x0a,
	0x14, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x64, 0x69, 0x66, 0x66,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x72, 0x61, 0x6e,
	0x6b, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x44, 0x69, 0x66, 0x66, 0x52, 0x61, 0x74, 0x65, 0x12, 0x38,
	0x0a, 0x19, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x64, 0x69, 0x66,
	0x66, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x15, 0x72, 0x61, 0x6e, 0x6b, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x44, 0x69, 0x66, 0x66,
	0x52, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x22, 0x96, 0x02, 0x0a, 0x04, 0x48, 0x61, 0x6e,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x0f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x52, 0x0e, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a,
	0x0f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x52,
	0x0e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x40, 0x0a, 0x0f, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x33, 0x52, 0x0e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2a, 0x90, 0x01, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x0a, 0x12, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x47, 0x41, 0x4d, 0x45, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x57, 0x41, 0x49, 0x54, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f,
	0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x10,
	0x02, 0x12, 0x13, 0x0a, 0x0f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x45, 0x58, 0x49, 0x54, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x10, 0x04, 0x12, 0x15, 0x0a,
	0x11, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x55,
	0x4c, 0x54, 0x10, 0x05, 0x2a, 0x6c, 0x0a, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x16, 0x0a, 0x12, 0x48, 0x41, 0x4e, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x48, 0x41, 0x4e,
	0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x48, 0x4f, 0x4c, 0x44, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x16, 0x0a, 0x12, 0x48, 0x41, 0x4e, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x4f, 0x50, 0x45, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x48, 0x41, 0x4e,
	0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x48, 0x52, 0x4f, 0x57, 0x49, 0x4e, 0x47,
	0x10, 0x03, 0x2a, 0x76, 0x0a, 0x0b, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x4f,
	0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x51, 0x55, 0x49, 0x45,
	0x53, 0x43, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x42, 0x4a, 0x45, 0x43,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4d, 0x4f, 0x56, 0x49, 0x4e, 0x47, 0x10, 0x02,
	0x12, 0x18, 0x0a, 0x14, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x46, 0x4f, 0x52, 0x43, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x2a, 0x42, 0x0a, 0x0a, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x42, 0x4a, 0x45,
	0x43, 0x54, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x1b, 0x0a, 0x17, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x4b, 0x49, 0x4e, 0x44,
	0x5f, 0x52, 0x45, 0x43, 0x54, 0x41, 0x4e, 0x47, 0x55, 0x4c, 0x41, 0x52, 0x10, 0x01, 0x42, 0xb4,
	0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b,
	0x2d, 0x4b, 0x69, 0x7a, 0x75, 0x6b, 0x75, 0x2f, 0x6b, 0x6f, 0x74, 0x61, 0x74, 0x75, 0x6e, 0x65,
	0x6b, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x61, 0x6d, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02,
	0x03, 0x47, 0x52, 0x58, 0xaa, 0x02, 0x0e, 0x47, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x0e, 0x47, 0x61, 0x6d, 0x65, 0x5c, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xe2, 0x02, 0x1a, 0x47, 0x61, 0x6d, 0x65, 0x5c, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x47, 0x61, 0x6d, 0x65, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_resources_game_proto_rawDescOnce sync.Once
	file_game_resources_game_proto_rawDescData = file_game_resources_game_proto_rawDesc
)

func file_game_resources_game_proto_rawDescGZIP() []byte {
	file_game_resources_game_proto_rawDescOnce.Do(func() {
		file_game_resources_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_resources_game_proto_rawDescData)
	})
	return file_game_resources_game_proto_rawDescData
}

var file_game_resources_game_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_game_resources_game_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_game_resources_game_proto_goTypes = []any{
	(GameState)(0),   // 0: game.resources.GameState
	(HandState)(0),   // 1: game.resources.HandState
	(ObjectState)(0), // 2: game.resources.ObjectState
	(ObjectKind)(0),  // 3: game.resources.ObjectKind
	(*Vector3)(nil),  // 4: game.resources.Vector3
	(*Player)(nil),   // 5: game.resources.Player
	(*Object)(nil),   // 6: game.resources.Object
	(*Stat)(nil),     // 7: game.resources.Stat
	(*Hand)(nil),     // 8: game.resources.Hand
}
var file_game_resources_game_proto_depIdxs = []int32{
	3, // 0: game.resources.Object.kinds:type_name -> game.resources.ObjectKind
	2, // 1: game.resources.Object.state:type_name -> game.resources.ObjectState
	4, // 2: game.resources.Object.position:type_name -> game.resources.Vector3
	4, // 3: game.resources.Object.size:type_name -> game.resources.Vector3
	1, // 4: game.resources.Hand.state:type_name -> game.resources.HandState
	4, // 5: game.resources.Hand.center_position:type_name -> game.resources.Vector3
	4, // 6: game.resources.Hand.action_position:type_name -> game.resources.Vector3
	4, // 7: game.resources.Hand.force_direction:type_name -> game.resources.Vector3
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_game_resources_game_proto_init() }
func file_game_resources_game_proto_init() {
	if File_game_resources_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_resources_game_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Vector3); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_resources_game_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Player); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_resources_game_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Object); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_resources_game_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Stat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_resources_game_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Hand); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_resources_game_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_resources_game_proto_goTypes,
		DependencyIndexes: file_game_resources_game_proto_depIdxs,
		EnumInfos:         file_game_resources_game_proto_enumTypes,
		MessageInfos:      file_game_resources_game_proto_msgTypes,
	}.Build()
	File_game_resources_game_proto = out.File
	file_game_resources_game_proto_rawDesc = nil
	file_game_resources_game_proto_goTypes = nil
	file_game_resources_game_proto_depIdxs = nil
}

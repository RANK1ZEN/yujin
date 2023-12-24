// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: internal/riotgrpc/proto/summoner.proto

package proto

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

type ByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ByNameRequest) Reset() {
	*x = ByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_riotgrpc_proto_summoner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByNameRequest) ProtoMessage() {}

func (x *ByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_riotgrpc_proto_summoner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByNameRequest.ProtoReflect.Descriptor instead.
func (*ByNameRequest) Descriptor() ([]byte, []int) {
	return file_internal_riotgrpc_proto_summoner_proto_rawDescGZIP(), []int{0}
}

func (x *ByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ByPuuidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Puuid string `protobuf:"bytes,1,opt,name=puuid,proto3" json:"puuid,omitempty"`
}

func (x *ByPuuidRequest) Reset() {
	*x = ByPuuidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_riotgrpc_proto_summoner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByPuuidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByPuuidRequest) ProtoMessage() {}

func (x *ByPuuidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_riotgrpc_proto_summoner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByPuuidRequest.ProtoReflect.Descriptor instead.
func (*ByPuuidRequest) Descriptor() ([]byte, []int) {
	return file_internal_riotgrpc_proto_summoner_proto_rawDescGZIP(), []int{1}
}

func (x *ByPuuidRequest) GetPuuid() string {
	if x != nil {
		return x.Puuid
	}
	return ""
}

type BySummonerIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SummonerId string `protobuf:"bytes,1,opt,name=summoner_id,json=summonerId,proto3" json:"summoner_id,omitempty"`
}

func (x *BySummonerIdRequest) Reset() {
	*x = BySummonerIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_riotgrpc_proto_summoner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BySummonerIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BySummonerIdRequest) ProtoMessage() {}

func (x *BySummonerIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_riotgrpc_proto_summoner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BySummonerIdRequest.ProtoReflect.Descriptor instead.
func (*BySummonerIdRequest) Descriptor() ([]byte, []int) {
	return file_internal_riotgrpc_proto_summoner_proto_rawDescGZIP(), []int{2}
}

func (x *BySummonerIdRequest) GetSummonerId() string {
	if x != nil {
		return x.SummonerId
	}
	return ""
}

type ByPuuidMatchlistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Puuid string `protobuf:"bytes,1,opt,name=puuid,proto3" json:"puuid,omitempty"`
	Start int32  `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	Count int32  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ByPuuidMatchlistRequest) Reset() {
	*x = ByPuuidMatchlistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_riotgrpc_proto_summoner_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByPuuidMatchlistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByPuuidMatchlistRequest) ProtoMessage() {}

func (x *ByPuuidMatchlistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_riotgrpc_proto_summoner_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByPuuidMatchlistRequest.ProtoReflect.Descriptor instead.
func (*ByPuuidMatchlistRequest) Descriptor() ([]byte, []int) {
	return file_internal_riotgrpc_proto_summoner_proto_rawDescGZIP(), []int{3}
}

func (x *ByPuuidMatchlistRequest) GetPuuid() string {
	if x != nil {
		return x.Puuid
	}
	return ""
}

func (x *ByPuuidMatchlistRequest) GetStart() int32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *ByPuuidMatchlistRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Summoner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Puuid         string `protobuf:"bytes,1,opt,name=puuid,proto3" json:"puuid,omitempty"`
	AccountId     string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	SummonerId    string `protobuf:"bytes,3,opt,name=summoner_id,json=summonerId,proto3" json:"summoner_id,omitempty"`
	Level         int64  `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
	ProfileIconId int32  `protobuf:"varint,5,opt,name=profile_icon_id,json=profileIconId,proto3" json:"profile_icon_id,omitempty"`
	Name          string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	LastRevision  int64  `protobuf:"varint,7,opt,name=last_revision,json=lastRevision,proto3" json:"last_revision,omitempty"`
}

func (x *Summoner) Reset() {
	*x = Summoner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_riotgrpc_proto_summoner_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Summoner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Summoner) ProtoMessage() {}

func (x *Summoner) ProtoReflect() protoreflect.Message {
	mi := &file_internal_riotgrpc_proto_summoner_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Summoner.ProtoReflect.Descriptor instead.
func (*Summoner) Descriptor() ([]byte, []int) {
	return file_internal_riotgrpc_proto_summoner_proto_rawDescGZIP(), []int{4}
}

func (x *Summoner) GetPuuid() string {
	if x != nil {
		return x.Puuid
	}
	return ""
}

func (x *Summoner) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Summoner) GetSummonerId() string {
	if x != nil {
		return x.SummonerId
	}
	return ""
}

func (x *Summoner) GetLevel() int64 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Summoner) GetProfileIconId() int32 {
	if x != nil {
		return x.ProfileIconId
	}
	return 0
}

func (x *Summoner) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Summoner) GetLastRevision() int64 {
	if x != nil {
		return x.LastRevision
	}
	return 0
}

type LeagueEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SummonerId   string `protobuf:"bytes,2,opt,name=summoner_id,json=summonerId,proto3" json:"summoner_id,omitempty"`
	SummonerName string `protobuf:"bytes,3,opt,name=summoner_name,json=summonerName,proto3" json:"summoner_name,omitempty"`
	Tier         string `protobuf:"bytes,5,opt,name=tier,proto3" json:"tier,omitempty"`
	Rank         string `protobuf:"bytes,6,opt,name=rank,proto3" json:"rank,omitempty"`
	Lp           int32  `protobuf:"varint,7,opt,name=lp,proto3" json:"lp,omitempty"`
	Wins         int32  `protobuf:"varint,8,opt,name=wins,proto3" json:"wins,omitempty"`
	Losses       int32  `protobuf:"varint,9,opt,name=losses,proto3" json:"losses,omitempty"`
}

func (x *LeagueEntry) Reset() {
	*x = LeagueEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_riotgrpc_proto_summoner_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeagueEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeagueEntry) ProtoMessage() {}

func (x *LeagueEntry) ProtoReflect() protoreflect.Message {
	mi := &file_internal_riotgrpc_proto_summoner_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeagueEntry.ProtoReflect.Descriptor instead.
func (*LeagueEntry) Descriptor() ([]byte, []int) {
	return file_internal_riotgrpc_proto_summoner_proto_rawDescGZIP(), []int{5}
}

func (x *LeagueEntry) GetSummonerId() string {
	if x != nil {
		return x.SummonerId
	}
	return ""
}

func (x *LeagueEntry) GetSummonerName() string {
	if x != nil {
		return x.SummonerName
	}
	return ""
}

func (x *LeagueEntry) GetTier() string {
	if x != nil {
		return x.Tier
	}
	return ""
}

func (x *LeagueEntry) GetRank() string {
	if x != nil {
		return x.Rank
	}
	return ""
}

func (x *LeagueEntry) GetLp() int32 {
	if x != nil {
		return x.Lp
	}
	return 0
}

func (x *LeagueEntry) GetWins() int32 {
	if x != nil {
		return x.Wins
	}
	return 0
}

func (x *LeagueEntry) GetLosses() int32 {
	if x != nil {
		return x.Losses
	}
	return 0
}

type Matchlist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchIds []string `protobuf:"bytes,1,rep,name=match_ids,json=matchIds,proto3" json:"match_ids,omitempty"`
}

func (x *Matchlist) Reset() {
	*x = Matchlist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_riotgrpc_proto_summoner_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Matchlist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Matchlist) ProtoMessage() {}

func (x *Matchlist) ProtoReflect() protoreflect.Message {
	mi := &file_internal_riotgrpc_proto_summoner_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Matchlist.ProtoReflect.Descriptor instead.
func (*Matchlist) Descriptor() ([]byte, []int) {
	return file_internal_riotgrpc_proto_summoner_proto_rawDescGZIP(), []int{6}
}

func (x *Matchlist) GetMatchIds() []string {
	if x != nil {
		return x.MatchIds
	}
	return nil
}

var File_internal_riotgrpc_proto_summoner_proto protoreflect.FileDescriptor

var file_internal_riotgrpc_proto_summoner_proto_rawDesc = []byte{
	0x0a, 0x26, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x69, 0x6f, 0x74, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0d, 0x42, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a,
	0x0e, 0x42, 0x79, 0x50, 0x75, 0x75, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x75, 0x75, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x13, 0x42, 0x79, 0x53, 0x75, 0x6d, 0x6d, 0x6f,
	0x6e, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5b, 0x0a,
	0x17, 0x42, 0x79, 0x50, 0x75, 0x75, 0x69, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xd7, 0x01, 0x0a, 0x08, 0x53,
	0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69,
	0x63, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x63, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0xb7, 0x01, 0x0a, 0x0b, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x6d, 0x6d, 0x6f,
	0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75,
	0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x61,
	0x6e, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x6c, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x6c, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x69, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x77, 0x69, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x73, 0x73, 0x65, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x6f, 0x73, 0x73, 0x65, 0x73, 0x22, 0x28,
	0x0a, 0x09, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x73, 0x32, 0xfb, 0x01, 0x0a, 0x0c, 0x52, 0x69, 0x6f,
	0x74, 0x53, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x06, 0x42, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x0e, 0x2e, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x22, 0x00,
	0x12, 0x27, 0x0a, 0x07, 0x42, 0x79, 0x50, 0x75, 0x75, 0x69, 0x64, 0x12, 0x0f, 0x2e, 0x42, 0x79,
	0x50, 0x75, 0x75, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x53,
	0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0c, 0x42, 0x79, 0x53,
	0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x2e, 0x42, 0x79, 0x53, 0x75,
	0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x09, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x53, 0x6f, 0x6c, 0x6f, 0x71, 0x12, 0x14, 0x2e, 0x42, 0x79, 0x53, 0x75, 0x6d,
	0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c,
	0x2e, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x00, 0x12, 0x36,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x18,
	0x2e, 0x42, 0x79, 0x50, 0x75, 0x75, 0x69, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x6c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x31, 0x7a, 0x65, 0x6e, 0x2f, 0x79, 0x75,
	0x6a, 0x69, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x69, 0x6f,
	0x74, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_internal_riotgrpc_proto_summoner_proto_rawDescOnce sync.Once
	file_internal_riotgrpc_proto_summoner_proto_rawDescData = file_internal_riotgrpc_proto_summoner_proto_rawDesc
)

func file_internal_riotgrpc_proto_summoner_proto_rawDescGZIP() []byte {
	file_internal_riotgrpc_proto_summoner_proto_rawDescOnce.Do(func() {
		file_internal_riotgrpc_proto_summoner_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_riotgrpc_proto_summoner_proto_rawDescData)
	})
	return file_internal_riotgrpc_proto_summoner_proto_rawDescData
}

var file_internal_riotgrpc_proto_summoner_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_internal_riotgrpc_proto_summoner_proto_goTypes = []interface{}{
	(*ByNameRequest)(nil),           // 0: ByNameRequest
	(*ByPuuidRequest)(nil),          // 1: ByPuuidRequest
	(*BySummonerIdRequest)(nil),     // 2: BySummonerIdRequest
	(*ByPuuidMatchlistRequest)(nil), // 3: ByPuuidMatchlistRequest
	(*Summoner)(nil),                // 4: Summoner
	(*LeagueEntry)(nil),             // 5: LeagueEntry
	(*Matchlist)(nil),               // 6: Matchlist
}
var file_internal_riotgrpc_proto_summoner_proto_depIdxs = []int32{
	0, // 0: RiotSummoner.ByName:input_type -> ByNameRequest
	1, // 1: RiotSummoner.ByPuuid:input_type -> ByPuuidRequest
	2, // 2: RiotSummoner.BySummonerId:input_type -> BySummonerIdRequest
	2, // 3: RiotSummoner.GetSoloq:input_type -> BySummonerIdRequest
	3, // 4: RiotSummoner.GetMatchlist:input_type -> ByPuuidMatchlistRequest
	4, // 5: RiotSummoner.ByName:output_type -> Summoner
	4, // 6: RiotSummoner.ByPuuid:output_type -> Summoner
	4, // 7: RiotSummoner.BySummonerId:output_type -> Summoner
	5, // 8: RiotSummoner.GetSoloq:output_type -> LeagueEntry
	6, // 9: RiotSummoner.GetMatchlist:output_type -> Matchlist
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_riotgrpc_proto_summoner_proto_init() }
func file_internal_riotgrpc_proto_summoner_proto_init() {
	if File_internal_riotgrpc_proto_summoner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_riotgrpc_proto_summoner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByNameRequest); i {
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
		file_internal_riotgrpc_proto_summoner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByPuuidRequest); i {
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
		file_internal_riotgrpc_proto_summoner_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BySummonerIdRequest); i {
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
		file_internal_riotgrpc_proto_summoner_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByPuuidMatchlistRequest); i {
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
		file_internal_riotgrpc_proto_summoner_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Summoner); i {
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
		file_internal_riotgrpc_proto_summoner_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeagueEntry); i {
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
		file_internal_riotgrpc_proto_summoner_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Matchlist); i {
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
			RawDescriptor: file_internal_riotgrpc_proto_summoner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_riotgrpc_proto_summoner_proto_goTypes,
		DependencyIndexes: file_internal_riotgrpc_proto_summoner_proto_depIdxs,
		MessageInfos:      file_internal_riotgrpc_proto_summoner_proto_msgTypes,
	}.Build()
	File_internal_riotgrpc_proto_summoner_proto = out.File
	file_internal_riotgrpc_proto_summoner_proto_rawDesc = nil
	file_internal_riotgrpc_proto_summoner_proto_goTypes = nil
	file_internal_riotgrpc_proto_summoner_proto_depIdxs = nil
}

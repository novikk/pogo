// Code generated by protoc-gen-go.
// source: pokemon.proto
// DO NOT EDIT!

/*
Package pokemon is a generated protocol buffer package.

It is generated from these files:
	pokemon.proto

It has these top-level messages:
	RequestEnvelop
	ResponseEnvelop
*/
package pokemon

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RequestEnvelop struct {
	Unknown1         *int32                     `protobuf:"varint,1,req,name=unknown1" json:"unknown1,omitempty"`
	RpcId            *int64                     `protobuf:"varint,3,opt,name=rpc_id" json:"rpc_id,omitempty"`
	Requests         []*RequestEnvelop_Requests `protobuf:"bytes,4,rep,name=requests" json:"requests,omitempty"`
	Unknown6         *RequestEnvelop_Unknown6   `protobuf:"bytes,6,opt,name=unknown6" json:"unknown6,omitempty"`
	Latitude         *uint64                    `protobuf:"fixed64,7,opt,name=latitude" json:"latitude,omitempty"`
	Longitude        *uint64                    `protobuf:"fixed64,8,opt,name=longitude" json:"longitude,omitempty"`
	Altitude         *uint64                    `protobuf:"fixed64,9,opt,name=altitude" json:"altitude,omitempty"`
	Auth             *RequestEnvelop_AuthInfo   `protobuf:"bytes,10,opt,name=auth" json:"auth,omitempty"`
	Unknown12        *int64                     `protobuf:"varint,12,opt,name=unknown12" json:"unknown12,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *RequestEnvelop) Reset()                    { *m = RequestEnvelop{} }
func (m *RequestEnvelop) String() string            { return proto.CompactTextString(m) }
func (*RequestEnvelop) ProtoMessage()               {}
func (*RequestEnvelop) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RequestEnvelop) GetUnknown1() int32 {
	if m != nil && m.Unknown1 != nil {
		return *m.Unknown1
	}
	return 0
}

func (m *RequestEnvelop) GetRpcId() int64 {
	if m != nil && m.RpcId != nil {
		return *m.RpcId
	}
	return 0
}

func (m *RequestEnvelop) GetRequests() []*RequestEnvelop_Requests {
	if m != nil {
		return m.Requests
	}
	return nil
}

func (m *RequestEnvelop) GetUnknown6() *RequestEnvelop_Unknown6 {
	if m != nil {
		return m.Unknown6
	}
	return nil
}

func (m *RequestEnvelop) GetLatitude() uint64 {
	if m != nil && m.Latitude != nil {
		return *m.Latitude
	}
	return 0
}

func (m *RequestEnvelop) GetLongitude() uint64 {
	if m != nil && m.Longitude != nil {
		return *m.Longitude
	}
	return 0
}

func (m *RequestEnvelop) GetAltitude() uint64 {
	if m != nil && m.Altitude != nil {
		return *m.Altitude
	}
	return 0
}

func (m *RequestEnvelop) GetAuth() *RequestEnvelop_AuthInfo {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *RequestEnvelop) GetUnknown12() int64 {
	if m != nil && m.Unknown12 != nil {
		return *m.Unknown12
	}
	return 0
}

type RequestEnvelop_Requests struct {
	Type             *int32                   `protobuf:"varint,1,req,name=type" json:"type,omitempty"`
	Message          *RequestEnvelop_Unknown3 `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *RequestEnvelop_Requests) Reset()                    { *m = RequestEnvelop_Requests{} }
func (m *RequestEnvelop_Requests) String() string            { return proto.CompactTextString(m) }
func (*RequestEnvelop_Requests) ProtoMessage()               {}
func (*RequestEnvelop_Requests) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *RequestEnvelop_Requests) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *RequestEnvelop_Requests) GetMessage() *RequestEnvelop_Unknown3 {
	if m != nil {
		return m.Message
	}
	return nil
}

type RequestEnvelop_Unknown3 struct {
	Unknown4         *string `protobuf:"bytes,1,req,name=unknown4" json:"unknown4,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RequestEnvelop_Unknown3) Reset()                    { *m = RequestEnvelop_Unknown3{} }
func (m *RequestEnvelop_Unknown3) String() string            { return proto.CompactTextString(m) }
func (*RequestEnvelop_Unknown3) ProtoMessage()               {}
func (*RequestEnvelop_Unknown3) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *RequestEnvelop_Unknown3) GetUnknown4() string {
	if m != nil && m.Unknown4 != nil {
		return *m.Unknown4
	}
	return ""
}

type RequestEnvelop_Unknown6 struct {
	Unknown1         *int32                            `protobuf:"varint,1,req,name=unknown1" json:"unknown1,omitempty"`
	Unknown2         *RequestEnvelop_Unknown6_Unknown2 `protobuf:"bytes,2,req,name=unknown2" json:"unknown2,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *RequestEnvelop_Unknown6) Reset()                    { *m = RequestEnvelop_Unknown6{} }
func (m *RequestEnvelop_Unknown6) String() string            { return proto.CompactTextString(m) }
func (*RequestEnvelop_Unknown6) ProtoMessage()               {}
func (*RequestEnvelop_Unknown6) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

func (m *RequestEnvelop_Unknown6) GetUnknown1() int32 {
	if m != nil && m.Unknown1 != nil {
		return *m.Unknown1
	}
	return 0
}

func (m *RequestEnvelop_Unknown6) GetUnknown2() *RequestEnvelop_Unknown6_Unknown2 {
	if m != nil {
		return m.Unknown2
	}
	return nil
}

type RequestEnvelop_Unknown6_Unknown2 struct {
	Unknown1         []byte `protobuf:"bytes,1,req,name=unknown1" json:"unknown1,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RequestEnvelop_Unknown6_Unknown2) Reset()         { *m = RequestEnvelop_Unknown6_Unknown2{} }
func (m *RequestEnvelop_Unknown6_Unknown2) String() string { return proto.CompactTextString(m) }
func (*RequestEnvelop_Unknown6_Unknown2) ProtoMessage()    {}
func (*RequestEnvelop_Unknown6_Unknown2) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 2, 0}
}

func (m *RequestEnvelop_Unknown6_Unknown2) GetUnknown1() []byte {
	if m != nil {
		return m.Unknown1
	}
	return nil
}

type RequestEnvelop_AuthInfo struct {
	Provider         *string                      `protobuf:"bytes,1,req,name=provider" json:"provider,omitempty"`
	Token            *RequestEnvelop_AuthInfo_JWT `protobuf:"bytes,2,req,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *RequestEnvelop_AuthInfo) Reset()                    { *m = RequestEnvelop_AuthInfo{} }
func (m *RequestEnvelop_AuthInfo) String() string            { return proto.CompactTextString(m) }
func (*RequestEnvelop_AuthInfo) ProtoMessage()               {}
func (*RequestEnvelop_AuthInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 3} }

func (m *RequestEnvelop_AuthInfo) GetProvider() string {
	if m != nil && m.Provider != nil {
		return *m.Provider
	}
	return ""
}

func (m *RequestEnvelop_AuthInfo) GetToken() *RequestEnvelop_AuthInfo_JWT {
	if m != nil {
		return m.Token
	}
	return nil
}

type RequestEnvelop_AuthInfo_JWT struct {
	Contents         *string `protobuf:"bytes,1,req,name=contents" json:"contents,omitempty"`
	Unknown13        *int32  `protobuf:"varint,2,req,name=unknown13" json:"unknown13,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RequestEnvelop_AuthInfo_JWT) Reset()         { *m = RequestEnvelop_AuthInfo_JWT{} }
func (m *RequestEnvelop_AuthInfo_JWT) String() string { return proto.CompactTextString(m) }
func (*RequestEnvelop_AuthInfo_JWT) ProtoMessage()    {}
func (*RequestEnvelop_AuthInfo_JWT) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 3, 0}
}

func (m *RequestEnvelop_AuthInfo_JWT) GetContents() string {
	if m != nil && m.Contents != nil {
		return *m.Contents
	}
	return ""
}

func (m *RequestEnvelop_AuthInfo_JWT) GetUnknown13() int32 {
	if m != nil && m.Unknown13 != nil {
		return *m.Unknown13
	}
	return 0
}

type ResponseEnvelop struct {
	Unknown1         *int32                     `protobuf:"varint,1,req,name=unknown1" json:"unknown1,omitempty"`
	RpcId            *int64                     `protobuf:"varint,2,opt,name=rpc_id" json:"rpc_id,omitempty"`
	ApiUrl           *string                    `protobuf:"bytes,3,opt,name=api_url" json:"api_url,omitempty"`
	Unknown6         *ResponseEnvelop_Unknown6  `protobuf:"bytes,6,opt,name=unknown6" json:"unknown6,omitempty"`
	Unknown7         *ResponseEnvelop_Unknown7  `protobuf:"bytes,7,opt,name=unknown7" json:"unknown7,omitempty"`
	Payload          []*ResponseEnvelop_Payload `protobuf:"bytes,100,rep,name=payload" json:"payload,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *ResponseEnvelop) Reset()                    { *m = ResponseEnvelop{} }
func (m *ResponseEnvelop) String() string            { return proto.CompactTextString(m) }
func (*ResponseEnvelop) ProtoMessage()               {}
func (*ResponseEnvelop) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ResponseEnvelop) GetUnknown1() int32 {
	if m != nil && m.Unknown1 != nil {
		return *m.Unknown1
	}
	return 0
}

func (m *ResponseEnvelop) GetRpcId() int64 {
	if m != nil && m.RpcId != nil {
		return *m.RpcId
	}
	return 0
}

func (m *ResponseEnvelop) GetApiUrl() string {
	if m != nil && m.ApiUrl != nil {
		return *m.ApiUrl
	}
	return ""
}

func (m *ResponseEnvelop) GetUnknown6() *ResponseEnvelop_Unknown6 {
	if m != nil {
		return m.Unknown6
	}
	return nil
}

func (m *ResponseEnvelop) GetUnknown7() *ResponseEnvelop_Unknown7 {
	if m != nil {
		return m.Unknown7
	}
	return nil
}

func (m *ResponseEnvelop) GetPayload() []*ResponseEnvelop_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

type ResponseEnvelop_Unknown6 struct {
	Unknown1         *int32                             `protobuf:"varint,1,req,name=unknown1" json:"unknown1,omitempty"`
	Unknown2         *ResponseEnvelop_Unknown6_Unknown2 `protobuf:"bytes,2,req,name=unknown2" json:"unknown2,omitempty"`
	XXX_unrecognized []byte                             `json:"-"`
}

func (m *ResponseEnvelop_Unknown6) Reset()                    { *m = ResponseEnvelop_Unknown6{} }
func (m *ResponseEnvelop_Unknown6) String() string            { return proto.CompactTextString(m) }
func (*ResponseEnvelop_Unknown6) ProtoMessage()               {}
func (*ResponseEnvelop_Unknown6) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *ResponseEnvelop_Unknown6) GetUnknown1() int32 {
	if m != nil && m.Unknown1 != nil {
		return *m.Unknown1
	}
	return 0
}

func (m *ResponseEnvelop_Unknown6) GetUnknown2() *ResponseEnvelop_Unknown6_Unknown2 {
	if m != nil {
		return m.Unknown2
	}
	return nil
}

type ResponseEnvelop_Unknown6_Unknown2 struct {
	Unknown1         []byte `protobuf:"bytes,1,req,name=unknown1" json:"unknown1,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ResponseEnvelop_Unknown6_Unknown2) Reset()         { *m = ResponseEnvelop_Unknown6_Unknown2{} }
func (m *ResponseEnvelop_Unknown6_Unknown2) String() string { return proto.CompactTextString(m) }
func (*ResponseEnvelop_Unknown6_Unknown2) ProtoMessage()    {}
func (*ResponseEnvelop_Unknown6_Unknown2) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0, 0}
}

func (m *ResponseEnvelop_Unknown6_Unknown2) GetUnknown1() []byte {
	if m != nil {
		return m.Unknown1
	}
	return nil
}

type ResponseEnvelop_Unknown7 struct {
	Unknown71        []byte `protobuf:"bytes,1,opt,name=unknown71" json:"unknown71,omitempty"`
	Unknown72        *int64 `protobuf:"varint,2,opt,name=unknown72" json:"unknown72,omitempty"`
	Unknown73        []byte `protobuf:"bytes,3,opt,name=unknown73" json:"unknown73,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ResponseEnvelop_Unknown7) Reset()                    { *m = ResponseEnvelop_Unknown7{} }
func (m *ResponseEnvelop_Unknown7) String() string            { return proto.CompactTextString(m) }
func (*ResponseEnvelop_Unknown7) ProtoMessage()               {}
func (*ResponseEnvelop_Unknown7) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

func (m *ResponseEnvelop_Unknown7) GetUnknown71() []byte {
	if m != nil {
		return m.Unknown71
	}
	return nil
}

func (m *ResponseEnvelop_Unknown7) GetUnknown72() int64 {
	if m != nil && m.Unknown72 != nil {
		return *m.Unknown72
	}
	return 0
}

func (m *ResponseEnvelop_Unknown7) GetUnknown73() []byte {
	if m != nil {
		return m.Unknown73
	}
	return nil
}

type ResponseEnvelop_Payload struct {
	Unknown1         *int32                   `protobuf:"varint,1,req,name=unknown1" json:"unknown1,omitempty"`
	Profile          *ResponseEnvelop_Profile `protobuf:"bytes,2,opt,name=profile" json:"profile,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *ResponseEnvelop_Payload) Reset()                    { *m = ResponseEnvelop_Payload{} }
func (m *ResponseEnvelop_Payload) String() string            { return proto.CompactTextString(m) }
func (*ResponseEnvelop_Payload) ProtoMessage()               {}
func (*ResponseEnvelop_Payload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 2} }

func (m *ResponseEnvelop_Payload) GetUnknown1() int32 {
	if m != nil && m.Unknown1 != nil {
		return *m.Unknown1
	}
	return 0
}

func (m *ResponseEnvelop_Payload) GetProfile() *ResponseEnvelop_Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type ResponseEnvelop_Profile struct {
	CreationTime     *int64                                 `protobuf:"varint,1,req,name=creation_time" json:"creation_time,omitempty"`
	Username         *string                                `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Team             *int32                                 `protobuf:"varint,5,opt,name=team" json:"team,omitempty"`
	Tutorial         []byte                                 `protobuf:"bytes,7,opt,name=tutorial" json:"tutorial,omitempty"`
	Avatar           *ResponseEnvelop_Profile_AvatarDetails `protobuf:"bytes,8,opt,name=avatar" json:"avatar,omitempty"`
	PokeStorage      *int32                                 `protobuf:"varint,9,opt,name=poke_storage" json:"poke_storage,omitempty"`
	ItemStorage      *int32                                 `protobuf:"varint,10,opt,name=item_storage" json:"item_storage,omitempty"`
	DailyBonus       *ResponseEnvelop_Profile_DailyBonus    `protobuf:"bytes,11,opt,name=daily_bonus" json:"daily_bonus,omitempty"`
	Unknown12        []byte                                 `protobuf:"bytes,12,opt,name=unknown12" json:"unknown12,omitempty"`
	Unknown13        []byte                                 `protobuf:"bytes,13,opt,name=unknown13" json:"unknown13,omitempty"`
	Currency         []*ResponseEnvelop_Profile_Currency    `protobuf:"bytes,14,rep,name=currency" json:"currency,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *ResponseEnvelop_Profile) Reset()                    { *m = ResponseEnvelop_Profile{} }
func (m *ResponseEnvelop_Profile) String() string            { return proto.CompactTextString(m) }
func (*ResponseEnvelop_Profile) ProtoMessage()               {}
func (*ResponseEnvelop_Profile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 3} }

func (m *ResponseEnvelop_Profile) GetCreationTime() int64 {
	if m != nil && m.CreationTime != nil {
		return *m.CreationTime
	}
	return 0
}

func (m *ResponseEnvelop_Profile) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *ResponseEnvelop_Profile) GetTeam() int32 {
	if m != nil && m.Team != nil {
		return *m.Team
	}
	return 0
}

func (m *ResponseEnvelop_Profile) GetTutorial() []byte {
	if m != nil {
		return m.Tutorial
	}
	return nil
}

func (m *ResponseEnvelop_Profile) GetAvatar() *ResponseEnvelop_Profile_AvatarDetails {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func (m *ResponseEnvelop_Profile) GetPokeStorage() int32 {
	if m != nil && m.PokeStorage != nil {
		return *m.PokeStorage
	}
	return 0
}

func (m *ResponseEnvelop_Profile) GetItemStorage() int32 {
	if m != nil && m.ItemStorage != nil {
		return *m.ItemStorage
	}
	return 0
}

func (m *ResponseEnvelop_Profile) GetDailyBonus() *ResponseEnvelop_Profile_DailyBonus {
	if m != nil {
		return m.DailyBonus
	}
	return nil
}

func (m *ResponseEnvelop_Profile) GetUnknown12() []byte {
	if m != nil {
		return m.Unknown12
	}
	return nil
}

func (m *ResponseEnvelop_Profile) GetUnknown13() []byte {
	if m != nil {
		return m.Unknown13
	}
	return nil
}

func (m *ResponseEnvelop_Profile) GetCurrency() []*ResponseEnvelop_Profile_Currency {
	if m != nil {
		return m.Currency
	}
	return nil
}

type ResponseEnvelop_Profile_AvatarDetails struct {
	Unknown2         *int32 `protobuf:"varint,2,opt,name=unknown2" json:"unknown2,omitempty"`
	Unknown3         *int32 `protobuf:"varint,3,opt,name=unknown3" json:"unknown3,omitempty"`
	Unknown9         *int32 `protobuf:"varint,9,opt,name=unknown9" json:"unknown9,omitempty"`
	Unknown10        *int32 `protobuf:"varint,10,opt,name=unknown10" json:"unknown10,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ResponseEnvelop_Profile_AvatarDetails) Reset()         { *m = ResponseEnvelop_Profile_AvatarDetails{} }
func (m *ResponseEnvelop_Profile_AvatarDetails) String() string { return proto.CompactTextString(m) }
func (*ResponseEnvelop_Profile_AvatarDetails) ProtoMessage()    {}
func (*ResponseEnvelop_Profile_AvatarDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 3, 0}
}

func (m *ResponseEnvelop_Profile_AvatarDetails) GetUnknown2() int32 {
	if m != nil && m.Unknown2 != nil {
		return *m.Unknown2
	}
	return 0
}

func (m *ResponseEnvelop_Profile_AvatarDetails) GetUnknown3() int32 {
	if m != nil && m.Unknown3 != nil {
		return *m.Unknown3
	}
	return 0
}

func (m *ResponseEnvelop_Profile_AvatarDetails) GetUnknown9() int32 {
	if m != nil && m.Unknown9 != nil {
		return *m.Unknown9
	}
	return 0
}

func (m *ResponseEnvelop_Profile_AvatarDetails) GetUnknown10() int32 {
	if m != nil && m.Unknown10 != nil {
		return *m.Unknown10
	}
	return 0
}

type ResponseEnvelop_Profile_DailyBonus struct {
	NextCollectTimestampMs              *int64 `protobuf:"varint,1,opt,name=NextCollectTimestampMs" json:"NextCollectTimestampMs,omitempty"`
	NextDefenderBonusCollectTimestampMs *int64 `protobuf:"varint,2,opt,name=NextDefenderBonusCollectTimestampMs" json:"NextDefenderBonusCollectTimestampMs,omitempty"`
	XXX_unrecognized                    []byte `json:"-"`
}

func (m *ResponseEnvelop_Profile_DailyBonus) Reset()         { *m = ResponseEnvelop_Profile_DailyBonus{} }
func (m *ResponseEnvelop_Profile_DailyBonus) String() string { return proto.CompactTextString(m) }
func (*ResponseEnvelop_Profile_DailyBonus) ProtoMessage()    {}
func (*ResponseEnvelop_Profile_DailyBonus) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 3, 1}
}

func (m *ResponseEnvelop_Profile_DailyBonus) GetNextCollectTimestampMs() int64 {
	if m != nil && m.NextCollectTimestampMs != nil {
		return *m.NextCollectTimestampMs
	}
	return 0
}

func (m *ResponseEnvelop_Profile_DailyBonus) GetNextDefenderBonusCollectTimestampMs() int64 {
	if m != nil && m.NextDefenderBonusCollectTimestampMs != nil {
		return *m.NextDefenderBonusCollectTimestampMs
	}
	return 0
}

type ResponseEnvelop_Profile_Currency struct {
	Type             *string `protobuf:"bytes,1,req,name=type" json:"type,omitempty"`
	Amount           *int32  `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ResponseEnvelop_Profile_Currency) Reset()         { *m = ResponseEnvelop_Profile_Currency{} }
func (m *ResponseEnvelop_Profile_Currency) String() string { return proto.CompactTextString(m) }
func (*ResponseEnvelop_Profile_Currency) ProtoMessage()    {}
func (*ResponseEnvelop_Profile_Currency) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 3, 2}
}

func (m *ResponseEnvelop_Profile_Currency) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *ResponseEnvelop_Profile_Currency) GetAmount() int32 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*RequestEnvelop)(nil), "pokemon.RequestEnvelop")
	proto.RegisterType((*RequestEnvelop_Requests)(nil), "pokemon.RequestEnvelop.Requests")
	proto.RegisterType((*RequestEnvelop_Unknown3)(nil), "pokemon.RequestEnvelop.Unknown3")
	proto.RegisterType((*RequestEnvelop_Unknown6)(nil), "pokemon.RequestEnvelop.Unknown6")
	proto.RegisterType((*RequestEnvelop_Unknown6_Unknown2)(nil), "pokemon.RequestEnvelop.Unknown6.Unknown2")
	proto.RegisterType((*RequestEnvelop_AuthInfo)(nil), "pokemon.RequestEnvelop.AuthInfo")
	proto.RegisterType((*RequestEnvelop_AuthInfo_JWT)(nil), "pokemon.RequestEnvelop.AuthInfo.JWT")
	proto.RegisterType((*ResponseEnvelop)(nil), "pokemon.ResponseEnvelop")
	proto.RegisterType((*ResponseEnvelop_Unknown6)(nil), "pokemon.ResponseEnvelop.Unknown6")
	proto.RegisterType((*ResponseEnvelop_Unknown6_Unknown2)(nil), "pokemon.ResponseEnvelop.Unknown6.Unknown2")
	proto.RegisterType((*ResponseEnvelop_Unknown7)(nil), "pokemon.ResponseEnvelop.Unknown7")
	proto.RegisterType((*ResponseEnvelop_Payload)(nil), "pokemon.ResponseEnvelop.Payload")
	proto.RegisterType((*ResponseEnvelop_Profile)(nil), "pokemon.ResponseEnvelop.Profile")
	proto.RegisterType((*ResponseEnvelop_Profile_AvatarDetails)(nil), "pokemon.ResponseEnvelop.Profile.AvatarDetails")
	proto.RegisterType((*ResponseEnvelop_Profile_DailyBonus)(nil), "pokemon.ResponseEnvelop.Profile.DailyBonus")
	proto.RegisterType((*ResponseEnvelop_Profile_Currency)(nil), "pokemon.ResponseEnvelop.Profile.Currency")
}

func init() { proto.RegisterFile("pokemon.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 685 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0x56, 0xf3, 0xe9, 0x4c, 0x9c, 0xb6, 0xaf, 0xf5, 0x82, 0x2c, 0x0b, 0xa1, 0x50, 0x38, 0xa4,
	0xad, 0x14, 0x51, 0x07, 0x35, 0x42, 0x20, 0x44, 0x69, 0x39, 0x00, 0xa2, 0x42, 0x55, 0x11, 0x70,
	0x8a, 0x96, 0x64, 0x5b, 0xac, 0xda, 0xbb, 0x66, 0x77, 0xdd, 0x36, 0x07, 0x7e, 0x16, 0x7f, 0x80,
	0x3f, 0xc5, 0x95, 0xfd, 0x4a, 0x6a, 0x37, 0x51, 0x9d, 0x93, 0xbd, 0xb3, 0x33, 0xcf, 0xcc, 0x3c,
	0xf3, 0xcc, 0x42, 0x27, 0xa5, 0x17, 0x38, 0xa1, 0xa4, 0x9f, 0x32, 0x2a, 0xa8, 0xd7, 0xb4, 0xc7,
	0xad, 0xbf, 0x35, 0x58, 0x3f, 0xc1, 0x3f, 0x33, 0xcc, 0xc5, 0x5b, 0x72, 0x89, 0x63, 0x9a, 0x7a,
	0x9b, 0xe0, 0x64, 0xe4, 0x82, 0xd0, 0x2b, 0xb2, 0xe7, 0xaf, 0x75, 0x2b, 0xbd, 0xba, 0xb7, 0x0e,
	0x0d, 0x96, 0x8e, 0x47, 0xd1, 0xc4, 0xaf, 0x76, 0xd7, 0x7a, 0x55, 0x2f, 0x04, 0x87, 0x99, 0x18,
	0xee, 0xd7, 0xba, 0xd5, 0x5e, 0x3b, 0xec, 0xf6, 0x67, 0xf8, 0x45, 0xb0, 0xd9, 0x91, 0xab, 0x18,
	0x8b, 0xba, 0xef, 0x37, 0x24, 0xca, 0x1d, 0x31, 0x9f, 0xad, 0x9f, 0xaa, 0x24, 0x46, 0x22, 0x12,
	0xd9, 0x04, 0xfb, 0x4d, 0x19, 0xd3, 0xf0, 0xfe, 0x83, 0x56, 0x4c, 0xc9, 0xb9, 0x31, 0x39, 0xda,
	0x24, 0x9d, 0x50, 0x6c, 0x9d, 0x5a, 0xda, 0xd2, 0x87, 0x1a, 0xca, 0xc4, 0x0f, 0x1f, 0xee, 0x4e,
	0x73, 0x20, 0x7d, 0xde, 0x91, 0x33, 0xaa, 0x40, 0x67, 0x0d, 0x87, 0xbe, 0xab, 0x3a, 0x0c, 0x3e,
	0x80, 0x33, 0xaf, 0xdc, 0x85, 0x9a, 0x98, 0xa6, 0xd8, 0x72, 0xb1, 0x07, 0xcd, 0x04, 0x73, 0x8e,
	0xce, 0xb1, 0x5f, 0x59, 0xa9, 0x8d, 0x41, 0xf0, 0x00, 0x9c, 0xd9, 0x7f, 0x8e, 0xdc, 0x67, 0x1a,
	0xb0, 0x15, 0x5c, 0xcd, 0x6f, 0xf7, 0x97, 0x50, 0xff, 0x62, 0x6e, 0x09, 0x65, 0xbe, 0x8a, 0xcc,
	0xb7, 0x5d, 0x46, 0xdb, 0xec, 0x27, 0xcc, 0x25, 0x0e, 0x17, 0xa0, 0xdd, 0xe0, 0x17, 0x38, 0x73,
	0x0a, 0xe4, 0xad, 0x14, 0xc6, 0x65, 0x34, 0xc1, 0xcc, 0x94, 0xe5, 0x0d, 0xa0, 0x2e, 0x64, 0x1e,
	0x62, 0xb3, 0x3e, 0x29, 0x63, 0xb1, 0xff, 0xfe, 0xcb, 0x69, 0xb0, 0x03, 0x55, 0xf9, 0x51, 0x68,
	0x63, 0x4a, 0x04, 0x26, 0x52, 0x1f, 0x06, 0x2d, 0x47, 0xf1, 0x40, 0x23, 0xd6, 0xb7, 0x7e, 0x3b,
	0xb0, 0x71, 0x82, 0x79, 0x4a, 0x09, 0xc7, 0xab, 0x48, 0xaf, 0xa2, 0xa5, 0xb7, 0x01, 0x4d, 0x94,
	0x46, 0xa3, 0x8c, 0xc5, 0x5a, 0x8b, 0xaa, 0xce, 0xdb, 0xba, 0x7a, 0x94, 0x2b, 0xb5, 0x00, 0x7f,
	0x23, 0xac, 0x9b, 0xa0, 0xa1, 0x16, 0xd6, 0x0a, 0x41, 0x43, 0x35, 0xf9, 0x14, 0x4d, 0x63, 0x8a,
	0x26, 0xfe, 0x64, 0x41, 0xf4, 0xc5, 0x98, 0x4f, 0xc6, 0x2f, 0xb8, 0xbe, 0x73, 0xb6, 0x2f, 0x17,
	0x66, 0xbb, 0x53, 0x5a, 0xfa, 0xaa, 0xc3, 0x3d, 0x9c, 0xdf, 0x0e, 0x73, 0xe4, 0x0f, 0xd5, 0xf5,
	0x5a, 0xcf, 0xcd, 0x9b, 0x42, 0xcb, 0x6c, 0xce, 0x34, 0xd0, 0xdc, 0xba, 0xc1, 0x31, 0x34, 0x6d,
	0x27, 0x4b, 0xaa, 0x57, 0x74, 0x30, 0x7a, 0x16, 0xc5, 0xcb, 0x16, 0xe1, 0x16, 0x1d, 0xc6, 0x2f,
	0xf8, 0x53, 0x93, 0x80, 0xe6, 0xdf, 0xbb, 0x07, 0x9d, 0x31, 0xc3, 0x72, 0xbb, 0x29, 0x19, 0x89,
	0x28, 0x31, 0xeb, 0x55, 0xd5, 0x79, 0x38, 0x66, 0x04, 0x25, 0x06, 0xb6, 0xa5, 0xd7, 0x0f, 0xa3,
	0xc4, 0xaf, 0xcb, 0x53, 0x5d, 0xdd, 0x8b, 0x4c, 0x50, 0x16, 0xa1, 0x58, 0x4f, 0xce, 0xf5, 0x5e,
	0x41, 0x03, 0x5d, 0x22, 0x81, 0x98, 0x7e, 0x0f, 0xda, 0x61, 0xbf, 0xac, 0x8c, 0xfe, 0x81, 0x76,
	0x3f, 0xc2, 0x02, 0x45, 0x31, 0xf7, 0xfe, 0x07, 0x57, 0x05, 0x8c, 0xb8, 0x44, 0x55, 0x5b, 0xdd,
	0xd2, 0x79, 0xa4, 0x35, 0x12, 0x38, 0x99, 0x5b, 0x41, 0x5b, 0x5f, 0x43, 0x7b, 0x22, 0x83, 0xa6,
	0xa3, 0xef, 0x94, 0x64, 0xdc, 0x6f, 0xeb, 0x84, 0xbb, 0xa5, 0x09, 0x8f, 0x54, 0xcc, 0x1b, 0x15,
	0xb2, 0xf8, 0xd6, 0xb8, 0xc5, 0xdd, 0xe8, 0x68, 0x93, 0xdc, 0xfa, 0x71, 0xc6, 0x18, 0x26, 0xe3,
	0xa9, 0xbf, 0xae, 0xb5, 0xb6, 0x5d, 0x9a, 0xe4, 0xd0, 0x06, 0x04, 0x5f, 0xa1, 0x53, 0xec, 0x70,
	0xb3, 0xa0, 0x33, 0xcb, 0xa2, 0xb5, 0x98, 0x51, 0xe7, 0x2d, 0xcf, 0x2d, 0x03, 0xb9, 0xb2, 0x9e,
	0x9a, 0xf6, 0x83, 0x6f, 0x00, 0xb9, 0x56, 0x1e, 0xc2, 0xfd, 0x63, 0x7c, 0x2d, 0x0e, 0x69, 0x1c,
	0xe3, 0xb1, 0x38, 0x95, 0x33, 0xe4, 0x02, 0x25, 0xe9, 0x47, 0xae, 0x35, 0x56, 0xf5, 0x76, 0xe1,
	0xb1, 0xba, 0x3f, 0xc2, 0x67, 0x98, 0xc8, 0x77, 0x45, 0x07, 0x2d, 0x71, 0xd6, 0xea, 0x0b, 0x7a,
	0xe0, 0xcc, 0x1a, 0x28, 0x3c, 0xb8, 0x2d, 0xf5, 0x02, 0xa0, 0x84, 0x66, 0x44, 0x98, 0xda, 0xff,
	0x05, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xdf, 0xfd, 0x94, 0xca, 0x06, 0x00, 0x00,
}

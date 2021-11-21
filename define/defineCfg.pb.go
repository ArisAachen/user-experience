// protobuf version

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.6.1
// source: protobuf/defineCfg.proto

package define

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

// Basic object struct for mul message
type Obj struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module string `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Obj) Reset() {
	*x = Obj{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_defineCfg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Obj) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Obj) ProtoMessage() {}

func (x *Obj) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_defineCfg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Obj.ProtoReflect.Descriptor instead.
func (*Obj) Descriptor() ([]byte, []int) {
	return file_protobuf_defineCfg_proto_rawDescGZIP(), []int{0}
}

func (x *Obj) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (x *Obj) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type HardwareOther struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Machine string `protobuf:"bytes,1,opt,name=machine,proto3" json:"machine,omitempty"`
	Apt     string `protobuf:"bytes,2,opt,name=apt,proto3" json:"apt,omitempty"`
	Active  string `protobuf:"bytes,3,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *HardwareOther) Reset() {
	*x = HardwareOther{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_defineCfg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HardwareOther) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HardwareOther) ProtoMessage() {}

func (x *HardwareOther) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_defineCfg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HardwareOther.ProtoReflect.Descriptor instead.
func (*HardwareOther) Descriptor() ([]byte, []int) {
	return file_protobuf_defineCfg_proto_rawDescGZIP(), []int{1}
}

func (x *HardwareOther) GetMachine() string {
	if x != nil {
		return x.Machine
	}
	return ""
}

func (x *HardwareOther) GetApt() string {
	if x != nil {
		return x.Apt
	}
	return ""
}

func (x *HardwareOther) GetActive() string {
	if x != nil {
		return x.Active
	}
	return ""
}

// save hardware info last time
type HardwareInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mac       string         `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
	OsType    string         `protobuf:"bytes,4,opt,name=os_type,json=osType,proto3" json:"os_type,omitempty"`
	OsVersion string         `protobuf:"bytes,5,opt,name=os_version,json=osVersion,proto3" json:"os_version,omitempty"`
	Version   string         `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
	UniId     string         `protobuf:"bytes,7,opt,name=uni_id,json=uniId,proto3" json:"uni_id,omitempty"`
	Cpu       *Obj           `protobuf:"bytes,8,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Board     *Obj           `protobuf:"bytes,9,opt,name=board,proto3" json:"board,omitempty"`
	Gpu       *Obj           `protobuf:"bytes,10,opt,name=gpu,proto3" json:"gpu,omitempty"`
	Memory    *Obj           `protobuf:"bytes,11,opt,name=memory,proto3" json:"memory,omitempty"`
	Disk      *Obj           `protobuf:"bytes,12,opt,name=disk,proto3" json:"disk,omitempty"`
	Network   *Obj           `protobuf:"bytes,13,opt,name=network,proto3" json:"network,omitempty"`
	Other     *HardwareOther `protobuf:"bytes,14,opt,name=other,proto3" json:"other,omitempty"`
}

func (x *HardwareInfo) Reset() {
	*x = HardwareInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_defineCfg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HardwareInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HardwareInfo) ProtoMessage() {}

func (x *HardwareInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_defineCfg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HardwareInfo.ProtoReflect.Descriptor instead.
func (*HardwareInfo) Descriptor() ([]byte, []int) {
	return file_protobuf_defineCfg_proto_rawDescGZIP(), []int{2}
}

func (x *HardwareInfo) GetMac() string {
	if x != nil {
		return x.Mac
	}
	return ""
}

func (x *HardwareInfo) GetOsType() string {
	if x != nil {
		return x.OsType
	}
	return ""
}

func (x *HardwareInfo) GetOsVersion() string {
	if x != nil {
		return x.OsVersion
	}
	return ""
}

func (x *HardwareInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *HardwareInfo) GetUniId() string {
	if x != nil {
		return x.UniId
	}
	return ""
}

func (x *HardwareInfo) GetCpu() *Obj {
	if x != nil {
		return x.Cpu
	}
	return nil
}

func (x *HardwareInfo) GetBoard() *Obj {
	if x != nil {
		return x.Board
	}
	return nil
}

func (x *HardwareInfo) GetGpu() *Obj {
	if x != nil {
		return x.Gpu
	}
	return nil
}

func (x *HardwareInfo) GetMemory() *Obj {
	if x != nil {
		return x.Memory
	}
	return nil
}

func (x *HardwareInfo) GetDisk() *Obj {
	if x != nil {
		return x.Disk
	}
	return nil
}

func (x *HardwareInfo) GetNetwork() *Obj {
	if x != nil {
		return x.Network
	}
	return nil
}

func (x *HardwareInfo) GetOther() *HardwareOther {
	if x != nil {
		return x.Other
	}
	return nil
}

// save post abstract
type PostDomain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UrlPath string `protobuf:"bytes,1,opt,name=url_path,json=urlPath,proto3" json:"url_path,omitempty"`
	Time    uint64 `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *PostDomain) Reset() {
	*x = PostDomain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_defineCfg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostDomain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostDomain) ProtoMessage() {}

func (x *PostDomain) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_defineCfg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostDomain.ProtoReflect.Descriptor instead.
func (*PostDomain) Descriptor() ([]byte, []int) {
	return file_protobuf_defineCfg_proto_rawDescGZIP(), []int{3}
}

func (x *PostDomain) GetUrlPath() string {
	if x != nil {
		return x.UrlPath
	}
	return ""
}

func (x *PostDomain) GetTime() uint64 {
	if x != nil {
		return x.Time
	}
	return 0
}

// save post interfaces info last time
type PostInterface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// all domains {"https://abc.com",12345}
	Domains      []*PostDomain `protobuf:"bytes,1,rep,name=domains,proto3" json:"domains,omitempty"`
	Base         string        `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
	Info         string        `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	Order        string        `protobuf:"bytes,4,opt,name=order,proto3" json:"order,omitempty"`
	Apt          string        `protobuf:"bytes,5,opt,name=apt,proto3" json:"apt,omitempty"`
	Use          string        `protobuf:"bytes,6,opt,name=use,proto3" json:"use,omitempty"`
	Update       string        `protobuf:"bytes,7,opt,name=update,proto3" json:"update,omitempty"`
	UpdateDomain string        `protobuf:"bytes,8,opt,name=update_domain,json=updateDomain,proto3" json:"update_domain,omitempty"`
	UpdatePath   string        `protobuf:"bytes,9,opt,name=update_path,json=updatePath,proto3" json:"update_path,omitempty"`
	General      string        `protobuf:"bytes,10,opt,name=general,proto3" json:"general,omitempty"`
}

func (x *PostInterface) Reset() {
	*x = PostInterface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_defineCfg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostInterface) ProtoMessage() {}

func (x *PostInterface) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_defineCfg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostInterface.ProtoReflect.Descriptor instead.
func (*PostInterface) Descriptor() ([]byte, []int) {
	return file_protobuf_defineCfg_proto_rawDescGZIP(), []int{4}
}

func (x *PostInterface) GetDomains() []*PostDomain {
	if x != nil {
		return x.Domains
	}
	return nil
}

func (x *PostInterface) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *PostInterface) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *PostInterface) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

func (x *PostInterface) GetApt() string {
	if x != nil {
		return x.Apt
	}
	return ""
}

func (x *PostInterface) GetUse() string {
	if x != nil {
		return x.Use
	}
	return ""
}

func (x *PostInterface) GetUpdate() string {
	if x != nil {
		return x.Update
	}
	return ""
}

func (x *PostInterface) GetUpdateDomain() string {
	if x != nil {
		return x.UpdateDomain
	}
	return ""
}

func (x *PostInterface) GetUpdatePath() string {
	if x != nil {
		return x.UpdatePath
	}
	return ""
}

func (x *PostInterface) GetGeneral() string {
	if x != nil {
		return x.General
	}
	return ""
}

// save system ref config, apt time, user-exp-enabled, develop-mode
type SysCfg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Develop bool   `protobuf:"varint,1,opt,name=develop,proto3" json:"develop,omitempty"`
	UserExp bool   `protobuf:"varint,2,opt,name=user_exp,json=userExp,proto3" json:"user_exp,omitempty"`
	Apt     uint64 `protobuf:"varint,3,opt,name=apt,proto3" json:"apt,omitempty"`
}

func (x *SysCfg) Reset() {
	*x = SysCfg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_defineCfg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysCfg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysCfg) ProtoMessage() {}

func (x *SysCfg) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_defineCfg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysCfg.ProtoReflect.Descriptor instead.
func (*SysCfg) Descriptor() ([]byte, []int) {
	return file_protobuf_defineCfg_proto_rawDescGZIP(), []int{5}
}

func (x *SysCfg) GetDevelop() bool {
	if x != nil {
		return x.Develop
	}
	return false
}

func (x *SysCfg) GetUserExp() bool {
	if x != nil {
		return x.UserExp
	}
	return false
}

func (x *SysCfg) GetApt() uint64 {
	if x != nil {
		return x.Apt
	}
	return 0
}

// RSAKey store rsa public key and private key
type RsaKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Public  string `protobuf:"bytes,1,opt,name=public,proto3" json:"public,omitempty"`
	Private string `protobuf:"bytes,2,opt,name=private,proto3" json:"private,omitempty"`
}

func (x *RsaKey) Reset() {
	*x = RsaKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_defineCfg_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RsaKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RsaKey) ProtoMessage() {}

func (x *RsaKey) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_defineCfg_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RsaKey.ProtoReflect.Descriptor instead.
func (*RsaKey) Descriptor() ([]byte, []int) {
	return file_protobuf_defineCfg_proto_rawDescGZIP(), []int{6}
}

func (x *RsaKey) GetPublic() string {
	if x != nil {
		return x.Public
	}
	return ""
}

func (x *RsaKey) GetPrivate() string {
	if x != nil {
		return x.Private
	}
	return ""
}

var File_protobuf_defineCfg_proto protoreflect.FileDescriptor

var file_protobuf_defineCfg_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x65, 0x43, 0x66, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x65, 0x22, 0x2d, 0x0a, 0x03, 0x4f, 0x62, 0x6a, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x53, 0x0a, 0x0d, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x4f, 0x74, 0x68,
	0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x61, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x84, 0x03, 0x0a, 0x0c, 0x48, 0x61, 0x72, 0x64, 0x77,
	0x61, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x61, 0x63, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x73, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x75,
	0x6e, 0x69, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x75, 0x6e, 0x69,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x52, 0x03, 0x63, 0x70,
	0x75, 0x12, 0x21, 0x0a, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x52, 0x05, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x03, 0x67, 0x70, 0x75, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x52, 0x03,
	0x67, 0x70, 0x75, 0x12, 0x23, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x4f, 0x62, 0x6a,
	0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x04, 0x64, 0x69, 0x73, 0x6b,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e,
	0x4f, 0x62, 0x6a, 0x52, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x12, 0x25, 0x0a, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x12, 0x2b, 0x0a, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72,
	0x65, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x22, 0x3b, 0x0a,
	0x0a, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x75,
	0x72, 0x6c, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75,
	0x72, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x97, 0x02, 0x0a, 0x0d, 0x50,
	0x6f, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x07,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x52, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61,
	0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x73,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x22, 0x4f, 0x0a, 0x06, 0x53, 0x79, 0x73, 0x43, 0x66, 0x67, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x65, 0x78, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x45, 0x78, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x03, 0x61, 0x70, 0x74, 0x22, 0x3a, 0x0a, 0x06, 0x52, 0x73, 0x61, 0x4b, 0x65, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_defineCfg_proto_rawDescOnce sync.Once
	file_protobuf_defineCfg_proto_rawDescData = file_protobuf_defineCfg_proto_rawDesc
)

func file_protobuf_defineCfg_proto_rawDescGZIP() []byte {
	file_protobuf_defineCfg_proto_rawDescOnce.Do(func() {
		file_protobuf_defineCfg_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_defineCfg_proto_rawDescData)
	})
	return file_protobuf_defineCfg_proto_rawDescData
}

var file_protobuf_defineCfg_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protobuf_defineCfg_proto_goTypes = []interface{}{
	(*Obj)(nil),           // 0: define.Obj
	(*HardwareOther)(nil), // 1: define.HardwareOther
	(*HardwareInfo)(nil),  // 2: define.HardwareInfo
	(*PostDomain)(nil),    // 3: define.PostDomain
	(*PostInterface)(nil), // 4: define.PostInterface
	(*SysCfg)(nil),        // 5: define.SysCfg
	(*RsaKey)(nil),        // 6: define.RsaKey
}
var file_protobuf_defineCfg_proto_depIdxs = []int32{
	0, // 0: define.HardwareInfo.cpu:type_name -> define.Obj
	0, // 1: define.HardwareInfo.board:type_name -> define.Obj
	0, // 2: define.HardwareInfo.gpu:type_name -> define.Obj
	0, // 3: define.HardwareInfo.memory:type_name -> define.Obj
	0, // 4: define.HardwareInfo.disk:type_name -> define.Obj
	0, // 5: define.HardwareInfo.network:type_name -> define.Obj
	1, // 6: define.HardwareInfo.other:type_name -> define.HardwareOther
	3, // 7: define.PostInterface.domains:type_name -> define.PostDomain
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_protobuf_defineCfg_proto_init() }
func file_protobuf_defineCfg_proto_init() {
	if File_protobuf_defineCfg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_defineCfg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Obj); i {
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
		file_protobuf_defineCfg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HardwareOther); i {
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
		file_protobuf_defineCfg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HardwareInfo); i {
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
		file_protobuf_defineCfg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostDomain); i {
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
		file_protobuf_defineCfg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostInterface); i {
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
		file_protobuf_defineCfg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysCfg); i {
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
		file_protobuf_defineCfg_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RsaKey); i {
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
			RawDescriptor: file_protobuf_defineCfg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_defineCfg_proto_goTypes,
		DependencyIndexes: file_protobuf_defineCfg_proto_depIdxs,
		MessageInfos:      file_protobuf_defineCfg_proto_msgTypes,
	}.Build()
	File_protobuf_defineCfg_proto = out.File
	file_protobuf_defineCfg_proto_rawDesc = nil
	file_protobuf_defineCfg_proto_goTypes = nil
	file_protobuf_defineCfg_proto_depIdxs = nil
}

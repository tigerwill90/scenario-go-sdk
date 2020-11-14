// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: proto/scenario/scenario.proto

package scenario

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Scenario struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unique identifier for the scenario
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// friendly name for the scenario
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Created string `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	// digest of the scenario
	Digest uint64 `protobuf:"varint,4,opt,name=digest,proto3" json:"digest,omitempty"`
	// actual hcl file
	File []byte `protobuf:"bytes,5,opt,name=file,proto3" json:"file,omitempty"`
	// recovery time objective is the excepted time to recover a backup
	Rto string `protobuf:"bytes,6,opt,name=rto,proto3" json:"rto,omitempty"`
	// time after which we consider the scenario has failed
	Timeout string `protobuf:"bytes,7,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// filter used for test case
	Filter map[string]*Scenario_Paths `protobuf:"bytes,8,rep,name=filter,proto3" json:"filter,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// docker image provided to create the test environment
	EnvImage string `protobuf:"bytes,9,opt,name=env_image,json=envImage,proto3" json:"env_image,omitempty"`
	// docker image version provided to create the test environment
	EnvVersion string `protobuf:"bytes,10,opt,name=env_version,json=envVersion,proto3" json:"env_version,omitempty"`
	// opaque config returned by the plugin after parsing the schema and validating the configuration
	EnvConfig []byte `protobuf:"bytes,11,opt,name=env_config,json=envConfig,proto3" json:"env_config,omitempty"`
	// all case and configuration
	Cases []*Scenario_TestCase `protobuf:"bytes,12,rep,name=cases,proto3" json:"cases,omitempty"`
}

func (x *Scenario) Reset() {
	*x = Scenario{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scenario_scenario_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scenario) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scenario) ProtoMessage() {}

func (x *Scenario) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scenario_scenario_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scenario.ProtoReflect.Descriptor instead.
func (*Scenario) Descriptor() ([]byte, []int) {
	return file_proto_scenario_scenario_proto_rawDescGZIP(), []int{0}
}

func (x *Scenario) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Scenario) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Scenario) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *Scenario) GetDigest() uint64 {
	if x != nil {
		return x.Digest
	}
	return 0
}

func (x *Scenario) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *Scenario) GetRto() string {
	if x != nil {
		return x.Rto
	}
	return ""
}

func (x *Scenario) GetTimeout() string {
	if x != nil {
		return x.Timeout
	}
	return ""
}

func (x *Scenario) GetFilter() map[string]*Scenario_Paths {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *Scenario) GetEnvImage() string {
	if x != nil {
		return x.EnvImage
	}
	return ""
}

func (x *Scenario) GetEnvVersion() string {
	if x != nil {
		return x.EnvVersion
	}
	return ""
}

func (x *Scenario) GetEnvConfig() []byte {
	if x != nil {
		return x.EnvConfig
	}
	return nil
}

func (x *Scenario) GetCases() []*Scenario_TestCase {
	if x != nil {
		return x.Cases
	}
	return nil
}

// multiple path for a filter key
type Scenario_Paths struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paths []string `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
}

func (x *Scenario_Paths) Reset() {
	*x = Scenario_Paths{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scenario_scenario_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scenario_Paths) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scenario_Paths) ProtoMessage() {}

func (x *Scenario_Paths) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scenario_scenario_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scenario_Paths.ProtoReflect.Descriptor instead.
func (*Scenario_Paths) Descriptor() ([]byte, []int) {
	return file_proto_scenario_scenario_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Scenario_Paths) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

// group of query to which compose a test case
type Scenario_TestCase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string                          `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	QueryGroups []*Scenario_TestCase_QueryGroup `protobuf:"bytes,2,rep,name=query_groups,json=queryGroups,proto3" json:"query_groups,omitempty"`
}

func (x *Scenario_TestCase) Reset() {
	*x = Scenario_TestCase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scenario_scenario_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scenario_TestCase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scenario_TestCase) ProtoMessage() {}

func (x *Scenario_TestCase) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scenario_scenario_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scenario_TestCase.ProtoReflect.Descriptor instead.
func (*Scenario_TestCase) Descriptor() ([]byte, []int) {
	return file_proto_scenario_scenario_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Scenario_TestCase) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Scenario_TestCase) GetQueryGroups() []*Scenario_TestCase_QueryGroup {
	if x != nil {
		return x.QueryGroups
	}
	return nil
}

// operation to perform on a filtered file with the submitted config
type Scenario_TestCase_Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op       string `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	Variable string `protobuf:"bytes,2,opt,name=variable,proto3" json:"variable,omitempty"`
	Filter   string `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	Config   []byte `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *Scenario_TestCase_Query) Reset() {
	*x = Scenario_TestCase_Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scenario_scenario_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scenario_TestCase_Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scenario_TestCase_Query) ProtoMessage() {}

func (x *Scenario_TestCase_Query) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scenario_scenario_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scenario_TestCase_Query.ProtoReflect.Descriptor instead.
func (*Scenario_TestCase_Query) Descriptor() ([]byte, []int) {
	return file_proto_scenario_scenario_proto_rawDescGZIP(), []int{0, 2, 0}
}

func (x *Scenario_TestCase_Query) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *Scenario_TestCase_Query) GetVariable() string {
	if x != nil {
		return x.Variable
	}
	return ""
}

func (x *Scenario_TestCase_Query) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *Scenario_TestCase_Query) GetConfig() []byte {
	if x != nil {
		return x.Config
	}
	return nil
}

// for_each operand can be use to process multiple nested query in a row
type Scenario_TestCase_QueryGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Assert string `protobuf:"bytes,1,opt,name=assert,proto3" json:"assert,omitempty"`
	// query can be nested in order to compose expressive test from small operation
	NestedQueries []*Scenario_TestCase_Query `protobuf:"bytes,2,rep,name=nested_queries,json=nestedQueries,proto3" json:"nested_queries,omitempty"`
}

func (x *Scenario_TestCase_QueryGroup) Reset() {
	*x = Scenario_TestCase_QueryGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scenario_scenario_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scenario_TestCase_QueryGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scenario_TestCase_QueryGroup) ProtoMessage() {}

func (x *Scenario_TestCase_QueryGroup) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scenario_scenario_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scenario_TestCase_QueryGroup.ProtoReflect.Descriptor instead.
func (*Scenario_TestCase_QueryGroup) Descriptor() ([]byte, []int) {
	return file_proto_scenario_scenario_proto_rawDescGZIP(), []int{0, 2, 1}
}

func (x *Scenario_TestCase_QueryGroup) GetAssert() string {
	if x != nil {
		return x.Assert
	}
	return ""
}

func (x *Scenario_TestCase_QueryGroup) GetNestedQueries() []*Scenario_TestCase_Query {
	if x != nil {
		return x.NestedQueries
	}
	return nil
}

var File_proto_scenario_scenario_proto protoreflect.FileDescriptor

var file_proto_scenario_scenario_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f,
	0x2f, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x22, 0xaf, 0x06, 0x0a, 0x08, 0x53, 0x63,
	0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x74, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x72, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x12, 0x36, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e, 0x53, 0x63, 0x65, 0x6e,
	0x61, 0x72, 0x69, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x76, 0x5f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x76,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x76, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x76, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x65, 0x6e, 0x76, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x31, 0x0a, 0x05, 0x63, 0x61, 0x73, 0x65, 0x73, 0x18, 0x0c,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e,
	0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73,
	0x65, 0x52, 0x05, 0x63, 0x61, 0x73, 0x65, 0x73, 0x1a, 0x1d, 0x0a, 0x05, 0x50, 0x61, 0x74, 0x68,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x1a, 0x53, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72,
	0x69, 0x6f, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e, 0x50, 0x61, 0x74, 0x68,
	0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xcc, 0x02, 0x0a,
	0x08, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x0c, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e, 0x53, 0x63, 0x65,
	0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x1a, 0x63, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x70, 0x12,
	0x1a, 0x0a, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x6e, 0x0a, 0x0a, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x73, 0x73,
	0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x73, 0x73, 0x65, 0x72,
	0x74, 0x12, 0x48, 0x0a, 0x0e, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x71, 0x75, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x63, 0x65, 0x6e,
	0x61, 0x72, 0x69, 0x6f, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x0d, 0x6e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x42, 0x10, 0x5a, 0x0e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_scenario_scenario_proto_rawDescOnce sync.Once
	file_proto_scenario_scenario_proto_rawDescData = file_proto_scenario_scenario_proto_rawDesc
)

func file_proto_scenario_scenario_proto_rawDescGZIP() []byte {
	file_proto_scenario_scenario_proto_rawDescOnce.Do(func() {
		file_proto_scenario_scenario_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_scenario_scenario_proto_rawDescData)
	})
	return file_proto_scenario_scenario_proto_rawDescData
}

var file_proto_scenario_scenario_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_scenario_scenario_proto_goTypes = []interface{}{
	(*Scenario)(nil),                     // 0: scenario.Scenario
	(*Scenario_Paths)(nil),               // 1: scenario.Scenario.Paths
	nil,                                  // 2: scenario.Scenario.FilterEntry
	(*Scenario_TestCase)(nil),            // 3: scenario.Scenario.TestCase
	(*Scenario_TestCase_Query)(nil),      // 4: scenario.Scenario.TestCase.Query
	(*Scenario_TestCase_QueryGroup)(nil), // 5: scenario.Scenario.TestCase.QueryGroup
}
var file_proto_scenario_scenario_proto_depIdxs = []int32{
	2, // 0: scenario.Scenario.filter:type_name -> scenario.Scenario.FilterEntry
	3, // 1: scenario.Scenario.cases:type_name -> scenario.Scenario.TestCase
	1, // 2: scenario.Scenario.FilterEntry.value:type_name -> scenario.Scenario.Paths
	5, // 3: scenario.Scenario.TestCase.query_groups:type_name -> scenario.Scenario.TestCase.QueryGroup
	4, // 4: scenario.Scenario.TestCase.QueryGroup.nested_queries:type_name -> scenario.Scenario.TestCase.Query
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_scenario_scenario_proto_init() }
func file_proto_scenario_scenario_proto_init() {
	if File_proto_scenario_scenario_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_scenario_scenario_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scenario); i {
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
		file_proto_scenario_scenario_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scenario_Paths); i {
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
		file_proto_scenario_scenario_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scenario_TestCase); i {
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
		file_proto_scenario_scenario_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scenario_TestCase_Query); i {
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
		file_proto_scenario_scenario_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scenario_TestCase_QueryGroup); i {
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
			RawDescriptor: file_proto_scenario_scenario_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_scenario_scenario_proto_goTypes,
		DependencyIndexes: file_proto_scenario_scenario_proto_depIdxs,
		MessageInfos:      file_proto_scenario_scenario_proto_msgTypes,
	}.Build()
	File_proto_scenario_scenario_proto = out.File
	file_proto_scenario_scenario_proto_rawDesc = nil
	file_proto_scenario_scenario_proto_goTypes = nil
	file_proto_scenario_scenario_proto_depIdxs = nil
}
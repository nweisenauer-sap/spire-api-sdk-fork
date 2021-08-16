// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: spire/api/types/selector.proto

package types

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

type SelectorMatch_MatchBehavior int32

const (
	// Indicates that the selectors in this match are equal to the
	// candidate selectors, independent of ordering.
	// Example:
	//   Given:
	//     - 'e1 { Selectors: ["a:1", "b:2", "c:3"]}'
	//     - 'e2 { Selectors: ["a:1", "b:2"]}'
	//     - 'e3 { Selectors: ["a:1]"}'
	//   Operation:
	//     - MATCH_EXACT ["a:1", "b:2"]
	//   Entries that match:
	//     - 'e2'
	SelectorMatch_MATCH_EXACT SelectorMatch_MatchBehavior = 0
	// Indicates that all candidates which have a non-empty subset
	// of the provided set of selectors will match.
	// Example:
	//   Given:
	//     - 'e1 { Selectors: ["a:1", "b:2", "c:3"]}'
	//     - 'e2 { Selectors: ["a:1", "b:2"]}'
	//     - 'e3 { Selectors: ["a:1]"}'
	//   Operation:
	//     - MATCH_SUBSET ["a:1"]
	//   Entries that match:
	//     - 'e1'
	SelectorMatch_MATCH_SUBSET SelectorMatch_MatchBehavior = 1
	// Indicates that all candidates which are a superset
	// of the provided selectors will match.
	// Example:
	//   Given:
	//     - 'e1 { Selectors: ["a:1", "b:2", "c:3"]}'
	//     - 'e2 { Selectors: ["a:1", "b:2"]}'
	//     - 'e3 { Selectors: ["a:1]"}'
	//   Operation:
	//     - MATCH_SUPERSET ["a:1", "b:2"]
	//   Entries that match:
	//     - 'e1'
	//     - 'e2'
	SelectorMatch_MATCH_SUPERSET SelectorMatch_MatchBehavior = 2
	// Indicates that all candidates which have at least one
	// of the provided set of selectors will match.
	// Example:
	//   Given:
	//     - 'e1 { Selectors: ["a:1", "b:2", "c:3"]}'
	//     - 'e2 { Selectors: ["a:1", "b:2"]}'
	//     - 'e3 { Selectors: ["a:1]"}'
	//   Operation:
	//     - MATCH_ANY ["a:1"]
	//   Entries that match:
	//     - 'e1'
	//     - 'e2'
	//     - 'e3'
	SelectorMatch_MATCH_ANY SelectorMatch_MatchBehavior = 3
)

// Enum value maps for SelectorMatch_MatchBehavior.
var (
	SelectorMatch_MatchBehavior_name = map[int32]string{
		0: "MATCH_EXACT",
		1: "MATCH_SUBSET",
		2: "MATCH_SUPERSET",
		3: "MATCH_ANY",
	}
	SelectorMatch_MatchBehavior_value = map[string]int32{
		"MATCH_EXACT":    0,
		"MATCH_SUBSET":   1,
		"MATCH_SUPERSET": 2,
		"MATCH_ANY":      3,
	}
)

func (x SelectorMatch_MatchBehavior) Enum() *SelectorMatch_MatchBehavior {
	p := new(SelectorMatch_MatchBehavior)
	*p = x
	return p
}

func (x SelectorMatch_MatchBehavior) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SelectorMatch_MatchBehavior) Descriptor() protoreflect.EnumDescriptor {
	return file_spire_api_types_selector_proto_enumTypes[0].Descriptor()
}

func (SelectorMatch_MatchBehavior) Type() protoreflect.EnumType {
	return &file_spire_api_types_selector_proto_enumTypes[0]
}

func (x SelectorMatch_MatchBehavior) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SelectorMatch_MatchBehavior.Descriptor instead.
func (SelectorMatch_MatchBehavior) EnumDescriptor() ([]byte, []int) {
	return file_spire_api_types_selector_proto_rawDescGZIP(), []int{1, 0}
}

type Selector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the selector. This is typically the name of the plugin that
	// produces the selector.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// The value of the selector.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Selector) Reset() {
	*x = Selector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_types_selector_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Selector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Selector) ProtoMessage() {}

func (x *Selector) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_types_selector_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Selector.ProtoReflect.Descriptor instead.
func (*Selector) Descriptor() ([]byte, []int) {
	return file_spire_api_types_selector_proto_rawDescGZIP(), []int{0}
}

func (x *Selector) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Selector) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SelectorMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The set of selectors to match on.
	Selectors []*Selector `protobuf:"bytes,1,rep,name=selectors,proto3" json:"selectors,omitempty"`
	// How to match the selectors.
	Match SelectorMatch_MatchBehavior `protobuf:"varint,2,opt,name=match,proto3,enum=spire.api.types.SelectorMatch_MatchBehavior" json:"match,omitempty"`
}

func (x *SelectorMatch) Reset() {
	*x = SelectorMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_types_selector_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectorMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectorMatch) ProtoMessage() {}

func (x *SelectorMatch) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_types_selector_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectorMatch.ProtoReflect.Descriptor instead.
func (*SelectorMatch) Descriptor() ([]byte, []int) {
	return file_spire_api_types_selector_proto_rawDescGZIP(), []int{1}
}

func (x *SelectorMatch) GetSelectors() []*Selector {
	if x != nil {
		return x.Selectors
	}
	return nil
}

func (x *SelectorMatch) GetMatch() SelectorMatch_MatchBehavior {
	if x != nil {
		return x.Match
	}
	return SelectorMatch_MATCH_EXACT
}

var File_spire_api_types_selector_proto protoreflect.FileDescriptor

var file_spire_api_types_selector_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x22, 0x34, 0x0a, 0x08, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xe3, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x12, 0x42, 0x0a, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2c, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x52,
	0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x22, 0x55, 0x0a, 0x0d, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x42,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x41, 0x54, 0x43, 0x48,
	0x5f, 0x45, 0x58, 0x41, 0x43, 0x54, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x41, 0x54, 0x43,
	0x48, 0x5f, 0x53, 0x55, 0x42, 0x53, 0x45, 0x54, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x41,
	0x54, 0x43, 0x48, 0x5f, 0x53, 0x55, 0x50, 0x45, 0x52, 0x53, 0x45, 0x54, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x41, 0x4e, 0x59, 0x10, 0x03, 0x42, 0x37, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x66,
	0x66, 0x65, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spire_api_types_selector_proto_rawDescOnce sync.Once
	file_spire_api_types_selector_proto_rawDescData = file_spire_api_types_selector_proto_rawDesc
)

func file_spire_api_types_selector_proto_rawDescGZIP() []byte {
	file_spire_api_types_selector_proto_rawDescOnce.Do(func() {
		file_spire_api_types_selector_proto_rawDescData = protoimpl.X.CompressGZIP(file_spire_api_types_selector_proto_rawDescData)
	})
	return file_spire_api_types_selector_proto_rawDescData
}

var file_spire_api_types_selector_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spire_api_types_selector_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_spire_api_types_selector_proto_goTypes = []interface{}{
	(SelectorMatch_MatchBehavior)(0), // 0: spire.api.types.SelectorMatch.MatchBehavior
	(*Selector)(nil),                 // 1: spire.api.types.Selector
	(*SelectorMatch)(nil),            // 2: spire.api.types.SelectorMatch
}
var file_spire_api_types_selector_proto_depIdxs = []int32{
	1, // 0: spire.api.types.SelectorMatch.selectors:type_name -> spire.api.types.Selector
	0, // 1: spire.api.types.SelectorMatch.match:type_name -> spire.api.types.SelectorMatch.MatchBehavior
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_spire_api_types_selector_proto_init() }
func file_spire_api_types_selector_proto_init() {
	if File_spire_api_types_selector_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spire_api_types_selector_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Selector); i {
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
		file_spire_api_types_selector_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectorMatch); i {
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
			RawDescriptor: file_spire_api_types_selector_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spire_api_types_selector_proto_goTypes,
		DependencyIndexes: file_spire_api_types_selector_proto_depIdxs,
		EnumInfos:         file_spire_api_types_selector_proto_enumTypes,
		MessageInfos:      file_spire_api_types_selector_proto_msgTypes,
	}.Build()
	File_spire_api_types_selector_proto = out.File
	file_spire_api_types_selector_proto_rawDesc = nil
	file_spire_api_types_selector_proto_goTypes = nil
	file_spire_api_types_selector_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: pkg/generator/changelog_generator.proto

package generator

import (
	semrel "github.com/go-semantic-release/semantic-release/pkg/semrel"
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

type ChangelogGeneratorConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commits       []*semrel.Commit `protobuf:"bytes,1,rep,name=Commits,proto3" json:"Commits,omitempty"`
	LatestRelease *semrel.Release  `protobuf:"bytes,2,opt,name=LatestRelease,proto3" json:"LatestRelease,omitempty"`
	NewVersion    string           `protobuf:"bytes,3,opt,name=NewVersion,proto3" json:"NewVersion,omitempty"`
}

func (x *ChangelogGeneratorConfig) Reset() {
	*x = ChangelogGeneratorConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_generator_changelog_generator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangelogGeneratorConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangelogGeneratorConfig) ProtoMessage() {}

func (x *ChangelogGeneratorConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_generator_changelog_generator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangelogGeneratorConfig.ProtoReflect.Descriptor instead.
func (*ChangelogGeneratorConfig) Descriptor() ([]byte, []int) {
	return file_pkg_generator_changelog_generator_proto_rawDescGZIP(), []int{0}
}

func (x *ChangelogGeneratorConfig) GetCommits() []*semrel.Commit {
	if x != nil {
		return x.Commits
	}
	return nil
}

func (x *ChangelogGeneratorConfig) GetLatestRelease() *semrel.Release {
	if x != nil {
		return x.LatestRelease
	}
	return nil
}

func (x *ChangelogGeneratorConfig) GetNewVersion() string {
	if x != nil {
		return x.NewVersion
	}
	return ""
}

type GenerateChangelog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GenerateChangelog) Reset() {
	*x = GenerateChangelog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_generator_changelog_generator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateChangelog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateChangelog) ProtoMessage() {}

func (x *GenerateChangelog) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_generator_changelog_generator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateChangelog.ProtoReflect.Descriptor instead.
func (*GenerateChangelog) Descriptor() ([]byte, []int) {
	return file_pkg_generator_changelog_generator_proto_rawDescGZIP(), []int{1}
}

type GenerateChangelog_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *ChangelogGeneratorConfig `protobuf:"bytes,1,opt,name=Config,proto3" json:"Config,omitempty"`
}

func (x *GenerateChangelog_Request) Reset() {
	*x = GenerateChangelog_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_generator_changelog_generator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateChangelog_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateChangelog_Request) ProtoMessage() {}

func (x *GenerateChangelog_Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_generator_changelog_generator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateChangelog_Request.ProtoReflect.Descriptor instead.
func (*GenerateChangelog_Request) Descriptor() ([]byte, []int) {
	return file_pkg_generator_changelog_generator_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GenerateChangelog_Request) GetConfig() *ChangelogGeneratorConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type GenerateChangelog_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Changelog string `protobuf:"bytes,1,opt,name=Changelog,proto3" json:"Changelog,omitempty"`
}

func (x *GenerateChangelog_Response) Reset() {
	*x = GenerateChangelog_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_generator_changelog_generator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateChangelog_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateChangelog_Response) ProtoMessage() {}

func (x *GenerateChangelog_Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_generator_changelog_generator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateChangelog_Response.ProtoReflect.Descriptor instead.
func (*GenerateChangelog_Response) Descriptor() ([]byte, []int) {
	return file_pkg_generator_changelog_generator_proto_rawDescGZIP(), []int{1, 1}
}

func (x *GenerateChangelog_Response) GetChangelog() string {
	if x != nil {
		return x.Changelog
	}
	return ""
}

var File_pkg_generator_changelog_generator_proto protoreflect.FileDescriptor

var file_pkg_generator_changelog_generator_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x73,
	0x65, 0x6d, 0x72, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x18, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f,
	0x67, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x21, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x73, 0x12, 0x2e, 0x0a, 0x0d, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x52, 0x0d, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4e, 0x65, 0x77, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x7b, 0x0a, 0x11, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x1a, 0x3c, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x28, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67,
	0x32, 0x5f, 0x0a, 0x18, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x43, 0x0a, 0x08,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2d, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x2d, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x2f, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x2d, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_generator_changelog_generator_proto_rawDescOnce sync.Once
	file_pkg_generator_changelog_generator_proto_rawDescData = file_pkg_generator_changelog_generator_proto_rawDesc
)

func file_pkg_generator_changelog_generator_proto_rawDescGZIP() []byte {
	file_pkg_generator_changelog_generator_proto_rawDescOnce.Do(func() {
		file_pkg_generator_changelog_generator_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_generator_changelog_generator_proto_rawDescData)
	})
	return file_pkg_generator_changelog_generator_proto_rawDescData
}

var file_pkg_generator_changelog_generator_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_generator_changelog_generator_proto_goTypes = []interface{}{
	(*ChangelogGeneratorConfig)(nil),   // 0: ChangelogGeneratorConfig
	(*GenerateChangelog)(nil),          // 1: GenerateChangelog
	(*GenerateChangelog_Request)(nil),  // 2: GenerateChangelog.Request
	(*GenerateChangelog_Response)(nil), // 3: GenerateChangelog.Response
	(*semrel.Commit)(nil),              // 4: Commit
	(*semrel.Release)(nil),             // 5: Release
}
var file_pkg_generator_changelog_generator_proto_depIdxs = []int32{
	4, // 0: ChangelogGeneratorConfig.Commits:type_name -> Commit
	5, // 1: ChangelogGeneratorConfig.LatestRelease:type_name -> Release
	0, // 2: GenerateChangelog.Request.Config:type_name -> ChangelogGeneratorConfig
	2, // 3: ChangelogGeneratorPlugin.Generate:input_type -> GenerateChangelog.Request
	3, // 4: ChangelogGeneratorPlugin.Generate:output_type -> GenerateChangelog.Response
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_generator_changelog_generator_proto_init() }
func file_pkg_generator_changelog_generator_proto_init() {
	if File_pkg_generator_changelog_generator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_generator_changelog_generator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangelogGeneratorConfig); i {
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
		file_pkg_generator_changelog_generator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateChangelog); i {
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
		file_pkg_generator_changelog_generator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateChangelog_Request); i {
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
		file_pkg_generator_changelog_generator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateChangelog_Response); i {
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
			RawDescriptor: file_pkg_generator_changelog_generator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_generator_changelog_generator_proto_goTypes,
		DependencyIndexes: file_pkg_generator_changelog_generator_proto_depIdxs,
		MessageInfos:      file_pkg_generator_changelog_generator_proto_msgTypes,
	}.Build()
	File_pkg_generator_changelog_generator_proto = out.File
	file_pkg_generator_changelog_generator_proto_rawDesc = nil
	file_pkg_generator_changelog_generator_proto_goTypes = nil
	file_pkg_generator_changelog_generator_proto_depIdxs = nil
}
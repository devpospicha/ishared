// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: stream.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MetricWrapper struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RawPayload    []byte                 `protobuf:"bytes,1,opt,name=raw_payload,json=rawPayload,proto3" json:"raw_payload,omitempty"` // serialized MetricPayload manually
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetricWrapper) Reset() {
	*x = MetricWrapper{}
	mi := &file_stream_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricWrapper) ProtoMessage() {}

func (x *MetricWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricWrapper.ProtoReflect.Descriptor instead.
func (*MetricWrapper) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{0}
}

func (x *MetricWrapper) GetRawPayload() []byte {
	if x != nil {
		return x.RawPayload
	}
	return nil
}

type ProcessWrapper struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RawPayload    []byte                 `protobuf:"bytes,1,opt,name=raw_payload,json=rawPayload,proto3" json:"raw_payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProcessWrapper) Reset() {
	*x = ProcessWrapper{}
	mi := &file_stream_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessWrapper) ProtoMessage() {}

func (x *ProcessWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessWrapper.ProtoReflect.Descriptor instead.
func (*ProcessWrapper) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{1}
}

func (x *ProcessWrapper) GetRawPayload() []byte {
	if x != nil {
		return x.RawPayload
	}
	return nil
}

type StreamPayload struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Payload:
	//
	//	*StreamPayload_Metric
	//	*StreamPayload_CommandRequest
	//	*StreamPayload_CommandResponse
	//	*StreamPayload_Process
	Payload       isStreamPayload_Payload `protobuf_oneof:"payload"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamPayload) Reset() {
	*x = StreamPayload{}
	mi := &file_stream_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamPayload) ProtoMessage() {}

func (x *StreamPayload) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamPayload.ProtoReflect.Descriptor instead.
func (*StreamPayload) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{2}
}

func (x *StreamPayload) GetPayload() isStreamPayload_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *StreamPayload) GetMetric() *MetricWrapper {
	if x != nil {
		if x, ok := x.Payload.(*StreamPayload_Metric); ok {
			return x.Metric
		}
	}
	return nil
}

func (x *StreamPayload) GetCommandRequest() *CommandRequest {
	if x != nil {
		if x, ok := x.Payload.(*StreamPayload_CommandRequest); ok {
			return x.CommandRequest
		}
	}
	return nil
}

func (x *StreamPayload) GetCommandResponse() *CommandResponse {
	if x != nil {
		if x, ok := x.Payload.(*StreamPayload_CommandResponse); ok {
			return x.CommandResponse
		}
	}
	return nil
}

func (x *StreamPayload) GetProcess() *ProcessWrapper {
	if x != nil {
		if x, ok := x.Payload.(*StreamPayload_Process); ok {
			return x.Process
		}
	}
	return nil
}

type isStreamPayload_Payload interface {
	isStreamPayload_Payload()
}

type StreamPayload_Metric struct {
	Metric *MetricWrapper `protobuf:"bytes,1,opt,name=metric,proto3,oneof"`
}

type StreamPayload_CommandRequest struct {
	CommandRequest *CommandRequest `protobuf:"bytes,2,opt,name=command_request,json=commandRequest,proto3,oneof"`
}

type StreamPayload_CommandResponse struct {
	CommandResponse *CommandResponse `protobuf:"bytes,3,opt,name=command_response,json=commandResponse,proto3,oneof"`
}

type StreamPayload_Process struct {
	Process *ProcessWrapper `protobuf:"bytes,5,opt,name=process,proto3,oneof"`
}

func (*StreamPayload_Metric) isStreamPayload_Payload() {}

func (*StreamPayload_CommandRequest) isStreamPayload_Payload() {}

func (*StreamPayload_CommandResponse) isStreamPayload_Payload() {}

func (*StreamPayload_Process) isStreamPayload_Payload() {}

type StreamResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	StatusCode    int32                  `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Command       *CommandRequest        `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"` // optional
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamResponse) Reset() {
	*x = StreamResponse{}
	mi := &file_stream_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponse) ProtoMessage() {}

func (x *StreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponse.ProtoReflect.Descriptor instead.
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{3}
}

func (x *StreamResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *StreamResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *StreamResponse) GetCommand() *CommandRequest {
	if x != nil {
		return x.Command
	}
	return nil
}

var File_stream_proto protoreflect.FileDescriptor

const file_stream_proto_rawDesc = "" +
	"\n" +
	"\fstream.proto\x12\x05proto\x1a\rcommand.proto\"0\n" +
	"\rMetricWrapper\x12\x1f\n" +
	"\vraw_payload\x18\x01 \x01(\fR\n" +
	"rawPayload\"1\n" +
	"\x0eProcessWrapper\x12\x1f\n" +
	"\vraw_payload\x18\x01 \x01(\fR\n" +
	"rawPayload\"\x84\x02\n" +
	"\rStreamPayload\x12.\n" +
	"\x06metric\x18\x01 \x01(\v2\x14.proto.MetricWrapperH\x00R\x06metric\x12@\n" +
	"\x0fcommand_request\x18\x02 \x01(\v2\x15.proto.CommandRequestH\x00R\x0ecommandRequest\x12C\n" +
	"\x10command_response\x18\x03 \x01(\v2\x16.proto.CommandResponseH\x00R\x0fcommandResponse\x121\n" +
	"\aprocess\x18\x05 \x01(\v2\x15.proto.ProcessWrapperH\x00R\aprocessB\t\n" +
	"\apayload\"z\n" +
	"\x0eStreamResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status\x12\x1f\n" +
	"\vstatus_code\x18\x02 \x01(\x05R\n" +
	"statusCode\x12/\n" +
	"\acommand\x18\x03 \x01(\v2\x15.proto.CommandRequestR\acommand2J\n" +
	"\rStreamService\x129\n" +
	"\x06Stream\x12\x14.proto.StreamPayload\x1a\x15.proto.StreamResponse(\x010\x01B.Z,github.com/devpospicha/ishared/protob\x06proto3"

var (
	file_stream_proto_rawDescOnce sync.Once
	file_stream_proto_rawDescData []byte
)

func file_stream_proto_rawDescGZIP() []byte {
	file_stream_proto_rawDescOnce.Do(func() {
		file_stream_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_stream_proto_rawDesc), len(file_stream_proto_rawDesc)))
	})
	return file_stream_proto_rawDescData
}

var file_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_stream_proto_goTypes = []any{
	(*MetricWrapper)(nil),   // 0: proto.MetricWrapper
	(*ProcessWrapper)(nil),  // 1: proto.ProcessWrapper
	(*StreamPayload)(nil),   // 2: proto.StreamPayload
	(*StreamResponse)(nil),  // 3: proto.StreamResponse
	(*CommandRequest)(nil),  // 4: proto.CommandRequest
	(*CommandResponse)(nil), // 5: proto.CommandResponse
}
var file_stream_proto_depIdxs = []int32{
	0, // 0: proto.StreamPayload.metric:type_name -> proto.MetricWrapper
	4, // 1: proto.StreamPayload.command_request:type_name -> proto.CommandRequest
	5, // 2: proto.StreamPayload.command_response:type_name -> proto.CommandResponse
	1, // 3: proto.StreamPayload.process:type_name -> proto.ProcessWrapper
	4, // 4: proto.StreamResponse.command:type_name -> proto.CommandRequest
	2, // 5: proto.StreamService.Stream:input_type -> proto.StreamPayload
	3, // 6: proto.StreamService.Stream:output_type -> proto.StreamResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_stream_proto_init() }
func file_stream_proto_init() {
	if File_stream_proto != nil {
		return
	}
	file_command_proto_init()
	file_stream_proto_msgTypes[2].OneofWrappers = []any{
		(*StreamPayload_Metric)(nil),
		(*StreamPayload_CommandRequest)(nil),
		(*StreamPayload_CommandResponse)(nil),
		(*StreamPayload_Process)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_stream_proto_rawDesc), len(file_stream_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stream_proto_goTypes,
		DependencyIndexes: file_stream_proto_depIdxs,
		MessageInfos:      file_stream_proto_msgTypes,
	}.Build()
	File_stream_proto = out.File
	file_stream_proto_goTypes = nil
	file_stream_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: Connection.proto

package connection

import (
	context "context"
	reflect "reflect"
	sync "sync"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk       []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	NChunk      int32  `protobuf:"varint,2,opt,name=nChunk,proto3" json:"nChunk,omitempty"`
	NombreLibro string `protobuf:"bytes,3,opt,name=NombreLibro,proto3" json:"NombreLibro,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Connection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_Connection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_Connection_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

func (x *Chunk) GetNChunk() int32 {
	if x != nil {
		return x.NChunk
	}
	return 0
}

func (x *Chunk) GetNombreLibro() string {
	if x != nil {
		return x.NombreLibro
	}
	return ""
}

type NombreLibro struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NombreLibro string `protobuf:"bytes,1,opt,name=NombreLibro,proto3" json:"NombreLibro,omitempty"`
}

func (x *NombreLibro) Reset() {
	*x = NombreLibro{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Connection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NombreLibro) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NombreLibro) ProtoMessage() {}

func (x *NombreLibro) ProtoReflect() protoreflect.Message {
	mi := &file_Connection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NombreLibro.ProtoReflect.Descriptor instead.
func (*NombreLibro) Descriptor() ([]byte, []int) {
	return file_Connection_proto_rawDescGZIP(), []int{1}
}

func (x *NombreLibro) GetNombreLibro() string {
	if x != nil {
		return x.NombreLibro
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Connection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_Connection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_Connection_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DivisionLibro struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NombreLibro string `protobuf:"bytes,1,opt,name=NombreLibro,proto3" json:"NombreLibro,omitempty"`
	NChunk      int32  `protobuf:"varint,2,opt,name=nChunk,proto3" json:"nChunk,omitempty"`
}

func (x *DivisionLibro) Reset() {
	*x = DivisionLibro{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Connection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DivisionLibro) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DivisionLibro) ProtoMessage() {}

func (x *DivisionLibro) ProtoReflect() protoreflect.Message {
	mi := &file_Connection_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DivisionLibro.ProtoReflect.Descriptor instead.
func (*DivisionLibro) Descriptor() ([]byte, []int) {
	return file_Connection_proto_rawDescGZIP(), []int{3}
}

func (x *DivisionLibro) GetNombreLibro() string {
	if x != nil {
		return x.NombreLibro
	}
	return ""
}

func (x *DivisionLibro) GetNChunk() int32 {
	if x != nil {
		return x.NChunk
	}
	return 0
}

type Distribucion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NombreLibro         string  `protobuf:"bytes,1,opt,name=NombreLibro,proto3" json:"NombreLibro,omitempty"`
	NumeroPar           int32   `protobuf:"varint,2,opt,name=numeroPar,proto3" json:"numeroPar,omitempty"`
	ListaDataNodesChunk []int32 `protobuf:"varint,3,rep,packed,name=listaDataNodesChunk,proto3" json:"listaDataNodesChunk,omitempty"`
}

func (x *Distribucion) Reset() {
	*x = Distribucion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Connection_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Distribucion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Distribucion) ProtoMessage() {}

func (x *Distribucion) ProtoReflect() protoreflect.Message {
	mi := &file_Connection_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Distribucion.ProtoReflect.Descriptor instead.
func (*Distribucion) Descriptor() ([]byte, []int) {
	return file_Connection_proto_rawDescGZIP(), []int{4}
}

func (x *Distribucion) GetNombreLibro() string {
	if x != nil {
		return x.NombreLibro
	}
	return ""
}

func (x *Distribucion) GetNumeroPar() int32 {
	if x != nil {
		return x.NumeroPar
	}
	return 0
}

func (x *Distribucion) GetListaDataNodesChunk() []int32 {
	if x != nil {
		return x.ListaDataNodesChunk
	}
	return nil
}

type Libros struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LibrosDisponibles []string `protobuf:"bytes,1,rep,name=librosDisponibles,proto3" json:"librosDisponibles,omitempty"`
}

func (x *Libros) Reset() {
	*x = Libros{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Connection_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Libros) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Libros) ProtoMessage() {}

func (x *Libros) ProtoReflect() protoreflect.Message {
	mi := &file_Connection_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Libros.ProtoReflect.Descriptor instead.
func (*Libros) Descriptor() ([]byte, []int) {
	return file_Connection_proto_rawDescGZIP(), []int{5}
}

func (x *Libros) GetLibrosDisponibles() []string {
	if x != nil {
		return x.LibrosDisponibles
	}
	return nil
}

var File_Connection_proto protoreflect.FileDescriptor

var file_Connection_proto_rawDesc = []byte{
	0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x57,
	0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x4c,
	0x69, 0x62, 0x72, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4e, 0x6f, 0x6d, 0x62,
	0x72, 0x65, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x22, 0x2f, 0x0a, 0x0b, 0x4e, 0x6f, 0x6d, 0x62, 0x72,
	0x65, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65,
	0x4c, 0x69, 0x62, 0x72, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4e, 0x6f, 0x6d,
	0x62, 0x72, 0x65, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x22, 0x23, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x49, 0x0a,
	0x0d, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x12, 0x20,
	0x0a, 0x0b, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x4c, 0x69, 0x62, 0x72, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x80, 0x01, 0x0a, 0x0c, 0x44, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x6f, 0x6d,
	0x62, 0x72, 0x65, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x75, 0x6d, 0x65, 0x72, 0x6f, 0x50, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x50, 0x61, 0x72, 0x12, 0x30, 0x0a, 0x13, 0x6c, 0x69, 0x73,
	0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x13, 0x6c, 0x69, 0x73, 0x74, 0x61, 0x44, 0x61, 0x74,
	0x61, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x36, 0x0a, 0x06, 0x4c,
	0x69, 0x62, 0x72, 0x6f, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x6c, 0x69, 0x62, 0x72, 0x6f, 0x73, 0x44,
	0x69, 0x73, 0x70, 0x6f, 0x6e, 0x69, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x11, 0x6c, 0x69, 0x62, 0x72, 0x6f, 0x73, 0x44, 0x69, 0x73, 0x70, 0x6f, 0x6e, 0x69, 0x62,
	0x6c, 0x65, 0x73, 0x32, 0xaf, 0x03, 0x0a, 0x11, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x72,
	0x69, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x45, 0x6e, 0x76,
	0x69, 0x61, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x13, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x00, 0x12, 0x4f, 0x0a, 0x18, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x55, 0x62,
	0x69, 0x63, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x6f, 0x12, 0x17,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x6d, 0x62,
	0x72, 0x65, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x63, 0x69, 0x6f,
	0x6e, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0d, 0x44, 0x65, 0x73, 0x63, 0x61, 0x72, 0x67, 0x61, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x1a,
	0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0e, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x50, 0x72, 0x6f,
	0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x63, 0x69, 0x6f, 0x6e,
	0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x11, 0x45, 0x6e, 0x76, 0x69, 0x61,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x63, 0x69, 0x6f, 0x6e, 0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a,
	0x19, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x73, 0x44,
	0x69, 0x73, 0x70, 0x6f, 0x6e, 0x69, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x13, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x62,
	0x72, 0x6f, 0x73, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x61, 0x74, 0x68,
	0x2f, 0x67, 0x65, 0x6e, 0x3b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Connection_proto_rawDescOnce sync.Once
	file_Connection_proto_rawDescData = file_Connection_proto_rawDesc
)

func file_Connection_proto_rawDescGZIP() []byte {
	file_Connection_proto_rawDescOnce.Do(func() {
		file_Connection_proto_rawDescData = protoimpl.X.CompressGZIP(file_Connection_proto_rawDescData)
	})
	return file_Connection_proto_rawDescData
}

var file_Connection_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_Connection_proto_goTypes = []interface{}{
	(*Chunk)(nil),         // 0: connection.Chunk
	(*NombreLibro)(nil),   // 1: connection.NombreLibro
	(*Message)(nil),       // 2: connection.Message
	(*DivisionLibro)(nil), // 3: connection.DivisionLibro
	(*Distribucion)(nil),  // 4: connection.Distribucion
	(*Libros)(nil),        // 5: connection.Libros
}
var file_Connection_proto_depIdxs = []int32{
	0, // 0: connection.MensajeriaService.EnviaChunks:input_type -> connection.Chunk
	1, // 1: connection.MensajeriaService.ConsultaUbicacionArchivo:input_type -> connection.NombreLibro
	3, // 2: connection.MensajeriaService.DescargaChunk:input_type -> connection.DivisionLibro
	4, // 3: connection.MensajeriaService.EnviaPropuesta:input_type -> connection.Distribucion
	4, // 4: connection.MensajeriaService.EnviaDistribucion:input_type -> connection.Distribucion
	2, // 5: connection.MensajeriaService.ConsultaLibrosDisponibles:input_type -> connection.Message
	2, // 6: connection.MensajeriaService.EnviaChunks:output_type -> connection.Message
	4, // 7: connection.MensajeriaService.ConsultaUbicacionArchivo:output_type -> connection.Distribucion
	0, // 8: connection.MensajeriaService.DescargaChunk:output_type -> connection.Chunk
	2, // 9: connection.MensajeriaService.EnviaPropuesta:output_type -> connection.Message
	2, // 10: connection.MensajeriaService.EnviaDistribucion:output_type -> connection.Message
	5, // 11: connection.MensajeriaService.ConsultaLibrosDisponibles:output_type -> connection.Libros
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Connection_proto_init() }
func file_Connection_proto_init() {
	if File_Connection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Connection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_Connection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NombreLibro); i {
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
		file_Connection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_Connection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DivisionLibro); i {
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
		file_Connection_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Distribucion); i {
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
		file_Connection_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Libros); i {
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
			RawDescriptor: file_Connection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Connection_proto_goTypes,
		DependencyIndexes: file_Connection_proto_depIdxs,
		MessageInfos:      file_Connection_proto_msgTypes,
	}.Build()
	File_Connection_proto = out.File
	file_Connection_proto_rawDesc = nil
	file_Connection_proto_goTypes = nil
	file_Connection_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MensajeriaServiceClient is the client API for MensajeriaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MensajeriaServiceClient interface {
	EnviaChunks(ctx context.Context, in *Chunk, opts ...grpc.CallOption) (*Message, error)
	ConsultaUbicacionArchivo(ctx context.Context, in *NombreLibro, opts ...grpc.CallOption) (*Distribucion, error)
	DescargaChunk(ctx context.Context, in *DivisionLibro, opts ...grpc.CallOption) (*Chunk, error)
	EnviaPropuesta(ctx context.Context, in *Distribucion, opts ...grpc.CallOption) (*Message, error)
	EnviaDistribucion(ctx context.Context, in *Distribucion, opts ...grpc.CallOption) (*Message, error)
	ConsultaLibrosDisponibles(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Libros, error)
}

type mensajeriaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMensajeriaServiceClient(cc grpc.ClientConnInterface) MensajeriaServiceClient {
	return &mensajeriaServiceClient{cc}
}

func (c *mensajeriaServiceClient) EnviaChunks(ctx context.Context, in *Chunk, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/connection.MensajeriaService/EnviaChunks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mensajeriaServiceClient) ConsultaUbicacionArchivo(ctx context.Context, in *NombreLibro, opts ...grpc.CallOption) (*Distribucion, error) {
	out := new(Distribucion)
	err := c.cc.Invoke(ctx, "/connection.MensajeriaService/ConsultaUbicacionArchivo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mensajeriaServiceClient) DescargaChunk(ctx context.Context, in *DivisionLibro, opts ...grpc.CallOption) (*Chunk, error) {
	out := new(Chunk)
	err := c.cc.Invoke(ctx, "/connection.MensajeriaService/DescargaChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mensajeriaServiceClient) EnviaPropuesta(ctx context.Context, in *Distribucion, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/connection.MensajeriaService/EnviaPropuesta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mensajeriaServiceClient) EnviaDistribucion(ctx context.Context, in *Distribucion, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/connection.MensajeriaService/EnviaDistribucion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mensajeriaServiceClient) ConsultaLibrosDisponibles(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Libros, error) {
	out := new(Libros)
	err := c.cc.Invoke(ctx, "/connection.MensajeriaService/ConsultaLibrosDisponibles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MensajeriaServiceServer is the server API for MensajeriaService service.
type MensajeriaServiceServer interface {
	EnviaChunks(context.Context, *Chunk) (*Message, error)
	ConsultaUbicacionArchivo(context.Context, *NombreLibro) (*Distribucion, error)
	DescargaChunk(context.Context, *DivisionLibro) (*Chunk, error)
	EnviaPropuesta(context.Context, *Distribucion) (*Message, error)
	EnviaDistribucion(context.Context, *Distribucion) (*Message, error)
	ConsultaLibrosDisponibles(context.Context, *Message) (*Libros, error)
}

// UnimplementedMensajeriaServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMensajeriaServiceServer struct {
}

func (*UnimplementedMensajeriaServiceServer) EnviaChunks(context.Context, *Chunk) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviaChunks not implemented")
}
func (*UnimplementedMensajeriaServiceServer) ConsultaUbicacionArchivo(context.Context, *NombreLibro) (*Distribucion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsultaUbicacionArchivo not implemented")
}
func (*UnimplementedMensajeriaServiceServer) DescargaChunk(context.Context, *DivisionLibro) (*Chunk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescargaChunk not implemented")
}
func (*UnimplementedMensajeriaServiceServer) EnviaPropuesta(context.Context, *Distribucion) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviaPropuesta not implemented")
}
func (*UnimplementedMensajeriaServiceServer) EnviaDistribucion(context.Context, *Distribucion) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviaDistribucion not implemented")
}
func (*UnimplementedMensajeriaServiceServer) ConsultaLibrosDisponibles(context.Context, *Message) (*Libros, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsultaLibrosDisponibles not implemented")
}

func RegisterMensajeriaServiceServer(s *grpc.Server, srv MensajeriaServiceServer) {
	s.RegisterService(&_MensajeriaService_serviceDesc, srv)
}

func _MensajeriaService_EnviaChunks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chunk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MensajeriaServiceServer).EnviaChunks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.MensajeriaService/EnviaChunks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MensajeriaServiceServer).EnviaChunks(ctx, req.(*Chunk))
	}
	return interceptor(ctx, in, info, handler)
}

func _MensajeriaService_ConsultaUbicacionArchivo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NombreLibro)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MensajeriaServiceServer).ConsultaUbicacionArchivo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.MensajeriaService/ConsultaUbicacionArchivo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MensajeriaServiceServer).ConsultaUbicacionArchivo(ctx, req.(*NombreLibro))
	}
	return interceptor(ctx, in, info, handler)
}

func _MensajeriaService_DescargaChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DivisionLibro)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MensajeriaServiceServer).DescargaChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.MensajeriaService/DescargaChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MensajeriaServiceServer).DescargaChunk(ctx, req.(*DivisionLibro))
	}
	return interceptor(ctx, in, info, handler)
}

func _MensajeriaService_EnviaPropuesta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Distribucion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MensajeriaServiceServer).EnviaPropuesta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.MensajeriaService/EnviaPropuesta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MensajeriaServiceServer).EnviaPropuesta(ctx, req.(*Distribucion))
	}
	return interceptor(ctx, in, info, handler)
}

func _MensajeriaService_EnviaDistribucion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Distribucion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MensajeriaServiceServer).EnviaDistribucion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.MensajeriaService/EnviaDistribucion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MensajeriaServiceServer).EnviaDistribucion(ctx, req.(*Distribucion))
	}
	return interceptor(ctx, in, info, handler)
}

func _MensajeriaService_ConsultaLibrosDisponibles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MensajeriaServiceServer).ConsultaLibrosDisponibles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.MensajeriaService/ConsultaLibrosDisponibles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MensajeriaServiceServer).ConsultaLibrosDisponibles(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _MensajeriaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "connection.MensajeriaService",
	HandlerType: (*MensajeriaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnviaChunks",
			Handler:    _MensajeriaService_EnviaChunks_Handler,
		},
		{
			MethodName: "ConsultaUbicacionArchivo",
			Handler:    _MensajeriaService_ConsultaUbicacionArchivo_Handler,
		},
		{
			MethodName: "DescargaChunk",
			Handler:    _MensajeriaService_DescargaChunk_Handler,
		},
		{
			MethodName: "EnviaPropuesta",
			Handler:    _MensajeriaService_EnviaPropuesta_Handler,
		},
		{
			MethodName: "EnviaDistribucion",
			Handler:    _MensajeriaService_EnviaDistribucion_Handler,
		},
		{
			MethodName: "ConsultaLibrosDisponibles",
			Handler:    _MensajeriaService_ConsultaLibrosDisponibles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Connection.proto",
}

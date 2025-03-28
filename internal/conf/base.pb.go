// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.1
// source: conf/base.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type GRPC struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Network       string                 `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout       *durationpb.Duration   `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GRPC) Reset() {
	*x = GRPC{}
	mi := &file_conf_base_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GRPC) ProtoMessage() {}

func (x *GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_conf_base_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GRPC.ProtoReflect.Descriptor instead.
func (*GRPC) Descriptor() ([]byte, []int) {
	return file_conf_base_proto_rawDescGZIP(), []int{0}
}

func (x *GRPC) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *GRPC) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *GRPC) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Database struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Driver        string                 `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Host          string                 `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Port          int32                  `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	Dbname        string                 `protobuf:"bytes,4,opt,name=dbname,proto3" json:"dbname,omitempty"`
	User          string                 `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	Password      string                 `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	NoSsl         bool                   `protobuf:"varint,7,opt,name=no_ssl,json=noSsl,proto3" json:"no_ssl,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Database) Reset() {
	*x = Database{}
	mi := &file_conf_base_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Database) ProtoMessage() {}

func (x *Database) ProtoReflect() protoreflect.Message {
	mi := &file_conf_base_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Database.ProtoReflect.Descriptor instead.
func (*Database) Descriptor() ([]byte, []int) {
	return file_conf_base_proto_rawDescGZIP(), []int{1}
}

func (x *Database) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Database) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Database) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Database) GetDbname() string {
	if x != nil {
		return x.Dbname
	}
	return ""
}

func (x *Database) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Database) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Database) GetNoSsl() bool {
	if x != nil {
		return x.NoSsl
	}
	return false
}

type S3 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Driver        string                 `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	EndPoint      string                 `protobuf:"bytes,2,opt,name=end_point,json=endPoint,proto3" json:"end_point,omitempty"`
	AccessKey     string                 `protobuf:"bytes,3,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	SecretKey     string                 `protobuf:"bytes,4,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	UseSsl        bool                   `protobuf:"varint,5,opt,name=use_ssl,json=useSsl,proto3" json:"use_ssl,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *S3) Reset() {
	*x = S3{}
	mi := &file_conf_base_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *S3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3) ProtoMessage() {}

func (x *S3) ProtoReflect() protoreflect.Message {
	mi := &file_conf_base_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3.ProtoReflect.Descriptor instead.
func (*S3) Descriptor() ([]byte, []int) {
	return file_conf_base_proto_rawDescGZIP(), []int{2}
}

func (x *S3) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *S3) GetEndPoint() string {
	if x != nil {
		return x.EndPoint
	}
	return ""
}

func (x *S3) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *S3) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *S3) GetUseSsl() bool {
	if x != nil {
		return x.UseSsl
	}
	return false
}

type MQ struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Driver        string                 `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MQ) Reset() {
	*x = MQ{}
	mi := &file_conf_base_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MQ) ProtoMessage() {}

func (x *MQ) ProtoReflect() protoreflect.Message {
	mi := &file_conf_base_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MQ.ProtoReflect.Descriptor instead.
func (*MQ) Descriptor() ([]byte, []int) {
	return file_conf_base_proto_rawDescGZIP(), []int{3}
}

func (x *MQ) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

type Auth struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PasswordSalt  string                 `protobuf:"bytes,1,opt,name=password_salt,json=passwordSalt,proto3" json:"password_salt,omitempty"`
	JwtIssuer     string                 `protobuf:"bytes,2,opt,name=jwt_issuer,json=jwtIssuer,proto3" json:"jwt_issuer,omitempty"`
	JwtSecret     string                 `protobuf:"bytes,3,opt,name=jwt_secret,json=jwtSecret,proto3" json:"jwt_secret,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Auth) Reset() {
	*x = Auth{}
	mi := &file_conf_base_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_conf_base_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_conf_base_proto_rawDescGZIP(), []int{4}
}

func (x *Auth) GetPasswordSalt() string {
	if x != nil {
		return x.PasswordSalt
	}
	return ""
}

func (x *Auth) GetJwtIssuer() string {
	if x != nil {
		return x.JwtIssuer
	}
	return ""
}

func (x *Auth) GetJwtSecret() string {
	if x != nil {
		return x.JwtSecret
	}
	return ""
}

type Cache struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Driver        string                 `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Db            int64                  `protobuf:"varint,3,opt,name=db,proto3" json:"db,omitempty"`
	User          string                 `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	Password      string                 `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Cache) Reset() {
	*x = Cache{}
	mi := &file_conf_base_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cache) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cache) ProtoMessage() {}

func (x *Cache) ProtoReflect() protoreflect.Message {
	mi := &file_conf_base_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cache.ProtoReflect.Descriptor instead.
func (*Cache) Descriptor() ([]byte, []int) {
	return file_conf_base_proto_rawDescGZIP(), []int{5}
}

func (x *Cache) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Cache) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Cache) GetDb() int64 {
	if x != nil {
		return x.Db
	}
	return 0
}

func (x *Cache) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Cache) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Consul struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Addr          string                 `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Token         string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Consul) Reset() {
	*x = Consul{}
	mi := &file_conf_base_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Consul) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Consul) ProtoMessage() {}

func (x *Consul) ProtoReflect() protoreflect.Message {
	mi := &file_conf_base_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Consul.ProtoReflect.Descriptor instead.
func (*Consul) Descriptor() ([]byte, []int) {
	return file_conf_base_proto_rawDescGZIP(), []int{6}
}

func (x *Consul) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Consul) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Sentry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Dsn           string                 `protobuf:"bytes,1,opt,name=dsn,proto3" json:"dsn,omitempty"`
	Environment   string                 `protobuf:"bytes,2,opt,name=environment,proto3" json:"environment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Sentry) Reset() {
	*x = Sentry{}
	mi := &file_conf_base_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Sentry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sentry) ProtoMessage() {}

func (x *Sentry) ProtoReflect() protoreflect.Message {
	mi := &file_conf_base_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sentry.ProtoReflect.Descriptor instead.
func (*Sentry) Descriptor() ([]byte, []int) {
	return file_conf_base_proto_rawDescGZIP(), []int{7}
}

func (x *Sentry) GetDsn() string {
	if x != nil {
		return x.Dsn
	}
	return ""
}

func (x *Sentry) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

type OTLP struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Protocol      string                 `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Endpoint      string                 `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Headers       string                 `protobuf:"bytes,3,opt,name=headers,proto3" json:"headers,omitempty"`
	GrpcInsecure  bool                   `protobuf:"varint,4,opt,name=grpc_insecure,json=grpcInsecure,proto3" json:"grpc_insecure,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OTLP) Reset() {
	*x = OTLP{}
	mi := &file_conf_base_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OTLP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OTLP) ProtoMessage() {}

func (x *OTLP) ProtoReflect() protoreflect.Message {
	mi := &file_conf_base_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OTLP.ProtoReflect.Descriptor instead.
func (*OTLP) Descriptor() ([]byte, []int) {
	return file_conf_base_proto_rawDescGZIP(), []int{8}
}

func (x *OTLP) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *OTLP) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *OTLP) GetHeaders() string {
	if x != nil {
		return x.Headers
	}
	return ""
}

func (x *OTLP) GetGrpcInsecure() bool {
	if x != nil {
		return x.GrpcInsecure
	}
	return false
}

type Search struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Driver        string                 `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Meili         *Search_MeiliSearch    `protobuf:"bytes,2,opt,name=meili,proto3,oneof" json:"meili,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Search) Reset() {
	*x = Search{}
	mi := &file_conf_base_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Search) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Search) ProtoMessage() {}

func (x *Search) ProtoReflect() protoreflect.Message {
	mi := &file_conf_base_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Search.ProtoReflect.Descriptor instead.
func (*Search) Descriptor() ([]byte, []int) {
	return file_conf_base_proto_rawDescGZIP(), []int{9}
}

func (x *Search) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Search) GetMeili() *Search_MeiliSearch {
	if x != nil {
		return x.Meili
	}
	return nil
}

type Search_MeiliSearch struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Addr          string                 `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	ApiKey        string                 `protobuf:"bytes,2,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Search_MeiliSearch) Reset() {
	*x = Search_MeiliSearch{}
	mi := &file_conf_base_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Search_MeiliSearch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Search_MeiliSearch) ProtoMessage() {}

func (x *Search_MeiliSearch) ProtoReflect() protoreflect.Message {
	mi := &file_conf_base_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Search_MeiliSearch.ProtoReflect.Descriptor instead.
func (*Search_MeiliSearch) Descriptor() ([]byte, []int) {
	return file_conf_base_proto_rawDescGZIP(), []int{9, 0}
}

func (x *Search_MeiliSearch) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Search_MeiliSearch) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

var File_conf_base_proto protoreflect.FileDescriptor

var file_conf_base_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a,
	0x04, 0x47, 0x52, 0x50, 0x43, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0xa9, 0x01, 0x0a, 0x08, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x62, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x15, 0x0a,
	0x06, 0x6e, 0x6f, 0x5f, 0x73, 0x73, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6e,
	0x6f, 0x53, 0x73, 0x6c, 0x22, 0x90, 0x01, 0x0a, 0x02, 0x53, 0x33, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x73, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x53, 0x73, 0x6c, 0x22, 0x1c, 0x0a, 0x02, 0x4d, 0x51, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x22, 0x69, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x23, 0x0a,
	0x0d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x73, 0x61, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x53, 0x61,
	0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6a, 0x77, 0x74, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x77, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x6a, 0x77, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x77, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x22, 0x73, 0x0a, 0x05, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x64, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x32, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3c, 0x0a, 0x06, 0x53, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x73, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x64, 0x73, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x7d, 0x0a, 0x04, 0x4f, 0x54, 0x4c, 0x50, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x49, 0x6e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x05, 0x6d, 0x65, 0x69,
	0x6c, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x4d, 0x65, 0x69,
	0x6c, 0x69, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x48, 0x00, 0x52, 0x05, 0x6d, 0x65, 0x69, 0x6c,
	0x69, 0x88, 0x01, 0x01, 0x1a, 0x3a, 0x0a, 0x0b, 0x4d, 0x65, 0x69, 0x6c, 0x69, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6d, 0x65, 0x69, 0x6c, 0x69, 0x42, 0x1e, 0x5a, 0x1c, 0x4c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_conf_base_proto_rawDescOnce sync.Once
	file_conf_base_proto_rawDescData []byte
)

func file_conf_base_proto_rawDescGZIP() []byte {
	file_conf_base_proto_rawDescOnce.Do(func() {
		file_conf_base_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_conf_base_proto_rawDesc), len(file_conf_base_proto_rawDesc)))
	})
	return file_conf_base_proto_rawDescData
}

var file_conf_base_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_conf_base_proto_goTypes = []any{
	(*GRPC)(nil),                // 0: kratos.api.GRPC
	(*Database)(nil),            // 1: kratos.api.Database
	(*S3)(nil),                  // 2: kratos.api.S3
	(*MQ)(nil),                  // 3: kratos.api.MQ
	(*Auth)(nil),                // 4: kratos.api.Auth
	(*Cache)(nil),               // 5: kratos.api.Cache
	(*Consul)(nil),              // 6: kratos.api.Consul
	(*Sentry)(nil),              // 7: kratos.api.Sentry
	(*OTLP)(nil),                // 8: kratos.api.OTLP
	(*Search)(nil),              // 9: kratos.api.Search
	(*Search_MeiliSearch)(nil),  // 10: kratos.api.Search.MeiliSearch
	(*durationpb.Duration)(nil), // 11: google.protobuf.Duration
}
var file_conf_base_proto_depIdxs = []int32{
	11, // 0: kratos.api.GRPC.timeout:type_name -> google.protobuf.Duration
	10, // 1: kratos.api.Search.meili:type_name -> kratos.api.Search.MeiliSearch
	2,  // [2:2] is the sub-list for method output_type
	2,  // [2:2] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_conf_base_proto_init() }
func file_conf_base_proto_init() {
	if File_conf_base_proto != nil {
		return
	}
	file_conf_base_proto_msgTypes[9].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_conf_base_proto_rawDesc), len(file_conf_base_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_base_proto_goTypes,
		DependencyIndexes: file_conf_base_proto_depIdxs,
		MessageInfos:      file_conf_base_proto_msgTypes,
	}.Build()
	File_conf_base_proto = out.File
	file_conf_base_proto_goTypes = nil
	file_conf_base_proto_depIdxs = nil
}

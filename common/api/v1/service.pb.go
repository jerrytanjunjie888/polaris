// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AliasType int32

const (
	AliasType_DEFAULT AliasType = 0
	AliasType_CL5SID  AliasType = 1
)

var AliasType_name = map[int32]string{
	0: "DEFAULT",
	1: "CL5SID",
}
var AliasType_value = map[string]int32{
	"DEFAULT": 0,
	"CL5SID":  1,
}

func (x AliasType) String() string {
	return proto.EnumName(AliasType_name, int32(x))
}
func (AliasType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_service_413f75d8eac84e7c, []int{0}
}

type HealthCheck_HealthCheckType int32

const (
	HealthCheck_UNKNOWN   HealthCheck_HealthCheckType = 0
	HealthCheck_HEARTBEAT HealthCheck_HealthCheckType = 1
)

var HealthCheck_HealthCheckType_name = map[int32]string{
	0: "UNKNOWN",
	1: "HEARTBEAT",
}
var HealthCheck_HealthCheckType_value = map[string]int32{
	"UNKNOWN":   0,
	"HEARTBEAT": 1,
}

func (x HealthCheck_HealthCheckType) String() string {
	return proto.EnumName(HealthCheck_HealthCheckType_name, int32(x))
}
func (HealthCheck_HealthCheckType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_service_413f75d8eac84e7c, []int{4, 0}
}

type Namespace struct {
	Name                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Comment              *wrappers.StringValue `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Owners               *wrappers.StringValue `protobuf:"bytes,3,opt,name=owners,proto3" json:"owners,omitempty"`
	Token                *wrappers.StringValue `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	Ctime                *wrappers.StringValue `protobuf:"bytes,5,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Mtime                *wrappers.StringValue `protobuf:"bytes,6,opt,name=mtime,proto3" json:"mtime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Namespace) Reset()         { *m = Namespace{} }
func (m *Namespace) String() string { return proto.CompactTextString(m) }
func (*Namespace) ProtoMessage()    {}
func (*Namespace) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_413f75d8eac84e7c, []int{0}
}
func (m *Namespace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Namespace.Unmarshal(m, b)
}
func (m *Namespace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Namespace.Marshal(b, m, deterministic)
}
func (dst *Namespace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Namespace.Merge(dst, src)
}
func (m *Namespace) XXX_Size() int {
	return xxx_messageInfo_Namespace.Size(m)
}
func (m *Namespace) XXX_DiscardUnknown() {
	xxx_messageInfo_Namespace.DiscardUnknown(m)
}

var xxx_messageInfo_Namespace proto.InternalMessageInfo

func (m *Namespace) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Namespace) GetComment() *wrappers.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *Namespace) GetOwners() *wrappers.StringValue {
	if m != nil {
		return m.Owners
	}
	return nil
}

func (m *Namespace) GetToken() *wrappers.StringValue {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *Namespace) GetCtime() *wrappers.StringValue {
	if m != nil {
		return m.Ctime
	}
	return nil
}

func (m *Namespace) GetMtime() *wrappers.StringValue {
	if m != nil {
		return m.Mtime
	}
	return nil
}

type Service struct {
	Name                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            *wrappers.StringValue `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Metadata             map[string]string     `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Ports                *wrappers.StringValue `protobuf:"bytes,4,opt,name=ports,proto3" json:"ports,omitempty"`
	Business             *wrappers.StringValue `protobuf:"bytes,5,opt,name=business,proto3" json:"business,omitempty"`
	Department           *wrappers.StringValue `protobuf:"bytes,6,opt,name=department,proto3" json:"department,omitempty"`
	CmdbMod1             *wrappers.StringValue `protobuf:"bytes,7,opt,name=cmdb_mod1,proto3" json:"cmdb_mod1,omitempty"`
	CmdbMod2             *wrappers.StringValue `protobuf:"bytes,8,opt,name=cmdb_mod2,proto3" json:"cmdb_mod2,omitempty"`
	CmdbMod3             *wrappers.StringValue `protobuf:"bytes,9,opt,name=cmdb_mod3,proto3" json:"cmdb_mod3,omitempty"`
	Comment              *wrappers.StringValue `protobuf:"bytes,10,opt,name=comment,proto3" json:"comment,omitempty"`
	Owners               *wrappers.StringValue `protobuf:"bytes,11,opt,name=owners,proto3" json:"owners,omitempty"`
	Token                *wrappers.StringValue `protobuf:"bytes,12,opt,name=token,proto3" json:"token,omitempty"`
	Ctime                *wrappers.StringValue `protobuf:"bytes,13,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Mtime                *wrappers.StringValue `protobuf:"bytes,14,opt,name=mtime,proto3" json:"mtime,omitempty"`
	Revision             *wrappers.StringValue `protobuf:"bytes,15,opt,name=revision,proto3" json:"revision,omitempty"`
	PlatformId           *wrappers.StringValue `protobuf:"bytes,16,opt,name=platform_id,proto3" json:"platform_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_413f75d8eac84e7c, []int{1}
}
func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (dst *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(dst, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Service) GetNamespace() *wrappers.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *Service) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Service) GetPorts() *wrappers.StringValue {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *Service) GetBusiness() *wrappers.StringValue {
	if m != nil {
		return m.Business
	}
	return nil
}

func (m *Service) GetDepartment() *wrappers.StringValue {
	if m != nil {
		return m.Department
	}
	return nil
}

func (m *Service) GetCmdbMod1() *wrappers.StringValue {
	if m != nil {
		return m.CmdbMod1
	}
	return nil
}

func (m *Service) GetCmdbMod2() *wrappers.StringValue {
	if m != nil {
		return m.CmdbMod2
	}
	return nil
}

func (m *Service) GetCmdbMod3() *wrappers.StringValue {
	if m != nil {
		return m.CmdbMod3
	}
	return nil
}

func (m *Service) GetComment() *wrappers.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *Service) GetOwners() *wrappers.StringValue {
	if m != nil {
		return m.Owners
	}
	return nil
}

func (m *Service) GetToken() *wrappers.StringValue {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *Service) GetCtime() *wrappers.StringValue {
	if m != nil {
		return m.Ctime
	}
	return nil
}

func (m *Service) GetMtime() *wrappers.StringValue {
	if m != nil {
		return m.Mtime
	}
	return nil
}

func (m *Service) GetRevision() *wrappers.StringValue {
	if m != nil {
		return m.Revision
	}
	return nil
}

func (m *Service) GetPlatformId() *wrappers.StringValue {
	if m != nil {
		return m.PlatformId
	}
	return nil
}

type ServiceAlias struct {
	Service              *wrappers.StringValue `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Namespace            *wrappers.StringValue `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Alias                *wrappers.StringValue `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias,omitempty"`
	Type                 AliasType             `protobuf:"varint,4,opt,name=type,proto3,enum=v1.AliasType" json:"type,omitempty"`
	Owners               *wrappers.StringValue `protobuf:"bytes,8,opt,name=owners,proto3" json:"owners,omitempty"`
	Comment              *wrappers.StringValue `protobuf:"bytes,9,opt,name=comment,proto3" json:"comment,omitempty"`
	ServiceToken         *wrappers.StringValue `protobuf:"bytes,5,opt,name=service_token,proto3" json:"service_token,omitempty"`
	Ctime                *wrappers.StringValue `protobuf:"bytes,6,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Mtime                *wrappers.StringValue `protobuf:"bytes,7,opt,name=mtime,proto3" json:"mtime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ServiceAlias) Reset()         { *m = ServiceAlias{} }
func (m *ServiceAlias) String() string { return proto.CompactTextString(m) }
func (*ServiceAlias) ProtoMessage()    {}
func (*ServiceAlias) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_413f75d8eac84e7c, []int{2}
}
func (m *ServiceAlias) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceAlias.Unmarshal(m, b)
}
func (m *ServiceAlias) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceAlias.Marshal(b, m, deterministic)
}
func (dst *ServiceAlias) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceAlias.Merge(dst, src)
}
func (m *ServiceAlias) XXX_Size() int {
	return xxx_messageInfo_ServiceAlias.Size(m)
}
func (m *ServiceAlias) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceAlias.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceAlias proto.InternalMessageInfo

func (m *ServiceAlias) GetService() *wrappers.StringValue {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ServiceAlias) GetNamespace() *wrappers.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ServiceAlias) GetAlias() *wrappers.StringValue {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *ServiceAlias) GetType() AliasType {
	if m != nil {
		return m.Type
	}
	return AliasType_DEFAULT
}

func (m *ServiceAlias) GetOwners() *wrappers.StringValue {
	if m != nil {
		return m.Owners
	}
	return nil
}

func (m *ServiceAlias) GetComment() *wrappers.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *ServiceAlias) GetServiceToken() *wrappers.StringValue {
	if m != nil {
		return m.ServiceToken
	}
	return nil
}

func (m *ServiceAlias) GetCtime() *wrappers.StringValue {
	if m != nil {
		return m.Ctime
	}
	return nil
}

func (m *ServiceAlias) GetMtime() *wrappers.StringValue {
	if m != nil {
		return m.Mtime
	}
	return nil
}

type Instance struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Service              *wrappers.StringValue `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Namespace            *wrappers.StringValue `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	VpcId                *wrappers.StringValue `protobuf:"bytes,21,opt,name=vpc_id,proto3" json:"vpc_id,omitempty"`
	Host                 *wrappers.StringValue `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	Port                 *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=port,proto3" json:"port,omitempty"`
	Protocol             *wrappers.StringValue `protobuf:"bytes,6,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Version              *wrappers.StringValue `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	Priority             *wrappers.UInt32Value `protobuf:"bytes,8,opt,name=priority,proto3" json:"priority,omitempty"`
	Weight               *wrappers.UInt32Value `protobuf:"bytes,9,opt,name=weight,proto3" json:"weight,omitempty"`
	EnableHealthCheck    *wrappers.BoolValue   `protobuf:"bytes,20,opt,name=enable_health_check,json=enableHealthCheck,proto3" json:"enable_health_check,omitempty"`
	HealthCheck          *HealthCheck          `protobuf:"bytes,10,opt,name=health_check,json=healthCheck,proto3" json:"health_check,omitempty"`
	Healthy              *wrappers.BoolValue   `protobuf:"bytes,11,opt,name=healthy,proto3" json:"healthy,omitempty"`
	Isolate              *wrappers.BoolValue   `protobuf:"bytes,12,opt,name=isolate,proto3" json:"isolate,omitempty"`
	Location             *Location             `protobuf:"bytes,13,opt,name=location,proto3" json:"location,omitempty"`
	Metadata             map[string]string     `protobuf:"bytes,14,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	LogicSet             *wrappers.StringValue `protobuf:"bytes,15,opt,name=logic_set,proto3" json:"logic_set,omitempty"`
	Ctime                *wrappers.StringValue `protobuf:"bytes,16,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Mtime                *wrappers.StringValue `protobuf:"bytes,17,opt,name=mtime,proto3" json:"mtime,omitempty"`
	Revision             *wrappers.StringValue `protobuf:"bytes,18,opt,name=revision,proto3" json:"revision,omitempty"`
	ServiceToken         *wrappers.StringValue `protobuf:"bytes,19,opt,name=service_token,proto3" json:"service_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Instance) Reset()         { *m = Instance{} }
func (m *Instance) String() string { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()    {}
func (*Instance) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_413f75d8eac84e7c, []int{3}
}
func (m *Instance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instance.Unmarshal(m, b)
}
func (m *Instance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instance.Marshal(b, m, deterministic)
}
func (dst *Instance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instance.Merge(dst, src)
}
func (m *Instance) XXX_Size() int {
	return xxx_messageInfo_Instance.Size(m)
}
func (m *Instance) XXX_DiscardUnknown() {
	xxx_messageInfo_Instance.DiscardUnknown(m)
}

var xxx_messageInfo_Instance proto.InternalMessageInfo

func (m *Instance) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Instance) GetService() *wrappers.StringValue {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *Instance) GetNamespace() *wrappers.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *Instance) GetVpcId() *wrappers.StringValue {
	if m != nil {
		return m.VpcId
	}
	return nil
}

func (m *Instance) GetHost() *wrappers.StringValue {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *Instance) GetPort() *wrappers.UInt32Value {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *Instance) GetProtocol() *wrappers.StringValue {
	if m != nil {
		return m.Protocol
	}
	return nil
}

func (m *Instance) GetVersion() *wrappers.StringValue {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *Instance) GetPriority() *wrappers.UInt32Value {
	if m != nil {
		return m.Priority
	}
	return nil
}

func (m *Instance) GetWeight() *wrappers.UInt32Value {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *Instance) GetEnableHealthCheck() *wrappers.BoolValue {
	if m != nil {
		return m.EnableHealthCheck
	}
	return nil
}

func (m *Instance) GetHealthCheck() *HealthCheck {
	if m != nil {
		return m.HealthCheck
	}
	return nil
}

func (m *Instance) GetHealthy() *wrappers.BoolValue {
	if m != nil {
		return m.Healthy
	}
	return nil
}

func (m *Instance) GetIsolate() *wrappers.BoolValue {
	if m != nil {
		return m.Isolate
	}
	return nil
}

func (m *Instance) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Instance) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Instance) GetLogicSet() *wrappers.StringValue {
	if m != nil {
		return m.LogicSet
	}
	return nil
}

func (m *Instance) GetCtime() *wrappers.StringValue {
	if m != nil {
		return m.Ctime
	}
	return nil
}

func (m *Instance) GetMtime() *wrappers.StringValue {
	if m != nil {
		return m.Mtime
	}
	return nil
}

func (m *Instance) GetRevision() *wrappers.StringValue {
	if m != nil {
		return m.Revision
	}
	return nil
}

func (m *Instance) GetServiceToken() *wrappers.StringValue {
	if m != nil {
		return m.ServiceToken
	}
	return nil
}

type HealthCheck struct {
	Type                 HealthCheck_HealthCheckType `protobuf:"varint,1,opt,name=type,proto3,enum=v1.HealthCheck_HealthCheckType" json:"type,omitempty"`
	Heartbeat            *HeartbeatHealthCheck       `protobuf:"bytes,2,opt,name=heartbeat,proto3" json:"heartbeat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_413f75d8eac84e7c, []int{4}
}
func (m *HealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck.Marshal(b, m, deterministic)
}
func (dst *HealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck.Merge(dst, src)
}
func (m *HealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck.Size(m)
}
func (m *HealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck proto.InternalMessageInfo

func (m *HealthCheck) GetType() HealthCheck_HealthCheckType {
	if m != nil {
		return m.Type
	}
	return HealthCheck_UNKNOWN
}

func (m *HealthCheck) GetHeartbeat() *HeartbeatHealthCheck {
	if m != nil {
		return m.Heartbeat
	}
	return nil
}

type HeartbeatHealthCheck struct {
	Ttl                  *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=ttl,proto3" json:"ttl,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *HeartbeatHealthCheck) Reset()         { *m = HeartbeatHealthCheck{} }
func (m *HeartbeatHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HeartbeatHealthCheck) ProtoMessage()    {}
func (*HeartbeatHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_413f75d8eac84e7c, []int{5}
}
func (m *HeartbeatHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatHealthCheck.Unmarshal(m, b)
}
func (m *HeartbeatHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatHealthCheck.Marshal(b, m, deterministic)
}
func (dst *HeartbeatHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatHealthCheck.Merge(dst, src)
}
func (m *HeartbeatHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HeartbeatHealthCheck.Size(m)
}
func (m *HeartbeatHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatHealthCheck proto.InternalMessageInfo

func (m *HeartbeatHealthCheck) GetTtl() *wrappers.UInt32Value {
	if m != nil {
		return m.Ttl
	}
	return nil
}

func init() {
	proto.RegisterType((*Namespace)(nil), "v1.Namespace")
	proto.RegisterType((*Service)(nil), "v1.Service")
	proto.RegisterMapType((map[string]string)(nil), "v1.Service.MetadataEntry")
	proto.RegisterType((*ServiceAlias)(nil), "v1.ServiceAlias")
	proto.RegisterType((*Instance)(nil), "v1.Instance")
	proto.RegisterMapType((map[string]string)(nil), "v1.Instance.MetadataEntry")
	proto.RegisterType((*HealthCheck)(nil), "v1.HealthCheck")
	proto.RegisterType((*HeartbeatHealthCheck)(nil), "v1.HeartbeatHealthCheck")
	proto.RegisterEnum("v1.AliasType", AliasType_name, AliasType_value)
	proto.RegisterEnum("v1.HealthCheck_HealthCheckType", HealthCheck_HealthCheckType_name, HealthCheck_HealthCheckType_value)
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_413f75d8eac84e7c) }

var fileDescriptor_service_413f75d8eac84e7c = []byte{
	// 892 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xd1, 0x6e, 0xdb, 0x36,
	0x14, 0xad, 0xed, 0xc4, 0xb6, 0xae, 0xec, 0xc4, 0x65, 0x33, 0x40, 0x33, 0x86, 0xad, 0x33, 0xf6,
	0x10, 0x0c, 0x9b, 0xba, 0xc8, 0x69, 0x50, 0x64, 0xc3, 0x00, 0xa7, 0x75, 0xd1, 0x6c, 0x99, 0x07,
	0x28, 0xce, 0xf6, 0x68, 0xd0, 0x12, 0x6b, 0x0b, 0x91, 0x44, 0x41, 0x62, 0x1c, 0xf8, 0x93, 0xf6,
	0x25, 0x7b, 0xdc, 0xd3, 0xbe, 0x64, 0x3f, 0x30, 0x90, 0x94, 0x64, 0xa6, 0xad, 0x0b, 0x5a, 0x7d,
	0xb3, 0xa4, 0x73, 0xae, 0xe4, 0x7b, 0xcf, 0xb9, 0x87, 0xd0, 0xcd, 0x48, 0xba, 0x0a, 0x3c, 0x62,
	0x27, 0x29, 0x65, 0x14, 0xd5, 0x57, 0x27, 0xfd, 0x2f, 0x17, 0x94, 0x2e, 0x42, 0xf2, 0x4c, 0xdc,
	0x99, 0xdf, 0xbd, 0x7d, 0x76, 0x9f, 0xe2, 0x24, 0x21, 0x69, 0x26, 0x31, 0x7d, 0x33, 0xa2, 0x3e,
	0x09, 0xe5, 0xc5, 0xe0, 0xef, 0x3a, 0x18, 0x13, 0x1c, 0x91, 0x2c, 0xc1, 0x1e, 0x41, 0x3f, 0xc0,
	0x5e, 0x8c, 0x23, 0x62, 0xd5, 0x9e, 0xd6, 0x8e, 0x4d, 0xe7, 0x0b, 0x5b, 0x56, 0xb2, 0x8b, 0x4a,
	0xf6, 0x35, 0x4b, 0x83, 0x78, 0xf1, 0x07, 0x0e, 0xef, 0x88, 0x2b, 0x90, 0xe8, 0x0c, 0x5a, 0x1e,
	0x8d, 0x22, 0x12, 0x33, 0xab, 0xae, 0x41, 0x2a, 0xc0, 0xe8, 0x14, 0x9a, 0xf4, 0x3e, 0x26, 0x69,
	0x66, 0x35, 0x34, 0x68, 0x39, 0x16, 0x39, 0xb0, 0xcf, 0xe8, 0x2d, 0x89, 0xad, 0x3d, 0x0d, 0x92,
	0x84, 0x72, 0x8e, 0xc7, 0x82, 0x88, 0x58, 0xfb, 0x3a, 0x1c, 0x01, 0xe5, 0x9c, 0x48, 0x70, 0x9a,
	0x3a, 0x1c, 0x01, 0x1d, 0xfc, 0xd3, 0x82, 0xd6, 0xb5, 0x1c, 0x46, 0x85, 0x3e, 0x9e, 0x83, 0x11,
	0x17, 0x63, 0xd0, 0xea, 0xe4, 0x06, 0x8e, 0x9e, 0x43, 0x3b, 0x22, 0x0c, 0xfb, 0x98, 0x61, 0xab,
	0xf1, 0xb4, 0x71, 0x6c, 0x3a, 0x9f, 0xdb, 0xab, 0x13, 0x3b, 0xff, 0x18, 0xfb, 0xb7, 0xfc, 0xd9,
	0x38, 0x66, 0xe9, 0xda, 0x2d, 0xa1, 0xfc, 0x4f, 0x26, 0x34, 0x65, 0x99, 0x5e, 0x33, 0x05, 0x14,
	0xbd, 0x80, 0xf6, 0xfc, 0x2e, 0x0b, 0x62, 0x92, 0x65, 0x5a, 0xfd, 0x2c, 0xd1, 0xe8, 0x27, 0x00,
	0x9f, 0x24, 0x38, 0x65, 0x42, 0x2b, 0x3a, 0x7d, 0x55, 0xf0, 0xbc, 0x3d, 0x5e, 0xe4, 0xcf, 0x67,
	0x11, 0xf5, 0x4f, 0xac, 0x96, 0x4e, 0x7b, 0x4a, 0xb8, 0xca, 0x75, 0xac, 0xf6, 0x2e, 0x5c, 0x47,
	0xe5, 0x0e, 0x2d, 0x63, 0x17, 0xee, 0x50, 0xb5, 0x06, 0x54, 0xb3, 0x86, 0x59, 0xc5, 0x1a, 0x9d,
	0x0a, 0xd6, 0xe8, 0x56, 0xb0, 0xc6, 0x81, 0xb6, 0x35, 0xb8, 0x6a, 0x52, 0xb2, 0x0a, 0xb2, 0x80,
	0xc6, 0xd6, 0xa1, 0x8e, 0x6a, 0x0a, 0x34, 0xfa, 0x19, 0xcc, 0x24, 0xc4, 0xec, 0x2d, 0x4d, 0xa3,
	0x59, 0xe0, 0x5b, 0x3d, 0x0d, 0xb2, 0x4a, 0xe8, 0xff, 0x08, 0xdd, 0x07, 0xf2, 0x47, 0x3d, 0x68,
	0xdc, 0x92, 0xb5, 0x30, 0xa6, 0xe1, 0xf2, 0x9f, 0xe8, 0x08, 0xf6, 0x57, 0x9c, 0x28, 0x5c, 0x67,
	0xb8, 0xf2, 0xe2, 0xbc, 0xfe, 0xa2, 0x36, 0xf8, 0xaf, 0x01, 0x9d, 0xdc, 0x44, 0xa3, 0x30, 0xc0,
	0x19, 0x9f, 0x68, 0xbe, 0x6e, 0xb5, 0x9c, 0x5d, 0x80, 0x3f, 0xc9, 0xdc, 0x0e, 0xec, 0x63, 0xfe,
	0x72, 0xad, 0x3d, 0x29, 0xa1, 0xe8, 0x6b, 0xd8, 0x63, 0xeb, 0x84, 0x08, 0x63, 0x1f, 0x38, 0x5d,
	0xbe, 0x0c, 0xc4, 0x1f, 0x98, 0xae, 0x13, 0xe2, 0x8a, 0x47, 0x8a, 0xc8, 0xda, 0x3b, 0x88, 0x4c,
	0x91, 0xb4, 0xb1, 0x8b, 0xa4, 0x2f, 0xca, 0x9c, 0x9a, 0x49, 0x91, 0xea, 0xec, 0x8e, 0x87, 0x94,
	0x8d, 0x58, 0x9b, 0x15, 0xc4, 0xda, 0xd2, 0xdf, 0xe3, 0xff, 0x1a, 0xd0, 0xbe, 0x8c, 0x33, 0x86,
	0x63, 0x8f, 0xa0, 0xef, 0xa0, 0x1e, 0xf8, 0x5a, 0xc3, 0xae, 0x07, 0xbe, 0xaa, 0x8f, 0x7a, 0x65,
	0x7d, 0x34, 0x76, 0xd3, 0xc7, 0x29, 0x34, 0x57, 0x89, 0xc7, 0xcd, 0xf1, 0x99, 0xce, 0x20, 0x25,
	0x96, 0x07, 0xd4, 0x92, 0x66, 0x4c, 0x6b, 0xf5, 0x0b, 0x24, 0x67, 0xf0, 0x08, 0xd8, 0x3a, 0xb9,
	0x9b, 0xcb, 0x98, 0x0d, 0x9d, 0x9c, 0xc1, 0x91, 0xdc, 0xf5, 0xe2, 0xa9, 0x47, 0x43, 0xad, 0x99,
	0x95, 0x68, 0xde, 0xc7, 0x15, 0x49, 0xc5, 0xba, 0xd0, 0x19, 0x5c, 0x01, 0x96, 0x6f, 0x0c, 0x68,
	0x1a, 0xb0, 0xf5, 0x56, 0x59, 0xab, 0xdf, 0x59, 0xa2, 0x79, 0x17, 0xef, 0x49, 0xb0, 0x58, 0x6e,
	0xd7, 0xb5, 0xca, 0xcb, 0xb1, 0xe8, 0x17, 0x78, 0x42, 0x62, 0x3c, 0x0f, 0xc9, 0x6c, 0x49, 0x70,
	0xc8, 0x96, 0x33, 0x6f, 0x49, 0xbc, 0x5b, 0xeb, 0x48, 0x94, 0xe8, 0xbf, 0x57, 0xe2, 0x82, 0xd2,
	0x50, 0x16, 0x78, 0x2c, 0x69, 0x6f, 0x04, 0xeb, 0x25, 0x27, 0x21, 0x07, 0x3a, 0x0f, 0x8a, 0xc8,
	0xc8, 0x38, 0xe4, 0xde, 0x55, 0x60, 0xae, 0xb9, 0x54, 0x38, 0xa7, 0xd0, 0x92, 0x97, 0xeb, 0x3c,
	0x2a, 0x3e, 0xf6, 0xce, 0x02, 0xca, 0x59, 0x41, 0x46, 0x43, 0xcc, 0x48, 0x9e, 0x15, 0x1f, 0x65,
	0xe5, 0x50, 0x74, 0x0c, 0xed, 0x90, 0x7a, 0x98, 0xf1, 0xa1, 0xc8, 0xb8, 0xe8, 0xf0, 0x6f, 0xbb,
	0xca, 0xef, 0xb9, 0xe5, 0x53, 0x74, 0xa6, 0x1c, 0x47, 0x0e, 0xc4, 0x71, 0xa4, 0xcf, 0x91, 0x85,
	0xa7, 0xb6, 0x9e, 0x47, 0xce, 0xc1, 0x08, 0xe9, 0x22, 0xf0, 0x66, 0x19, 0x61, 0x5a, 0x31, 0xb1,
	0x81, 0x6f, 0x96, 0x43, 0xaf, 0xc2, 0x72, 0x78, 0x5c, 0x2d, 0xc9, 0xd0, 0x4e, 0x49, 0xf6, 0xde,
	0x0a, 0x7c, 0xb2, 0xf3, 0x0a, 0xfc, 0xb4, 0x34, 0xfb, 0xab, 0x06, 0xa6, 0x2a, 0xb8, 0x61, 0x1e,
	0x12, 0x35, 0x11, 0x12, 0x5f, 0xbd, 0x23, 0x34, 0xf5, 0xb7, 0x12, 0x1b, 0x67, 0x60, 0x2c, 0x09,
	0x4e, 0xd9, 0x9c, 0xe0, 0xe2, 0xc0, 0x6f, 0xe5, 0x4c, 0x79, 0x53, 0xd5, 0xea, 0x06, 0x3a, 0xf8,
	0x1e, 0x0e, 0xdf, 0x29, 0x88, 0x4c, 0x68, 0xdd, 0x4c, 0x7e, 0x9d, 0xfc, 0xfe, 0xe7, 0xa4, 0xf7,
	0x08, 0x75, 0xc1, 0x78, 0x33, 0x1e, 0xb9, 0xd3, 0x8b, 0xf1, 0x68, 0xda, 0xab, 0x0d, 0x5e, 0xc3,
	0xd1, 0x87, 0x2a, 0x22, 0x1b, 0x1a, 0x8c, 0x85, 0x5b, 0xf7, 0xb1, 0xea, 0x51, 0x0e, 0xfc, 0xf6,
	0x1b, 0x30, 0xca, 0xe0, 0xe3, 0x2f, 0x7c, 0x35, 0x7e, 0x3d, 0xba, 0xb9, 0x9a, 0xf6, 0x1e, 0x21,
	0x80, 0xe6, 0xcb, 0xab, 0xe7, 0xd7, 0x97, 0xaf, 0x7a, 0xb5, 0x79, 0x53, 0x14, 0x18, 0xfe, 0x1f,
	0x00, 0x00, 0xff, 0xff, 0xc0, 0x18, 0xe5, 0x5a, 0x4c, 0x0d, 0x00, 0x00,
}
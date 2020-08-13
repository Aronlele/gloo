// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/secret.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//Certain features such as the AWS Lambda option require the use of secrets for authentication, configuration of SSL Certificates, and other data that should not be stored in plaintext configuration.
//
//Gloo runs an independent (goroutine) controller to monitor secrets. Secrets are stored in their own secret storage layer. Gloo can monitor secrets stored in the following secret storage services:
//
//- Kubernetes Secrets
//- Hashicorp Vault
//- Plaintext files (recommended only for testing)
//- Secrets must adhere to a structure, specified by the option that requires them.
//
//Gloo's secret backend can be configured in Gloo's bootstrap options
type Secret struct {
	// Types that are valid to be assigned to Kind:
	//	*Secret_Aws
	//	*Secret_Azure
	//	*Secret_Tls
	//	*Secret_Oauth
	//	*Secret_ApiKey
	//	*Secret_Header
	//	*Secret_Extensions
	Kind isSecret_Kind `protobuf_oneof:"kind"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Secret) Reset()         { *m = Secret{} }
func (m *Secret) String() string { return proto.CompactTextString(m) }
func (*Secret) ProtoMessage()    {}
func (*Secret) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2f79c35f1213791, []int{0}
}
func (m *Secret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Secret.Unmarshal(m, b)
}
func (m *Secret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Secret.Marshal(b, m, deterministic)
}
func (m *Secret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Secret.Merge(m, src)
}
func (m *Secret) XXX_Size() int {
	return xxx_messageInfo_Secret.Size(m)
}
func (m *Secret) XXX_DiscardUnknown() {
	xxx_messageInfo_Secret.DiscardUnknown(m)
}

var xxx_messageInfo_Secret proto.InternalMessageInfo

type isSecret_Kind interface {
	isSecret_Kind()
	Equal(interface{}) bool
}

type Secret_Aws struct {
	Aws *AwsSecret `protobuf:"bytes,1,opt,name=aws,proto3,oneof" json:"aws,omitempty"`
}
type Secret_Azure struct {
	Azure *AzureSecret `protobuf:"bytes,2,opt,name=azure,proto3,oneof" json:"azure,omitempty"`
}
type Secret_Tls struct {
	Tls *TlsSecret `protobuf:"bytes,3,opt,name=tls,proto3,oneof" json:"tls,omitempty"`
}
type Secret_Oauth struct {
	Oauth *v1.OauthSecret `protobuf:"bytes,5,opt,name=oauth,proto3,oneof" json:"oauth,omitempty"`
}
type Secret_ApiKey struct {
	ApiKey *v1.ApiKeySecret `protobuf:"bytes,6,opt,name=api_key,json=apiKey,proto3,oneof" json:"api_key,omitempty"`
}
type Secret_Header struct {
	Header *HeaderSecret `protobuf:"bytes,8,opt,name=header,proto3,oneof" json:"header,omitempty"`
}
type Secret_Extensions struct {
	Extensions *Extensions `protobuf:"bytes,4,opt,name=extensions,proto3,oneof" json:"extensions,omitempty"`
}

func (*Secret_Aws) isSecret_Kind()        {}
func (*Secret_Azure) isSecret_Kind()      {}
func (*Secret_Tls) isSecret_Kind()        {}
func (*Secret_Oauth) isSecret_Kind()      {}
func (*Secret_ApiKey) isSecret_Kind()     {}
func (*Secret_Header) isSecret_Kind()     {}
func (*Secret_Extensions) isSecret_Kind() {}

func (m *Secret) GetKind() isSecret_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Secret) GetAws() *AwsSecret {
	if x, ok := m.GetKind().(*Secret_Aws); ok {
		return x.Aws
	}
	return nil
}

func (m *Secret) GetAzure() *AzureSecret {
	if x, ok := m.GetKind().(*Secret_Azure); ok {
		return x.Azure
	}
	return nil
}

func (m *Secret) GetTls() *TlsSecret {
	if x, ok := m.GetKind().(*Secret_Tls); ok {
		return x.Tls
	}
	return nil
}

func (m *Secret) GetOauth() *v1.OauthSecret {
	if x, ok := m.GetKind().(*Secret_Oauth); ok {
		return x.Oauth
	}
	return nil
}

func (m *Secret) GetApiKey() *v1.ApiKeySecret {
	if x, ok := m.GetKind().(*Secret_ApiKey); ok {
		return x.ApiKey
	}
	return nil
}

func (m *Secret) GetHeader() *HeaderSecret {
	if x, ok := m.GetKind().(*Secret_Header); ok {
		return x.Header
	}
	return nil
}

func (m *Secret) GetExtensions() *Extensions {
	if x, ok := m.GetKind().(*Secret_Extensions); ok {
		return x.Extensions
	}
	return nil
}

func (m *Secret) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Secret) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Secret_Aws)(nil),
		(*Secret_Azure)(nil),
		(*Secret_Tls)(nil),
		(*Secret_Oauth)(nil),
		(*Secret_ApiKey)(nil),
		(*Secret_Header)(nil),
		(*Secret_Extensions)(nil),
	}
}

//
//
//There are two ways of providing AWS secrets:
//
//- Method 1: `glooctl create secret aws`
//
// ```
// glooctl create secret aws --name aws-secret-from-glooctl \
//     --namespace default \
//     --access-key $ACC \
//     --secret-key $SEC
// ```
//
//will produce a Kubernetes resource similar to this (note the `aws` field and `resource_kind` annotation):
//
// ```
// apiVersion: v1
// data:
//   aws: base64EncodedStringForMachineConsumption
// kind: Secret
// metadata:
//   annotations:
//     resource_kind: '*v1.Secret'
//   creationTimestamp: "2019-08-23T15:10:20Z"
//   name: aws-secret-from-glooctl
//   namespace: default
//   resourceVersion: "592637"
//   selfLink: /api/v1/namespaces/default/secrets/secret-e2e
//   uid: 1f8c147f-c5b8-11e9-bbf3-42010a8001bc
// type: Opaque
// ```
//
// - Method 2: `kubectl apply -f resource-file.yaml`
//   - If using a git-ops flow, or otherwise creating secrets from yaml files, you may prefer to provide AWS credentials
//   using the format below, with `aws_access_key_id` and `aws_secret_access_key` fields.
//   - This circumvents the need for the annotation, which are not supported by some tools such as
//   [godaddy/kubernetes-external-secrets](https://github.com/godaddy/kubernetes-external-secrets)
//
// ```yaml
// # a sample aws secret resource-file.yaml
// apiVersion: v1
// data:
//   aws_access_key_id: some-id
//   aws_secret_access_key: some-secret
// kind: Secret
// metadata:
//   name: aws-secret-abcd
//   namespace: default
// ```
//
type AwsSecret struct {
	// provided by `glooctl create secret aws`
	AccessKey string `protobuf:"bytes,1,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	// provided by `glooctl create secret aws`
	SecretKey string `protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	// provided by `glooctl create secret aws`
	SessionToken         string   `protobuf:"bytes,3,opt,name=session_token,json=sessionToken,proto3" json:"session_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AwsSecret) Reset()         { *m = AwsSecret{} }
func (m *AwsSecret) String() string { return proto.CompactTextString(m) }
func (*AwsSecret) ProtoMessage()    {}
func (*AwsSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2f79c35f1213791, []int{1}
}
func (m *AwsSecret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsSecret.Unmarshal(m, b)
}
func (m *AwsSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsSecret.Marshal(b, m, deterministic)
}
func (m *AwsSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsSecret.Merge(m, src)
}
func (m *AwsSecret) XXX_Size() int {
	return xxx_messageInfo_AwsSecret.Size(m)
}
func (m *AwsSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsSecret.DiscardUnknown(m)
}

var xxx_messageInfo_AwsSecret proto.InternalMessageInfo

func (m *AwsSecret) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *AwsSecret) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func (m *AwsSecret) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

type AzureSecret struct {
	// provided by `glooctl create secret azure`
	ApiKeys              map[string]string `protobuf:"bytes,1,rep,name=api_keys,json=apiKeys,proto3" json:"api_keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AzureSecret) Reset()         { *m = AzureSecret{} }
func (m *AzureSecret) String() string { return proto.CompactTextString(m) }
func (*AzureSecret) ProtoMessage()    {}
func (*AzureSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2f79c35f1213791, []int{2}
}
func (m *AzureSecret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureSecret.Unmarshal(m, b)
}
func (m *AzureSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureSecret.Marshal(b, m, deterministic)
}
func (m *AzureSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureSecret.Merge(m, src)
}
func (m *AzureSecret) XXX_Size() int {
	return xxx_messageInfo_AzureSecret.Size(m)
}
func (m *AzureSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureSecret.DiscardUnknown(m)
}

var xxx_messageInfo_AzureSecret proto.InternalMessageInfo

func (m *AzureSecret) GetApiKeys() map[string]string {
	if m != nil {
		return m.ApiKeys
	}
	return nil
}

//
//Note that the annotation `resource_kind: '*v1.Secret'` is needed for Gloo to find this secret.
//Glooctl adds it by default when the tls secret is created via `glooctl create secret tls`.
type TlsSecret struct {
	// provided by `glooctl create secret tls`
	CertChain string `protobuf:"bytes,1,opt,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
	// provided by `glooctl create secret tls`
	PrivateKey string `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// provided by `glooctl create secret tls`
	RootCa               string   `protobuf:"bytes,3,opt,name=root_ca,json=rootCa,proto3" json:"root_ca,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TlsSecret) Reset()         { *m = TlsSecret{} }
func (m *TlsSecret) String() string { return proto.CompactTextString(m) }
func (*TlsSecret) ProtoMessage()    {}
func (*TlsSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2f79c35f1213791, []int{3}
}
func (m *TlsSecret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TlsSecret.Unmarshal(m, b)
}
func (m *TlsSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TlsSecret.Marshal(b, m, deterministic)
}
func (m *TlsSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TlsSecret.Merge(m, src)
}
func (m *TlsSecret) XXX_Size() int {
	return xxx_messageInfo_TlsSecret.Size(m)
}
func (m *TlsSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_TlsSecret.DiscardUnknown(m)
}

var xxx_messageInfo_TlsSecret proto.InternalMessageInfo

func (m *TlsSecret) GetCertChain() string {
	if m != nil {
		return m.CertChain
	}
	return ""
}

func (m *TlsSecret) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *TlsSecret) GetRootCa() string {
	if m != nil {
		return m.RootCa
	}
	return ""
}

type HeaderSecret struct {
	// A collection of header name to header value mappings, each representing an additional header that could be added to a request.
	// Provided by `glooctl create secret header`
	Headers              map[string]string `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HeaderSecret) Reset()         { *m = HeaderSecret{} }
func (m *HeaderSecret) String() string { return proto.CompactTextString(m) }
func (*HeaderSecret) ProtoMessage()    {}
func (*HeaderSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2f79c35f1213791, []int{4}
}
func (m *HeaderSecret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderSecret.Unmarshal(m, b)
}
func (m *HeaderSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderSecret.Marshal(b, m, deterministic)
}
func (m *HeaderSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderSecret.Merge(m, src)
}
func (m *HeaderSecret) XXX_Size() int {
	return xxx_messageInfo_HeaderSecret.Size(m)
}
func (m *HeaderSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderSecret.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderSecret proto.InternalMessageInfo

func (m *HeaderSecret) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func init() {
	proto.RegisterType((*Secret)(nil), "gloo.solo.io.Secret")
	proto.RegisterType((*AwsSecret)(nil), "gloo.solo.io.AwsSecret")
	proto.RegisterType((*AzureSecret)(nil), "gloo.solo.io.AzureSecret")
	proto.RegisterMapType((map[string]string)(nil), "gloo.solo.io.AzureSecret.ApiKeysEntry")
	proto.RegisterType((*TlsSecret)(nil), "gloo.solo.io.TlsSecret")
	proto.RegisterType((*HeaderSecret)(nil), "gloo.solo.io.HeaderSecret")
	proto.RegisterMapType((map[string]string)(nil), "gloo.solo.io.HeaderSecret.HeadersEntry")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/secret.proto", fileDescriptor_c2f79c35f1213791)
}

var fileDescriptor_c2f79c35f1213791 = []byte{
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xad, 0x9b, 0xc4, 0x69, 0x6e, 0x82, 0x84, 0x46, 0x55, 0x6b, 0x22, 0xb5, 0x45, 0xe1, 0x55,
	0x81, 0xb0, 0x69, 0x61, 0x51, 0x22, 0x16, 0xa4, 0x55, 0xa5, 0x4a, 0x08, 0x21, 0x85, 0xae, 0xd8,
	0x44, 0x53, 0x67, 0x94, 0x0c, 0x71, 0x3d, 0xd6, 0xcc, 0xa4, 0x0f, 0x96, 0x5d, 0xc3, 0x7f, 0xf0,
	0x09, 0x7c, 0x02, 0x5f, 0x81, 0x04, 0x7f, 0xd0, 0x05, 0x7b, 0x74, 0x67, 0xc6, 0x8e, 0x83, 0x08,
	0x82, 0x9d, 0xef, 0x3d, 0xe7, 0x5c, 0xcf, 0x99, 0x7b, 0x34, 0xf0, 0x7c, 0xc4, 0xf5, 0x78, 0x7a,
	0x12, 0xc6, 0xe2, 0x34, 0x52, 0x22, 0x11, 0x8f, 0xb9, 0x88, 0x46, 0x89, 0x10, 0x51, 0x26, 0xc5,
	0x7b, 0x16, 0x6b, 0x65, 0x2b, 0x9a, 0xf1, 0xe8, 0x6c, 0x27, 0x52, 0x2c, 0x96, 0x4c, 0x87, 0x99,
	0x14, 0x5a, 0x90, 0x16, 0x22, 0x21, 0x8a, 0x42, 0x2e, 0xda, 0xab, 0x23, 0x31, 0x12, 0x06, 0x88,
	0xf0, 0xcb, 0x72, 0xda, 0x84, 0x5d, 0x68, 0xdb, 0x64, 0x17, 0x4e, 0xd7, 0x7e, 0xb8, 0x78, 0x3e,
	0xbb, 0xd0, 0x2c, 0x55, 0x5c, 0xa4, 0xca, 0x71, 0x0f, 0xff, 0xc2, 0x4d, 0x35, 0x93, 0x99, 0xe4,
	0x8a, 0x45, 0x22, 0xd3, 0xa8, 0x41, 0x39, 0x9d, 0xea, 0xb1, 0x9b, 0x84, 0x9f, 0x6e, 0xcc, 0xa6,
	0xb1, 0x36, 0xe1, 0x3a, 0x17, 0x9f, 0x32, 0x4d, 0x87, 0x54, 0xd3, 0x45, 0x78, 0x5e, 0x5b, 0xbc,
	0xf3, 0xbd, 0x02, 0xfe, 0x5b, 0xe3, 0x9d, 0x3c, 0x82, 0x0a, 0x3d, 0x57, 0x81, 0x77, 0xdb, 0xdb,
	0x6e, 0xee, 0xae, 0x87, 0xe5, 0x3b, 0x08, 0x7b, 0xe7, 0xca, 0xb2, 0x8e, 0x96, 0xfa, 0xc8, 0x22,
	0x3b, 0x50, 0xa3, 0x1f, 0xa6, 0x92, 0x05, 0xcb, 0x86, 0x7e, 0xeb, 0x37, 0x3a, 0x42, 0x85, 0xc0,
	0x32, 0x71, 0xbe, 0x4e, 0x54, 0x50, 0xf9, 0xd3, 0xfc, 0xe3, 0xa4, 0x34, 0x5f, 0x27, 0x8a, 0xbc,
	0x80, 0x9a, 0x40, 0x9b, 0x41, 0xcd, 0xd0, 0xef, 0x86, 0xb3, 0x4b, 0x99, 0x57, 0xbe, 0x41, 0xd6,
	0xec, 0x57, 0x46, 0x44, 0x5e, 0x42, 0x9d, 0x66, 0x7c, 0x30, 0x61, 0x97, 0x81, 0x6f, 0xf4, 0xf7,
	0x16, 0xea, 0x7b, 0x19, 0x7f, 0xc5, 0x2e, 0x8b, 0x01, 0x3e, 0x35, 0x35, 0x79, 0x06, 0xfe, 0x98,
	0xd1, 0x21, 0x93, 0xc1, 0x8a, 0x19, 0xd0, 0x9e, 0x57, 0x1d, 0x19, 0x6c, 0xa6, 0xb2, 0x5c, 0xd2,
	0x05, 0x98, 0x2d, 0x3a, 0xa8, 0x1a, 0x65, 0x30, 0xaf, 0x3c, 0x2c, 0xf0, 0xa3, 0xa5, 0x7e, 0x89,
	0x4d, 0xf6, 0x60, 0x25, 0xdf, 0x5d, 0x50, 0x37, 0xca, 0xb5, 0x30, 0x16, 0x92, 0x15, 0xca, 0xd7,
	0x0e, 0xdd, 0xaf, 0x7e, 0xfd, 0xb6, 0xb5, 0xd4, 0x2f, 0xd8, 0xdd, 0xb5, 0xab, 0xeb, 0x6a, 0x0d,
	0x2a, 0x8a, 0xc5, 0x57, 0xd7, 0xd5, 0x06, 0xa9, 0xdb, 0x2c, 0xab, 0x7d, 0x1f, 0xaa, 0x13, 0x9e,
	0x0e, 0x3b, 0x29, 0x34, 0x8a, 0xfd, 0x91, 0x0d, 0x00, 0x1a, 0xc7, 0x4c, 0x29, 0x73, 0x3b, 0xb8,
	0xec, 0x46, 0xbf, 0x61, 0x3b, 0xe8, 0x7b, 0x03, 0xc0, 0xca, 0x0d, 0xbc, 0x6c, 0x61, 0xdb, 0x41,
	0xf8, 0x0e, 0xdc, 0x50, 0x4c, 0xe1, 0x81, 0x07, 0x5a, 0x4c, 0x58, 0x6a, 0xb6, 0xd9, 0xe8, 0xb7,
	0x5c, 0xf3, 0x18, 0x7b, 0x9d, 0x8f, 0x1e, 0x34, 0x4b, 0x09, 0x20, 0x3d, 0x58, 0x71, 0xdb, 0xc0,
	0x74, 0x55, 0xb6, 0x9b, 0xbb, 0xf7, 0x17, 0xc6, 0xc5, 0xed, 0x43, 0x1d, 0xa6, 0x5a, 0x5e, 0xf6,
	0xeb, 0x76, 0x1b, 0xaa, 0xdd, 0x85, 0x56, 0x19, 0x20, 0x37, 0xa1, 0x32, 0x3b, 0x3e, 0x7e, 0x92,
	0x55, 0xa8, 0x9d, 0xd1, 0x64, 0xca, 0xdc, 0x99, 0x6d, 0xd1, 0x5d, 0xde, 0xf3, 0x3a, 0x43, 0x68,
	0x14, 0xf1, 0x42, 0x7f, 0x31, 0x93, 0x7a, 0x10, 0x8f, 0x29, 0x4f, 0x73, 0xfb, 0xd8, 0x39, 0xc0,
	0x06, 0xd9, 0x82, 0x66, 0x26, 0xf9, 0x19, 0xd5, 0xac, 0xe4, 0x1f, 0x5c, 0x0b, 0x2f, 0x60, 0x1d,
	0xea, 0x52, 0x08, 0x3d, 0x88, 0xa9, 0xb3, 0xee, 0x63, 0x79, 0x40, 0x3b, 0x9f, 0x3c, 0x68, 0x95,
	0x53, 0x41, 0x7a, 0x50, 0xb7, 0xa9, 0xc8, 0x4d, 0x3f, 0x58, 0x1c, 0x21, 0x57, 0xe4, 0xae, 0x9d,
	0x0e, 0x5d, 0x97, 0x81, 0xff, 0x71, 0xbd, 0xdf, 0xfd, 0xf2, 0xb3, 0xea, 0x7d, 0xfe, 0xb1, 0xe9,
	0xbd, 0x7b, 0xf2, 0x6f, 0x0f, 0x61, 0x36, 0x19, 0xb9, 0x37, 0xe2, 0xc4, 0x37, 0x6f, 0xc3, 0xd3,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xfe, 0x86, 0x1a, 0x43, 0x05, 0x00, 0x00,
}

func (this *Secret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret)
	if !ok {
		that2, ok := that.(Secret)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Kind == nil {
		if this.Kind != nil {
			return false
		}
	} else if this.Kind == nil {
		return false
	} else if !this.Kind.Equal(that1.Kind) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Secret_Aws) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_Aws)
	if !ok {
		that2, ok := that.(Secret_Aws)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Aws.Equal(that1.Aws) {
		return false
	}
	return true
}
func (this *Secret_Azure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_Azure)
	if !ok {
		that2, ok := that.(Secret_Azure)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Azure.Equal(that1.Azure) {
		return false
	}
	return true
}
func (this *Secret_Tls) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_Tls)
	if !ok {
		that2, ok := that.(Secret_Tls)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Tls.Equal(that1.Tls) {
		return false
	}
	return true
}
func (this *Secret_Oauth) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_Oauth)
	if !ok {
		that2, ok := that.(Secret_Oauth)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Oauth.Equal(that1.Oauth) {
		return false
	}
	return true
}
func (this *Secret_ApiKey) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_ApiKey)
	if !ok {
		that2, ok := that.(Secret_ApiKey)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ApiKey.Equal(that1.ApiKey) {
		return false
	}
	return true
}
func (this *Secret_Header) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_Header)
	if !ok {
		that2, ok := that.(Secret_Header)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Header.Equal(that1.Header) {
		return false
	}
	return true
}
func (this *Secret_Extensions) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_Extensions)
	if !ok {
		that2, ok := that.(Secret_Extensions)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Extensions.Equal(that1.Extensions) {
		return false
	}
	return true
}
func (this *AwsSecret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AwsSecret)
	if !ok {
		that2, ok := that.(AwsSecret)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.AccessKey != that1.AccessKey {
		return false
	}
	if this.SecretKey != that1.SecretKey {
		return false
	}
	if this.SessionToken != that1.SessionToken {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AzureSecret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AzureSecret)
	if !ok {
		that2, ok := that.(AzureSecret)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.ApiKeys) != len(that1.ApiKeys) {
		return false
	}
	for i := range this.ApiKeys {
		if this.ApiKeys[i] != that1.ApiKeys[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TlsSecret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TlsSecret)
	if !ok {
		that2, ok := that.(TlsSecret)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.CertChain != that1.CertChain {
		return false
	}
	if this.PrivateKey != that1.PrivateKey {
		return false
	}
	if this.RootCa != that1.RootCa {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *HeaderSecret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HeaderSecret)
	if !ok {
		that2, ok := that.(HeaderSecret)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Headers) != len(that1.Headers) {
		return false
	}
	for i := range this.Headers {
		if this.Headers[i] != that1.Headers[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

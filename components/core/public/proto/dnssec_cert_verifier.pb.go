// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: dnssec_cert_verifier.proto

package proto

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

type SecurityState int32

const (
	// Check error code for more details.
	SecurityState_BOGUS SecurityState = 0
	// Certificate was validated successfully.
	SecurityState_SECURE SecurityState = 1
	// Fallback to WebPKI validation or fail.
	SecurityState_INSECURE SecurityState = 2
)

// Enum value maps for SecurityState.
var (
	SecurityState_name = map[int32]string{
		0: "BOGUS",
		1: "SECURE",
		2: "INSECURE",
	}
	SecurityState_value = map[string]int32{
		"BOGUS":    0,
		"SECURE":   1,
		"INSECURE": 2,
	}
)

func (x SecurityState) Enum() *SecurityState {
	p := new(SecurityState)
	*p = x
	return p
}

func (x SecurityState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SecurityState) Descriptor() protoreflect.EnumDescriptor {
	return file_dnssec_cert_verifier_proto_enumTypes[0].Descriptor()
}

func (SecurityState) Type() protoreflect.EnumType {
	return &file_dnssec_cert_verifier_proto_enumTypes[0]
}

func (x SecurityState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SecurityState.Descriptor instead.
func (SecurityState) EnumDescriptor() ([]byte, []int) {
	return file_dnssec_cert_verifier_proto_rawDescGZIP(), []int{0}
}

type ErrorCode int32

const (
	ErrorCode_UNKNOWN_ERROR ErrorCode = 0
	// DNSSEC errors
	ErrorCode_ERR_DNSSEC_BOGUS                        ErrorCode = 1
	ErrorCode_ERR_DNSSEC_SIGNATURE_EXPIRED            ErrorCode = 2
	ErrorCode_ERR_DNSSEC_SIGNATURE_MISSING            ErrorCode = 3
	ErrorCode_ERR_DNSSEC_DNSKEY_MISSING               ErrorCode = 4
	ErrorCode_ERR_DNSSEC_NSEC_MISSING                 ErrorCode = 5
	ErrorCode_ERR_DNSSEC_PINNED_KEY_NOT_IN_CERT_CHAIN ErrorCode = 6
	ErrorCode_ERR_DNSSEC_FETCH_FAILED                 ErrorCode = 7
	ErrorCode_ERR_DNSSEC_FETCH_TIMED_OUT              ErrorCode = 8
	// HNS errors
	ErrorCode_ERR_HNS_IS_SYNCING             ErrorCode = 9
	ErrorCode_ERR_HNS_NO_PEERS               ErrorCode = 10
	ErrorCode_ERR_HNS_PEER_TIMED_OUT         ErrorCode = 11
	ErrorCode_ERR_HNS_REQUEST_FAILED         ErrorCode = 12
	ErrorCode_ERR_HNS_HIP5_HANDLER_TIMED_OUT ErrorCode = 13
	ErrorCode_ERR_HNS_HIP5_HANDLER_FAILED    ErrorCode = 14
	// Trust service communication errors
	ErrorCode_ERR_TRUST_SERVICE_REQUEST_FAILED    ErrorCode = 15
	ErrorCode_ERR_TRUST_SERVICE_REQUEST_TIMED_OUT ErrorCode = 16
	ErrorCode_ERR_TRUST_SERVICE_REQUEST_INVALID   ErrorCode = 17
	ErrorCode_ERR_TRUST_SERVICE_RESPONSE_INVALID  ErrorCode = 18
	// Some net errors from chromium
	ErrorCode_ERR_DNS_SECURE_RESOLVER_HOSTNAME_RESOLUTION_FAILED ErrorCode = 19
	ErrorCode_ERR_DNS_TIMED_OUT                                  ErrorCode = 20
	ErrorCode_ERR_DNS_SERVER_FAILED                              ErrorCode = 21
	ErrorCode_ERR_DNS_MALFORMED_RESPONSE                         ErrorCode = 22
	ErrorCode_ERR_DNS_REQUEST_CANCELLED                          ErrorCode = 23
	// Cert errors from chromium
	ErrorCode_ERR_CERT_COMMON_NAME_INVALID        ErrorCode = 24
	ErrorCode_ERR_CERT_DATE_INVALID               ErrorCode = 25
	ErrorCode_ERR_CERT_AUTHORITY_INVALID          ErrorCode = 26
	ErrorCode_ERR_CERT_REVOKED                    ErrorCode = 27
	ErrorCode_ERR_CERT_INVALID                    ErrorCode = 28
	ErrorCode_ERR_CERT_WEAK_SIGNATURE_ALGORITHM   ErrorCode = 29
	ErrorCode_ERR_CERT_NON_UNIQUE_NAME            ErrorCode = 30
	ErrorCode_ERR_CERT_WEAK_KEY                   ErrorCode = 31
	ErrorCode_ERR_CERT_NAME_CONSTRAINT_VIOLATION  ErrorCode = 32
	ErrorCode_ERR_CERT_VALIDITY_TOO_LONG          ErrorCode = 33
	ErrorCode_ERR_CERT_KNOWN_INTERCEPTION_BLOCKED ErrorCode = 34
	// Some other generic errors
	ErrorCode_ERR_FAILED     ErrorCode = 35
	ErrorCode_ERR_ABORTED    ErrorCode = 36
	ErrorCode_ERR_UNEXPECTED ErrorCode = 37
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0:  "UNKNOWN_ERROR",
		1:  "ERR_DNSSEC_BOGUS",
		2:  "ERR_DNSSEC_SIGNATURE_EXPIRED",
		3:  "ERR_DNSSEC_SIGNATURE_MISSING",
		4:  "ERR_DNSSEC_DNSKEY_MISSING",
		5:  "ERR_DNSSEC_NSEC_MISSING",
		6:  "ERR_DNSSEC_PINNED_KEY_NOT_IN_CERT_CHAIN",
		7:  "ERR_DNSSEC_FETCH_FAILED",
		8:  "ERR_DNSSEC_FETCH_TIMED_OUT",
		9:  "ERR_HNS_IS_SYNCING",
		10: "ERR_HNS_NO_PEERS",
		11: "ERR_HNS_PEER_TIMED_OUT",
		12: "ERR_HNS_REQUEST_FAILED",
		13: "ERR_HNS_HIP5_HANDLER_TIMED_OUT",
		14: "ERR_HNS_HIP5_HANDLER_FAILED",
		15: "ERR_TRUST_SERVICE_REQUEST_FAILED",
		16: "ERR_TRUST_SERVICE_REQUEST_TIMED_OUT",
		17: "ERR_TRUST_SERVICE_REQUEST_INVALID",
		18: "ERR_TRUST_SERVICE_RESPONSE_INVALID",
		19: "ERR_DNS_SECURE_RESOLVER_HOSTNAME_RESOLUTION_FAILED",
		20: "ERR_DNS_TIMED_OUT",
		21: "ERR_DNS_SERVER_FAILED",
		22: "ERR_DNS_MALFORMED_RESPONSE",
		23: "ERR_DNS_REQUEST_CANCELLED",
		24: "ERR_CERT_COMMON_NAME_INVALID",
		25: "ERR_CERT_DATE_INVALID",
		26: "ERR_CERT_AUTHORITY_INVALID",
		27: "ERR_CERT_REVOKED",
		28: "ERR_CERT_INVALID",
		29: "ERR_CERT_WEAK_SIGNATURE_ALGORITHM",
		30: "ERR_CERT_NON_UNIQUE_NAME",
		31: "ERR_CERT_WEAK_KEY",
		32: "ERR_CERT_NAME_CONSTRAINT_VIOLATION",
		33: "ERR_CERT_VALIDITY_TOO_LONG",
		34: "ERR_CERT_KNOWN_INTERCEPTION_BLOCKED",
		35: "ERR_FAILED",
		36: "ERR_ABORTED",
		37: "ERR_UNEXPECTED",
	}
	ErrorCode_value = map[string]int32{
		"UNKNOWN_ERROR":                                      0,
		"ERR_DNSSEC_BOGUS":                                   1,
		"ERR_DNSSEC_SIGNATURE_EXPIRED":                       2,
		"ERR_DNSSEC_SIGNATURE_MISSING":                       3,
		"ERR_DNSSEC_DNSKEY_MISSING":                          4,
		"ERR_DNSSEC_NSEC_MISSING":                            5,
		"ERR_DNSSEC_PINNED_KEY_NOT_IN_CERT_CHAIN":            6,
		"ERR_DNSSEC_FETCH_FAILED":                            7,
		"ERR_DNSSEC_FETCH_TIMED_OUT":                         8,
		"ERR_HNS_IS_SYNCING":                                 9,
		"ERR_HNS_NO_PEERS":                                   10,
		"ERR_HNS_PEER_TIMED_OUT":                             11,
		"ERR_HNS_REQUEST_FAILED":                             12,
		"ERR_HNS_HIP5_HANDLER_TIMED_OUT":                     13,
		"ERR_HNS_HIP5_HANDLER_FAILED":                        14,
		"ERR_TRUST_SERVICE_REQUEST_FAILED":                   15,
		"ERR_TRUST_SERVICE_REQUEST_TIMED_OUT":                16,
		"ERR_TRUST_SERVICE_REQUEST_INVALID":                  17,
		"ERR_TRUST_SERVICE_RESPONSE_INVALID":                 18,
		"ERR_DNS_SECURE_RESOLVER_HOSTNAME_RESOLUTION_FAILED": 19,
		"ERR_DNS_TIMED_OUT":                                  20,
		"ERR_DNS_SERVER_FAILED":                              21,
		"ERR_DNS_MALFORMED_RESPONSE":                         22,
		"ERR_DNS_REQUEST_CANCELLED":                          23,
		"ERR_CERT_COMMON_NAME_INVALID":                       24,
		"ERR_CERT_DATE_INVALID":                              25,
		"ERR_CERT_AUTHORITY_INVALID":                         26,
		"ERR_CERT_REVOKED":                                   27,
		"ERR_CERT_INVALID":                                   28,
		"ERR_CERT_WEAK_SIGNATURE_ALGORITHM":                  29,
		"ERR_CERT_NON_UNIQUE_NAME":                           30,
		"ERR_CERT_WEAK_KEY":                                  31,
		"ERR_CERT_NAME_CONSTRAINT_VIOLATION":                 32,
		"ERR_CERT_VALIDITY_TOO_LONG":                         33,
		"ERR_CERT_KNOWN_INTERCEPTION_BLOCKED":                34,
		"ERR_FAILED":                                         35,
		"ERR_ABORTED":                                        36,
		"ERR_UNEXPECTED":                                     37,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_dnssec_cert_verifier_proto_enumTypes[1].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_dnssec_cert_verifier_proto_enumTypes[1]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_dnssec_cert_verifier_proto_rawDescGZIP(), []int{1}
}

type CertVerifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string       `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port string       `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	Cert *Certificate `protobuf:"bytes,3,opt,name=cert,proto3" json:"cert,omitempty"`
}

func (x *CertVerifyRequest) Reset() {
	*x = CertVerifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dnssec_cert_verifier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertVerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertVerifyRequest) ProtoMessage() {}

func (x *CertVerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dnssec_cert_verifier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertVerifyRequest.ProtoReflect.Descriptor instead.
func (*CertVerifyRequest) Descriptor() ([]byte, []int) {
	return file_dnssec_cert_verifier_proto_rawDescGZIP(), []int{0}
}

func (x *CertVerifyRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *CertVerifyRequest) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *CertVerifyRequest) GetCert() *Certificate {
	if x != nil {
		return x.Cert
	}
	return nil
}

type CertVerifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VerifiedCert   *Certificate  `protobuf:"bytes,1,opt,name=verified_cert,json=verifiedCert,proto3" json:"verified_cert,omitempty"`
	State          SecurityState `protobuf:"varint,2,opt,name=state,proto3,enum=dnssec_cert_verifier.SecurityState" json:"state,omitempty"`
	Code           ErrorCode     `protobuf:"varint,3,opt,name=code,proto3,enum=dnssec_cert_verifier.ErrorCode" json:"code,omitempty"`
	AdditionalInfo string        `protobuf:"bytes,4,opt,name=additional_info,json=additionalInfo,proto3" json:"additional_info,omitempty"`
}

func (x *CertVerifyResponse) Reset() {
	*x = CertVerifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dnssec_cert_verifier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertVerifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertVerifyResponse) ProtoMessage() {}

func (x *CertVerifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dnssec_cert_verifier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertVerifyResponse.ProtoReflect.Descriptor instead.
func (*CertVerifyResponse) Descriptor() ([]byte, []int) {
	return file_dnssec_cert_verifier_proto_rawDescGZIP(), []int{1}
}

func (x *CertVerifyResponse) GetVerifiedCert() *Certificate {
	if x != nil {
		return x.VerifiedCert
	}
	return nil
}

func (x *CertVerifyResponse) GetState() SecurityState {
	if x != nil {
		return x.State
	}
	return SecurityState_BOGUS
}

func (x *CertVerifyResponse) GetCode() ErrorCode {
	if x != nil {
		return x.Code
	}
	return ErrorCode_UNKNOWN_ERROR
}

func (x *CertVerifyResponse) GetAdditionalInfo() string {
	if x != nil {
		return x.AdditionalInfo
	}
	return ""
}

type Certificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DerCerts [][]byte `protobuf:"bytes,1,rep,name=der_certs,json=derCerts,proto3" json:"der_certs,omitempty"`
}

func (x *Certificate) Reset() {
	*x = Certificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dnssec_cert_verifier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Certificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Certificate) ProtoMessage() {}

func (x *Certificate) ProtoReflect() protoreflect.Message {
	mi := &file_dnssec_cert_verifier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Certificate.ProtoReflect.Descriptor instead.
func (*Certificate) Descriptor() ([]byte, []int) {
	return file_dnssec_cert_verifier_proto_rawDescGZIP(), []int{2}
}

func (x *Certificate) GetDerCerts() [][]byte {
	if x != nil {
		return x.DerCerts
	}
	return nil
}

var File_dnssec_cert_verifier_proto protoreflect.FileDescriptor

var file_dnssec_cert_verifier_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x64, 0x6e, 0x73, 0x73, 0x65, 0x63, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x64, 0x6e,
	0x73, 0x73, 0x65, 0x63, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x22, 0x72, 0x0a, 0x11, 0x43, 0x65, 0x72, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x35, 0x0a, 0x04, 0x63, 0x65, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x64, 0x6e, 0x73, 0x73, 0x65, 0x63, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x52, 0x04, 0x63, 0x65, 0x72, 0x74, 0x22, 0xf5, 0x01, 0x0a, 0x12, 0x43, 0x65, 0x72, 0x74, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a,
	0x0d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x6e, 0x73, 0x73, 0x65, 0x63, 0x5f, 0x63, 0x65,
	0x72, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x43, 0x65, 0x72, 0x74, 0x12, 0x39, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x64, 0x6e, 0x73, 0x73, 0x65, 0x63, 0x5f, 0x63, 0x65,
	0x72, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x33, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f,
	0x2e, 0x64, 0x6e, 0x73, 0x73, 0x65, 0x63, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2a,
	0x0a, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c,
	0x52, 0x08, 0x64, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x73, 0x2a, 0x34, 0x0a, 0x0d, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x42,
	0x4f, 0x47, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x43, 0x55, 0x52, 0x45,
	0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x53, 0x45, 0x43, 0x55, 0x52, 0x45, 0x10, 0x02,
	0x2a, 0x9b, 0x09, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x11,
	0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x52, 0x52, 0x5f, 0x44, 0x4e, 0x53, 0x53, 0x45, 0x43, 0x5f,
	0x42, 0x4f, 0x47, 0x55, 0x53, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x45, 0x52, 0x52, 0x5f, 0x44,
	0x4e, 0x53, 0x53, 0x45, 0x43, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f,
	0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c, 0x45, 0x52, 0x52,
	0x5f, 0x44, 0x4e, 0x53, 0x53, 0x45, 0x43, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x54, 0x55, 0x52,
	0x45, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x45,
	0x52, 0x52, 0x5f, 0x44, 0x4e, 0x53, 0x53, 0x45, 0x43, 0x5f, 0x44, 0x4e, 0x53, 0x4b, 0x45, 0x59,
	0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x52,
	0x52, 0x5f, 0x44, 0x4e, 0x53, 0x53, 0x45, 0x43, 0x5f, 0x4e, 0x53, 0x45, 0x43, 0x5f, 0x4d, 0x49,
	0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x2b, 0x0a, 0x27, 0x45, 0x52, 0x52, 0x5f, 0x44,
	0x4e, 0x53, 0x53, 0x45, 0x43, 0x5f, 0x50, 0x49, 0x4e, 0x4e, 0x45, 0x44, 0x5f, 0x4b, 0x45, 0x59,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x5f, 0x43, 0x48, 0x41,
	0x49, 0x4e, 0x10, 0x06, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x5f, 0x44, 0x4e, 0x53, 0x53,
	0x45, 0x43, 0x5f, 0x46, 0x45, 0x54, 0x43, 0x48, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10,
	0x07, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x52, 0x52, 0x5f, 0x44, 0x4e, 0x53, 0x53, 0x45, 0x43, 0x5f,
	0x46, 0x45, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x44, 0x5f, 0x4f, 0x55, 0x54, 0x10,
	0x08, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x52, 0x52, 0x5f, 0x48, 0x4e, 0x53, 0x5f, 0x49, 0x53, 0x5f,
	0x53, 0x59, 0x4e, 0x43, 0x49, 0x4e, 0x47, 0x10, 0x09, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x52, 0x52,
	0x5f, 0x48, 0x4e, 0x53, 0x5f, 0x4e, 0x4f, 0x5f, 0x50, 0x45, 0x45, 0x52, 0x53, 0x10, 0x0a, 0x12,
	0x1a, 0x0a, 0x16, 0x45, 0x52, 0x52, 0x5f, 0x48, 0x4e, 0x53, 0x5f, 0x50, 0x45, 0x45, 0x52, 0x5f,
	0x54, 0x49, 0x4d, 0x45, 0x44, 0x5f, 0x4f, 0x55, 0x54, 0x10, 0x0b, 0x12, 0x1a, 0x0a, 0x16, 0x45,
	0x52, 0x52, 0x5f, 0x48, 0x4e, 0x53, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x0c, 0x12, 0x22, 0x0a, 0x1e, 0x45, 0x52, 0x52, 0x5f, 0x48,
	0x4e, 0x53, 0x5f, 0x48, 0x49, 0x50, 0x35, 0x5f, 0x48, 0x41, 0x4e, 0x44, 0x4c, 0x45, 0x52, 0x5f,
	0x54, 0x49, 0x4d, 0x45, 0x44, 0x5f, 0x4f, 0x55, 0x54, 0x10, 0x0d, 0x12, 0x1f, 0x0a, 0x1b, 0x45,
	0x52, 0x52, 0x5f, 0x48, 0x4e, 0x53, 0x5f, 0x48, 0x49, 0x50, 0x35, 0x5f, 0x48, 0x41, 0x4e, 0x44,
	0x4c, 0x45, 0x52, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x0e, 0x12, 0x24, 0x0a, 0x20,
	0x45, 0x52, 0x52, 0x5f, 0x54, 0x52, 0x55, 0x53, 0x54, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43,
	0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x10, 0x0f, 0x12, 0x27, 0x0a, 0x23, 0x45, 0x52, 0x52, 0x5f, 0x54, 0x52, 0x55, 0x53, 0x54, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f,
	0x54, 0x49, 0x4d, 0x45, 0x44, 0x5f, 0x4f, 0x55, 0x54, 0x10, 0x10, 0x12, 0x25, 0x0a, 0x21, 0x45,
	0x52, 0x52, 0x5f, 0x54, 0x52, 0x55, 0x53, 0x54, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x10, 0x11, 0x12, 0x26, 0x0a, 0x22, 0x45, 0x52, 0x52, 0x5f, 0x54, 0x52, 0x55, 0x53, 0x54, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45,
	0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x12, 0x12, 0x36, 0x0a, 0x32, 0x45, 0x52,
	0x52, 0x5f, 0x44, 0x4e, 0x53, 0x5f, 0x53, 0x45, 0x43, 0x55, 0x52, 0x45, 0x5f, 0x52, 0x45, 0x53,
	0x4f, 0x4c, 0x56, 0x45, 0x52, 0x5f, 0x48, 0x4f, 0x53, 0x54, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x52,
	0x45, 0x53, 0x4f, 0x4c, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x10, 0x13, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x52, 0x52, 0x5f, 0x44, 0x4e, 0x53, 0x5f, 0x54, 0x49,
	0x4d, 0x45, 0x44, 0x5f, 0x4f, 0x55, 0x54, 0x10, 0x14, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x52, 0x52,
	0x5f, 0x44, 0x4e, 0x53, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x15, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x52, 0x52, 0x5f, 0x44, 0x4e, 0x53, 0x5f,
	0x4d, 0x41, 0x4c, 0x46, 0x4f, 0x52, 0x4d, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e,
	0x53, 0x45, 0x10, 0x16, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x52, 0x52, 0x5f, 0x44, 0x4e, 0x53, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45,
	0x44, 0x10, 0x17, 0x12, 0x20, 0x0a, 0x1c, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x5f,
	0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0x18, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x45, 0x52,
	0x54, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x19,
	0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x5f, 0x41, 0x55, 0x54,
	0x48, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x1a,
	0x12, 0x14, 0x0a, 0x10, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x5f, 0x52, 0x45, 0x56,
	0x4f, 0x4b, 0x45, 0x44, 0x10, 0x1b, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x45,
	0x52, 0x54, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x1c, 0x12, 0x25, 0x0a, 0x21,
	0x45, 0x52, 0x52, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x5f, 0x57, 0x45, 0x41, 0x4b, 0x5f, 0x53, 0x49,
	0x47, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x41, 0x4c, 0x47, 0x4f, 0x52, 0x49, 0x54, 0x48,
	0x4d, 0x10, 0x1d, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x5f,
	0x4e, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x49, 0x51, 0x55, 0x45, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10,
	0x1e, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x5f, 0x57, 0x45,
	0x41, 0x4b, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x1f, 0x12, 0x26, 0x0a, 0x22, 0x45, 0x52, 0x52, 0x5f,
	0x43, 0x45, 0x52, 0x54, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x54, 0x52,
	0x41, 0x49, 0x4e, 0x54, 0x5f, 0x56, 0x49, 0x4f, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x20,
	0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x5f, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x4f, 0x4e, 0x47, 0x10, 0x21,
	0x12, 0x27, 0x0a, 0x23, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x5f, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x43, 0x45, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x22, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x52, 0x52,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x23, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x52, 0x52,
	0x5f, 0x41, 0x42, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x24, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x52,
	0x52, 0x5f, 0x55, 0x4e, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x25, 0x32, 0x71,
	0x0a, 0x0c, 0x43, 0x65, 0x72, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x61,
	0x0a, 0x0a, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x65, 0x72, 0x74, 0x12, 0x27, 0x2e, 0x64,
	0x6e, 0x73, 0x73, 0x65, 0x63, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x64, 0x6e, 0x73, 0x73, 0x65, 0x63, 0x5f, 0x63,
	0x65, 0x72, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x43, 0x65, 0x72,
	0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x35, 0x48, 0x03, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x6d, 0x70, 0x65, 0x72, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x69, 0x6e, 0x63, 0x2f,
	0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dnssec_cert_verifier_proto_rawDescOnce sync.Once
	file_dnssec_cert_verifier_proto_rawDescData = file_dnssec_cert_verifier_proto_rawDesc
)

func file_dnssec_cert_verifier_proto_rawDescGZIP() []byte {
	file_dnssec_cert_verifier_proto_rawDescOnce.Do(func() {
		file_dnssec_cert_verifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_dnssec_cert_verifier_proto_rawDescData)
	})
	return file_dnssec_cert_verifier_proto_rawDescData
}

var file_dnssec_cert_verifier_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_dnssec_cert_verifier_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_dnssec_cert_verifier_proto_goTypes = []interface{}{
	(SecurityState)(0),         // 0: dnssec_cert_verifier.SecurityState
	(ErrorCode)(0),             // 1: dnssec_cert_verifier.ErrorCode
	(*CertVerifyRequest)(nil),  // 2: dnssec_cert_verifier.CertVerifyRequest
	(*CertVerifyResponse)(nil), // 3: dnssec_cert_verifier.CertVerifyResponse
	(*Certificate)(nil),        // 4: dnssec_cert_verifier.Certificate
}
var file_dnssec_cert_verifier_proto_depIdxs = []int32{
	4, // 0: dnssec_cert_verifier.CertVerifyRequest.cert:type_name -> dnssec_cert_verifier.Certificate
	4, // 1: dnssec_cert_verifier.CertVerifyResponse.verified_cert:type_name -> dnssec_cert_verifier.Certificate
	0, // 2: dnssec_cert_verifier.CertVerifyResponse.state:type_name -> dnssec_cert_verifier.SecurityState
	1, // 3: dnssec_cert_verifier.CertVerifyResponse.code:type_name -> dnssec_cert_verifier.ErrorCode
	2, // 4: dnssec_cert_verifier.CertVerifier.VerifyCert:input_type -> dnssec_cert_verifier.CertVerifyRequest
	3, // 5: dnssec_cert_verifier.CertVerifier.VerifyCert:output_type -> dnssec_cert_verifier.CertVerifyResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_dnssec_cert_verifier_proto_init() }
func file_dnssec_cert_verifier_proto_init() {
	if File_dnssec_cert_verifier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dnssec_cert_verifier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertVerifyRequest); i {
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
		file_dnssec_cert_verifier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertVerifyResponse); i {
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
		file_dnssec_cert_verifier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Certificate); i {
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
			RawDescriptor: file_dnssec_cert_verifier_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dnssec_cert_verifier_proto_goTypes,
		DependencyIndexes: file_dnssec_cert_verifier_proto_depIdxs,
		EnumInfos:         file_dnssec_cert_verifier_proto_enumTypes,
		MessageInfos:      file_dnssec_cert_verifier_proto_msgTypes,
	}.Build()
	File_dnssec_cert_verifier_proto = out.File
	file_dnssec_cert_verifier_proto_rawDesc = nil
	file_dnssec_cert_verifier_proto_goTypes = nil
	file_dnssec_cert_verifier_proto_depIdxs = nil
}

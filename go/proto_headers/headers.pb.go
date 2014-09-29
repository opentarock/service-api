// Code generated by protoc-gen-gogo.
// source: headers.proto
// DO NOT EDIT!

/*
Package proto_headers is a generated protocol buffer package.

It is generated from these files:
	headers.proto

It has these top-level messages:
	AuthorizationHeader
	RequestCorrelationHeader
*/
package proto_headers

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type AuthorizationHeader struct {
	UserId           *uint64 `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	AccessToken      *string `protobuf:"bytes,2,req,name=access_token" json:"access_token,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AuthorizationHeader) Reset()         { *m = AuthorizationHeader{} }
func (m *AuthorizationHeader) String() string { return proto.CompactTextString(m) }
func (*AuthorizationHeader) ProtoMessage()    {}

func (m *AuthorizationHeader) GetUserId() uint64 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *AuthorizationHeader) GetAccessToken() string {
	if m != nil && m.AccessToken != nil {
		return *m.AccessToken
	}
	return ""
}

type RequestCorrelationHeader struct {
	CorrelationId    *string `protobuf:"bytes,1,req,name=correlation_id" json:"correlation_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RequestCorrelationHeader) Reset()         { *m = RequestCorrelationHeader{} }
func (m *RequestCorrelationHeader) String() string { return proto.CompactTextString(m) }
func (*RequestCorrelationHeader) ProtoMessage()    {}

func (m *RequestCorrelationHeader) GetCorrelationId() string {
	if m != nil && m.CorrelationId != nil {
		return *m.CorrelationId
	}
	return ""
}

func init() {
}

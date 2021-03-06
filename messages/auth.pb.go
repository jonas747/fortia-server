// Code generated by protoc-gen-go.
// source: auth.proto
// DO NOT EDIT!

package messages

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Login struct {
	Username         *string `protobuf:"bytes,1,req" json:"Username,omitempty"`
	Password         *string `protobuf:"bytes,2,req" json:"Password,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Login) Reset()         { *m = Login{} }
func (m *Login) String() string { return proto.CompactTextString(m) }
func (*Login) ProtoMessage()    {}

func (m *Login) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *Login) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

type LoginResp struct {
	Ok               *bool   `protobuf:"varint,1,req" json:"Ok,omitempty"`
	Err              *string `protobuf:"bytes,2,opt" json:"Err,omitempty"`
	Modhash          *string `protobuf:"bytes,3,opt" json:"Modhash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LoginResp) Reset()         { *m = LoginResp{} }
func (m *LoginResp) String() string { return proto.CompactTextString(m) }
func (*LoginResp) ProtoMessage()    {}

func (m *LoginResp) GetOk() bool {
	if m != nil && m.Ok != nil {
		return *m.Ok
	}
	return false
}

func (m *LoginResp) GetErr() string {
	if m != nil && m.Err != nil {
		return *m.Err
	}
	return ""
}

func (m *LoginResp) GetModhash() string {
	if m != nil && m.Modhash != nil {
		return *m.Modhash
	}
	return ""
}

type Register struct {
	Username         *string `protobuf:"bytes,1,req" json:"Username,omitempty"`
	Password         *string `protobuf:"bytes,2,req" json:"Password,omitempty"`
	Email            *string `protobuf:"bytes,3,req" json:"Email,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Register) Reset()         { *m = Register{} }
func (m *Register) String() string { return proto.CompactTextString(m) }
func (*Register) ProtoMessage()    {}

func (m *Register) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *Register) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *Register) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

type ChangePassword struct {
	OldPassword      *string `protobuf:"bytes,1,req" json:"OldPassword,omitempty"`
	Modhash          *string `protobuf:"bytes,2,req" json:"Modhash,omitempty"`
	NewPassword      *string `protobuf:"bytes,3,req" json:"NewPassword,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ChangePassword) Reset()         { *m = ChangePassword{} }
func (m *ChangePassword) String() string { return proto.CompactTextString(m) }
func (*ChangePassword) ProtoMessage()    {}

func (m *ChangePassword) GetOldPassword() string {
	if m != nil && m.OldPassword != nil {
		return *m.OldPassword
	}
	return ""
}

func (m *ChangePassword) GetModhash() string {
	if m != nil && m.Modhash != nil {
		return *m.Modhash
	}
	return ""
}

func (m *ChangePassword) GetNewPassword() string {
	if m != nil && m.NewPassword != nil {
		return *m.NewPassword
	}
	return ""
}

type ChangeEmail struct {
	Modhash          *string `protobuf:"bytes,1,req" json:"Modhash,omitempty"`
	NewEmail         *string `protobuf:"bytes,2,req" json:"NewEmail,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ChangeEmail) Reset()         { *m = ChangeEmail{} }
func (m *ChangeEmail) String() string { return proto.CompactTextString(m) }
func (*ChangeEmail) ProtoMessage()    {}

func (m *ChangeEmail) GetModhash() string {
	if m != nil && m.Modhash != nil {
		return *m.Modhash
	}
	return ""
}

func (m *ChangeEmail) GetNewEmail() string {
	if m != nil && m.NewEmail != nil {
		return *m.NewEmail
	}
	return ""
}

func init() {
}

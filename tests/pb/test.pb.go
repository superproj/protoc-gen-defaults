// Copyright 2021 Linka Cloud  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: tests/pb/test.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	_ "go.linka.cloud/protoc-gen-defaults/defaults"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Test_Type int32

const (
	Test_NONE Test_Type = 0
	Test_ONE  Test_Type = 1
	Test_TWO  Test_Type = 2
)

// Enum value maps for Test_Type.
var (
	Test_Type_name = map[int32]string{
		0: "NONE",
		1: "ONE",
		2: "TWO",
	}
	Test_Type_value = map[string]int32{
		"NONE": 0,
		"ONE":  1,
		"TWO":  2,
	}
)

func (x Test_Type) Enum() *Test_Type {
	p := new(Test_Type)
	*p = x
	return p
}

func (x Test_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Test_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_pb_test_proto_enumTypes[0].Descriptor()
}

func (Test_Type) Type() protoreflect.EnumType {
	return &file_tests_pb_test_proto_enumTypes[0]
}

func (x Test_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Test_Type.Descriptor instead.
func (Test_Type) EnumDescriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{0, 0}
}

type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringField          string                  `protobuf:"bytes,1,opt,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
	NumberField          int64                   `protobuf:"varint,2,opt,name=number_field,json=numberField,proto3" json:"number_field,omitempty"`
	BoolField            bool                    `protobuf:"varint,3,opt,name=bool_field,json=boolField,proto3" json:"bool_field,omitempty"`
	EnumField            Test_Type               `protobuf:"varint,4,opt,name=enum_field,json=enumField,proto3,enum=go.linka.cloud.protoc_gen_defaults.test.Test_Type" json:"enum_field,omitempty"`
	MessageField         *Test                   `protobuf:"bytes,5,opt,name=message_field,json=messageField,proto3" json:"message_field,omitempty"`
	RepeatedStringField  []string                `protobuf:"bytes,6,rep,name=repeated_string_field,json=repeatedStringField,proto3" json:"repeated_string_field,omitempty"`
	RepeatedMessageField []Test_Type             `protobuf:"varint,7,rep,packed,name=repeated_message_field,json=repeatedMessageField,proto3,enum=go.linka.cloud.protoc_gen_defaults.test.Test_Type" json:"repeated_message_field,omitempty"`
	NumberValueField     *wrapperspb.Int64Value  `protobuf:"bytes,8,opt,name=number_value_field,json=numberValueField,proto3" json:"number_value_field,omitempty"`
	StringValueField     *wrapperspb.StringValue `protobuf:"bytes,9,opt,name=string_value_field,json=stringValueField,proto3" json:"string_value_field,omitempty"`
	BoolValueField       *wrapperspb.BoolValue   `protobuf:"bytes,10,opt,name=bool_value_field,json=boolValueField,proto3" json:"bool_value_field,omitempty"`
	TimeValueField       *timestamppb.Timestamp  `protobuf:"bytes,11,opt,name=time_value_field,json=timeValueField,proto3" json:"time_value_field,omitempty"`
	DurationValueField   *durationpb.Duration    `protobuf:"bytes,12,opt,name=duration_value_field,json=durationValueField,proto3" json:"duration_value_field,omitempty"`
	// Types that are assignable to Oneof:
	//	*Test_One
	//	*Test_Two
	//	*Test_Three
	//	*Test_Four
	Oneof                     isTest_Oneof                  `protobuf_oneof:"oneof"`
	Descriptor_               *descriptorpb.DescriptorProto `protobuf:"bytes,17,opt,name=descriptor,proto3" json:"descriptor,omitempty"`
	TimeValueFieldWithDefault *timestamppb.Timestamp        `protobuf:"bytes,18,opt,name=time_value_field_with_default,json=timeValueFieldWithDefault,proto3" json:"time_value_field_with_default,omitempty"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_tests_pb_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{0}
}

func (x *Test) GetStringField() string {
	if x != nil {
		return x.StringField
	}
	return ""
}

func (x *Test) GetNumberField() int64 {
	if x != nil {
		return x.NumberField
	}
	return 0
}

func (x *Test) GetBoolField() bool {
	if x != nil {
		return x.BoolField
	}
	return false
}

func (x *Test) GetEnumField() Test_Type {
	if x != nil {
		return x.EnumField
	}
	return Test_NONE
}

func (x *Test) GetMessageField() *Test {
	if x != nil {
		return x.MessageField
	}
	return nil
}

func (x *Test) GetRepeatedStringField() []string {
	if x != nil {
		return x.RepeatedStringField
	}
	return nil
}

func (x *Test) GetRepeatedMessageField() []Test_Type {
	if x != nil {
		return x.RepeatedMessageField
	}
	return nil
}

func (x *Test) GetNumberValueField() *wrapperspb.Int64Value {
	if x != nil {
		return x.NumberValueField
	}
	return nil
}

func (x *Test) GetStringValueField() *wrapperspb.StringValue {
	if x != nil {
		return x.StringValueField
	}
	return nil
}

func (x *Test) GetBoolValueField() *wrapperspb.BoolValue {
	if x != nil {
		return x.BoolValueField
	}
	return nil
}

func (x *Test) GetTimeValueField() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeValueField
	}
	return nil
}

func (x *Test) GetDurationValueField() *durationpb.Duration {
	if x != nil {
		return x.DurationValueField
	}
	return nil
}

func (m *Test) GetOneof() isTest_Oneof {
	if m != nil {
		return m.Oneof
	}
	return nil
}

func (x *Test) GetOne() *OneOfOne {
	if x, ok := x.GetOneof().(*Test_One); ok {
		return x.One
	}
	return nil
}

func (x *Test) GetTwo() *OneOfTwo {
	if x, ok := x.GetOneof().(*Test_Two); ok {
		return x.Two
	}
	return nil
}

func (x *Test) GetThree() *OneOfThree {
	if x, ok := x.GetOneof().(*Test_Three); ok {
		return x.Three
	}
	return nil
}

func (x *Test) GetFour() Test_Type {
	if x, ok := x.GetOneof().(*Test_Four); ok {
		return x.Four
	}
	return Test_NONE
}

func (x *Test) GetDescriptor_() *descriptorpb.DescriptorProto {
	if x != nil {
		return x.Descriptor_
	}
	return nil
}

func (x *Test) GetTimeValueFieldWithDefault() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeValueFieldWithDefault
	}
	return nil
}

type isTest_Oneof interface {
	isTest_Oneof()
}

type Test_One struct {
	One *OneOfOne `protobuf:"bytes,13,opt,name=one,proto3,oneof"`
}

type Test_Two struct {
	Two *OneOfTwo `protobuf:"bytes,14,opt,name=two,proto3,oneof"`
}

type Test_Three struct {
	Three *OneOfThree `protobuf:"bytes,15,opt,name=three,proto3,oneof"`
}

type Test_Four struct {
	Four Test_Type `protobuf:"varint,16,opt,name=four,proto3,enum=go.linka.cloud.protoc_gen_defaults.test.Test_Type,oneof"`
}

func (*Test_One) isTest_Oneof() {}

func (*Test_Two) isTest_Oneof() {}

func (*Test_Three) isTest_Oneof() {}

func (*Test_Four) isTest_Oneof() {}

type OneOfOne struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringField string `protobuf:"bytes,1,opt,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
}

func (x *OneOfOne) Reset() {
	*x = OneOfOne{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneOfOne) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneOfOne) ProtoMessage() {}

func (x *OneOfOne) ProtoReflect() protoreflect.Message {
	mi := &file_tests_pb_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneOfOne.ProtoReflect.Descriptor instead.
func (*OneOfOne) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{1}
}

func (x *OneOfOne) GetStringField() string {
	if x != nil {
		return x.StringField
	}
	return ""
}

type OneOfTwo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringField string `protobuf:"bytes,1,opt,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
}

func (x *OneOfTwo) Reset() {
	*x = OneOfTwo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneOfTwo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneOfTwo) ProtoMessage() {}

func (x *OneOfTwo) ProtoReflect() protoreflect.Message {
	mi := &file_tests_pb_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneOfTwo.ProtoReflect.Descriptor instead.
func (*OneOfTwo) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{2}
}

func (x *OneOfTwo) GetStringField() string {
	if x != nil {
		return x.StringField
	}
	return ""
}

type OneOfThree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringField string `protobuf:"bytes,1,opt,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
}

func (x *OneOfThree) Reset() {
	*x = OneOfThree{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneOfThree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneOfThree) ProtoMessage() {}

func (x *OneOfThree) ProtoReflect() protoreflect.Message {
	mi := &file_tests_pb_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneOfThree.ProtoReflect.Descriptor instead.
func (*OneOfThree) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{3}
}

func (x *OneOfThree) GetStringField() string {
	if x != nil {
		return x.StringField
	}
	return ""
}

var File_tests_pb_test_proto protoreflect.FileDescriptor

var file_tests_pb_test_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x67, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x61, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e,
	0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x2f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x0b, 0x0a, 0x04, 0x54,
	0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0x9a, 0x49, 0x0e, 0x72, 0x0c,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0b, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x28, 0x0a, 0x0c, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x05, 0x9a, 0x49, 0x02, 0x20, 0x2a, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x24, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42, 0x05, 0x9a, 0x49, 0x02, 0x68, 0x01, 0x52, 0x09,
	0x62, 0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x59, 0x0a, 0x0a, 0x65, 0x6e, 0x75,
	0x6d, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e,
	0x67, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x06, 0x9a, 0x49, 0x03, 0x80, 0x01, 0x02, 0x52, 0x09, 0x65, 0x6e, 0x75, 0x6d, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x5c, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f,
	0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x42, 0x08, 0x9a, 0x49, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x13, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x68, 0x0a, 0x16, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b,
	0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67,
	0x65, 0x6e, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x14, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x50, 0x0a, 0x12, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x05, 0x9a, 0x49, 0x02, 0x20, 0x2b,
	0x52, 0x10, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x5d, 0x0a, 0x12, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x11, 0x9a, 0x49,
	0x0e, 0x72, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x10, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x4b, 0x0a, 0x10, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x05, 0x9a, 0x49, 0x02, 0x68, 0x00, 0x52, 0x0e,
	0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x4f,
	0x0a, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x09, 0x9a, 0x49, 0x06, 0xb2, 0x01, 0x03, 0x6e, 0x6f, 0x77, 0x52,
	0x0e, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x56, 0x0a, 0x14, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x9a, 0x49, 0x06, 0xaa, 0x01, 0x03,
	0x34, 0x32, 0x73, 0x52, 0x12, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x4f, 0x0a, 0x03, 0x6f, 0x6e, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x61, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e,
	0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4f,
	0x6e, 0x65, 0x4f, 0x66, 0x4f, 0x6e, 0x65, 0x42, 0x08, 0x9a, 0x49, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x48, 0x00, 0x52, 0x03, 0x6f, 0x6e, 0x65, 0x12, 0x4f, 0x0a, 0x03, 0x74, 0x77, 0x6f, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x61,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65,
	0x6e, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x54, 0x77, 0x6f, 0x42, 0x08, 0x9a, 0x49, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x48, 0x00, 0x52, 0x03, 0x74, 0x77, 0x6f, 0x12, 0x55, 0x0a, 0x05, 0x74, 0x68, 0x72,
	0x65, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x2e, 0x6c, 0x69,
	0x6e, 0x6b, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x54, 0x68, 0x72, 0x65, 0x65, 0x42, 0x08, 0x9a,
	0x49, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x05, 0x74, 0x68, 0x72, 0x65, 0x65,
	0x12, 0x50, 0x0a, 0x04, 0x66, 0x6f, 0x75, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32,
	0x2e, 0x67, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x06, 0x9a, 0x49, 0x03, 0x80, 0x01, 0x01, 0x48, 0x00, 0x52, 0x04, 0x66, 0x6f,
	0x75, 0x72, 0x12, 0x4c, 0x0a, 0x0a, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0a, 0x9a, 0x49, 0x07, 0x8a, 0x01, 0x04,
	0x08, 0x01, 0x10, 0x01, 0x52, 0x0a, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x12, 0x78, 0x0a, 0x1d, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x1a, 0x9a, 0x49, 0x17, 0xb2, 0x01, 0x14, 0x31, 0x39, 0x35, 0x32, 0x2d,
	0x30, 0x33, 0x2d, 0x31, 0x31, 0x54, 0x30, 0x30, 0x3a, 0x30, 0x30, 0x3a, 0x30, 0x30, 0x5a, 0x52,
	0x19, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x57,
	0x69, 0x74, 0x68, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x22, 0x22, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03,
	0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x57, 0x4f, 0x10, 0x02, 0x42, 0x0f,
	0x0a, 0x05, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x06, 0x9a, 0x49, 0x03, 0x74, 0x77, 0x6f, 0x22,
	0x40, 0x0a, 0x08, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x4f, 0x6e, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x11, 0x9a, 0x49, 0x0e, 0x72, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x22, 0x40, 0x0a, 0x08, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x54, 0x77, 0x6f, 0x12, 0x34, 0x0a,
	0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x11, 0x9a, 0x49, 0x0e, 0x72, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x22, 0x47, 0x0a, 0x0a, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x54, 0x68, 0x72, 0x65,
	0x65, 0x12, 0x34, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0x9a, 0x49, 0x0e, 0x72, 0x0c, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x03, 0x98, 0x49, 0x01, 0x42, 0x2d, 0x5a, 0x2b,
	0x67, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tests_pb_test_proto_rawDescOnce sync.Once
	file_tests_pb_test_proto_rawDescData = file_tests_pb_test_proto_rawDesc
)

func file_tests_pb_test_proto_rawDescGZIP() []byte {
	file_tests_pb_test_proto_rawDescOnce.Do(func() {
		file_tests_pb_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_pb_test_proto_rawDescData)
	})
	return file_tests_pb_test_proto_rawDescData
}

var file_tests_pb_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tests_pb_test_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tests_pb_test_proto_goTypes = []interface{}{
	(Test_Type)(0),                       // 0: go.linka.cloud.protoc_gen_defaults.test.Test.Type
	(*Test)(nil),                         // 1: go.linka.cloud.protoc_gen_defaults.test.Test
	(*OneOfOne)(nil),                     // 2: go.linka.cloud.protoc_gen_defaults.test.OneOfOne
	(*OneOfTwo)(nil),                     // 3: go.linka.cloud.protoc_gen_defaults.test.OneOfTwo
	(*OneOfThree)(nil),                   // 4: go.linka.cloud.protoc_gen_defaults.test.OneOfThree
	(*wrapperspb.Int64Value)(nil),        // 5: google.protobuf.Int64Value
	(*wrapperspb.StringValue)(nil),       // 6: google.protobuf.StringValue
	(*wrapperspb.BoolValue)(nil),         // 7: google.protobuf.BoolValue
	(*timestamppb.Timestamp)(nil),        // 8: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),          // 9: google.protobuf.Duration
	(*descriptorpb.DescriptorProto)(nil), // 10: google.protobuf.DescriptorProto
}
var file_tests_pb_test_proto_depIdxs = []int32{
	0,  // 0: go.linka.cloud.protoc_gen_defaults.test.Test.enum_field:type_name -> go.linka.cloud.protoc_gen_defaults.test.Test.Type
	1,  // 1: go.linka.cloud.protoc_gen_defaults.test.Test.message_field:type_name -> go.linka.cloud.protoc_gen_defaults.test.Test
	0,  // 2: go.linka.cloud.protoc_gen_defaults.test.Test.repeated_message_field:type_name -> go.linka.cloud.protoc_gen_defaults.test.Test.Type
	5,  // 3: go.linka.cloud.protoc_gen_defaults.test.Test.number_value_field:type_name -> google.protobuf.Int64Value
	6,  // 4: go.linka.cloud.protoc_gen_defaults.test.Test.string_value_field:type_name -> google.protobuf.StringValue
	7,  // 5: go.linka.cloud.protoc_gen_defaults.test.Test.bool_value_field:type_name -> google.protobuf.BoolValue
	8,  // 6: go.linka.cloud.protoc_gen_defaults.test.Test.time_value_field:type_name -> google.protobuf.Timestamp
	9,  // 7: go.linka.cloud.protoc_gen_defaults.test.Test.duration_value_field:type_name -> google.protobuf.Duration
	2,  // 8: go.linka.cloud.protoc_gen_defaults.test.Test.one:type_name -> go.linka.cloud.protoc_gen_defaults.test.OneOfOne
	3,  // 9: go.linka.cloud.protoc_gen_defaults.test.Test.two:type_name -> go.linka.cloud.protoc_gen_defaults.test.OneOfTwo
	4,  // 10: go.linka.cloud.protoc_gen_defaults.test.Test.three:type_name -> go.linka.cloud.protoc_gen_defaults.test.OneOfThree
	0,  // 11: go.linka.cloud.protoc_gen_defaults.test.Test.four:type_name -> go.linka.cloud.protoc_gen_defaults.test.Test.Type
	10, // 12: go.linka.cloud.protoc_gen_defaults.test.Test.descriptor:type_name -> google.protobuf.DescriptorProto
	8,  // 13: go.linka.cloud.protoc_gen_defaults.test.Test.time_value_field_with_default:type_name -> google.protobuf.Timestamp
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_tests_pb_test_proto_init() }
func file_tests_pb_test_proto_init() {
	if File_tests_pb_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_pb_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
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
		file_tests_pb_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneOfOne); i {
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
		file_tests_pb_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneOfTwo); i {
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
		file_tests_pb_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneOfThree); i {
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
	file_tests_pb_test_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Test_One)(nil),
		(*Test_Two)(nil),
		(*Test_Three)(nil),
		(*Test_Four)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_pb_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_pb_test_proto_goTypes,
		DependencyIndexes: file_tests_pb_test_proto_depIdxs,
		EnumInfos:         file_tests_pb_test_proto_enumTypes,
		MessageInfos:      file_tests_pb_test_proto_msgTypes,
	}.Build()
	File_tests_pb_test_proto = out.File
	file_tests_pb_test_proto_rawDesc = nil
	file_tests_pb_test_proto_goTypes = nil
	file_tests_pb_test_proto_depIdxs = nil
}

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
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: tests/pb/test.proto

package pb

import (
	_ "go.linka.cloud/protoc-gen-defaults/defaults"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
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

type TestOptional_Type int32

const (
	TestOptional_NONE TestOptional_Type = 0
	TestOptional_ONE  TestOptional_Type = 1
	TestOptional_TWO  TestOptional_Type = 2
)

// Enum value maps for TestOptional_Type.
var (
	TestOptional_Type_name = map[int32]string{
		0: "NONE",
		1: "ONE",
		2: "TWO",
	}
	TestOptional_Type_value = map[string]int32{
		"NONE": 0,
		"ONE":  1,
		"TWO":  2,
	}
)

func (x TestOptional_Type) Enum() *TestOptional_Type {
	p := new(TestOptional_Type)
	*p = x
	return p
}

func (x TestOptional_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestOptional_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_pb_test_proto_enumTypes[1].Descriptor()
}

func (TestOptional_Type) Type() protoreflect.EnumType {
	return &file_tests_pb_test_proto_enumTypes[1]
}

func (x TestOptional_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestOptional_Type.Descriptor instead.
func (TestOptional_Type) EnumDescriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{1, 0}
}

type TestUnexported_Type int32

const (
	TestUnexported_NONE TestUnexported_Type = 0
	TestUnexported_ONE  TestUnexported_Type = 1
	TestUnexported_TWO  TestUnexported_Type = 2
)

// Enum value maps for TestUnexported_Type.
var (
	TestUnexported_Type_name = map[int32]string{
		0: "NONE",
		1: "ONE",
		2: "TWO",
	}
	TestUnexported_Type_value = map[string]int32{
		"NONE": 0,
		"ONE":  1,
		"TWO":  2,
	}
)

func (x TestUnexported_Type) Enum() *TestUnexported_Type {
	p := new(TestUnexported_Type)
	*p = x
	return p
}

func (x TestUnexported_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestUnexported_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_pb_test_proto_enumTypes[2].Descriptor()
}

func (TestUnexported_Type) Type() protoreflect.EnumType {
	return &file_tests_pb_test_proto_enumTypes[2]
}

func (x TestUnexported_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestUnexported_Type.Descriptor instead.
func (TestUnexported_Type) EnumDescriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{2, 0}
}

type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringField          string                  `protobuf:"bytes,1,opt,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
	NumberField          int64                   `protobuf:"varint,2,opt,name=number_field,json=numberField,proto3" json:"number_field,omitempty"`
	BoolField            bool                    `protobuf:"varint,3,opt,name=bool_field,json=boolField,proto3" json:"bool_field,omitempty"`
	EnumField            Test_Type               `protobuf:"varint,4,opt,name=enum_field,json=enumField,proto3,enum=tests.Test_Type" json:"enum_field,omitempty"`
	MessageField         *Test                   `protobuf:"bytes,5,opt,name=message_field,json=messageField,proto3" json:"message_field,omitempty"`
	RepeatedStringField  []string                `protobuf:"bytes,6,rep,name=repeated_string_field,json=repeatedStringField,proto3" json:"repeated_string_field,omitempty"`
	RepeatedEnumField    []Test_Type             `protobuf:"varint,7,rep,packed,name=repeated_enum_field,json=repeatedEnumField,proto3,enum=tests.Test_Type" json:"repeated_enum_field,omitempty"`
	RepeatedMessageField []*Test                 `protobuf:"bytes,8,rep,name=repeated_message_field,json=repeatedMessageField,proto3" json:"repeated_message_field,omitempty"`
	NumberValueField     *wrapperspb.Int64Value  `protobuf:"bytes,9,opt,name=number_value_field,json=numberValueField,proto3" json:"number_value_field,omitempty"`
	StringValueField     *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=string_value_field,json=stringValueField,proto3" json:"string_value_field,omitempty"`
	BoolValueField       *wrapperspb.BoolValue   `protobuf:"bytes,11,opt,name=bool_value_field,json=boolValueField,proto3" json:"bool_value_field,omitempty"`
	TimeValueField       *timestamppb.Timestamp  `protobuf:"bytes,12,opt,name=time_value_field,json=timeValueField,proto3" json:"time_value_field,omitempty"`
	DurationValueField   *durationpb.Duration    `protobuf:"bytes,13,opt,name=duration_value_field,json=durationValueField,proto3" json:"duration_value_field,omitempty"`
	// Types that are assignable to Oneof:
	//	*Test_One
	//	*Test_Two
	//	*Test_Three
	//	*Test_Four
	Oneof                     isTest_Oneof                  `protobuf_oneof:"oneof"`
	Descriptor_               *descriptorpb.DescriptorProto `protobuf:"bytes,18,opt,name=descriptor,proto3" json:"descriptor,omitempty"`
	TimeValueFieldWithDefault *timestamppb.Timestamp        `protobuf:"bytes,19,opt,name=time_value_field_with_default,json=timeValueFieldWithDefault,proto3" json:"time_value_field_with_default,omitempty"`
	Bytes                     []byte                        `protobuf:"bytes,20,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Xid                       string                        `protobuf:"bytes,21,opt,name=xid,proto3" json:"xid,omitempty"`
	Uuid                      string                        `protobuf:"bytes,22,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XidPointer                *string                       `protobuf:"bytes,23,opt,name=xid_pointer,json=xidPointer,proto3,oneof" json:"xid_pointer,omitempty"`
	UuidPointer               *string                       `protobuf:"bytes,24,opt,name=uuid_pointer,json=uuidPointer,proto3,oneof" json:"uuid_pointer,omitempty"`
	XidValue                  *wrapperspb.StringValue       `protobuf:"bytes,25,opt,name=xid_value,json=xidValue,proto3" json:"xid_value,omitempty"`
	UuidValue                 *wrapperspb.StringValue       `protobuf:"bytes,26,opt,name=uuid_value,json=uuidValue,proto3" json:"uuid_value,omitempty"`
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

func (x *Test) GetRepeatedEnumField() []Test_Type {
	if x != nil {
		return x.RepeatedEnumField
	}
	return nil
}

func (x *Test) GetRepeatedMessageField() []*Test {
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

func (x *Test) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *Test) GetXid() string {
	if x != nil {
		return x.Xid
	}
	return ""
}

func (x *Test) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Test) GetXidPointer() string {
	if x != nil && x.XidPointer != nil {
		return *x.XidPointer
	}
	return ""
}

func (x *Test) GetUuidPointer() string {
	if x != nil && x.UuidPointer != nil {
		return *x.UuidPointer
	}
	return ""
}

func (x *Test) GetXidValue() *wrapperspb.StringValue {
	if x != nil {
		return x.XidValue
	}
	return nil
}

func (x *Test) GetUuidValue() *wrapperspb.StringValue {
	if x != nil {
		return x.UuidValue
	}
	return nil
}

type isTest_Oneof interface {
	isTest_Oneof()
}

type Test_One struct {
	One *OneOfOne `protobuf:"bytes,14,opt,name=one,proto3,oneof"`
}

type Test_Two struct {
	Two *OneOfTwo `protobuf:"bytes,15,opt,name=two,proto3,oneof"`
}

type Test_Three struct {
	Three *OneOfThree `protobuf:"bytes,16,opt,name=three,proto3,oneof"`
}

type Test_Four struct {
	Four Test_Type `protobuf:"varint,17,opt,name=four,proto3,enum=tests.Test_Type,oneof"`
}

func (*Test_One) isTest_Oneof() {}

func (*Test_Two) isTest_Oneof() {}

func (*Test_Three) isTest_Oneof() {}

func (*Test_Four) isTest_Oneof() {}

type TestOptional struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringField *string            `protobuf:"bytes,1,opt,name=string_field,json=stringField,proto3,oneof" json:"string_field,omitempty"`
	NumberField *int64             `protobuf:"varint,2,opt,name=number_field,json=numberField,proto3,oneof" json:"number_field,omitempty"`
	BoolField   *bool              `protobuf:"varint,3,opt,name=bool_field,json=boolField,proto3,oneof" json:"bool_field,omitempty"`
	EnumField   *TestOptional_Type `protobuf:"varint,4,opt,name=enum_field,json=enumField,proto3,enum=tests.TestOptional_Type,oneof" json:"enum_field,omitempty"`
}

func (x *TestOptional) Reset() {
	*x = TestOptional{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestOptional) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestOptional) ProtoMessage() {}

func (x *TestOptional) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TestOptional.ProtoReflect.Descriptor instead.
func (*TestOptional) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{1}
}

func (x *TestOptional) GetStringField() string {
	if x != nil && x.StringField != nil {
		return *x.StringField
	}
	return ""
}

func (x *TestOptional) GetNumberField() int64 {
	if x != nil && x.NumberField != nil {
		return *x.NumberField
	}
	return 0
}

func (x *TestOptional) GetBoolField() bool {
	if x != nil && x.BoolField != nil {
		return *x.BoolField
	}
	return false
}

func (x *TestOptional) GetEnumField() TestOptional_Type {
	if x != nil && x.EnumField != nil {
		return *x.EnumField
	}
	return TestOptional_NONE
}

type TestUnexported struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringField *string              `protobuf:"bytes,1,opt,name=string_field,json=stringField,proto3,oneof" json:"string_field,omitempty"`
	NumberField *int64               `protobuf:"varint,2,opt,name=number_field,json=numberField,proto3,oneof" json:"number_field,omitempty"`
	BoolField   *bool                `protobuf:"varint,3,opt,name=bool_field,json=boolField,proto3,oneof" json:"bool_field,omitempty"`
	EnumField   *TestUnexported_Type `protobuf:"varint,4,opt,name=enum_field,json=enumField,proto3,enum=tests.TestUnexported_Type,oneof" json:"enum_field,omitempty"`
}

func (x *TestUnexported) Reset() {
	*x = TestUnexported{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestUnexported) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestUnexported) ProtoMessage() {}

func (x *TestUnexported) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TestUnexported.ProtoReflect.Descriptor instead.
func (*TestUnexported) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{2}
}

func (x *TestUnexported) GetStringField() string {
	if x != nil && x.StringField != nil {
		return *x.StringField
	}
	return ""
}

func (x *TestUnexported) GetNumberField() int64 {
	if x != nil && x.NumberField != nil {
		return *x.NumberField
	}
	return 0
}

func (x *TestUnexported) GetBoolField() bool {
	if x != nil && x.BoolField != nil {
		return *x.BoolField
	}
	return false
}

func (x *TestUnexported) GetEnumField() TestUnexported_Type {
	if x != nil && x.EnumField != nil {
		return *x.EnumField
	}
	return TestUnexported_NONE
}

var File_tests_pb_test_proto protoreflect.FileDescriptor

var file_tests_pb_test_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x65, 0x73, 0x74, 0x73, 0x1a, 0x17, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x2f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f,
	0x70, 0x62, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c,
	0x0d, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0x9a,
	0x49, 0x0e, 0x72, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x28, 0x0a,
	0x0c, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x05, 0x9a, 0x49, 0x02, 0x20, 0x2a, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x24, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42, 0x05, 0x9a, 0x49, 0x02,
	0x68, 0x01, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x37, 0x0a,
	0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x10, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x42, 0x06, 0x9a, 0x49, 0x03, 0x80, 0x01, 0x02, 0x52, 0x09, 0x65, 0x6e, 0x75,
	0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x3a, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x42, 0x08, 0x9a, 0x49, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x13, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x40, 0x0a, 0x13, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x11, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45,
	0x6e, 0x75, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x4b, 0x0a, 0x16, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x42, 0x08, 0x9a, 0x49, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52,
	0x14, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x50, 0x0a, 0x12, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x05,
	0x9a, 0x49, 0x02, 0x20, 0x2b, 0x52, 0x10, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x5d, 0x0a, 0x12, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x11, 0x9a, 0x49, 0x0e, 0x72, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x4b, 0x0a, 0x10, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x05, 0x9a, 0x49,
	0x02, 0x68, 0x00, 0x52, 0x0e, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x4f, 0x0a, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x09, 0x9a, 0x49, 0x06, 0xb2, 0x01,
	0x03, 0x6e, 0x6f, 0x77, 0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x56, 0x0a, 0x14, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x9a,
	0x49, 0x06, 0xaa, 0x01, 0x03, 0x34, 0x32, 0x77, 0x52, 0x12, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2f, 0x0a, 0x03,
	0x6f, 0x6e, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x73, 0x2e, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x4f, 0x6e, 0x65, 0x42, 0x0a, 0x9a, 0x49, 0x07, 0x8a,
	0x01, 0x04, 0x08, 0x01, 0x10, 0x01, 0x48, 0x00, 0x52, 0x03, 0x6f, 0x6e, 0x65, 0x12, 0x2f, 0x0a,
	0x03, 0x74, 0x77, 0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x73, 0x2e, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x54, 0x77, 0x6f, 0x42, 0x0a, 0x9a, 0x49, 0x07,
	0x8a, 0x01, 0x04, 0x08, 0x01, 0x10, 0x01, 0x48, 0x00, 0x52, 0x03, 0x74, 0x77, 0x6f, 0x12, 0x35,
	0x0a, 0x05, 0x74, 0x68, 0x72, 0x65, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x54, 0x68, 0x72, 0x65, 0x65,
	0x42, 0x0a, 0x9a, 0x49, 0x07, 0x8a, 0x01, 0x04, 0x08, 0x01, 0x10, 0x01, 0x48, 0x00, 0x52, 0x05,
	0x74, 0x68, 0x72, 0x65, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x66, 0x6f, 0x75, 0x72, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0x9a, 0x49, 0x03, 0x80, 0x01, 0x01, 0x48, 0x00, 0x52,
	0x04, 0x66, 0x6f, 0x75, 0x72, 0x12, 0x4c, 0x0a, 0x0a, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0a, 0x9a, 0x49, 0x07,
	0x8a, 0x01, 0x04, 0x08, 0x01, 0x10, 0x01, 0x52, 0x0a, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x12, 0x78, 0x0a, 0x1d, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x1a, 0x9a, 0x49, 0x17, 0xb2, 0x01, 0x14, 0x31, 0x39,
	0x35, 0x32, 0x2d, 0x30, 0x33, 0x2d, 0x31, 0x31, 0x54, 0x30, 0x30, 0x3a, 0x30, 0x30, 0x3a, 0x30,
	0x30, 0x5a, 0x52, 0x19, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a,
	0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x07, 0x9a, 0x49,
	0x04, 0x7a, 0x02, 0x3f, 0x3f, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x03,
	0x78, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x9a, 0x49, 0x05, 0x72, 0x03,
	0x78, 0x69, 0x64, 0x52, 0x03, 0x78, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x9a, 0x49, 0x06, 0x72, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x0b, 0x78, 0x69, 0x64, 0x5f, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x9a, 0x49,
	0x05, 0x72, 0x03, 0x78, 0x69, 0x64, 0x48, 0x01, 0x52, 0x0a, 0x78, 0x69, 0x64, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x0c, 0x75, 0x75, 0x69, 0x64, 0x5f,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x9a,
	0x49, 0x06, 0x72, 0x04, 0x75, 0x75, 0x69, 0x64, 0x48, 0x02, 0x52, 0x0b, 0x75, 0x75, 0x69, 0x64,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x43, 0x0a, 0x09, 0x78, 0x69,
	0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x08, 0x9a, 0x49, 0x05,
	0x72, 0x03, 0x78, 0x69, 0x64, 0x52, 0x08, 0x78, 0x69, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x46, 0x0a, 0x0a, 0x75, 0x75, 0x69, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x1a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x09, 0x9a, 0x49, 0x06, 0x72, 0x04, 0x75, 0x75, 0x69, 0x64, 0x52, 0x09, 0x75, 0x75,
	0x69, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x22, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x4e, 0x45,
	0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x57, 0x4f, 0x10, 0x02, 0x42, 0x0f, 0x0a, 0x05, 0x6f,
	0x6e, 0x65, 0x6f, 0x66, 0x12, 0x06, 0x9a, 0x49, 0x03, 0x74, 0x77, 0x6f, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x78, 0x69, 0x64, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x42, 0x0f, 0x0a, 0x0d,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x22, 0xcd, 0x02,
	0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x39,
	0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0x9a, 0x49, 0x0e, 0x72, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0c, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x05, 0x9a, 0x49, 0x02, 0x20, 0x2a, 0x48, 0x01, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42, 0x05, 0x9a, 0x49,
	0x02, 0x68, 0x01, 0x48, 0x02, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x44, 0x0a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x06, 0x9a, 0x49, 0x03, 0x80, 0x01, 0x02, 0x48, 0x03, 0x52, 0x09, 0x65, 0x6e, 0x75,
	0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x22, 0x22, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4f,
	0x4e, 0x45, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x57, 0x4f, 0x10, 0x02, 0x42, 0x0f, 0x0a,
	0x0d, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0f,
	0x0a, 0x0d, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0xd6, 0x02,
	0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x55, 0x6e, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64,
	0x12, 0x39, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0x9a, 0x49, 0x0e, 0x72, 0x0c, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0c, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x05, 0x9a, 0x49, 0x02, 0x20, 0x2a, 0x48, 0x01, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0a, 0x62, 0x6f,
	0x6f, 0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42, 0x05,
	0x9a, 0x49, 0x02, 0x68, 0x01, 0x48, 0x02, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x88, 0x01, 0x01, 0x12, 0x46, 0x0a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x55, 0x6e, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0x9a, 0x49, 0x03, 0x80, 0x01, 0x02, 0x48, 0x03, 0x52,
	0x09, 0x65, 0x6e, 0x75, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x22, 0x22, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x07, 0x0a, 0x03, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x57, 0x4f, 0x10,
	0x02, 0x3a, 0x03, 0xa8, 0x49, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x62, 0x6f, 0x6f,
	0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x65, 0x6e, 0x75, 0x6d,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x6f, 0x2e, 0x6c, 0x69, 0x6e,
	0x6b, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_tests_pb_test_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_tests_pb_test_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tests_pb_test_proto_goTypes = []interface{}{
	(Test_Type)(0),                       // 0: tests.Test.Type
	(TestOptional_Type)(0),               // 1: tests.TestOptional.Type
	(TestUnexported_Type)(0),             // 2: tests.TestUnexported.Type
	(*Test)(nil),                         // 3: tests.Test
	(*TestOptional)(nil),                 // 4: tests.TestOptional
	(*TestUnexported)(nil),               // 5: tests.TestUnexported
	(*wrapperspb.Int64Value)(nil),        // 6: google.protobuf.Int64Value
	(*wrapperspb.StringValue)(nil),       // 7: google.protobuf.StringValue
	(*wrapperspb.BoolValue)(nil),         // 8: google.protobuf.BoolValue
	(*timestamppb.Timestamp)(nil),        // 9: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),          // 10: google.protobuf.Duration
	(*OneOfOne)(nil),                     // 11: tests.OneOfOne
	(*OneOfTwo)(nil),                     // 12: tests.OneOfTwo
	(*OneOfThree)(nil),                   // 13: tests.OneOfThree
	(*descriptorpb.DescriptorProto)(nil), // 14: google.protobuf.DescriptorProto
}
var file_tests_pb_test_proto_depIdxs = []int32{
	0,  // 0: tests.Test.enum_field:type_name -> tests.Test.Type
	3,  // 1: tests.Test.message_field:type_name -> tests.Test
	0,  // 2: tests.Test.repeated_enum_field:type_name -> tests.Test.Type
	3,  // 3: tests.Test.repeated_message_field:type_name -> tests.Test
	6,  // 4: tests.Test.number_value_field:type_name -> google.protobuf.Int64Value
	7,  // 5: tests.Test.string_value_field:type_name -> google.protobuf.StringValue
	8,  // 6: tests.Test.bool_value_field:type_name -> google.protobuf.BoolValue
	9,  // 7: tests.Test.time_value_field:type_name -> google.protobuf.Timestamp
	10, // 8: tests.Test.duration_value_field:type_name -> google.protobuf.Duration
	11, // 9: tests.Test.one:type_name -> tests.OneOfOne
	12, // 10: tests.Test.two:type_name -> tests.OneOfTwo
	13, // 11: tests.Test.three:type_name -> tests.OneOfThree
	0,  // 12: tests.Test.four:type_name -> tests.Test.Type
	14, // 13: tests.Test.descriptor:type_name -> google.protobuf.DescriptorProto
	9,  // 14: tests.Test.time_value_field_with_default:type_name -> google.protobuf.Timestamp
	7,  // 15: tests.Test.xid_value:type_name -> google.protobuf.StringValue
	7,  // 16: tests.Test.uuid_value:type_name -> google.protobuf.StringValue
	1,  // 17: tests.TestOptional.enum_field:type_name -> tests.TestOptional.Type
	2,  // 18: tests.TestUnexported.enum_field:type_name -> tests.TestUnexported.Type
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_tests_pb_test_proto_init() }
func file_tests_pb_test_proto_init() {
	if File_tests_pb_test_proto != nil {
		return
	}
	file_tests_pb_types_proto_init()
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
			switch v := v.(*TestOptional); i {
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
			switch v := v.(*TestUnexported); i {
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
	file_tests_pb_test_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_tests_pb_test_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_pb_test_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
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

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

// Code generated by protoc-gen-defaults. DO NOT EDIT.

package pb

import (
	"github.com/google/uuid"
	"github.com/rs/xid"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	_ *timestamppb.Timestamp
	_ *durationpb.Duration
	_ *wrapperspb.BoolValue
)

func (x *Test) Default() {
	if x.StringField == "" {
		x.StringField = "string_field"
	}
	if x.NumberField == 0 {
		x.NumberField = 42
	}
	if x.BoolField == false {
		x.BoolField = true
	}
	if x.EnumField == 0 {
		x.EnumField = 2
	}
	if v, ok := interface{}(x.MessageField).(interface{ Default() }); ok && x.MessageField != nil {
		v.Default()
	}
	for _, v := range x.RepeatedMessageField {
		if v, ok := interface{}(v).(interface{ Default() }); ok && v != nil {
			v.Default()
		}
	}
	if x.NumberValueField == nil {
		x.NumberValueField = &wrapperspb.Int64Value{Value: 43}
	}
	if x.StringValueField == nil {
		x.StringValueField = &wrapperspb.StringValue{Value: "string_value"}
	}
	if x.BoolValueField == nil {
		x.BoolValueField = &wrapperspb.BoolValue{Value: false}
	}
	if x.TimeValueField == nil {
		x.TimeValueField = timestamppb.Now()
	}
	if x.DurationValueField == nil {
		x.DurationValueField = durationpb.New(25401600000000000)
	}
	if x.Oneof == nil {
		x.Oneof = &Test_Two{}
	}
	switch x := x.Oneof.(type) {
	case *Test_One:
		if x.One == nil {
			x.One = &OneOfOne{}
		}
		if v, ok := interface{}(x.One).(interface{ Default() }); ok && x.One != nil {
			v.Default()
		}
	case *Test_Two:
		if x.Two == nil {
			x.Two = &OneOfTwo{}
		}
		if v, ok := interface{}(x.Two).(interface{ Default() }); ok && x.Two != nil {
			v.Default()
		}
	case *Test_Three:
		if x.Three == nil {
			x.Three = &OneOfThree{}
		}
		if v, ok := interface{}(x.Three).(interface{ Default() }); ok && x.Three != nil {
			v.Default()
		}
	case *Test_Four:
		if x.Four == 0 {
			x.Four = 1
		}
	}
	if x.Descriptor_ == nil {
		x.Descriptor_ = &descriptorpb.DescriptorProto{}
	}
	if v, ok := interface{}(x.Descriptor_).(interface{ Default() }); ok && x.Descriptor_ != nil {
		v.Default()
	}
	if x.TimeValueFieldWithDefault == nil {
		x.TimeValueFieldWithDefault = &timestamppb.Timestamp{Seconds: -562032000, Nanos: 0}

	}
	if len(x.Bytes) == 0 {
		x.Bytes = []byte("??")
	}
	if x.Xid == "" {
		x.Xid = xid.New().String()
	}
	if x.Uuid == "" {
		x.Uuid = uuid.New().String()
	}
	if x.XidPointer == nil {
		v := string(xid.New().String())
		x.XidPointer = &v
	}
	if x.UuidPointer == nil {
		v := string(uuid.New().String())
		x.UuidPointer = &v
	}
	if x.XidValue == nil {
		x.XidValue = &wrapperspb.StringValue{Value: xid.New().String()}
	}
	if x.UuidValue == nil {
		x.UuidValue = &wrapperspb.StringValue{Value: uuid.New().String()}
	}
}

func (x *TestOptional) Default() {
	if x.StringField == nil {
		v := string("string_field")
		x.StringField = &v
	}
	if x.NumberField == nil {
		v := int64(42)
		x.NumberField = &v
	}
	if x.BoolField == nil {
		v := bool(true)
		x.BoolField = &v
	}
	if x.EnumField == nil {
		v := TestOptional_Type(2)
		x.EnumField = &v
	}
}

func (x *TestUnexported) _Default() {
	if x.StringField == nil {
		v := string("string_field")
		x.StringField = &v
	}
	if x.NumberField == nil {
		v := int64(42)
		x.NumberField = &v
	}
	if x.BoolField == nil {
		v := bool(true)
		x.BoolField = &v
	}
	if x.EnumField == nil {
		v := TestUnexported_Type(2)
		x.EnumField = &v
	}
}

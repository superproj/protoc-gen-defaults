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

package tests

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rs/xid"
	assert2 "github.com/stretchr/testify/assert"
	require2 "github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"go.linka.cloud/protoc-gen-defaults/defaults"
	"go.linka.cloud/protoc-gen-defaults/tests/pb"
)

var (
	expect = &pb.Test{
		StringField:          "string_field",
		NumberField:          42,
		BoolField:            true,
		EnumField:            2,
		MessageField:         nil,
		RepeatedStringField:  nil,
		RepeatedMessageField: nil,
		NumberValueField:     wrapperspb.Int64(43),
		StringValueField:     wrapperspb.String("string_value"),
		BoolValueField:       wrapperspb.Bool(false),
		DurationValueField:   durationpb.New(25401600000000000),
		Oneof: &pb.Test_Two{
			Two: &pb.OneOfTwo{
				StringField: "string_field",
			},
		},
		Descriptor_:               &descriptorpb.DescriptorProto{},
		TimeValueFieldWithDefault: &timestamppb.Timestamp{Seconds: -562032000},
		Bytes:                     []byte("??"),
	}
)

func init() {
	m := proto.Clone(expect).(*pb.Test)
	expect.RepeatedMessageField = []*pb.Test{m}
}

func testAnCleanUpUuidXid(t *testing.T, test *pb.Test) {
	assert := assert2.New(t)
	require := require2.New(t)

	require.NotNil(test.UuidValue)
	require.NotNil(test.UuidPointer)
	for _, v := range []string{test.UuidValue.Value, *test.UuidPointer, test.Uuid} {
		_, err := uuid.Parse(v)
		assert.NoError(err)
	}
	test.Uuid = ""
	test.UuidValue = nil
	test.UuidPointer = nil

	require.NotNil(test.XidValue)
	require.NotNil(test.XidPointer)
	for _, v := range []string{test.XidValue.Value, *test.XidPointer, test.Xid} {
		_, err := xid.FromString(v)
		assert.NoError(err)
	}
	test.Xid = ""
	test.XidValue = nil
	test.XidPointer = nil
}

func TestDefaults(t *testing.T) {
	assert := assert2.New(t)
	require := require2.New(t)
	now := timestamppb.Now()

	test := &pb.Test{RepeatedMessageField: []*pb.Test{{}}}
	// test := &pb.Test{}
	test.Default()
	require.NotNil(test.TimeValueField)
	assert.InDelta(now.Seconds, test.TimeValueField.Seconds, 1)
	test.TimeValueField = nil
	require.Len(test.RepeatedMessageField, 1)
	require.NotNil(test.RepeatedMessageField[0])
	require.NotNil(test.RepeatedMessageField[0].TimeValueField)
	assert.InDelta(now.Seconds, test.RepeatedMessageField[0].TimeValueField.Seconds, 1)
	test.RepeatedMessageField[0].TimeValueField = nil

	testAnCleanUpUuidXid(t, test)
	testAnCleanUpUuidXid(t, test.RepeatedMessageField[0])

	assert.Equal(expect.String(), test.String())

	_, generated := interface{}(&pb.OneOfOne{}).(interface{ Default() })
	assert.False(generated)
}

func TestDefaultsReflect(t *testing.T) {
	assert := assert2.New(t)
	require := require2.New(t)
	now := timestamppb.Now()

	test := &pb.Test{RepeatedMessageField: []*pb.Test{{}}}
	defaults.Apply(test)

	require.NotNil(test.TimeValueField)
	assert.InDelta(now.Seconds, test.TimeValueField.Seconds, 1)
	test.TimeValueField = nil

	require.Len(test.RepeatedMessageField, 1)
	require.NotNil(test.RepeatedMessageField[0])
	require.NotNil(test.RepeatedMessageField[0].TimeValueField)
	assert.InDelta(now.Seconds, test.RepeatedMessageField[0].TimeValueField.Seconds, 1)
	test.RepeatedMessageField[0].TimeValueField = nil

	testAnCleanUpUuidXid(t, test)
	testAnCleanUpUuidXid(t, test.RepeatedMessageField[0])

	assert.Equal(expect.String(), test.String())

	_, generated := interface{}(&pb.OneOfOne{}).(interface{ Default() })
	assert.False(generated)

	expect.StringField = "other"
	test = &pb.Test{StringField: "other"}
	defaults.Apply(test)
	require.NotNil(test.TimeValueField)
	assert.InDelta(now.Seconds, test.TimeValueField.Seconds, 1)
	test.TimeValueField = nil

	require.Len(test.RepeatedMessageField, 0)

	testAnCleanUpUuidXid(t, test)
	expect.RepeatedMessageField = nil

	assert.Equal(expect.String(), test.String())
}

func TestDefaultsReflectOptionals(t *testing.T) {
	assert := assert2.New(t)

	enum := pb.TestOptional_TWO
	expect := &pb.TestOptional{
		StringField: proto.String("string_field"),
		NumberField: proto.Int64(42),
		BoolField:   proto.Bool(true),
		EnumField:   &enum,
	}

	test := &pb.TestOptional{}
	defaults.Apply(test)

	assert.True(proto.Equal(expect, test))

	expect.StringField = proto.String("other")
	test = &pb.TestOptional{
		StringField: proto.String("other"),
	}
	defaults.Apply(test)

	assert.True(proto.Equal(expect, test))

}

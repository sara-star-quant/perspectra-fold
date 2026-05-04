//go:build tools
// +build tools

package coreclient

// Deprecated: retained for reference after switching to generated stubs.

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func coreDescriptor() (protoreflect.FileDescriptor, error) {
	file := &descriptorpb.FileDescriptorProto{
		Syntax:  proto.String("proto3"),
		Name:    proto.String("core_compute.proto"),
		Package: proto.String("mdqc.core.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: proto.String("CosineSimilarityRequest"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:   proto.String("a"),
						Number: proto.Int32(1),
						Label:  descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
						Type:   descriptorpb.FieldDescriptorProto_TYPE_DOUBLE.Enum(),
					},
					{
						Name:   proto.String("b"),
						Number: proto.Int32(2),
						Label:  descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
						Type:   descriptorpb.FieldDescriptorProto_TYPE_DOUBLE.Enum(),
					},
				},
			},
			{
				Name: proto.String("CosineSimilarityResponse"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:   proto.String("value"),
						Number: proto.Int32(1),
						Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
						Type:   descriptorpb.FieldDescriptorProto_TYPE_DOUBLE.Enum(),
					},
				},
			},
		},
		Service: []*descriptorpb.ServiceDescriptorProto{
			{
				Name: proto.String("CoreCompute"),
				Method: []*descriptorpb.MethodDescriptorProto{
					{
						Name:       proto.String("CosineSimilarity"),
						InputType:  proto.String(".mdqc.core.v1.CosineSimilarityRequest"),
						OutputType: proto.String(".mdqc.core.v1.CosineSimilarityResponse"),
					},
				},
			},
		},
	}

	return protodesc.NewFile(file, nil)
}

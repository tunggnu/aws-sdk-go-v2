// Code generated by smithy-go-codegen DO NOT EDIT.

package ebs

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ebs/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpCompleteSnapshot struct {
}

func (*awsRestjson1_serializeOpCompleteSnapshot) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpCompleteSnapshot) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CompleteSnapshotInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/snapshots/completion/{SnapshotId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsCompleteSnapshotInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsCompleteSnapshotInput(v *CompleteSnapshotInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ChangedBlocksCount != nil {
		locationName := "X-Amz-Changedblockscount"
		encoder.SetHeader(locationName).Integer(*v.ChangedBlocksCount)
	}

	if v.Checksum != nil && len(*v.Checksum) > 0 {
		locationName := "X-Amz-Checksum"
		encoder.SetHeader(locationName).String(*v.Checksum)
	}

	if len(v.ChecksumAggregationMethod) > 0 {
		locationName := "X-Amz-Checksum-Aggregation-Method"
		encoder.SetHeader(locationName).String(string(v.ChecksumAggregationMethod))
	}

	if len(v.ChecksumAlgorithm) > 0 {
		locationName := "X-Amz-Checksum-Algorithm"
		encoder.SetHeader(locationName).String(string(v.ChecksumAlgorithm))
	}

	if v.SnapshotId == nil || len(*v.SnapshotId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member SnapshotId must not be empty")}
	}
	if v.SnapshotId != nil {
		if err := encoder.SetURI("SnapshotId").String(*v.SnapshotId); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpGetSnapshotBlock struct {
}

func (*awsRestjson1_serializeOpGetSnapshotBlock) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetSnapshotBlock) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetSnapshotBlockInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/snapshots/{SnapshotId}/blocks/{BlockIndex}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsGetSnapshotBlockInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetSnapshotBlockInput(v *GetSnapshotBlockInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.BlockIndex == nil {
		return &smithy.SerializationError{Err: fmt.Errorf("input member BlockIndex must not be empty")}
	}
	if v.BlockIndex != nil {
		if err := encoder.SetURI("BlockIndex").Integer(*v.BlockIndex); err != nil {
			return err
		}
	}

	if v.BlockToken != nil {
		encoder.SetQuery("blockToken").String(*v.BlockToken)
	}

	if v.SnapshotId == nil || len(*v.SnapshotId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member SnapshotId must not be empty")}
	}
	if v.SnapshotId != nil {
		if err := encoder.SetURI("SnapshotId").String(*v.SnapshotId); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpListChangedBlocks struct {
}

func (*awsRestjson1_serializeOpListChangedBlocks) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListChangedBlocks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListChangedBlocksInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/snapshots/{SecondSnapshotId}/changedblocks")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListChangedBlocksInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListChangedBlocksInput(v *ListChangedBlocksInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.FirstSnapshotId != nil {
		encoder.SetQuery("firstSnapshotId").String(*v.FirstSnapshotId)
	}

	if v.MaxResults != nil {
		encoder.SetQuery("maxResults").Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		encoder.SetQuery("pageToken").String(*v.NextToken)
	}

	if v.SecondSnapshotId == nil || len(*v.SecondSnapshotId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member SecondSnapshotId must not be empty")}
	}
	if v.SecondSnapshotId != nil {
		if err := encoder.SetURI("SecondSnapshotId").String(*v.SecondSnapshotId); err != nil {
			return err
		}
	}

	if v.StartingBlockIndex != nil {
		encoder.SetQuery("startingBlockIndex").Integer(*v.StartingBlockIndex)
	}

	return nil
}

type awsRestjson1_serializeOpListSnapshotBlocks struct {
}

func (*awsRestjson1_serializeOpListSnapshotBlocks) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListSnapshotBlocks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListSnapshotBlocksInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/snapshots/{SnapshotId}/blocks")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListSnapshotBlocksInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListSnapshotBlocksInput(v *ListSnapshotBlocksInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.MaxResults != nil {
		encoder.SetQuery("maxResults").Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		encoder.SetQuery("pageToken").String(*v.NextToken)
	}

	if v.SnapshotId == nil || len(*v.SnapshotId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member SnapshotId must not be empty")}
	}
	if v.SnapshotId != nil {
		if err := encoder.SetURI("SnapshotId").String(*v.SnapshotId); err != nil {
			return err
		}
	}

	if v.StartingBlockIndex != nil {
		encoder.SetQuery("startingBlockIndex").Integer(*v.StartingBlockIndex)
	}

	return nil
}

type awsRestjson1_serializeOpPutSnapshotBlock struct {
}

func (*awsRestjson1_serializeOpPutSnapshotBlock) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpPutSnapshotBlock) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*PutSnapshotBlockInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/snapshots/{SnapshotId}/blocks/{BlockIndex}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "PUT"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsPutSnapshotBlockInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if !restEncoder.HasHeader("Content-Type") {
		ctx = smithyhttp.SetIsContentTypeDefaultValue(ctx, true)
		restEncoder.SetHeader("Content-Type").String("application/octet-stream")
	}

	if input.BlockData != nil {
		payload := input.BlockData
		if request, err = request.SetStream(payload); err != nil {
			return out, metadata, &smithy.SerializationError{Err: err}
		}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsPutSnapshotBlockInput(v *PutSnapshotBlockInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.BlockIndex == nil {
		return &smithy.SerializationError{Err: fmt.Errorf("input member BlockIndex must not be empty")}
	}
	if v.BlockIndex != nil {
		if err := encoder.SetURI("BlockIndex").Integer(*v.BlockIndex); err != nil {
			return err
		}
	}

	if v.Checksum != nil && len(*v.Checksum) > 0 {
		locationName := "X-Amz-Checksum"
		encoder.SetHeader(locationName).String(*v.Checksum)
	}

	if len(v.ChecksumAlgorithm) > 0 {
		locationName := "X-Amz-Checksum-Algorithm"
		encoder.SetHeader(locationName).String(string(v.ChecksumAlgorithm))
	}

	if v.DataLength != nil {
		locationName := "X-Amz-Data-Length"
		encoder.SetHeader(locationName).Integer(*v.DataLength)
	}

	if v.Progress != nil {
		locationName := "X-Amz-Progress"
		encoder.SetHeader(locationName).Integer(*v.Progress)
	}

	if v.SnapshotId == nil || len(*v.SnapshotId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member SnapshotId must not be empty")}
	}
	if v.SnapshotId != nil {
		if err := encoder.SetURI("SnapshotId").String(*v.SnapshotId); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpStartSnapshot struct {
}

func (*awsRestjson1_serializeOpStartSnapshot) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpStartSnapshot) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*StartSnapshotInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/snapshots")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentStartSnapshotInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsStartSnapshotInput(v *StartSnapshotInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentStartSnapshotInput(v *StartSnapshotInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ClientToken != nil {
		ok := object.Key("ClientToken")
		ok.String(*v.ClientToken)
	}

	if v.Description != nil {
		ok := object.Key("Description")
		ok.String(*v.Description)
	}

	if v.Encrypted != nil {
		ok := object.Key("Encrypted")
		ok.Boolean(*v.Encrypted)
	}

	if v.KmsKeyArn != nil {
		ok := object.Key("KmsKeyArn")
		ok.String(*v.KmsKeyArn)
	}

	if v.ParentSnapshotId != nil {
		ok := object.Key("ParentSnapshotId")
		ok.String(*v.ParentSnapshotId)
	}

	if v.Tags != nil {
		ok := object.Key("Tags")
		if err := awsRestjson1_serializeDocumentTags(v.Tags, ok); err != nil {
			return err
		}
	}

	if v.Timeout != nil {
		ok := object.Key("Timeout")
		ok.Integer(*v.Timeout)
	}

	if v.VolumeSize != nil {
		ok := object.Key("VolumeSize")
		ok.Long(*v.VolumeSize)
	}

	return nil
}

func awsRestjson1_serializeDocumentTag(v *types.Tag, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Key != nil {
		ok := object.Key("Key")
		ok.String(*v.Key)
	}

	if v.Value != nil {
		ok := object.Key("Value")
		ok.String(*v.Value)
	}

	return nil
}

func awsRestjson1_serializeDocumentTags(v []types.Tag, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentTag(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

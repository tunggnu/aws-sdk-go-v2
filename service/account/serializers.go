// Code generated by smithy-go-codegen DO NOT EDIT.

package account

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/account/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpDeleteAlternateContact struct {
}

func (*awsRestjson1_serializeOpDeleteAlternateContact) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDeleteAlternateContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteAlternateContactInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/deleteAlternateContact")
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
	if err := awsRestjson1_serializeOpDocumentDeleteAlternateContactInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsDeleteAlternateContactInput(v *DeleteAlternateContactInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentDeleteAlternateContactInput(v *DeleteAlternateContactInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AccountId != nil {
		ok := object.Key("AccountId")
		ok.String(*v.AccountId)
	}

	if len(v.AlternateContactType) > 0 {
		ok := object.Key("AlternateContactType")
		ok.String(string(v.AlternateContactType))
	}

	return nil
}

type awsRestjson1_serializeOpDisableRegion struct {
}

func (*awsRestjson1_serializeOpDisableRegion) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDisableRegion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DisableRegionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/disableRegion")
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
	if err := awsRestjson1_serializeOpDocumentDisableRegionInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsDisableRegionInput(v *DisableRegionInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentDisableRegionInput(v *DisableRegionInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AccountId != nil {
		ok := object.Key("AccountId")
		ok.String(*v.AccountId)
	}

	if v.RegionName != nil {
		ok := object.Key("RegionName")
		ok.String(*v.RegionName)
	}

	return nil
}

type awsRestjson1_serializeOpEnableRegion struct {
}

func (*awsRestjson1_serializeOpEnableRegion) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpEnableRegion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*EnableRegionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/enableRegion")
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
	if err := awsRestjson1_serializeOpDocumentEnableRegionInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsEnableRegionInput(v *EnableRegionInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentEnableRegionInput(v *EnableRegionInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AccountId != nil {
		ok := object.Key("AccountId")
		ok.String(*v.AccountId)
	}

	if v.RegionName != nil {
		ok := object.Key("RegionName")
		ok.String(*v.RegionName)
	}

	return nil
}

type awsRestjson1_serializeOpGetAlternateContact struct {
}

func (*awsRestjson1_serializeOpGetAlternateContact) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetAlternateContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetAlternateContactInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/getAlternateContact")
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
	if err := awsRestjson1_serializeOpDocumentGetAlternateContactInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsGetAlternateContactInput(v *GetAlternateContactInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentGetAlternateContactInput(v *GetAlternateContactInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AccountId != nil {
		ok := object.Key("AccountId")
		ok.String(*v.AccountId)
	}

	if len(v.AlternateContactType) > 0 {
		ok := object.Key("AlternateContactType")
		ok.String(string(v.AlternateContactType))
	}

	return nil
}

type awsRestjson1_serializeOpGetContactInformation struct {
}

func (*awsRestjson1_serializeOpGetContactInformation) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetContactInformation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetContactInformationInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/getContactInformation")
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
	if err := awsRestjson1_serializeOpDocumentGetContactInformationInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsGetContactInformationInput(v *GetContactInformationInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentGetContactInformationInput(v *GetContactInformationInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AccountId != nil {
		ok := object.Key("AccountId")
		ok.String(*v.AccountId)
	}

	return nil
}

type awsRestjson1_serializeOpGetRegionOptStatus struct {
}

func (*awsRestjson1_serializeOpGetRegionOptStatus) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetRegionOptStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetRegionOptStatusInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/getRegionOptStatus")
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
	if err := awsRestjson1_serializeOpDocumentGetRegionOptStatusInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsGetRegionOptStatusInput(v *GetRegionOptStatusInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentGetRegionOptStatusInput(v *GetRegionOptStatusInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AccountId != nil {
		ok := object.Key("AccountId")
		ok.String(*v.AccountId)
	}

	if v.RegionName != nil {
		ok := object.Key("RegionName")
		ok.String(*v.RegionName)
	}

	return nil
}

type awsRestjson1_serializeOpListRegions struct {
}

func (*awsRestjson1_serializeOpListRegions) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListRegions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListRegionsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/listRegions")
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
	if err := awsRestjson1_serializeOpDocumentListRegionsInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsListRegionsInput(v *ListRegionsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentListRegionsInput(v *ListRegionsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AccountId != nil {
		ok := object.Key("AccountId")
		ok.String(*v.AccountId)
	}

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	if v.RegionOptStatusContains != nil {
		ok := object.Key("RegionOptStatusContains")
		if err := awsRestjson1_serializeDocumentRegionOptStatusList(v.RegionOptStatusContains, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpPutAlternateContact struct {
}

func (*awsRestjson1_serializeOpPutAlternateContact) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpPutAlternateContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*PutAlternateContactInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/putAlternateContact")
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
	if err := awsRestjson1_serializeOpDocumentPutAlternateContactInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsPutAlternateContactInput(v *PutAlternateContactInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentPutAlternateContactInput(v *PutAlternateContactInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AccountId != nil {
		ok := object.Key("AccountId")
		ok.String(*v.AccountId)
	}

	if len(v.AlternateContactType) > 0 {
		ok := object.Key("AlternateContactType")
		ok.String(string(v.AlternateContactType))
	}

	if v.EmailAddress != nil {
		ok := object.Key("EmailAddress")
		ok.String(*v.EmailAddress)
	}

	if v.Name != nil {
		ok := object.Key("Name")
		ok.String(*v.Name)
	}

	if v.PhoneNumber != nil {
		ok := object.Key("PhoneNumber")
		ok.String(*v.PhoneNumber)
	}

	if v.Title != nil {
		ok := object.Key("Title")
		ok.String(*v.Title)
	}

	return nil
}

type awsRestjson1_serializeOpPutContactInformation struct {
}

func (*awsRestjson1_serializeOpPutContactInformation) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpPutContactInformation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*PutContactInformationInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/putContactInformation")
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
	if err := awsRestjson1_serializeOpDocumentPutContactInformationInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsPutContactInformationInput(v *PutContactInformationInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentPutContactInformationInput(v *PutContactInformationInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AccountId != nil {
		ok := object.Key("AccountId")
		ok.String(*v.AccountId)
	}

	if v.ContactInformation != nil {
		ok := object.Key("ContactInformation")
		if err := awsRestjson1_serializeDocumentContactInformation(v.ContactInformation, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentContactInformation(v *types.ContactInformation, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AddressLine1 != nil {
		ok := object.Key("AddressLine1")
		ok.String(*v.AddressLine1)
	}

	if v.AddressLine2 != nil {
		ok := object.Key("AddressLine2")
		ok.String(*v.AddressLine2)
	}

	if v.AddressLine3 != nil {
		ok := object.Key("AddressLine3")
		ok.String(*v.AddressLine3)
	}

	if v.City != nil {
		ok := object.Key("City")
		ok.String(*v.City)
	}

	if v.CompanyName != nil {
		ok := object.Key("CompanyName")
		ok.String(*v.CompanyName)
	}

	if v.CountryCode != nil {
		ok := object.Key("CountryCode")
		ok.String(*v.CountryCode)
	}

	if v.DistrictOrCounty != nil {
		ok := object.Key("DistrictOrCounty")
		ok.String(*v.DistrictOrCounty)
	}

	if v.FullName != nil {
		ok := object.Key("FullName")
		ok.String(*v.FullName)
	}

	if v.PhoneNumber != nil {
		ok := object.Key("PhoneNumber")
		ok.String(*v.PhoneNumber)
	}

	if v.PostalCode != nil {
		ok := object.Key("PostalCode")
		ok.String(*v.PostalCode)
	}

	if v.StateOrRegion != nil {
		ok := object.Key("StateOrRegion")
		ok.String(*v.StateOrRegion)
	}

	if v.WebsiteUrl != nil {
		ok := object.Key("WebsiteUrl")
		ok.String(*v.WebsiteUrl)
	}

	return nil
}

func awsRestjson1_serializeDocumentRegionOptStatusList(v []types.RegionOptStatus, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(string(v[i]))
	}
	return nil
}

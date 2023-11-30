// Code generated by smithy-go-codegen DO NOT EDIT.

package arczonalshift

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/arczonalshift/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpCancelZonalShift struct {
}

func (*awsRestjson1_serializeOpCancelZonalShift) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpCancelZonalShift) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CancelZonalShiftInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/zonalshifts/{zonalShiftId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "DELETE"
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

	if err := awsRestjson1_serializeOpHttpBindingsCancelZonalShiftInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsCancelZonalShiftInput(v *CancelZonalShiftInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ZonalShiftId == nil || len(*v.ZonalShiftId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member zonalShiftId must not be empty")}
	}
	if v.ZonalShiftId != nil {
		if err := encoder.SetURI("zonalShiftId").String(*v.ZonalShiftId); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpCreatePracticeRunConfiguration struct {
}

func (*awsRestjson1_serializeOpCreatePracticeRunConfiguration) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpCreatePracticeRunConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreatePracticeRunConfigurationInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/configuration")
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
	if err := awsRestjson1_serializeOpDocumentCreatePracticeRunConfigurationInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsCreatePracticeRunConfigurationInput(v *CreatePracticeRunConfigurationInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentCreatePracticeRunConfigurationInput(v *CreatePracticeRunConfigurationInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.BlockedDates != nil {
		ok := object.Key("blockedDates")
		if err := awsRestjson1_serializeDocumentBlockedDates(v.BlockedDates, ok); err != nil {
			return err
		}
	}

	if v.BlockedWindows != nil {
		ok := object.Key("blockedWindows")
		if err := awsRestjson1_serializeDocumentBlockedWindows(v.BlockedWindows, ok); err != nil {
			return err
		}
	}

	if v.BlockingAlarms != nil {
		ok := object.Key("blockingAlarms")
		if err := awsRestjson1_serializeDocumentControlConditions(v.BlockingAlarms, ok); err != nil {
			return err
		}
	}

	if v.OutcomeAlarms != nil {
		ok := object.Key("outcomeAlarms")
		if err := awsRestjson1_serializeDocumentControlConditions(v.OutcomeAlarms, ok); err != nil {
			return err
		}
	}

	if v.ResourceIdentifier != nil {
		ok := object.Key("resourceIdentifier")
		ok.String(*v.ResourceIdentifier)
	}

	return nil
}

type awsRestjson1_serializeOpDeletePracticeRunConfiguration struct {
}

func (*awsRestjson1_serializeOpDeletePracticeRunConfiguration) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDeletePracticeRunConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeletePracticeRunConfigurationInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/configuration/{resourceIdentifier}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "DELETE"
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

	if err := awsRestjson1_serializeOpHttpBindingsDeletePracticeRunConfigurationInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsDeletePracticeRunConfigurationInput(v *DeletePracticeRunConfigurationInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ResourceIdentifier == nil || len(*v.ResourceIdentifier) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member resourceIdentifier must not be empty")}
	}
	if v.ResourceIdentifier != nil {
		if err := encoder.SetURI("resourceIdentifier").String(*v.ResourceIdentifier); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpGetManagedResource struct {
}

func (*awsRestjson1_serializeOpGetManagedResource) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetManagedResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetManagedResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/managedresources/{resourceIdentifier}")
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

	if err := awsRestjson1_serializeOpHttpBindingsGetManagedResourceInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetManagedResourceInput(v *GetManagedResourceInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ResourceIdentifier == nil || len(*v.ResourceIdentifier) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member resourceIdentifier must not be empty")}
	}
	if v.ResourceIdentifier != nil {
		if err := encoder.SetURI("resourceIdentifier").String(*v.ResourceIdentifier); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpListAutoshifts struct {
}

func (*awsRestjson1_serializeOpListAutoshifts) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListAutoshifts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListAutoshiftsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/autoshifts")
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

	if err := awsRestjson1_serializeOpHttpBindingsListAutoshiftsInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListAutoshiftsInput(v *ListAutoshiftsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.MaxResults != nil {
		encoder.SetQuery("maxResults").Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		encoder.SetQuery("nextToken").String(*v.NextToken)
	}

	if len(v.Status) > 0 {
		encoder.SetQuery("status").String(string(v.Status))
	}

	return nil
}

type awsRestjson1_serializeOpListManagedResources struct {
}

func (*awsRestjson1_serializeOpListManagedResources) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListManagedResources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListManagedResourcesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/managedresources")
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

	if err := awsRestjson1_serializeOpHttpBindingsListManagedResourcesInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListManagedResourcesInput(v *ListManagedResourcesInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.MaxResults != nil {
		encoder.SetQuery("maxResults").Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		encoder.SetQuery("nextToken").String(*v.NextToken)
	}

	return nil
}

type awsRestjson1_serializeOpListZonalShifts struct {
}

func (*awsRestjson1_serializeOpListZonalShifts) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListZonalShifts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListZonalShiftsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/zonalshifts")
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

	if err := awsRestjson1_serializeOpHttpBindingsListZonalShiftsInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListZonalShiftsInput(v *ListZonalShiftsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.MaxResults != nil {
		encoder.SetQuery("maxResults").Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		encoder.SetQuery("nextToken").String(*v.NextToken)
	}

	if v.ResourceIdentifier != nil {
		encoder.SetQuery("resourceIdentifier").String(*v.ResourceIdentifier)
	}

	if len(v.Status) > 0 {
		encoder.SetQuery("status").String(string(v.Status))
	}

	return nil
}

type awsRestjson1_serializeOpStartZonalShift struct {
}

func (*awsRestjson1_serializeOpStartZonalShift) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpStartZonalShift) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*StartZonalShiftInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/zonalshifts")
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
	if err := awsRestjson1_serializeOpDocumentStartZonalShiftInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsStartZonalShiftInput(v *StartZonalShiftInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentStartZonalShiftInput(v *StartZonalShiftInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AwayFrom != nil {
		ok := object.Key("awayFrom")
		ok.String(*v.AwayFrom)
	}

	if v.Comment != nil {
		ok := object.Key("comment")
		ok.String(*v.Comment)
	}

	if v.ExpiresIn != nil {
		ok := object.Key("expiresIn")
		ok.String(*v.ExpiresIn)
	}

	if v.ResourceIdentifier != nil {
		ok := object.Key("resourceIdentifier")
		ok.String(*v.ResourceIdentifier)
	}

	return nil
}

type awsRestjson1_serializeOpUpdatePracticeRunConfiguration struct {
}

func (*awsRestjson1_serializeOpUpdatePracticeRunConfiguration) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpUpdatePracticeRunConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdatePracticeRunConfigurationInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/configuration/{resourceIdentifier}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "PATCH"
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

	if err := awsRestjson1_serializeOpHttpBindingsUpdatePracticeRunConfigurationInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentUpdatePracticeRunConfigurationInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsUpdatePracticeRunConfigurationInput(v *UpdatePracticeRunConfigurationInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ResourceIdentifier == nil || len(*v.ResourceIdentifier) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member resourceIdentifier must not be empty")}
	}
	if v.ResourceIdentifier != nil {
		if err := encoder.SetURI("resourceIdentifier").String(*v.ResourceIdentifier); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentUpdatePracticeRunConfigurationInput(v *UpdatePracticeRunConfigurationInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.BlockedDates != nil {
		ok := object.Key("blockedDates")
		if err := awsRestjson1_serializeDocumentBlockedDates(v.BlockedDates, ok); err != nil {
			return err
		}
	}

	if v.BlockedWindows != nil {
		ok := object.Key("blockedWindows")
		if err := awsRestjson1_serializeDocumentBlockedWindows(v.BlockedWindows, ok); err != nil {
			return err
		}
	}

	if v.BlockingAlarms != nil {
		ok := object.Key("blockingAlarms")
		if err := awsRestjson1_serializeDocumentControlConditions(v.BlockingAlarms, ok); err != nil {
			return err
		}
	}

	if v.OutcomeAlarms != nil {
		ok := object.Key("outcomeAlarms")
		if err := awsRestjson1_serializeDocumentControlConditions(v.OutcomeAlarms, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpUpdateZonalAutoshiftConfiguration struct {
}

func (*awsRestjson1_serializeOpUpdateZonalAutoshiftConfiguration) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpUpdateZonalAutoshiftConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdateZonalAutoshiftConfigurationInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/managedresources/{resourceIdentifier}")
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

	if err := awsRestjson1_serializeOpHttpBindingsUpdateZonalAutoshiftConfigurationInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentUpdateZonalAutoshiftConfigurationInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsUpdateZonalAutoshiftConfigurationInput(v *UpdateZonalAutoshiftConfigurationInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ResourceIdentifier == nil || len(*v.ResourceIdentifier) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member resourceIdentifier must not be empty")}
	}
	if v.ResourceIdentifier != nil {
		if err := encoder.SetURI("resourceIdentifier").String(*v.ResourceIdentifier); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentUpdateZonalAutoshiftConfigurationInput(v *UpdateZonalAutoshiftConfigurationInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.ZonalAutoshiftStatus) > 0 {
		ok := object.Key("zonalAutoshiftStatus")
		ok.String(string(v.ZonalAutoshiftStatus))
	}

	return nil
}

type awsRestjson1_serializeOpUpdateZonalShift struct {
}

func (*awsRestjson1_serializeOpUpdateZonalShift) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpUpdateZonalShift) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdateZonalShiftInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/zonalshifts/{zonalShiftId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "PATCH"
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

	if err := awsRestjson1_serializeOpHttpBindingsUpdateZonalShiftInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentUpdateZonalShiftInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsUpdateZonalShiftInput(v *UpdateZonalShiftInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ZonalShiftId == nil || len(*v.ZonalShiftId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member zonalShiftId must not be empty")}
	}
	if v.ZonalShiftId != nil {
		if err := encoder.SetURI("zonalShiftId").String(*v.ZonalShiftId); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentUpdateZonalShiftInput(v *UpdateZonalShiftInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Comment != nil {
		ok := object.Key("comment")
		ok.String(*v.Comment)
	}

	if v.ExpiresIn != nil {
		ok := object.Key("expiresIn")
		ok.String(*v.ExpiresIn)
	}

	return nil
}

func awsRestjson1_serializeDocumentBlockedDates(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentBlockedWindows(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentControlCondition(v *types.ControlCondition, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AlarmIdentifier != nil {
		ok := object.Key("alarmIdentifier")
		ok.String(*v.AlarmIdentifier)
	}

	if len(v.Type) > 0 {
		ok := object.Key("type")
		ok.String(string(v.Type))
	}

	return nil
}

func awsRestjson1_serializeDocumentControlConditions(v []types.ControlCondition, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentControlCondition(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

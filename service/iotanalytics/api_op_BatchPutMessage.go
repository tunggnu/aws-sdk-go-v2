// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sends messages to a channel.
func (c *Client) BatchPutMessage(ctx context.Context, params *BatchPutMessageInput, optFns ...func(*Options)) (*BatchPutMessageOutput, error) {
	if params == nil {
		params = &BatchPutMessageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchPutMessage", params, optFns, addOperationBatchPutMessageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchPutMessageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchPutMessageInput struct {

	// The name of the channel where the messages are sent.
	//
	// This member is required.
	ChannelName *string

	// The list of messages to be sent. Each message has the format: { "messageId":
	// "string", "payload": "string"}. The field names of message payloads (data) that
	// you send to AWS IoT Analytics:
	//
	// * Must contain only alphanumeric characters and
	// undescores (_). No other special characters are allowed.
	//
	// * Must begin with an
	// alphabetic character or single underscore (_).
	//
	// * Cannot contain hyphens (-).
	//
	// *
	// In regular expression terms:
	// "^[A-Za-z_]([A-Za-z0-9]*|[A-Za-z0-9][A-Za-z0-9_]*)$".
	//
	// * Cannot be more than 255
	// characters.
	//
	// * Are case insensitive. (Fields named foo and FOO in the same
	// payload are considered duplicates.)
	//
	// For example, {"temp_01": 29} or
	// {"_temp_01": 29} are valid, but {"temp-01": 29}, {"01_temp": 29} or
	// {"__temp_01": 29} are invalid in message payloads.
	//
	// This member is required.
	Messages []types.Message
}

type BatchPutMessageOutput struct {

	// A list of any errors encountered when sending the messages to the channel.
	BatchPutMessageErrorEntries []types.BatchPutMessageErrorEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchPutMessageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchPutMessage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchPutMessage{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpBatchPutMessageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchPutMessage(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opBatchPutMessage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotanalytics",
		OperationName: "BatchPutMessage",
	}
}
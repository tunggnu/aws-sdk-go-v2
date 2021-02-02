// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables the Baidu channel for an application or updates the status and settings
// of the Baidu channel for an application.
func (c *Client) UpdateBaiduChannel(ctx context.Context, params *UpdateBaiduChannelInput, optFns ...func(*Options)) (*UpdateBaiduChannelOutput, error) {
	if params == nil {
		params = &UpdateBaiduChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateBaiduChannel", params, optFns, addOperationUpdateBaiduChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateBaiduChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateBaiduChannelInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string

	// Specifies the status and settings of the Baidu (Baidu Cloud Push) channel for an
	// application.
	//
	// This member is required.
	BaiduChannelRequest *types.BaiduChannelRequest
}

type UpdateBaiduChannelOutput struct {

	// Provides information about the status and settings of the Baidu (Baidu Cloud
	// Push) channel for an application.
	//
	// This member is required.
	BaiduChannelResponse *types.BaiduChannelResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateBaiduChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateBaiduChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateBaiduChannel{}, middleware.After)
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
	if err = addOpUpdateBaiduChannelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBaiduChannel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateBaiduChannel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "UpdateBaiduChannel",
	}
}
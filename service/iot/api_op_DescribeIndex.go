// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes a search index.
func (c *Client) DescribeIndex(ctx context.Context, params *DescribeIndexInput, optFns ...func(*Options)) (*DescribeIndexOutput, error) {
	if params == nil {
		params = &DescribeIndexInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeIndex", params, optFns, addOperationDescribeIndexMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeIndexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeIndexInput struct {

	// The index name.
	//
	// This member is required.
	IndexName *string
}

type DescribeIndexOutput struct {

	// The index name.
	IndexName *string

	// The index status.
	IndexStatus types.IndexStatus

	// Contains a value that specifies the type of indexing performed. Valid values
	// are:
	//
	// * REGISTRY – Your thing index contains only registry data.
	//
	// *
	// REGISTRY_AND_SHADOW - Your thing index contains registry data and shadow
	// data.
	//
	// * REGISTRY_AND_CONNECTIVITY_STATUS - Your thing index contains registry
	// data and thing connectivity status data.
	//
	// *
	// REGISTRY_AND_SHADOW_AND_CONNECTIVITY_STATUS - Your thing index contains registry
	// data, shadow data, and thing connectivity status data.
	Schema *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeIndexMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeIndex{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeIndex{}, middleware.After)
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
	if err = addOpDescribeIndexValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeIndex(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeIndex(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeIndex",
	}
}
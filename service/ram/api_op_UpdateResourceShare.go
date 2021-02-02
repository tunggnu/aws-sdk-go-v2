// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified resource share that you own.
func (c *Client) UpdateResourceShare(ctx context.Context, params *UpdateResourceShareInput, optFns ...func(*Options)) (*UpdateResourceShareOutput, error) {
	if params == nil {
		params = &UpdateResourceShareInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateResourceShare", params, optFns, addOperationUpdateResourceShareMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateResourceShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateResourceShareInput struct {

	// The Amazon Resource Name (ARN) of the resource share.
	//
	// This member is required.
	ResourceShareArn *string

	// Indicates whether principals outside your AWS organization can be associated
	// with a resource share.
	AllowExternalPrincipals *bool

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string

	// The name of the resource share.
	Name *string
}

type UpdateResourceShareOutput struct {

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string

	// Information about the resource share.
	ResourceShare *types.ResourceShare

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateResourceShareMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateResourceShare{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateResourceShare{}, middleware.After)
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
	if err = addOpUpdateResourceShareValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateResourceShare(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateResourceShare(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "UpdateResourceShare",
	}
}
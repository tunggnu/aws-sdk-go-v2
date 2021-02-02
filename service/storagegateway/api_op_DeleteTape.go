// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified virtual tape. This operation is only supported in the tape
// gateway type.
func (c *Client) DeleteTape(ctx context.Context, params *DeleteTapeInput, optFns ...func(*Options)) (*DeleteTapeOutput, error) {
	if params == nil {
		params = &DeleteTapeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTape", params, optFns, addOperationDeleteTapeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTapeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DeleteTapeInput
type DeleteTapeInput struct {

	// The unique Amazon Resource Name (ARN) of the gateway that the virtual tape to
	// delete is associated with. Use the ListGateways operation to return a list of
	// gateways for your account and AWS Region.
	//
	// This member is required.
	GatewayARN *string

	// The Amazon Resource Name (ARN) of the virtual tape to delete.
	//
	// This member is required.
	TapeARN *string

	// Set to TRUE to delete an archived tape that belongs to a custom pool with tape
	// retention lock. Only archived tapes with tape retention lock set to governance
	// can be deleted. Archived tapes with tape retention lock set to compliance can't
	// be deleted.
	BypassGovernanceRetention bool
}

// DeleteTapeOutput
type DeleteTapeOutput struct {

	// The Amazon Resource Name (ARN) of the deleted virtual tape.
	TapeARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteTapeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteTape{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteTape{}, middleware.After)
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
	if err = addOpDeleteTapeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTape(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteTape(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DeleteTape",
	}
}
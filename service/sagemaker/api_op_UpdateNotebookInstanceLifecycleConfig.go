// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a notebook instance lifecycle configuration created with the
// CreateNotebookInstanceLifecycleConfig API.
func (c *Client) UpdateNotebookInstanceLifecycleConfig(ctx context.Context, params *UpdateNotebookInstanceLifecycleConfigInput, optFns ...func(*Options)) (*UpdateNotebookInstanceLifecycleConfigOutput, error) {
	if params == nil {
		params = &UpdateNotebookInstanceLifecycleConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateNotebookInstanceLifecycleConfig", params, optFns, addOperationUpdateNotebookInstanceLifecycleConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateNotebookInstanceLifecycleConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateNotebookInstanceLifecycleConfigInput struct {

	// The name of the lifecycle configuration.
	//
	// This member is required.
	NotebookInstanceLifecycleConfigName *string

	// The shell script that runs only once, when you create a notebook instance. The
	// shell script must be a base64-encoded string.
	OnCreate []types.NotebookInstanceLifecycleHook

	// The shell script that runs every time you start a notebook instance, including
	// when you create the notebook instance. The shell script must be a base64-encoded
	// string.
	OnStart []types.NotebookInstanceLifecycleHook
}

type UpdateNotebookInstanceLifecycleConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateNotebookInstanceLifecycleConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateNotebookInstanceLifecycleConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateNotebookInstanceLifecycleConfig{}, middleware.After)
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
	if err = addOpUpdateNotebookInstanceLifecycleConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNotebookInstanceLifecycleConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateNotebookInstanceLifecycleConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "UpdateNotebookInstanceLifecycleConfig",
	}
}
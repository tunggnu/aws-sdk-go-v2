// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new image recipe. Image recipes define how images are configured,
// tested, and assessed.
func (c *Client) CreateImageRecipe(ctx context.Context, params *CreateImageRecipeInput, optFns ...func(*Options)) (*CreateImageRecipeOutput, error) {
	if params == nil {
		params = &CreateImageRecipeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateImageRecipe", params, optFns, addOperationCreateImageRecipeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateImageRecipeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateImageRecipeInput struct {

	// The idempotency token used to make this request idempotent.
	//
	// This member is required.
	ClientToken *string

	// The components of the image recipe.
	//
	// This member is required.
	Components []types.ComponentConfiguration

	// The name of the image recipe.
	//
	// This member is required.
	Name *string

	// The parent image of the image recipe. The value of the string can be the ARN of
	// the parent image or an AMI ID. The format for the ARN follows this example:
	// arn:aws:imagebuilder:us-west-2:aws:image/windows-server-2016-english-full-base-x86/xxxx.x.x.
	// You can provide the specific version that you want to use, or you can use a
	// wildcard in all of the fields. If you enter an AMI ID for the string value, you
	// must have access to the AMI, and the AMI must be in the same Region in which you
	// are using Image Builder.
	//
	// This member is required.
	ParentImage *string

	// The semantic version of the image recipe.
	//
	// This member is required.
	SemanticVersion *string

	// The block device mappings of the image recipe.
	BlockDeviceMappings []types.InstanceBlockDeviceMapping

	// The description of the image recipe.
	Description *string

	// The tags of the image recipe.
	Tags map[string]string

	// The working directory to be used during build and test workflows.
	WorkingDirectory *string
}

type CreateImageRecipeOutput struct {

	// The idempotency token used to make this request idempotent.
	ClientToken *string

	// The Amazon Resource Name (ARN) of the image recipe that was created by this
	// request.
	ImageRecipeArn *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateImageRecipeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateImageRecipe{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateImageRecipe{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateImageRecipeMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateImageRecipeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateImageRecipe(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateImageRecipe struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateImageRecipe) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateImageRecipe) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateImageRecipeInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateImageRecipeInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateImageRecipeMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateImageRecipe{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateImageRecipe(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "CreateImageRecipe",
	}
}
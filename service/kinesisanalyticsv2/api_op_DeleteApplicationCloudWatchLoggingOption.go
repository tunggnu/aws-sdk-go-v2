// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an Amazon CloudWatch log stream from an Kinesis Data Analytics
// application.
func (c *Client) DeleteApplicationCloudWatchLoggingOption(ctx context.Context, params *DeleteApplicationCloudWatchLoggingOptionInput, optFns ...func(*Options)) (*DeleteApplicationCloudWatchLoggingOptionOutput, error) {
	if params == nil {
		params = &DeleteApplicationCloudWatchLoggingOptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteApplicationCloudWatchLoggingOption", params, optFns, addOperationDeleteApplicationCloudWatchLoggingOptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteApplicationCloudWatchLoggingOptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteApplicationCloudWatchLoggingOptionInput struct {

	// The application name.
	//
	// This member is required.
	ApplicationName *string

	// The CloudWatchLoggingOptionId of the Amazon CloudWatch logging option to delete.
	// You can get the CloudWatchLoggingOptionId by using the DescribeApplication
	// operation.
	//
	// This member is required.
	CloudWatchLoggingOptionId *string

	// The version ID of the application. You can retrieve the application version ID
	// using DescribeApplication.
	//
	// This member is required.
	CurrentApplicationVersionId *int64
}

type DeleteApplicationCloudWatchLoggingOptionOutput struct {

	// The application's Amazon Resource Name (ARN).
	ApplicationARN *string

	// The version ID of the application. Kinesis Data Analytics updates the
	// ApplicationVersionId each time you change the CloudWatch logging options.
	ApplicationVersionId *int64

	// The descriptions of the remaining CloudWatch logging options for the
	// application.
	CloudWatchLoggingOptionDescriptions []types.CloudWatchLoggingOptionDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteApplicationCloudWatchLoggingOptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteApplicationCloudWatchLoggingOption{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteApplicationCloudWatchLoggingOption{}, middleware.After)
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
	if err = addOpDeleteApplicationCloudWatchLoggingOptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteApplicationCloudWatchLoggingOption(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteApplicationCloudWatchLoggingOption(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "DeleteApplicationCloudWatchLoggingOption",
	}
}
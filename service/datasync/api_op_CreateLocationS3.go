// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an endpoint for an Amazon S3 bucket. For more information, see
// https://docs.aws.amazon.com/datasync/latest/userguide/create-locations-cli.html#create-location-s3-cli
// in the AWS DataSync User Guide.
func (c *Client) CreateLocationS3(ctx context.Context, params *CreateLocationS3Input, optFns ...func(*Options)) (*CreateLocationS3Output, error) {
	if params == nil {
		params = &CreateLocationS3Input{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLocationS3", params, optFns, addOperationCreateLocationS3Middlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLocationS3Output)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateLocationS3Request
type CreateLocationS3Input struct {

	// The ARN of the Amazon S3 bucket. If the bucket is on an AWS Outpost, this must
	// be an access point ARN.
	//
	// This member is required.
	S3BucketArn *string

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM)
	// role that is used to access an Amazon S3 bucket. For detailed information about
	// using such a role, see Creating a Location for Amazon S3 in the AWS DataSync
	// User Guide.
	//
	// This member is required.
	S3Config *types.S3Config

	// If you are using DataSync on an AWS Outpost, specify the Amazon Resource Names
	// (ARNs) of the DataSync agents deployed on your Outpost. For more information
	// about launching a DataSync agent on an AWS Outpost, see outposts-agent.
	AgentArns []string

	// The Amazon S3 storage class that you want to store your files in when this
	// location is used as a task destination. For buckets in AWS Regions, the storage
	// class defaults to Standard. For buckets on AWS Outposts, the storage class
	// defaults to AWS S3 Outposts. For more information about S3 storage classes, see
	// Amazon S3 Storage Classes (http://aws.amazon.com/s3/storage-classes/). Some
	// storage classes have behaviors that can affect your S3 storage cost. For
	// detailed information, see using-storage-classes.
	S3StorageClass types.S3StorageClass

	// A subdirectory in the Amazon S3 bucket. This subdirectory in Amazon S3 is used
	// to read data from the S3 source location or write data to the S3 destination.
	Subdirectory *string

	// The key-value pair that represents the tag that you want to add to the location.
	// The value can be an empty string. We recommend using tags to name your
	// resources.
	Tags []types.TagListEntry
}

// CreateLocationS3Response
type CreateLocationS3Output struct {

	// The Amazon Resource Name (ARN) of the source Amazon S3 bucket location that is
	// created.
	LocationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateLocationS3Middlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLocationS3{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLocationS3{}, middleware.After)
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
	if err = addOpCreateLocationS3ValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLocationS3(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLocationS3(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "CreateLocationS3",
	}
}
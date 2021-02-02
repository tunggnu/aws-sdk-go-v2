// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the properties associated with a sentiment detection job. Use this
// operation to get the status of a detection job.
func (c *Client) DescribeSentimentDetectionJob(ctx context.Context, params *DescribeSentimentDetectionJobInput, optFns ...func(*Options)) (*DescribeSentimentDetectionJobOutput, error) {
	if params == nil {
		params = &DescribeSentimentDetectionJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSentimentDetectionJob", params, optFns, addOperationDescribeSentimentDetectionJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSentimentDetectionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSentimentDetectionJobInput struct {

	// The identifier that Amazon Comprehend generated for the job. The operation
	// returns this identifier in its response.
	//
	// This member is required.
	JobId *string
}

type DescribeSentimentDetectionJobOutput struct {

	// An object that contains the properties associated with a sentiment detection
	// job.
	SentimentDetectionJobProperties *types.SentimentDetectionJobProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeSentimentDetectionJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSentimentDetectionJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSentimentDetectionJob{}, middleware.After)
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
	if err = addOpDescribeSentimentDetectionJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSentimentDetectionJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeSentimentDetectionJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "DescribeSentimentDetectionJob",
	}
}
// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a campaign by deploying a solution version. When a client calls the
// GetRecommendations
// (https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetRecommendations.html)
// and GetPersonalizedRanking
// (https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetPersonalizedRanking.html)
// APIs, a campaign is specified in the request. Minimum Provisioned TPS and
// Auto-Scaling A transaction is a single GetRecommendations or
// GetPersonalizedRanking call. Transactions per second (TPS) is the throughput and
// unit of billing for Amazon Personalize. The minimum provisioned TPS
// (minProvisionedTPS) specifies the baseline throughput provisioned by Amazon
// Personalize, and thus, the minimum billing charge. If your TPS increases beyond
// minProvisionedTPS, Amazon Personalize auto-scales the provisioned capacity up
// and down, but never below minProvisionedTPS, to maintain a 70% utilization.
// There's a short time delay while the capacity is increased that might cause loss
// of transactions. It's recommended to start with a low minProvisionedTPS, track
// your usage using Amazon CloudWatch metrics, and then increase the
// minProvisionedTPS as necessary. Status A campaign can be in one of the following
// states:
//
// * CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// *
// DELETE PENDING > DELETE IN_PROGRESS
//
// To get the campaign status, call
// DescribeCampaign. Wait until the status of the campaign is ACTIVE before asking
// the campaign for recommendations. Related APIs
//
// * ListCampaigns
//
// *
// DescribeCampaign
//
// * UpdateCampaign
//
// * DeleteCampaign
func (c *Client) CreateCampaign(ctx context.Context, params *CreateCampaignInput, optFns ...func(*Options)) (*CreateCampaignOutput, error) {
	if params == nil {
		params = &CreateCampaignInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCampaign", params, optFns, addOperationCreateCampaignMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCampaignOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCampaignInput struct {

	// Specifies the requested minimum provisioned transactions (recommendations) per
	// second that Amazon Personalize will support.
	//
	// This member is required.
	MinProvisionedTPS *int32

	// A name for the new campaign. The campaign name must be unique within your
	// account.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the solution version to deploy.
	//
	// This member is required.
	SolutionVersionArn *string

	// The configuration details of a campaign.
	CampaignConfig *types.CampaignConfig
}

type CreateCampaignOutput struct {

	// The Amazon Resource Name (ARN) of the campaign.
	CampaignArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateCampaignMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCampaign{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCampaign{}, middleware.After)
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
	if err = addOpCreateCampaignValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCampaign(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCampaign(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "CreateCampaign",
	}
}
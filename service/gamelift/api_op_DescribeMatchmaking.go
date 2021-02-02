// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves one or more matchmaking tickets. Use this operation to retrieve ticket
// information, including--after a successful match is made--connection information
// for the resulting new game session. To request matchmaking tickets, provide a
// list of up to 10 ticket IDs. If the request is successful, a ticket object is
// returned for each requested ID that currently exists. This operation is not
// designed to be continually called to track matchmaking ticket status. This
// practice can cause you to exceed your API limit, which results in errors.
// Instead, as a best practice, set up an Amazon Simple Notification Service (SNS)
// to receive notifications, and provide the topic ARN in the matchmaking
// configuration. Continuously poling ticket status with DescribeMatchmaking should
// only be used for games in development with low matchmaking usage. Learn more
// Add FlexMatch to a Game Client
// (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-client.html)
// Set Up FlexMatch Event Notification
// (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-notification.html)
// Related operations
//
// * StartMatchmaking
//
// * DescribeMatchmaking
//
// *
// StopMatchmaking
//
// * AcceptMatch
//
// * StartMatchBackfill
func (c *Client) DescribeMatchmaking(ctx context.Context, params *DescribeMatchmakingInput, optFns ...func(*Options)) (*DescribeMatchmakingOutput, error) {
	if params == nil {
		params = &DescribeMatchmakingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMatchmaking", params, optFns, addOperationDescribeMatchmakingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMatchmakingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DescribeMatchmakingInput struct {

	// A unique identifier for a matchmaking ticket. You can include up to 10 ID
	// values.
	//
	// This member is required.
	TicketIds []string
}

// Represents the returned data in response to a request operation.
type DescribeMatchmakingOutput struct {

	// A collection of existing matchmaking ticket objects matching the request.
	TicketList []types.MatchmakingTicket

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeMatchmakingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMatchmaking{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMatchmaking{}, middleware.After)
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
	if err = addOpDescribeMatchmakingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMatchmaking(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeMatchmaking(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeMatchmaking",
	}
}